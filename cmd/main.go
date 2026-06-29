package main

import (
	"time"

	"github.100xBugShipper/rogue_like/internal/canvas"
	gameState "github.100xBugShipper/rogue_like/internal/game_state"
	playerInputs "github.100xBugShipper/rogue_like/internal/inputs"
	"github.100xBugShipper/rogue_like/internal/renderer"
	"github.100xBugShipper/rogue_like/internal/snake"
	"github.100xBugShipper/rogue_like/internal/world"
	"github.com/theprimeagen/vim-with-me/pkg/assert"
)

const (
	ROW = 30
	COL = 80
)

func main() {
	assert.Assert(ROW > 0 && COL > 0, "Rows or Columns cant be zero")
	canvasMap := make([][]string, ROW)

	gameWorld := &world.World{
		Canvas: canvasMap,
		Snake:  snake.CreateSnake(),
	}

	gs := gameState.GameState{
		World: gameWorld,
	}

	canvas.CreateCanvas(ROW, COL, gameWorld)
	gs.SpawnSnake()
	renderer.RenderGameMap(*gameWorld)

	playerInputs := playerInputs.CreatePlayerInputObj()

	go playerInputs.DetectKeys()

	ticker := time.NewTicker(500 * time.Millisecond)
	// Game Loop
	for {
		<-ticker.C
		gs.MutateWorld(*gameWorld, playerInputs.MoveChan)
		go gs.AutoMove()
		renderer.ClearScreen()
		renderer.RenderGameMap(*gameWorld)
	}
}
