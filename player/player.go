package player

type Player struct {
	Symbol    string
	X         int
	Y         int
	Health    int
	Power     int
	KillCount int
}

func CreatePlayer() *Player {
	randRow := 10
	randCol := 10

	return &Player{
		Symbol:    "@",
		X:         randRow,
		Y:         randCol,
		Health:    100,
		Power:     0,
		KillCount: 0,
	}
}
