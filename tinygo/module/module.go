package module

import (
	msgpack "github.com/wapc/tinygo-msgpack"
	wapc "github.com/wapc/wapc-guest-tinygo"
)

type Host struct {
	binding string
}

func NewHost(binding string) *Host {
	return &Host{
		binding: binding,
	}
}

func (h *Host) TestFunction(required Required, optional Optional, maps Maps, lists Lists) (Tests, error) {
	inputArgs := TestFunctionArgs{
		Required: required,
		Optional: optional,
		Maps:     maps,
		Lists:    lists,
	}
	payload, err := wapc.HostCall(
		h.binding,
		"tests",
		"testFunction",
		inputArgs.ToBuffer(),
	)
	if err != nil {
		return Tests{}, err
	}
	decoder := msgpack.NewDecoder(payload)
	return DecodeTests(&decoder)
}

func (h *Host) TestUnary(tests Tests) (Tests, error) {
	payload, err := wapc.HostCall(h.binding, "tests", "testUnary", tests.ToBuffer())
	if err != nil {
		return Tests{}, err
	}
	decoder := msgpack.NewDecoder(payload)
	return DecodeTests(&decoder)
}

func (h *Host) TestDecode(tests Tests) (string, error) {
	payload, err := wapc.HostCall(h.binding, "tests", "testDecode", tests.ToBuffer())
	if err != nil {
		return "", err
	}
	decoder := msgpack.NewDecoder(payload)
	ret, err := decoder.ReadString()
	return ret, err
}

type Handlers struct {
	TestFunction func(required Required, optional Optional, maps Maps, lists Lists) (Tests, error)
	TestUnary    func(tests Tests) (Tests, error)
	TestDecode   func(tests Tests) (string, error)
}

func (h Handlers) Register() {
	if h.TestFunction != nil {
		testFunctionHandler = h.TestFunction
		wapc.RegisterFunction("testFunction", testFunctionWrapper)
	}
	if h.TestUnary != nil {
		testUnaryHandler = h.TestUnary
		wapc.RegisterFunction("testUnary", testUnaryWrapper)
	}
	if h.TestDecode != nil {
		testDecodeHandler = h.TestDecode
		wapc.RegisterFunction("testDecode", testDecodeWrapper)
	}
}

var (
	testFunctionHandler func(required Required, optional Optional, maps Maps, lists Lists) (Tests, error)
	testUnaryHandler    func(tests Tests) (Tests, error)
	testDecodeHandler   func(tests Tests) (string, error)
)

func testFunctionWrapper(payload []byte) ([]byte, error) {
	decoder := msgpack.NewDecoder(payload)
	var inputArgs TestFunctionArgs
	inputArgs.Decode(&decoder)
	response, err := testFunctionHandler(inputArgs.Required, inputArgs.Optional, inputArgs.Maps, inputArgs.Lists)
	if err != nil {
		return nil, err
	}
	return response.ToBuffer(), nil
}

func testUnaryWrapper(payload []byte) ([]byte, error) {
	decoder := msgpack.NewDecoder(payload)
	var request Tests
	request.Decode(&decoder)
	response, err := testUnaryHandler(request)
	if err != nil {
		return nil, err
	}
	return response.ToBuffer(), nil
}

func testDecodeWrapper(payload []byte) ([]byte, error) {
	decoder := msgpack.NewDecoder(payload)
	var request Tests
	request.Decode(&decoder)
	response, err := testDecodeHandler(request)
	if err != nil {
		return nil, err
	}
	var sizer msgpack.Sizer
	sizer.WriteString(response)

	ua := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(ua)
	encoder.WriteString(response)

	return ua, nil
}

type TestFunctionArgs struct {
	Required Required
	Optional Optional
	Maps     Maps
	Lists    Lists
}

func DecodeTestFunctionArgsNullable(decoder *msgpack.Decoder) (*TestFunctionArgs, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeTestFunctionArgs(decoder)
	return &decoded, err
}

func DecodeTestFunctionArgs(decoder *msgpack.Decoder) (TestFunctionArgs, error) {
	var o TestFunctionArgs
	err := o.Decode(decoder)
	return o, err
}

