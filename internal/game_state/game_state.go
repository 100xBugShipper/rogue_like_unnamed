package gameState

import (
	"fmt"
	"math/rand/v2"
	"os"

	log "github.100xBugShipper/rogue_like/internal/logger"
	"github.100xBugShipper/rogue_like/internal/queue"
	world "github.100xBugShipper/rogue_like/internal/world"
)

type GameState struct {
	World *world.World
	Cords queue.Cords
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

func (gs *GameState) SpawnFood() {
	randRow := rand.IntN(29)
	randCol := rand.IntN(79)

	gs.World.Canvas[randRow][randCol] = "$"
}

func (gs *GameState) wallDetection(row, col int, canvas [][]string) bool {
	return canvas[row][col] == "#"
}

func (gs *GameState) isValidMove(gameWorld world.World, moveChan chan string) (string, bool) {
	select {
	case item, ok := <-moveChan:
		if ok {
			switch item {
			case "w":
				if gameWorld.Canvas[gs.World.Snake.X-1][gs.World.Snake.Y] != "#" {
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
	default:
		return "", false
	}
	return "", false
}

func clearPreviousPosition(x, y int, canvas [][]string) {
	canvas[x][y] = "."
}

func movePlayer(x, y int, canvas [][]string) {
	canvas[x][y] = "@"
}

func (gs *GameState) foodAhead() bool {
	if gs.World.Canvas[gs.World.Snake.X+1][gs.World.Snake.Y] == "$" {
		return true
	}
	if gs.World.Canvas[gs.World.Snake.X][gs.World.Snake.Y+1] == "$" {
		return true
	}
	if gs.World.Canvas[gs.World.Snake.X-1][gs.World.Snake.Y] == "$" {
		return true
	}
	if gs.World.Canvas[gs.World.Snake.X][gs.World.Snake.Y-1] == "$" {
		return true
	}
	return false
}

func (gs *GameState) MoveSnake() {
	clearPreviousPosition(gs.World.Snake.X, gs.World.Snake.Y, gs.World.Canvas)

	switch gs.World.Snake.Direction {
	case "up":
		gs.World.Snake.X--
	case "down":
		gs.World.Snake.X++
	case "left":
		gs.World.Snake.Y--
	case "right":
		gs.World.Snake.Y++
	}

	movePlayer(gs.World.Snake.X, gs.World.Snake.Y, gs.World.Canvas)
}

func (gs *GameState) MutateWorld(gameWorld world.World, moveChan chan string) {
	move, isValidMove := gs.isValidMove(gameWorld, moveChan)
	clearPreviousPosition(gs.World.Snake.X, gs.World.Snake.Y, gs.World.Canvas)

	if isValidMove {
		switch move {
		case "x++":
			gs.World.Snake.Direction = "down"
		case "x--":
			gs.World.Snake.Direction = "up"
		case "y++":
			gs.World.Snake.Direction = "right"
		case "y--":
			gs.World.Snake.Direction = "left"
		}
	}

	if gs.foodAhead() {
		gs.World.Snake.Grow()
		gs.SpawnFood()
		log.WriteToFile(*gs.World.Snake)
	}
	gs.MoveSnake()
}
