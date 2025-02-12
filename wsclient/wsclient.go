package wsclient

import (
	"awesomeProject4/structs"
	"encoding/json"
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	conn *websocket.Conn
}

// новое соединение
func NewWebSocketClient(host string) (*WebSocketClient, error) {
	u := url.URL{Scheme: "ws", Host: host, Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	return &WebSocketClient{conn: conn}, nil
}

// слушает обновления от сервера
func (ws *WebSocketClient) ListenForUpdates(gameState *structs.GameState) {
	for {
		_, message, err := ws.conn.ReadMessage()
		if err != nil {
			log.Println("Ошибка чтения из WebSocket, переподключение:", err)
			ws.Reconnect()
			continue
		}
		// Декодируем общий формат
		var baseMessage struct {
			Type    string          `json:"type"`
			Payload json.RawMessage `json:"payload"`
		}
		err = json.Unmarshal(message, &baseMessage)
		if err != nil {
			log.Println("Ошибка разбора JSON:", err)
			continue
		}

		// Обрабатываем "init"
		// TODO: прорпботать state machine событий
		if baseMessage.Type == "init" {
			var newPlayer structs.Player
			err := json.Unmarshal(baseMessage.Payload, &newPlayer)
			if err != nil {
				log.Println("Ошибка разбора игрока:", err)
				continue
			}

			gameState.Players = append(gameState.Players, newPlayer)
			log.Printf("Добавлен игрок %d (%s) в координаты (%.2f, %.2f)", newPlayer.ID, newPlayer.Class, newPlayer.X, newPlayer.Y)
		}
	}
}

// закрывает соединение
func (ws *WebSocketClient) Close() {
	ws.conn.Close()
}

// Reconnect пытается переподключиться при разрыве соединения
func (ws *WebSocketClient) Reconnect() {
	for {
		log.Println("Попытка переподключения к WebSocket...")
		time.Sleep(3 * time.Second) // Задержка перед попыткой

		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/ws"}
		conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Println("Не удалось переподключиться:", err)
			continue
		}

		ws.conn = conn
		log.Println("Успешное переподключение к WebSocket!")
		return
	}
}
