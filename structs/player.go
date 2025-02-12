package structs

type GameState struct {
	Players []Player `json:"players"`
}

type Player struct {
	Class string  `json:"class"`
	ID    string  `json:"id"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}
