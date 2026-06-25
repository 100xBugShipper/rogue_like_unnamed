package player

type Player struct {
	X         int
	Y         int
	Health    int
	Power     int
	KillCount int
}

func SpawnPlayer(gameCanvas *[][]string) *Player {
	player := "@"
	randRow := 10
	randCol := 10

	if !wallDetection(randRow, randCol, *gameCanvas) {
		(*gameCanvas)[randRow][randCol] = player
	}

	return &Player{
		X:         randRow,
		Y:         randCol,
		Health:    100,
		Power:     0,
		KillCount: 0,
	}
}

func wallDetection(row, col int, canvas [][]string) bool {
	return canvas[row][col] == "#"
}
