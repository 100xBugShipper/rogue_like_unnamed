package renderer

import (
	"fmt"

	"github.100xBugShipper/rogue_like/internal/snake"
	"github.100xBugShipper/rogue_like/internal/world"
)

func RenderGameMap(gameMap world.World) {

	for i := 0; i <= len(gameMap.Canvas) - 1; i++ {
		for j := 0; j < len(gameMap.Canvas[i]); j++ {
			fmt.Print(gameMap.Canvas[i][j])
		}
		fmt.Println()
	}
}

func ClearScreen() {
	fmt.Println("\033[H\033[2J")
}

func DrawBody(snk *snake.Snake, gameMap *world.World) {
	for _, item := range snk.SnakeQueue.SnakeBody {
		gameMap.Canvas[item.X][item.Y] = "O"
		gameMap.Canvas[item.X][item.Y] = "."
	}
}
