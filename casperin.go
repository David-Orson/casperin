package main

import (
	"github.com/David-Orson/casperin/engine"
	"github.com/David-Orson/casperin/evaluation"
	"github.com/David-Orson/casperin/uci"
	"os"
)

func main() {
	if len(os.Args) > 0 && os.Args[1] == "tune" {
		evaluation.Tune()
		return
	}
	e := engine.NewEngine()
	uci := uci.NewUciProtocol(e)
	uci.Run()
}
