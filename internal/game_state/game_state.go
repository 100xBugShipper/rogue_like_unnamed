package gameState

import (
	"fmt"
	"math/rand/v2"
	"os"

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
	gs.World.Snake.FoodX = rand.IntN(29)
	gs.World.Snake.FoodY = rand.IntN(79)

	gs.World.Canvas[gs.World.Snake.FoodX][gs.World.Snake.FoodY] = "$"
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

func (gs *GameState) ateFood(x, y int) bool {

	if gs.World.Canvas[x][y] == "$" {
		return true
	}

	return false
}

func (gs *GameState) MoveSnake() {

	// Start creating the body of the snake..
	oldHead := queue.Cords {
		X: gs.World.Snake.X,
		Y: gs.World.Snake.Y,
	}

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

	//Old head becomes body
	gs.World.Snake.SnakeQueue.Append(oldHead)

	if !gs.ateFood(gs.World.Snake.X, gs.World.Snake.Y) {
		// If snake didnt eat food, remove body.. only head stays
		gs.World.Snake.Dequeue()
	}else {
		gs.SpawnFood()
	}

	movePlayer(gs.World.Snake.X, gs.World.Snake.Y, gs.World.Canvas)
}

// HACK: Clear everything and print everything
func (gs *GameState) Draw() {
    // Clear everything except walls
    for i := range gs.World.Canvas {
        for j := range gs.World.Canvas[i] {
            if gs.World.Canvas[i][j] != "#" {
                gs.World.Canvas[i][j] = "."
            }
        }
    }
    for _, part := range gs.World.Snake.SnakeQueue.SnakeBody {
        gs.World.Canvas[part.X][part.Y] = "O"
    }

    gs.World.Canvas[gs.World.Snake.X][gs.World.Snake.Y] = "@"
	gs.World.Canvas[gs.World.Snake.FoodX][gs.World.Snake.FoodY] = "$"
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

	gs.MoveSnake()
	gs.Draw()
}
