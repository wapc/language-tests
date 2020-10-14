package main

import (
	"strconv"

	"github.com/wapc/language-tests/tinygo/module"
)

func main() {
	module.Handlers{
		TestFunction: testFunction,
		TestUnary:    testUnary,
		TestDecode:   testDecode,
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

func testDecode(tests module.Tests) (string, error) {
	ret := "{\n"
	ret += strconv.FormatBool(tests.Required.BoolValue) + "\n"
	ret += strconv.FormatUint(uint64(tests.Required.U8Value), 10) + "\n"
	ret += strconv.FormatUint(uint64(tests.Required.U16Value), 10) + "\n"
	ret += strconv.FormatUint(uint64(tests.Required.U32Value), 10) + "\n"
	ret += strconv.FormatUint(tests.Required.U64Value, 10) + "\n"
	ret += strconv.FormatInt(int64(tests.Required.S8Value), 10) + "\n"
	ret += strconv.FormatInt(int64(tests.Required.S16Value), 10) + "\n"
	ret += strconv.FormatInt(int64(tests.Required.S32Value), 10) + "\n"
	ret += strconv.FormatInt(tests.Required.S64Value, 10) + "\n"
	ret += strconv.FormatFloat(float64(tests.Required.F32Value), 'e', 16, 64) + "\n"
	ret += strconv.FormatFloat(tests.Required.F64Value, 'e', 16, 64) + "\n"
	ret += tests.Required.StringValue + "\n"
	ret += string(tests.Required.BytesValue) + "\n"
	ret += "}"
	return ret, nil
}
