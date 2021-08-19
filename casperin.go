package main

import (
	"os"

	"github.com/David-Orson/casperin/engine"
	"github.com/David-Orson/casperin/tuning"
	"github.com/David-Orson/casperin/uci"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "tune":
			tuning.Tune()
		case "trace-tune":
			tuning.TraceTune()
		case "bench":
			engine.Benchmark()
		}
		return
	}
	engine := engine.NewEngine()
	uci := uci.NewUciProtocol(&engine)
	uci.Run()
}
