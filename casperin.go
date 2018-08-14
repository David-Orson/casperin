package main

import (
	"github.com/David-Orson/casperin/engine"
	"github.com/David-Orson/casperin/uci"
)

func main() {
	e := engine.NewEngine()
	uci := uci.NewUciProtocol(e)
	uci.Run()
}
