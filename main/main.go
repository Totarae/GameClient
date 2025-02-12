package main

import (
	"awesomeProject4/structs"
	"awesomeProject4/wsclient"
	"image/color"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/imdraw"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var gameState structs.GameState

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Game Client",
		Bounds: pixel.R(0, 0, 800, 800),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatalf("Ошибка создания окна: %v", err)
	}

	// Подключение к WebSocket
	ws, err := wsclient.NewWebSocketClient("localhost:8080")
	if err != nil {
		log.Fatalf("Ошибка подключения к WebSocket: %v", err)
	}
	defer ws.Close()

	// Запускаем получение обновлений
	go ws.ListenForUpdates(&gameState)

	imd := imdraw.New(nil)

	for !win.Closed() {
		// Рисуем координатную сетку
		imd.Clear()
		imd.Color = color.RGBA{R: 100, G: 100, B: 100, A: 255}
		// Сраное округление
		for x := 0.0; x <= 800; x += 40 {
			imd.Push(pixel.V(x, 0), pixel.V(x, 1000))
			imd.Line(3)
		}
		for y := 0.0; y <= 800; y += 40 {
			imd.Push(pixel.V(0, y), pixel.V(1000, y))
			imd.Line(3)
		}

		// Убираем повторную очистку, чтобы не затирать сетку
		win.Clear(colornames.Gray)
		imd.Draw(win)

		imd.Clear()
		for _, player := range gameState.Players {
			cellX := float64(int(player.X/40) * 40)
			cellY := float64(int(player.Y/40) * 40)
			imd.Color = getClassColor(player.Class)
			log.Printf("Класс: %s игрока %s цвет: %v", player.Class, player.ID, imd.Color)
			imd.Push(pixel.V(cellX, cellY), pixel.V(cellX+40, cellY+40)) // пихаем в сетку
			imd.Rectangle(0)
		}
		imd.Draw(win)

		win.Update()
		time.Sleep(time.Millisecond * 16) // Ограничение FPS
	}
}

func init() {
	logFileName := time.Now().Format("02.01.06") + ".txt"
	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка создания файла лога: %v", err)
	}
	log.SetOutput(file)
}

func main() {
	pixelgl.Run(run)
}

func getClassColor(class string) color.RGBA {
	switch {
	case strings.HasPrefix(class, "Warrior"):
		log.Printf("Класс: Warrior")
		return color.RGBA{R: 255, G: 0, B: 0, A: 255} // Красный
	case strings.HasPrefix(class, "Thief"):
		log.Printf("Класс: Thief")
		return color.RGBA{R: 0, G: 0, B: 255, A: 255} // Синий
	case strings.HasPrefix(class, "Mage"):
		log.Printf("Класс: Mage")
		return color.RGBA{R: 128, G: 0, B: 128, A: 255} // Фиолетовый
	case strings.HasPrefix(class, "Archer"):
		log.Printf("Класс: Archer")
		return color.RGBA{R: 0, G: 255, B: 0, A: 255} // Зеленый
	default:
		return color.RGBA{R: 0, G: 0, B: 0, A: 255} // Белый (по умолчанию)
	}
}
