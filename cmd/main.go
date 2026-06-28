package main

import (
	"time"

	"github.100xBugShipper/rogue_like/internal/canvas"
	gameState "github.100xBugShipper/rogue_like/internal/game_state"
	playerInputs "github.100xBugShipper/rogue_like/internal/inputs"
	"github.100xBugShipper/rogue_like/internal/snake"
	"github.100xBugShipper/rogue_like/internal/renderer"
	"github.100xBugShipper/rogue_like/internal/world"
)

const (
	ROW = 15
	COL = 40
)

func main() {
	canvasMap := make([][]string, ROW)

	gameWorld := &world.World {
		Canvas: canvasMap,
		Snake: snake.CreateSnake(),
	}

	gs := gameState.GameState{
		World: gameWorld,
	}

	canvas.CreateCanvas(ROW, COL, gameWorld)
	gs.SpawnSnake()
	renderer.RenderGameMap(*gameWorld)

	playerInputs := playerInputs.CreatePlayerInputObj()

	go playerInputs.DetectKeys()

	ticker := time.NewTicker(1000 * time.Millisecond)
	//Game Loop
	for {
		<- ticker.C
		gs.MutateWorld(*gameWorld, playerInputs.MoveChan)
		renderer.ClearScreen()
		renderer.RenderGameMap(*gameWorld)
	}

}
