package log

import (
	"encoding/json"
	"os"

	"github.100xBugShipper/rogue_like/internal/snake"
)

func WriteToFile(payload snake.Snake) {
	for _, v := range payload.SnakeQueue.SnakeBody {
		data, _ := json.Marshal(v)
		_ = os.WriteFile("logging.json", data, 0o644)
	}
}
