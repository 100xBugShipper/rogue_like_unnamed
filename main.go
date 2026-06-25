package main

import (

	"github.100xBugShipper/rogue_like/canvas"
	"github.100xBugShipper/rogue_like/renderer"
)

func main() {
	mapCanvas := canvas.CreateCanvas(15, 40)
	renderer.RenderGameMap(mapCanvas)
}
