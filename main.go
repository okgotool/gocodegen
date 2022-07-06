package main

import (
	"os"

	"github.com/okgotool/gocodegen/codegen"
)

func main() {

	// configFile := "./config/gen.yaml"
	configFile := "./config/gen_simple.yaml"

	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	codegen.StartGen(configFile)
}
