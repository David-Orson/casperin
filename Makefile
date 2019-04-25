# DO NOT USE THIS
# This file is provided to make it possible to compile project outside of GOPATH in OpenBench instance

EXE = casperin
ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

casperin:
	mkdir -p "${GOPATH}/src/github.com/David-Orson"
	ln -s -f "$(ROOT_DIR)" "${GOPATH}/src/github.com/David-Orson/casperin"
	go build -o $(EXE) casperin.go
	rm "${GOPATH}/src/github.com/David-Orson/casperin"
