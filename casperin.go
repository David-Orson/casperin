package main

import (
	"github.com/David-Orson/casperin/engine"
	"github.com/David-Orson/casperin/evaluation"
	"github.com/David-Orson/casperin/uci"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "tune":
			evaluation.Tune()
		case "trace-tune":
			evaluation.TraceTune()
		case "bench":
			engine.Benchmark()
		}
		return
	}
	uci := uci.NewUciProtocol(engine.NewEngine())
	uci.Run()
}
