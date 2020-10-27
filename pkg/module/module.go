package module

import (
	"context"

	"github.com/vmihailenco/msgpack/v4"
	"github.com/wapc/wapc-go"
)

type Module struct {
	instance *wapc.Instance
}

func New(instance *wapc.Instance) *Module {
	return &Module{
		instance: instance,
	}
}

func (m *Module) TestFunction(ctx context.Context, required Required, optional Optional, maps Maps, lists Lists) (Tests, error) {
	var ret Tests
	inputArgs := TestFunctionArgs{
		Required: required,
		Optional: optional,
		Maps:     maps,
		Lists:    lists,
	}
	inputPayload, err := msgpack.Marshal(&inputArgs)
	if err != nil {
		return ret, err
	}
	payload, err := m.instance.Invoke(
		ctx,
		"testFunction",
		inputPayload,
	)
	if err != nil {
		return ret, err
	}
	err = msgpack.Unmarshal(payload, &ret)
	return ret, err
}

func (m *Module) TestUnary(ctx context.Context, tests Tests) (Tests, error) {
	var ret Tests
	inputPayload, err := msgpack.Marshal(&tests)
	if err != nil {
		return ret, err
	}
	payload, err := m.instance.Invoke(ctx, "testUnary", inputPayload)
	if err != nil {
		return ret, err
	}
	err = msgpack.Unmarshal(payload, &ret)
	return ret, err
}

func (m *Module) TestDecode(ctx context.Context, tests Tests) (string, error) {
	var ret string
	inputPayload, err := msgpack.Marshal(&tests)
	if err != nil {
		return ret, err
	}
	payload, err := m.instance.Invoke(ctx, "testDecode", inputPayload)
	if err != nil {
		return ret, err
	}
	err = msgpack.Unmarshal(payload, &ret)
	return ret, err
}

type TestFunctionArgs struct {
	Required Required `msgpack:"required"`
	Optional Optional `msgpack:"optional"`
	Maps     Maps     `msgpack:"maps"`
	Lists    Lists    `msgpack:"lists"`
}

type Tests struct {
	Required Required `msgpack:"required"`
	Optional Optional `msgpack:"optional"`
	Maps     Maps     `msgpack:"maps"`
	Lists    Lists    `msgpack:"lists"`
}

type Required struct {
	BoolValue   bool    `msgpack:"boolValue"`
	U8Value     uint8   `msgpack:"u8Value"`
	U16Value    uint16  `msgpack:"u16Value"`
	U32Value    uint32  `msgpack:"u32Value"`
	U64Value    uint64  `msgpack:"u64Value"`
	S8Value     int8    `msgpack:"s8Value"`
	S16Value    int16   `msgpack:"s16Value"`
	S32Value    int32   `msgpack:"s32Value"`
	S64Value    int64   `msgpack:"s64Value"`
	F32Value    float32 `msgpack:"f32Value"`
	F64Value    float64 `msgpack:"f64Value"`
	StringValue string  `msgpack:"stringValue"`
	BytesValue  []byte  `msgpack:"bytesValue"`
	ObjectValue Thing   `msgpack:"objectValue"`
}

type Optional struct {
	BoolValue   *bool    `msgpack:"boolValue"`
	U8Value     *uint8   `msgpack:"u8Value"`
	U16Value    *uint16  `msgpack:"u16Value"`
	U32Value    *uint32  `msgpack:"u32Value"`
	U64Value    *uint64  `msgpack:"u64Value"`
	S8Value     *int8    `msgpack:"s8Value"`
	S16Value    *int16   `msgpack:"s16Value"`
	S32Value    *int32   `msgpack:"s32Value"`
	S64Value    *int64   `msgpack:"s64Value"`
	F32Value    *float32 `msgpack:"f32Value"`
	F64Value    *float64 `msgpack:"f64Value"`
	StringValue *string  `msgpack:"stringValue"`
	BytesValue  []byte   `msgpack:"bytesValue"`
	ObjectValue *Thing   `msgpack:"objectValue"`
}

type Maps struct {
	MapStringPrimative map[uint32]string `msgpack:"mapStringPrimative"`
	MapU64Primative    map[uint32]uint64 `msgpack:"mapU64Primative"`
}

type Lists struct {
	ListStrings         []string `msgpack:"listStrings"`
	ListU64s            []uint64 `msgpack:"listU64s"`
	ListObjects         []Thing  `msgpack:"listObjects"`
	ListObjectsOptional []*Thing `msgpack:"listObjectsOptional"`
}

type Thing struct {
	Value string `msgpack:"value"`
}
