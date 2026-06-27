package renderer

import (
	"fmt"
	"os"
	"os/exec"

	"github.100xBugShipper/rogue_like/world"
)

func RenderGameMap(gameMap world.World) {
	for i := 0; i <= len(*gameMap.Canvas) - 1; i++ {
		for j := 0; j < len((*gameMap.Canvas)[i]); j++ {
			fmt.Print((*gameMap.Canvas)[i][j])
		}
		fmt.Println()
	}
}

func ClearScreen(clearObj *exec.Cmd) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout

	cmd.Run()
}
