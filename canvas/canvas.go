package canvas

var CanvasMap [][]string

func createMemoryMap(row, col int) *[][]string {
	matrixMap := make([][]string, row)

	for i := range matrixMap {
		matrixMap[i] = make([]string, col)
	}

	return &matrixMap
}

func CreateCanvas(row, col int) {
	CanvasMap = *createMemoryMap(row, col)
	for i := range row {
		for j := range col {
			if i == 0 || j == 0 || i == row-1 || j == col-1 {
				CanvasMap[i][j] = "#"
			} else {
				CanvasMap[i][j] = "."
			}
		}
	}
}
