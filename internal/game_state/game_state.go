package gameState

import (
	"fmt"
	"os"

	world "github.100xBugShipper/rogue_like/internal/world"
)

type GameState struct {
	World *world.World
}

func CreateGameState() *GameState {
	return &GameState{
		World: &world.World{},
	}
}

func (gs *GameState) SpawnSnake() {
	if !gs.wallDetection(gs.World.Snake.X, gs.World.Snake.Y, gs.World.Canvas) {
		gs.World.Canvas[gs.World.Snake.X][gs.World.Snake.Y] = gs.World.Snake.Symbol
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
			if gameWorld.Canvas[gs.World.Snake.X+1][gs.World.Snake.Y] != "#" {
				return "x--", true
			}
		case "s":
			if gameWorld.Canvas[gs.World.Snake.X+1][gs.World.Snake.Y] != "#" {
				return "x++", true
			}
		case "d":
			if gameWorld.Canvas[gs.World.Snake.X][gs.World.Snake.Y+1] != "#" {
				return "y++", true
			}
		case "a":
			if gameWorld.Canvas[gs.World.Snake.X][gs.World.Snake.Y-1] != "#" {
				return "y--", true
			}
		case "q":
			fmt.Println("Thanks for playing")
			os.Exit(0)
		}
	} else if (gameWorld.Canvas[gs.World.Snake.X][gs.World.Snake.Y] == "#") ||
		(gameWorld.Canvas[gs.World.Snake.X][gs.World.Snake.Y] == "@") {
		fmt.Println("GAME OVER")
		os.Exit(1)
	}

	return "", false
}

func clearPreviousPosition(x, y int, canvas [][]string) {
	canvas[x][y] = "."
}

func movePlayer(x, y int, canvas [][]string) {
	canvas[x][y] = "@"
}

func (gs *GameState) AutoMove() {
	gs.World.Snake.SnakeQueue.Pop()
	gs.World.Snake.SnakeQueue.AddToHead()
}

func (gs *GameState) MutateWorld(gameWorld world.World, moveChan chan string) {
	move, isValidMove := gs.isValidMove(gameWorld, moveChan)

	clearPreviousPosition(gs.World.Snake.X, gs.World.Snake.Y, gameWorld.Canvas)
	if isValidMove {
		switch move {
		case "x++":
			gs.World.Snake.X++
		case "x--":
			gs.World.Snake.X--
		case "y++":
			gs.World.Snake.Y++
		case "y--":
			gs.World.Snake.Y--
		default:
		}
		movePlayer(gs.World.Snake.X, gs.World.Snake.Y, gameWorld.Canvas)
	}
}
