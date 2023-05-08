package cookie

import (
	"os"
	_ "pick-up/backend/configs"
)

var appKey string
var appKeyBytes []byte

func init() {
	appKey = os.Getenv("APP_KEY")
	appKeyBytes = []byte(appKey)
}
