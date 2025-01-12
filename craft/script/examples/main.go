package main

import (
	"github.com/flint92/play-with-compiler/craft/script"
	"os"
)

func main() {
	script.NewScript().Run(os.Stdin, os.Stdout)
}
