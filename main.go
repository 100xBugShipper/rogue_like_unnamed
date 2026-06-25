package main

import (
	"github.100xBugShipper/rogue_like/canvas"
	gameState "github.100xBugShipper/rogue_like/game_state"
	playerInputs "github.100xBugShipper/rogue_like/inputs"
	"github.100xBugShipper/rogue_like/player"
	"github.100xBugShipper/rogue_like/renderer"
)

var ROW = 15
var COL = 40

func main() {
	player := player.CreatePlayer()
	gs := gameState.GameState{
		Player: player,
	}
	canvas.CreateCanvas(ROW, COL)
	gs.SpawnPlayer()
	renderer.RenderGameMap()

	playerInputs := playerInputs.CreatePlayerInputObj()

	playerInputs.Wg.Add(1)
	go playerInputs.DetectKeys()

	for {
		gs.MutateWorld(&canvas.CanvasMap, playerInputs.MoveChan)
		renderer.RenderGameMap()
	}
}