func (o *TestFunctionArgs) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "required":
			o.Required, err = DecodeRequired(decoder)
		case "optional":
			o.Optional, err = DecodeOptional(decoder)
		case "maps":
			o.Maps, err = DecodeMaps(decoder)
		case "lists":
			o.Lists, err = DecodeLists(decoder)
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *TestFunctionArgs) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(4)
	sizer.WriteString("required")
	o.Required.Size(sizer)
	sizer.WriteString("optional")
	o.Optional.Size(sizer)
	sizer.WriteString("maps")
	o.Maps.Size(sizer)
	sizer.WriteString("lists")
	o.Lists.Size(sizer)
}

func (o *TestFunctionArgs) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(4)
	encoder.WriteString("required")
	o.Required.Encode(encoder)
	encoder.WriteString("optional")
	o.Optional.Encode(encoder)
	encoder.WriteString("maps")
	o.Maps.Encode(encoder)
	encoder.WriteString("lists")
	o.Lists.Encode(encoder)
}

func (o *TestFunctionArgs) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

type Tests struct {
	Required Required
	Optional Optional
	Maps     Maps
	Lists    Lists
}

func DecodeTestsNullable(decoder *msgpack.Decoder) (*Tests, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeTests(decoder)
	return &decoded, err
}

func DecodeTests(decoder *msgpack.Decoder) (Tests, error) {
	var o Tests
	err := o.Decode(decoder)
	return o, err
}

func (o *Tests) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "required":
			o.Required, err = DecodeRequired(decoder)
		case "optional":
			o.Optional, err = DecodeOptional(decoder)
		case "maps":
			o.Maps, err = DecodeMaps(decoder)
		case "lists":
			o.Lists, err = DecodeLists(decoder)
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Tests) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(4)
	sizer.WriteString("required")
	o.Required.Size(sizer)
	sizer.WriteString("optional")
	o.Optional.Size(sizer)
	sizer.WriteString("maps")
	o.Maps.Size(sizer)
	sizer.WriteString("lists")
	o.Lists.Size(sizer)
}

func (o *Tests) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(4)
	encoder.WriteString("required")
	o.Required.Encode(encoder)
	encoder.WriteString("optional")
	o.Optional.Encode(encoder)
	encoder.WriteString("maps")
	o.Maps.Encode(encoder)
	encoder.WriteString("lists")
	o.Lists.Encode(encoder)
}

func (o *Tests) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

type Required struct {
	BoolValue   bool
	U8Value     uint8
	U16Value    uint16
	U32Value    uint32
	U64Value    uint64
	S8Value     int8
	S16Value    int16
	S32Value    int32
	S64Value    int64
	F32Value    float32
	F64Value    float64
	StringValue string
	BytesValue  []byte
	ObjectValue Thing
}

func DecodeRequiredNullable(decoder *msgpack.Decoder) (*Required, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeRequired(decoder)
	return &decoded, err
}

func DecodeRequired(decoder *msgpack.Decoder) (Required, error) {
	var o Required
	err := o.Decode(decoder)
	return o, err
}

func (o *Required) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "boolValue":
			o.BoolValue, err = decoder.ReadBool()
		case "u8Value":
			o.U8Value, err = decoder.ReadUint8()
		case "u16Value":
			o.U16Value, err = decoder.ReadUint16()
		case "u32Value":
			o.U32Value, err = decoder.ReadUint32()
		case "u64Value":
			o.U64Value, err = decoder.ReadUint64()
		case "s8Value":
			o.S8Value, err = decoder.ReadInt8()
		case "s16Value":
			o.S16Value, err = decoder.ReadInt16()
		case "s32Value":
			o.S32Value, err = decoder.ReadInt32()
		case "s64Value":
			o.S64Value, err = decoder.ReadInt64()
		case "f32Value":
			o.F32Value, err = decoder.ReadFloat32()
		case "f64Value":
			o.F64Value, err = decoder.ReadFloat64()
		case "stringValue":
			o.StringValue, err = decoder.ReadString()
		case "bytesValue":
			o.BytesValue, err = decoder.ReadByteArray()
		case "objectValue":
			o.ObjectValue, err = DecodeThing(decoder)
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Required) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(14)
	sizer.WriteString("boolValue")
	sizer.WriteBool(o.BoolValue)
	sizer.WriteString("u8Value")
	sizer.WriteUint8(o.U8Value)
	sizer.WriteString("u16Value")
	sizer.WriteUint16(o.U16Value)
	sizer.WriteString("u32Value")
	sizer.WriteUint32(o.U32Value)
	sizer.WriteString("u64Value")
	sizer.WriteUint64(o.U64Value)
	sizer.WriteString("s8Value")
	sizer.WriteInt8(o.S8Value)
	sizer.WriteString("s16Value")
	sizer.WriteInt16(o.S16Value)
	sizer.WriteString("s32Value")
	sizer.WriteInt32(o.S32Value)
	sizer.WriteString("s64Value")
	sizer.WriteInt64(o.S64Value)
	sizer.WriteString("f32Value")
	sizer.WriteFloat32(o.F32Value)
	sizer.WriteString("f64Value")
	sizer.WriteFloat64(o.F64Value)
	sizer.WriteString("stringValue")
	sizer.WriteString(o.StringValue)
	sizer.WriteString("bytesValue")
	sizer.WriteByteArray(o.BytesValue)
	sizer.WriteString("objectValue")
	o.ObjectValue.Size(sizer)
}

