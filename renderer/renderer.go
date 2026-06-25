package renderer

import (
	"fmt"

	"github.100xBugShipper/rogue_like/canvas"
)

func RenderGameMap() {
	gameWorld := canvas.CanvasMap
	for i := 0; i <= len(gameWorld) - 1; i++ {
		for j := 0; j < len(gameWorld[i]); j++ {
			fmt.Print(gameWorld[i][j])
		}
		fmt.Println()
	}
}
