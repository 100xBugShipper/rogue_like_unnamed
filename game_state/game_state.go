package gameState

import (
	"github.100xBugShipper/rogue_like/player"
	world "github.100xBugShipper/rogue_like/world"
)

type GameState struct {
	Player *player.Player
}

func (gs *GameState) SpawnPlayer(gameWorld *world.World) {
	if !gs.wallDetection(gs.Player.X, gs.Player.Y, gameWorld.Canvas) {
		gameWorld.Canvas[gs.Player.X][gs.Player.Y] = gs.Player.Symbol
	}
}

func (gs *GameState) wallDetection(row, col int, canvas [][]string) bool {
	return canvas[row][col] == "#"
}

func (gs *GameState) isValidMove(gameWorld world.World, moveChan chan string) (string, bool) {
	for item := range moveChan {
		switch item {
		case "w":
			if gameWorld.Canvas[gs.Player.X][gs.Player.Y+1] != "#" {
				return "y++", true
			}
		case "s":
			if gameWorld.Canvas[gs.Player.X][gs.Player.Y-1] != "#" {
				return "y--", true
			}
		case "d":
			if gameWorld.Canvas[gs.Player.X+1][gs.Player.Y] != "#" {
				return "x++", true
			}
		case "a":
			if gameWorld.Canvas[gs.Player.X-1][gs.Player.Y] != "#" {
				return "x--", true
			}
		}
	}

	return "", false
}

func clearPreviousPosition(x, y int, canvas *[][]string) {
	(*canvas)[x][y] = "."
}

func movePlayer(x, y int, canvas [][]string) [][]string {
	canvas[x][y] = "@"
	return canvas
}

func (gs *GameState) MutateWorld(gameWorld world.World, moveChan chan string) world.World {
	move, isValidMove := gs.isValidMove(gameWorld, moveChan)
	xPos := gs.Player.X
	yPos := gs.Player.Y
	var updatedCanvas world.World

	clearPreviousPosition(xPos, yPos, &gameWorld.Canvas)
	if isValidMove {
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
		updatedCanvas.Canvas = movePlayer(xPos, yPos, gameWorld.Canvas)
	}
	return updatedCanvas
}