func (o *Required) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(14)
	encoder.WriteString("boolValue")
	encoder.WriteBool(o.BoolValue)
	encoder.WriteString("u8Value")
	encoder.WriteUint8(o.U8Value)
	encoder.WriteString("u16Value")
	encoder.WriteUint16(o.U16Value)
	encoder.WriteString("u32Value")
	encoder.WriteUint32(o.U32Value)
	encoder.WriteString("u64Value")
	encoder.WriteUint64(o.U64Value)
	encoder.WriteString("s8Value")
	encoder.WriteInt8(o.S8Value)
	encoder.WriteString("s16Value")
	encoder.WriteInt16(o.S16Value)
	encoder.WriteString("s32Value")
	encoder.WriteInt32(o.S32Value)
	encoder.WriteString("s64Value")
	encoder.WriteInt64(o.S64Value)
	encoder.WriteString("f32Value")
	encoder.WriteFloat32(o.F32Value)
	encoder.WriteString("f64Value")
	encoder.WriteFloat64(o.F64Value)
	encoder.WriteString("stringValue")
	encoder.WriteString(o.StringValue)
	encoder.WriteString("bytesValue")
	encoder.WriteByteArray(o.BytesValue)
	encoder.WriteString("objectValue")
	o.ObjectValue.Encode(encoder)
}

func (o *Required) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

type Optional struct {
	BoolValue   *bool
	U8Value     *uint8
	U16Value    *uint16
	U32Value    *uint32
	U64Value    *uint64
	S8Value     *int8
	S16Value    *int16
	S32Value    *int32
	S64Value    *int64
	F32Value    *float32
	F64Value    *float64
	StringValue *string
	BytesValue  []byte
	ObjectValue *Thing
}

func DecodeOptionalNullable(decoder *msgpack.Decoder) (*Optional, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeOptional(decoder)
	return &decoded, err
}

func DecodeOptional(decoder *msgpack.Decoder) (Optional, error) {
	var o Optional
	err := o.Decode(decoder)
	return o, err
}

