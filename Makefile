# DO NOT USE THIS
# This file is provided to make it possible to compile project outside of GOPATH in OpenBench instance

EXE = casperin

casperin:
	mkdir -p "${GOPATH}/src/github.com/David-Orson"
	ln -s -f `pwd` "${GOPATH}/src/github.com/David-Orson/"
	go build -o $(EXE) casperin.go
