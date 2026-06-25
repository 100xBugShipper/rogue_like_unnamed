package gameState

import (
	"github.100xBugShipper/rogue_like/canvas"
	"github.100xBugShipper/rogue_like/player"
)

type GameState struct {
	Player *player.Player
}

func (gs *GameState) SpawnPlayer() {
	if !gs.wallDetection(gs.Player.X, gs.Player.Y, canvas.CanvasMap) {
		canvas.CanvasMap[gs.Player.X][gs.Player.Y] = gs.Player.Symbol
	}
}

func (gs *GameState) wallDetection(row, col int, canvas [][]string) bool {
	return canvas[row][col] == "#"
}

func (gs *GameState) isValidMove(moveChan chan string) (string, bool) {
	for item := range moveChan {
		switch item {
		case "w":
			if canvas.CanvasMap[gs.Player.X][gs.Player.Y+1] != "#" {
				return "y++", true
			}
		case "s":
			if canvas.CanvasMap[gs.Player.X][gs.Player.Y-1] != "#" {
				return "y--", true
			}
		case "d":
			if canvas.CanvasMap[gs.Player.X+1][gs.Player.Y] != "#" {
				return "x++", true
			}
		case "a":
			if canvas.CanvasMap[gs.Player.X-1][gs.Player.Y] != "#" {
				return "x--", true
			}
		}
	}

	return "", false
}

func clearPreviousPosition(x, y int, canvas *[][]string) {
	(*canvas)[x][y] = "."
}

func movePlayer(x, y int, canvas *[][]string) {
	(*canvas)[x][y] = "@"
}

func (gs *GameState) MutateWorld(canvas *[][]string, moveChan chan string) {
	move, isValidMove := gs.isValidMove(moveChan)
	xPos := gs.Player.X
	yPos := gs.Player.Y

	if isValidMove {
		clearPreviousPosition(xPos, yPos, canvas)
		switch move {
		case "x++":
			xPos++
		case "x--":
			xPos--
		case "y++":
			yPos++
		case "y--":
			yPos--
		}
		movePlayer(xPos, yPos, canvas)
	}
}
