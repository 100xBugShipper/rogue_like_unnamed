package main

import (

	"github.100xBugShipper/rogue_like/canvas"
	gameState "github.100xBugShipper/rogue_like/game_state"
	playerInputs "github.100xBugShipper/rogue_like/inputs"
	"github.100xBugShipper/rogue_like/player"
	"github.100xBugShipper/rogue_like/renderer"
	"github.100xBugShipper/rogue_like/world"
)

const (
	ROW = 15
	COL = 40
)

func main() {
	canvasMap := make([][]string, ROW)

	player := player.CreatePlayer()
	gs := gameState.GameState{
		Player: player,
	}
	gameWorld := &world.World {
		Canvas: canvasMap,
	}
	canvas.CreateCanvas(ROW, COL, gameWorld)
	gs.SpawnPlayer(gameWorld)
	renderer.RenderGameMap(*gameWorld)

	playerInputs := playerInputs.CreatePlayerInputObj()

	go playerInputs.DetectKeys()
	for {
		gs.MutateWorld(*gameWorld, playerInputs.MoveChan)
		renderer.ClearScreen()
		renderer.RenderGameMap(*gameWorld)
	}

}