func (o *Optional) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "boolValue":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.BoolValue = nil
				} else {
					var nonNil bool
					nonNil, err = decoder.ReadBool()
					o.BoolValue = &nonNil
				}
			}
		case "u8Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.U8Value = nil
				} else {
					var nonNil uint8
					nonNil, err = decoder.ReadUint8()
					o.U8Value = &nonNil
				}
			}
		case "u16Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.U16Value = nil
				} else {
					var nonNil uint16
					nonNil, err = decoder.ReadUint16()
					o.U16Value = &nonNil
				}
			}
		case "u32Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.U32Value = nil
				} else {
					var nonNil uint32
					nonNil, err = decoder.ReadUint32()
					o.U32Value = &nonNil
				}
			}
		case "u64Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.U64Value = nil
				} else {
					var nonNil uint64
					nonNil, err = decoder.ReadUint64()
					o.U64Value = &nonNil
				}
			}
		case "s8Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.S8Value = nil
				} else {
					var nonNil int8
					nonNil, err = decoder.ReadInt8()
					o.S8Value = &nonNil
				}
			}
		case "s16Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.S16Value = nil
				} else {
					var nonNil int16
					nonNil, err = decoder.ReadInt16()
					o.S16Value = &nonNil
				}
			}
		case "s32Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.S32Value = nil
				} else {
					var nonNil int32
					nonNil, err = decoder.ReadInt32()
					o.S32Value = &nonNil
				}
			}
		case "s64Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.S64Value = nil
				} else {
					var nonNil int64
					nonNil, err = decoder.ReadInt64()
					o.S64Value = &nonNil
				}
			}
		case "f32Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.F32Value = nil
				} else {
					var nonNil float32
					nonNil, err = decoder.ReadFloat32()
					o.F32Value = &nonNil
				}
			}
		case "f64Value":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.F64Value = nil
				} else {
					var nonNil float64
					nonNil, err = decoder.ReadFloat64()
					o.F64Value = &nonNil
				}
			}
		case "stringValue":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.StringValue = nil
				} else {
					var nonNil string
					nonNil, err = decoder.ReadString()
					o.StringValue = &nonNil
				}
			}
		case "bytesValue":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.BytesValue = nil
				} else {
					var nonNil []byte
					nonNil, err = decoder.ReadByteArray()
					o.BytesValue = nonNil
				}
			}
		case "objectValue":
			isNil, err := decoder.IsNextNil()
			if err == nil {
				if isNil {
					o.ObjectValue = nil
				} else {
					var nonNil Thing
					nonNil, err = DecodeThing(decoder)
					o.ObjectValue = &nonNil
				}
			}
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Optional) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(14)
	sizer.WriteString("boolValue")
	if o.BoolValue == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteBool(*o.BoolValue)
	}
	sizer.WriteString("u8Value")
	if o.U8Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteUint8(*o.U8Value)
	}
	sizer.WriteString("u16Value")
	if o.U16Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteUint16(*o.U16Value)
	}
	sizer.WriteString("u32Value")
	if o.U32Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteUint32(*o.U32Value)
	}
	sizer.WriteString("u64Value")
	if o.U64Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteUint64(*o.U64Value)
	}
	sizer.WriteString("s8Value")
	if o.S8Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteInt8(*o.S8Value)
	}
	sizer.WriteString("s16Value")
	if o.S16Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteInt16(*o.S16Value)
	}
	sizer.WriteString("s32Value")
	if o.S32Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteInt32(*o.S32Value)
	}
	sizer.WriteString("s64Value")
	if o.S64Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteInt64(*o.S64Value)
	}
	sizer.WriteString("f32Value")
	if o.F32Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteFloat32(*o.F32Value)
	}
	sizer.WriteString("f64Value")
	if o.F64Value == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteFloat64(*o.F64Value)
	}
	sizer.WriteString("stringValue")
	if o.StringValue == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteString(*o.StringValue)
	}
	sizer.WriteString("bytesValue")
	if o.BytesValue == nil {
		sizer.WriteNil()
	} else {
		sizer.WriteByteArray(o.BytesValue)
	}
	sizer.WriteString("objectValue")
	if o.ObjectValue == nil {
		sizer.WriteNil()
	} else {
		o.ObjectValue.Size(sizer)
	}
}

func (o *Optional) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(14)
	encoder.WriteString("boolValue")
	if o.BoolValue == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteBool(*o.BoolValue)
	}
	encoder.WriteString("u8Value")
	if o.U8Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteUint8(*o.U8Value)
	}
	encoder.WriteString("u16Value")
	if o.U16Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteUint16(*o.U16Value)
	}
	encoder.WriteString("u32Value")
	if o.U32Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteUint32(*o.U32Value)
	}
	encoder.WriteString("u64Value")
	if o.U64Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteUint64(*o.U64Value)
	}
	encoder.WriteString("s8Value")
	if o.S8Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteInt8(*o.S8Value)
	}
	encoder.WriteString("s16Value")
	if o.S16Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteInt16(*o.S16Value)
	}
	encoder.WriteString("s32Value")
	if o.S32Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteInt32(*o.S32Value)
	}
	encoder.WriteString("s64Value")
	if o.S64Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteInt64(*o.S64Value)
	}
	encoder.WriteString("f32Value")
	if o.F32Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteFloat32(*o.F32Value)
	}
	encoder.WriteString("f64Value")
	if o.F64Value == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteFloat64(*o.F64Value)
	}
	encoder.WriteString("stringValue")
	if o.StringValue == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteString(*o.StringValue)
	}
	encoder.WriteString("bytesValue")
	if o.BytesValue == nil {
		encoder.WriteNil()
	} else {
		encoder.WriteByteArray(o.BytesValue)
	}
	encoder.WriteString("objectValue")
	if o.ObjectValue == nil {
		encoder.WriteNil()
	} else {
		o.ObjectValue.Encode(encoder)
	}
}

