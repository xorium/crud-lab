package main

import (
	"os"
)

const (
	ModeModelStubs = "model-stubs"
)

func main() {
	if len(os.Args) < 2 {
		panic("arguments is empty")
	}
	mode := os.Args[1]
	args := make([]string, 0)
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}
	cwd, _ := os.Getwd()

	switch mode {
	case ModeModelStubs:
		runModelStubsGenerator(cwd, args)
	}
}
