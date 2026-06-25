package main

import (
	"github.100xBugShipper/rogue_like/canvas"
	"github.100xBugShipper/rogue_like/player"
	"github.100xBugShipper/rogue_like/renderer"
)

var ROW int
var COL int

func main() {
	ROW = 15
	COL = 40
	mapCanvas := canvas.CreateCanvas(ROW, COL)
	player.SpawnPlayer(&mapCanvas)
	renderer.RenderGameMap(mapCanvas)
}