func (o *Optional) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

type Maps struct {
	MapStringPrimative map[uint32]string
	MapU64Primative    map[uint32]uint64
}

func DecodeMapsNullable(decoder *msgpack.Decoder) (*Maps, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeMaps(decoder)
	return &decoded, err
}

func DecodeMaps(decoder *msgpack.Decoder) (Maps, error) {
	var o Maps
	err := o.Decode(decoder)
	return o, err
}

func (o *Maps) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "mapStringPrimative":
			mapSize, err := decoder.ReadMapSize()
			if err != nil {
				return err
			}
			o.MapStringPrimative = make(map[uint32]string, mapSize)
			for mapSize > 0 {
				mapSize--
				key, err := decoder.ReadUint32()
				if err != nil {
					return err
				}
				value, err := decoder.ReadString()
				if err != nil {
					return err
				}
				o.MapStringPrimative[key] = value
			}
		case "mapU64Primative":
			mapSize, err := decoder.ReadMapSize()
			if err != nil {
				return err
			}
			o.MapU64Primative = make(map[uint32]uint64, mapSize)
			for mapSize > 0 {
				mapSize--
				key, err := decoder.ReadUint32()
				if err != nil {
					return err
				}
				value, err := decoder.ReadUint64()
				if err != nil {
					return err
				}
				o.MapU64Primative[key] = value
			}
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Maps) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(2)
	sizer.WriteString("mapStringPrimative")
	sizer.WriteMapSize(uint32(len(o.MapStringPrimative)))
	if o.MapStringPrimative != nil { // TinyGo bug: ranging over nil maps panics.
		for k, v := range o.MapStringPrimative {
			sizer.WriteUint32(k)
			sizer.WriteString(v)
		}
	}
	sizer.WriteString("mapU64Primative")
	sizer.WriteMapSize(uint32(len(o.MapU64Primative)))
	if o.MapU64Primative != nil { // TinyGo bug: ranging over nil maps panics.
		for k, v := range o.MapU64Primative {
			sizer.WriteUint32(k)
			sizer.WriteUint64(v)
		}
	}
}

func (o *Maps) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(2)
	encoder.WriteString("mapStringPrimative")
	encoder.WriteMapSize(uint32(len(o.MapStringPrimative)))
	if o.MapStringPrimative != nil { // TinyGo bug: ranging over nil maps panics.
		for k, v := range o.MapStringPrimative {
			encoder.WriteUint32(k)
			encoder.WriteString(v)
		}
	}
	encoder.WriteString("mapU64Primative")
	encoder.WriteMapSize(uint32(len(o.MapU64Primative)))
	if o.MapU64Primative != nil { // TinyGo bug: ranging over nil maps panics.
		for k, v := range o.MapU64Primative {
			encoder.WriteUint32(k)
			encoder.WriteUint64(v)
		}
	}
}

func (o *Maps) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

type Lists struct {
	ListStrings         []string
	ListU64s            []uint64
	ListObjects         []Thing
	ListObjectsOptional []*Thing
}

func DecodeListsNullable(decoder *msgpack.Decoder) (*Lists, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeLists(decoder)
	return &decoded, err
}

func DecodeLists(decoder *msgpack.Decoder) (Lists, error) {
	var o Lists
	err := o.Decode(decoder)
	return o, err
}

