package main

import (
	"github.100xBugShipper/rogue_like/canvas"
	gameState "github.100xBugShipper/rogue_like/game_state"
	playerInputs "github.100xBugShipper/rogue_like/inputs"
	"github.100xBugShipper/rogue_like/player"
	"github.100xBugShipper/rogue_like/renderer"
	"github.100xBugShipper/rogue_like/world"
)

var ROW = 15
var COL = 40

func main() {
	player := player.CreatePlayer()
	gs := gameState.GameState{
		Player: player,
	}
	gameWorld := &world.World {
		Canvas: make([][]string, ROW),
	}
	canvas.CreateCanvas(ROW, COL, gameWorld)
	gs.SpawnPlayer(gameWorld)
	renderer.RenderGameMap(*gameWorld)

	playerInputs := playerInputs.CreatePlayerInputObj()

	go playerInputs.DetectKeys()

	for {
		gs.MutateWorld(*gameWorld, playerInputs.MoveChan)
		renderer.RenderGameMap(*gameWorld)
	}
}
