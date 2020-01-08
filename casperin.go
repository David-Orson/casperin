package main

import (
	"github.com/David-Orson/casperin/engine"
	"github.com/David-Orson/casperin/evaluation"
	"github.com/David-Orson/casperin/uci"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "tune" {
			evaluation.Tune()
		} else if os.Args[1] == "bench" {
			engine.Benchmark()
		}
		return
	}
	uci := uci.NewUciProtocol(engine.NewEngine())
	uci.Run()
}
