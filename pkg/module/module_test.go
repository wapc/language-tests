package module_test

import (
	"context"
	"io/ioutil"
	"math"
	"strings"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/wapc/wapc-go"

	"github.com/wapc/language-tests/pkg/module"
)

func TestTinyGo(t *testing.T) {
	testLanguage(t, "../../build/tinygo.wasm")
}

func TestAssemblyScript(t *testing.T) {
	testLanguage(t, "../../build/assemblyscript.wasm")
}

func TestRust(t *testing.T) {
	testLanguage(t, "../../build/rust.wasm")
}

func testLanguage(t *testing.T, wasmFile string) {
	wapcModule, err := getModule(wasmFile)
	require.NoError(t, err, "could load Wasm module")
	defer wapcModule.Close()
	wapcInstance, err := wapcModule.Instantiate()
	require.NoError(t, err, "could instantiate module")
	m := module.New(wapcInstance)
	testEcho(t, m)
	testDecode(t, m)
}

func getModule(wasmFile string) (*wapc.Module, error) {
	code, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		return nil, err
	}
	wapcModule, err := wapc.New(code, func(ctx context.Context, binding, namespace, operation string, payload []byte) ([]byte, error) {
		return []byte("unimplemented"), nil
	})
	if err != nil {
		return nil, err
	}
	wapcModule.SetLogger(wapc.Println) // Send __console_log calls to stardard out
	wapcModule.SetWriter(wapc.Print)   // Send WASI fd_write calls to stardard out

	return wapcModule, nil
}

func testEcho(t *testing.T, m *module.Module) {
	ctx := context.Background()
	expected := module.Tests{
		Required: module.Required{
			BoolValue:   true,
			U8Value:     math.MaxUint8,
			U16Value:    math.MaxUint16,
			U32Value:    math.MaxUint32,
			U64Value:    math.MaxUint64,
			S8Value:     math.MinInt8,
			S16Value:    math.MinInt16,
			S32Value:    math.MinInt32,
			S64Value:    math.MinInt64,
			F32Value:    math.MaxFloat32,
			F64Value:    math.MaxFloat64,
			StringValue: "test",
			BytesValue:  []byte("test"),
			ObjectValue: module.Thing{
				Value: "test",
			},
		},
		Optional: module.Optional{
			U8Value:     pointer.ToUint8(math.MaxUint8),
			U16Value:    pointer.ToUint16(math.MaxUint16),
			U32Value:    pointer.ToUint32(math.MaxUint32),
			U64Value:    pointer.ToUint64(math.MaxUint64),
			S8Value:     pointer.ToInt8(math.MinInt8),
			S16Value:    pointer.ToInt16(math.MinInt16),
			S32Value:    pointer.ToInt32(math.MinInt32),
			S64Value:    pointer.ToInt64(math.MinInt64),
			F32Value:    pointer.ToFloat32(math.MaxFloat32),
			F64Value:    pointer.ToFloat64(math.MaxFloat64),
			StringValue: pointer.ToString("test"),
			BytesValue:  []byte("test"),
			ObjectValue: &module.Thing{
				Value: "test",
			},
		},
		Maps: module.Maps{
			MapStringPrimative: map[uint32]string{
				1234: "test",
			},
			MapU64Primative: map[uint32]uint64{
				5678: 01234,
			},
		},
		Lists: module.Lists{
			ListStrings:         []string{"test"},
			ListU64s:            []uint64{1234},
			ListObjects:         []module.Thing{{Value: "test"}},
			ListObjectsOptional: []*module.Thing{{Value: "test"}},
		},
	}
	actual, err := m.TestFunction(ctx, expected.Required, expected.Optional, expected.Maps, expected.Lists)
	require.NoError(t, err, "could not invoke testFunction")

	assert.Equal(t, expected.Required, actual.Required, "mismatch with required fields")
	assert.Equal(t, expected.Optional, actual.Optional, "mismatch with optional fields")
	assert.Equal(t, expected.Maps, actual.Maps, "mismatch with map fields")
	assert.Equal(t, expected.Lists, actual.Lists, "mismatch with list fields")

	actual, err = m.TestUnary(ctx, expected)
	require.NoError(t, err, "could not invoke testUnary")

	assert.Equal(t, expected.Required, actual.Required, "mismatch with required fields")
	assert.Equal(t, expected.Optional, actual.Optional, "mismatch with optional fields")
	assert.Equal(t, expected.Maps, actual.Maps, "mismatch with map fields")
	assert.Equal(t, expected.Lists, actual.Lists, "mismatch with list fields")
}

func testDecode(t *testing.T, m *module.Module) {
	ctx := context.Background()
	expected := module.Tests{
		Required: module.Required{
			BoolValue:   true,
			U8Value:     math.MaxUint8,
			U16Value:    math.MaxUint16,
			U32Value:    math.MaxUint32,
			U64Value:    math.MaxUint64,
			S8Value:     math.MinInt8,
			S16Value:    math.MinInt16,
			S32Value:    math.MinInt32,
			S64Value:    math.MinInt64,
			F32Value:    math.MaxFloat32,
			F64Value:    math.MaxFloat64,
			StringValue: "test",
			BytesValue:  []byte("test"),
			ObjectValue: module.Thing{
				Value: "test",
			},
		},
		Optional: module.Optional{
			U8Value:     pointer.ToUint8(math.MaxUint8),
			U16Value:    pointer.ToUint16(math.MaxUint16),
			U32Value:    pointer.ToUint32(math.MaxUint32),
			U64Value:    pointer.ToUint64(math.MaxUint64),
			S8Value:     pointer.ToInt8(math.MinInt8),
			S16Value:    pointer.ToInt16(math.MinInt16),
			S32Value:    pointer.ToInt32(math.MinInt32),
			S64Value:    pointer.ToInt64(math.MinInt64),
			F32Value:    pointer.ToFloat32(math.MaxFloat32),
			F64Value:    pointer.ToFloat64(math.MaxFloat64),
			StringValue: pointer.ToString("test"),
			BytesValue:  []byte("test"),
			ObjectValue: &module.Thing{
				Value: "test",
			},
		},
		Maps: module.Maps{
			MapStringPrimative: map[uint32]string{
				1234: "test",
			},
			MapU64Primative: map[uint32]uint64{
				5678: 01234,
			},
		},
		Lists: module.Lists{
			ListStrings:         []string{"test"},
			ListU64s:            []uint64{1234},
			ListObjects:         []module.Thing{{Value: "test"}},
			ListObjectsOptional: []*module.Thing{{Value: "test"}},
		},
	}
	actual, err := m.TestDecode(ctx, expected)
	require.NoError(t, err, "could not invoke testFunction")

	// Adjust for float variances
	actual = strings.Replace(actual, "3.4028234663852886e+38", "3.4028234663852887e+38", -1)
	actual = strings.Replace(actual, "3.4028234663852886e38", "3.4028234663852887e+38", -1)
	actual = strings.Replace(actual, "1.7976931348623157e308", "1.7976931348623157e+308", -1)

	assert.Equal(t, `{
true
255
65535
4294967295
18446744073709551615
-128
-32768
-2147483648
-9223372036854775808
3.4028234663852887e+38
1.7976931348623157e+308
test
test
}`, actual)
}
