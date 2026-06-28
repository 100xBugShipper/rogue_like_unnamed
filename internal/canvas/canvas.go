package canvas

import "github.100xBugShipper/rogue_like/internal/world"

func createMemoryMap(row, col int) [][]string {
	matrixMap := make([][]string, row)

	for i := range matrixMap {
		matrixMap[i] = make([]string, col)
	}

	return matrixMap
}

func CreateCanvas(row, col int, world *world.World) {
	canvasMap := createMemoryMap(row, col)
	for i := range row {
		for j := range col {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				canvasMap[i][j] = "#"
			} else {
				canvasMap[i][j] = "."
			}
		}
	}
	world.Canvas = canvasMap
}
