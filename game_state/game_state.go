package gameState

import (
	"os"

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
	item, ok := <-moveChan
	if ok {
		switch item {
		case "w":
			if gameWorld.Canvas[gs.Player.X+1][gs.Player.Y] != "#" {
				return "x--", true
			}
		case "s":
			if gameWorld.Canvas[gs.Player.X+1][gs.Player.Y] != "#" {
				return "x++", true
			}
		case "d":
			if gameWorld.Canvas[gs.Player.X][gs.Player.Y+1] != "#" {
				return "y++", true
			}
		case "a":
			if gameWorld.Canvas[gs.Player.X][gs.Player.Y-1] != "#" {
				return "y--", true
			}
		}
	} else {
		os.Exit(0)
	}

	return "", false
}

func clearPreviousPosition(x, y int, canvas [][]string) {
	canvas[x][y] = "."
}

func movePlayer(x, y int, canvas [][]string) {
	canvas[x][y] = "@"
}

func (gs *GameState) MutateWorld(gameWorld world.World, moveChan chan string) {
	move, isValidMove := gs.isValidMove(gameWorld, moveChan)

	clearPreviousPosition(gs.Player.X, gs.Player.Y, gameWorld.Canvas)
	if isValidMove {
		switch move {
		case "x++":
			gs.Player.X++
		case "x--":
			gs.Player.X--
		case "y++":
			gs.Player.Y++
		case "y--":
			gs.Player.Y--
		}
		movePlayer(gs.Player.X, gs.Player.Y, gameWorld.Canvas)
	}
}
