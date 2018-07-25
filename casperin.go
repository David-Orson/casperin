package main

import (
	"fmt"
	"github.com/David-Orson/casperin/backend"
)

func main() {
	backend.InitBB()
	position := backend.InitialPosition
	fmt.Println(backend.Perft(&position, 6))
}
