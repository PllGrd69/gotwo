package main

import (
	"github.com/PllGrd69/gotwo/routers"
)

func main() {

	r := routers.SetupRouter()
	r.Run("localhost:8085")
}
