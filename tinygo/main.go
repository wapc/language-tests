package main

import (
	"github.com/wapc/language-tests/tinygo/module"
)

func main() {
	module.Handlers{
		TestFunction: testFunction,
		TestUnary:    testUnary,
	}.Register()
}

func testFunction(required module.Required, optional module.Optional, maps module.Maps, lists module.Lists) (module.Tests, error) {
	// Echo arguments
	return module.Tests{
		Required: required,
		Optional: optional,
		Maps:     maps,
		Lists:    lists,
	}, nil
}

func testUnary(tests module.Tests) (module.Tests, error) {
	// Echo input
	return tests, nil
}