func (o *Lists) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "listStrings":
			listSize, err := decoder.ReadArraySize()
			if err != nil {
				return err
			}
			o.ListStrings = make([]string, 0, listSize)
			for listSize > 0 {
				listSize--
				var nonNilItem string
				nonNilItem, err = decoder.ReadString()
				if err != nil {
					return err
				}
				o.ListStrings = append(o.ListStrings, nonNilItem)
			}
		case "listU64s":
			listSize, err := decoder.ReadArraySize()
			if err != nil {
				return err
			}
			o.ListU64s = make([]uint64, 0, listSize)
			for listSize > 0 {
				listSize--
				var nonNilItem uint64
				nonNilItem, err = decoder.ReadUint64()
				if err != nil {
					return err
				}
				o.ListU64s = append(o.ListU64s, nonNilItem)
			}
		case "listObjects":
			listSize, err := decoder.ReadArraySize()
			if err != nil {
				return err
			}
			o.ListObjects = make([]Thing, 0, listSize)
			for listSize > 0 {
				listSize--
				var nonNilItem Thing
				nonNilItem, err = DecodeThing(decoder)
				if err != nil {
					return err
				}
				o.ListObjects = append(o.ListObjects, nonNilItem)
			}
		case "listObjectsOptional":
			listSize, err := decoder.ReadArraySize()
			if err != nil {
				return err
			}
			o.ListObjectsOptional = make([]*Thing, 0, listSize)
			for listSize > 0 {
				listSize--
				var nonNilItem *Thing
				isNil, err := decoder.IsNextNil()
				if err == nil {
					if isNil {
						nonNilItem = nil
					} else {
						var nonNil Thing
						nonNil, err = DecodeThing(decoder)
						nonNilItem = &nonNil
					}
				}
				if err != nil {
					return err
				}
				o.ListObjectsOptional = append(o.ListObjectsOptional, nonNilItem)
			}
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Lists) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(4)
	sizer.WriteString("listStrings")
	sizer.WriteArraySize(uint32(len(o.ListStrings)))
	for _, v := range o.ListStrings {
		sizer.WriteString(v)
	}
	sizer.WriteString("listU64s")
	sizer.WriteArraySize(uint32(len(o.ListU64s)))
	for _, v := range o.ListU64s {
		sizer.WriteUint64(v)
	}
	sizer.WriteString("listObjects")
	sizer.WriteArraySize(uint32(len(o.ListObjects)))
	for _, v := range o.ListObjects {
		v.Size(sizer)
	}
	sizer.WriteString("listObjectsOptional")
	sizer.WriteArraySize(uint32(len(o.ListObjectsOptional)))
	for _, v := range o.ListObjectsOptional {
		if v == nil {
			sizer.WriteNil()
		} else {
			v.Size(sizer)
		}
	}
}

func (o *Lists) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(4)
	encoder.WriteString("listStrings")
	encoder.WriteArraySize(uint32(len(o.ListStrings)))
	for _, v := range o.ListStrings {
		encoder.WriteString(v)
	}
	encoder.WriteString("listU64s")
	encoder.WriteArraySize(uint32(len(o.ListU64s)))
	for _, v := range o.ListU64s {
		encoder.WriteUint64(v)
	}
	encoder.WriteString("listObjects")
	encoder.WriteArraySize(uint32(len(o.ListObjects)))
	for _, v := range o.ListObjects {
		v.Encode(encoder)
	}
	encoder.WriteString("listObjectsOptional")
	encoder.WriteArraySize(uint32(len(o.ListObjectsOptional)))
	for _, v := range o.ListObjectsOptional {
		if v == nil {
			encoder.WriteNil()
		} else {
			v.Encode(encoder)
		}
	}
}

func (o *Lists) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}

type Thing struct {
	Value string
}

func DecodeThingNullable(decoder *msgpack.Decoder) (*Thing, error) {
	if isNil, err := decoder.IsNextNil(); isNil || err != nil {
		return nil, err
	}
	decoded, err := DecodeThing(decoder)
	return &decoded, err
}

func DecodeThing(decoder *msgpack.Decoder) (Thing, error) {
	var o Thing
	err := o.Decode(decoder)
	return o, err
}

func (o *Thing) Decode(decoder *msgpack.Decoder) error {
	numFields, err := decoder.ReadMapSize()
	if err != nil {
		return err
	}

	for numFields > 0 {
		numFields--
		field, err := decoder.ReadString()
		if err != nil {
			return err
		}
		switch field {
		case "value":
			o.Value, err = decoder.ReadString()
		default:
			err = decoder.Skip()
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (o *Thing) Size(sizer *msgpack.Sizer) {
	if o == nil {
		sizer.WriteNil()
		return
	}
	sizer.WriteMapSize(1)
	sizer.WriteString("value")
	sizer.WriteString(o.Value)
}

func (o *Thing) Encode(encoder *msgpack.Encoder) {
	if o == nil {
		encoder.WriteNil()
		return
	}
	encoder.WriteMapSize(1)
	encoder.WriteString("value")
	encoder.WriteString(o.Value)
}

func (o *Thing) ToBuffer() []byte {
	var sizer msgpack.Sizer
	o.Size(&sizer)
	buffer := make([]byte, sizer.Len())
	encoder := msgpack.NewEncoder(buffer)
	o.Encode(&encoder)
	return buffer
}
