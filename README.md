

### Проект: Графический клиент для игры на Go

## Описание

Гграфический клиент для многопользовательской игры 
Клиент подключается к серверу через WebSockets, 
получает обновления игрового состояния и визуализирует его с использованием PixelGL.

## Основные возможности

WebSocket-соединение – клиент получает обновления состояния игры в реальном времени

Рендеринг игрового поля – сетка и игровые персонажи отрисовываются с помощью PixelGL

Обновление позиций игроков – клиент получает координаты персонажей по сетке

Обработка разрыва соединения – автоматическая переподключаемость при обрыве WebSocket

Логирование – файлы логов создаются автоматически для отладки

## Используемые технологии

Go – основной язык реализации.

PixelGL – библиотека для отрисовки графики.

Gorilla WebSocket – реализация WebSocket-клиента.

colornames – работа с цветами.

## Важное

```
go build -x .
```
Для билда дополнитлеьно
Поставить поддержку opengl и MSYS2
