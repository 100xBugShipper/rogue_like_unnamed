package world

import (
	"github.100xBugShipper/rogue_like/internal/snake"
)

type World struct {
	Canvas [][]string
	Snake *snake.Snake
}
