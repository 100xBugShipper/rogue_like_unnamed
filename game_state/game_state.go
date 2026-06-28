package gameState

import (
	"os"

	world "github.100xBugShipper/rogue_like/world"
)

type GameState struct {
	World *world.World
}

func CreateGameState() *GameState {
	return &GameState {
		World: &world.World{},
	}
}

func (gs *GameState) SpawnPlayer() {
	if !gs.wallDetection(gs.World.Player.X, gs.World.Player.Y, gs.World.Canvas) {
		gs.World.Canvas[gs.World.Player.X][gs.World.Player.Y] = gs.World.Player.Symbol
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
			if gameWorld.Canvas[gs.World.Player.X+1][gs.World.Player.Y] != "#" {
				return "x--", true
			}
		case "s":
			if gameWorld.Canvas[gs.World.Player.X+1][gs.World.Player.Y] != "#" {
				return "x++", true
			}
		case "d":
			if gameWorld.Canvas[gs.World.Player.X][gs.World.Player.Y+1] != "#" {
				return "y++", true
			}
		case "a":
			if gameWorld.Canvas[gs.World.Player.X][gs.World.Player.Y-1] != "#" {
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

	clearPreviousPosition(gs.World.Player.X, gs.World.Player.Y, gameWorld.Canvas)
	if isValidMove {
		switch move {
		case "x++":
			gs.World.Player.X++
		case "x--":
			gs.World.Player.X--
		case "y++":
			gs.World.Player.Y++
		case "y--":
			gs.World.Player.Y--
		}
		movePlayer(gs.World.Player.X, gs.World.Player.Y, gameWorld.Canvas)
	}
}
