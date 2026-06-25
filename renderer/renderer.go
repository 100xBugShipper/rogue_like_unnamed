package renderer

import "fmt"

func RenderGameMap(gameState [][]string) {
	for i := 0; i <= len(gameState) - 1; i++ {
		for j := 0; j < len(gameState[i]); j++ {
			fmt.Print(gameState[i][j])
		}
		fmt.Println()
	}
}
