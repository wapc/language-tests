import { Decoder, Writer, Encoder, Sizer, Value } from "@wapc/as-msgpack";
import { register, hostCall } from "wapc-guest-as";

export class Host {
  binding: string;

  constructor(binding: string) {
    this.binding = binding;
  }

  testFunction(
    required: Required,
    optional: Optional,
    maps: Maps,
    lists: Lists
  ): Tests {
    const inputArgs = new TestFunctionArgs();
    inputArgs.required = required;
    inputArgs.optional = optional;
    inputArgs.maps = maps;
    inputArgs.lists = lists;
    const payload = hostCall(
      this.binding,
      "tests",
      "testFunction",
      inputArgs.toBuffer()
    );
    const decoder = new Decoder(payload);
    return Tests.decode(decoder);
  }

  testUnary(tests: Tests): Tests {
    const payload = hostCall(
      this.binding,
      "tests",
      "testUnary",
      tests.toBuffer()
    );
    const decoder = new Decoder(payload);
    return Tests.decode(decoder);
  }

  testDecode(tests: Tests): string {
    const payload = hostCall(
      this.binding,
      "tests",
      "testDecode",
      tests.toBuffer()
    );
    const decoder = new Decoder(payload);
    const ret = decoder.readString();
    return ret;
  }
}

export class Handlers {
  static registerTestFunction(
    handler: (
      required: Required,
      optional: Optional,
      maps: Maps,
      lists: Lists
    ) => Tests
  ): void {
    testFunctionHandler = handler;
    register("testFunction", testFunctionWrapper);
  }

  static registerTestUnary(handler: (tests: Tests) => Tests): void {
    testUnaryHandler = handler;
    register("testUnary", testUnaryWrapper);
  }

  static registerTestDecode(handler: (tests: Tests) => string): void {
    testDecodeHandler = handler;
    register("testDecode", testDecodeWrapper);
  }
}

var testFunctionHandler: (
  required: Required,
  optional: Optional,
  maps: Maps,
  lists: Lists
) => Tests;
function testFunctionWrapper(payload: ArrayBuffer): ArrayBuffer {
  const decoder = new Decoder(payload);
  const inputArgs = new TestFunctionArgs();
  inputArgs.decode(decoder);
  const response = testFunctionHandler(
    inputArgs.required,
    inputArgs.optional,
    inputArgs.maps,
    inputArgs.lists
  );
  return response.toBuffer();
}

var testUnaryHandler: (tests: Tests) => Tests;
function testUnaryWrapper(payload: ArrayBuffer): ArrayBuffer {
  const decoder = new Decoder(payload);
  const request = new Tests();
  request.decode(decoder);
  const response = testUnaryHandler(request);
  return response.toBuffer();
}

var testDecodeHandler: (tests: Tests) => string;
function testDecodeWrapper(payload: ArrayBuffer): ArrayBuffer {
  const decoder = new Decoder(payload);
  const request = new Tests();
  request.decode(decoder);
  const response = testDecodeHandler(request);
  const sizer = new Sizer();
  sizer.writeString(response);
  const ua = new ArrayBuffer(sizer.length);
  const encoder = new Encoder(ua);
  encoder.writeString(response);
  return ua;
}

export class TestFunctionArgs {
  required: Required = new Required();
  optional: Optional = new Optional();
  maps: Maps = new Maps();
  lists: Lists = new Lists();

  static decodeNullable(decoder: Decoder): TestFunctionArgs | null {
    if (decoder.isNextNil()) return null;
    return TestFunctionArgs.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): TestFunctionArgs {
    const o = new TestFunctionArgs();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "required") {
        this.required = Required.decode(decoder);
      } else if (field == "optional") {
        this.optional = Optional.decode(decoder);
      } else if (field == "maps") {
        this.maps = Maps.decode(decoder);
      } else if (field == "lists") {
        this.lists = Lists.decode(decoder);
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(4);
    encoder.writeString("required");
    this.required.encode(encoder);
    encoder.writeString("optional");
    this.optional.encode(encoder);
    encoder.writeString("maps");
    this.maps.encode(encoder);
    encoder.writeString("lists");
    this.lists.encode(encoder);
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }
}

export class Tests {
  required: Required = new Required();
  optional: Optional = new Optional();
  maps: Maps = new Maps();
  lists: Lists = new Lists();

  static decodeNullable(decoder: Decoder): Tests | null {
    if (decoder.isNextNil()) return null;
    return Tests.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): Tests {
    const o = new Tests();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "required") {
        this.required = Required.decode(decoder);
      } else if (field == "optional") {
        this.optional = Optional.decode(decoder);
      } else if (field == "maps") {
        this.maps = Maps.decode(decoder);
      } else if (field == "lists") {
        this.lists = Lists.decode(decoder);
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(4);
    encoder.writeString("required");
    this.required.encode(encoder);
    encoder.writeString("optional");
    this.optional.encode(encoder);
    encoder.writeString("maps");
    this.maps.encode(encoder);
    encoder.writeString("lists");
    this.lists.encode(encoder);
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }

  static newBuilder(): TestsBuilder {
    return new TestsBuilder();
  }
}

export class TestsBuilder {
  instance: Tests = new Tests();

  withRequired(required: Required): TestsBuilder {
    this.instance.required = required;
    return this;
  }

  withOptional(optional: Optional): TestsBuilder {
    this.instance.optional = optional;
    return this;
  }

  withMaps(maps: Maps): TestsBuilder {
    this.instance.maps = maps;
    return this;
  }

  withLists(lists: Lists): TestsBuilder {
    this.instance.lists = lists;
    return this;
  }

  build(): Tests {
    return this.instance;
  }
}

export class Required {
  boolValue: bool = false;
  u8Value: u8 = 0;
  u16Value: u16 = 0;
  u32Value: u32 = 0;
  u64Value: u64 = 0;
  s8Value: i8 = 0;
  s16Value: i16 = 0;
  s32Value: i32 = 0;
  s64Value: i64 = 0;
  f32Value: f32 = 0;
  f64Value: f64 = 0;
  stringValue: string = "";
  bytesValue: ArrayBuffer = new ArrayBuffer(0);
  objectValue: Thing = new Thing();

  static decodeNullable(decoder: Decoder): Required | null {
    if (decoder.isNextNil()) return null;
    return Required.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): Required {
    const o = new Required();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "boolValue") {
        this.boolValue = decoder.readBool();
      } else if (field == "u8Value") {
        this.u8Value = decoder.readUInt8();
      } else if (field == "u16Value") {
        this.u16Value = decoder.readUInt16();
      } else if (field == "u32Value") {
        this.u32Value = decoder.readUInt32();
      } else if (field == "u64Value") {
        this.u64Value = decoder.readUInt64();
      } else if (field == "s8Value") {
        this.s8Value = decoder.readInt8();
      } else if (field == "s16Value") {
        this.s16Value = decoder.readInt16();
      } else if (field == "s32Value") {
        this.s32Value = decoder.readInt32();
      } else if (field == "s64Value") {
        this.s64Value = decoder.readInt64();
      } else if (field == "f32Value") {
        this.f32Value = decoder.readFloat32();
      } else if (field == "f64Value") {
        this.f64Value = decoder.readFloat64();
      } else if (field == "stringValue") {
        this.stringValue = decoder.readString();
      } else if (field == "bytesValue") {
        this.bytesValue = decoder.readByteArray();
      } else if (field == "objectValue") {
        this.objectValue = Thing.decode(decoder);
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(14);
    encoder.writeString("boolValue");
    encoder.writeBool(this.boolValue);
    encoder.writeString("u8Value");
    encoder.writeUInt8(this.u8Value);
    encoder.writeString("u16Value");
    encoder.writeUInt16(this.u16Value);
    encoder.writeString("u32Value");
    encoder.writeUInt32(this.u32Value);
    encoder.writeString("u64Value");
    encoder.writeUInt64(this.u64Value);
    encoder.writeString("s8Value");
    encoder.writeInt8(this.s8Value);
    encoder.writeString("s16Value");
    encoder.writeInt16(this.s16Value);
    encoder.writeString("s32Value");
    encoder.writeInt32(this.s32Value);
    encoder.writeString("s64Value");
    encoder.writeInt64(this.s64Value);
    encoder.writeString("f32Value");
    encoder.writeFloat32(this.f32Value);
    encoder.writeString("f64Value");
    encoder.writeFloat64(this.f64Value);
    encoder.writeString("stringValue");
    encoder.writeString(this.stringValue);
    encoder.writeString("bytesValue");
    encoder.writeByteArray(this.bytesValue);
    encoder.writeString("objectValue");
    this.objectValue.encode(encoder);
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }

  static newBuilder(): RequiredBuilder {
    return new RequiredBuilder();
  }
}

export class RequiredBuilder {
  instance: Required = new Required();

  withBoolValue(boolValue: bool): RequiredBuilder {
    this.instance.boolValue = boolValue;
    return this;
  }

  withU8Value(u8Value: u8): RequiredBuilder {
    this.instance.u8Value = u8Value;
    return this;
  }

  withU16Value(u16Value: u16): RequiredBuilder {
    this.instance.u16Value = u16Value;
    return this;
  }

  withU32Value(u32Value: u32): RequiredBuilder {
    this.instance.u32Value = u32Value;
    return this;
  }

  withU64Value(u64Value: u64): RequiredBuilder {
    this.instance.u64Value = u64Value;
    return this;
  }

  withS8Value(s8Value: i8): RequiredBuilder {
    this.instance.s8Value = s8Value;
    return this;
  }

  withS16Value(s16Value: i16): RequiredBuilder {
    this.instance.s16Value = s16Value;
    return this;
  }

  withS32Value(s32Value: i32): RequiredBuilder {
    this.instance.s32Value = s32Value;
    return this;
  }

  withS64Value(s64Value: i64): RequiredBuilder {
    this.instance.s64Value = s64Value;
    return this;
  }

  withF32Value(f32Value: f32): RequiredBuilder {
    this.instance.f32Value = f32Value;
    return this;
  }

  withF64Value(f64Value: f64): RequiredBuilder {
    this.instance.f64Value = f64Value;
    return this;
  }

  withStringValue(stringValue: string): RequiredBuilder {
    this.instance.stringValue = stringValue;
    return this;
  }

  withBytesValue(bytesValue: ArrayBuffer): RequiredBuilder {
    this.instance.bytesValue = bytesValue;
    return this;
  }

  withObjectValue(objectValue: Thing): RequiredBuilder {
    this.instance.objectValue = objectValue;
    return this;
  }

  build(): Required {
    return this.instance;
  }
}

export class Optional {
  boolValue: Value<bool> | null = null;
  u8Value: Value<u8> | null = null;
  u16Value: Value<u16> | null = null;
  u32Value: Value<u32> | null = null;
  u64Value: Value<u64> | null = null;
  s8Value: Value<i8> | null = null;
  s16Value: Value<i16> | null = null;
  s32Value: Value<i32> | null = null;
  s64Value: Value<i64> | null = null;
  f32Value: Value<f32> | null = null;
  f64Value: Value<f64> | null = null;
  stringValue: Value<string> | null = null;
  bytesValue: ArrayBuffer | null = null;
  objectValue: Thing | null = null;

  static decodeNullable(decoder: Decoder): Optional | null {
    if (decoder.isNextNil()) return null;
    return Optional.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): Optional {
    const o = new Optional();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "boolValue") {
        if (decoder.isNextNil()) {
          this.boolValue = null;
        } else {
          this.boolValue = new Value(decoder.readBool());
        }
      } else if (field == "u8Value") {
        if (decoder.isNextNil()) {
          this.u8Value = null;
        } else {
          this.u8Value = new Value(decoder.readUInt8());
        }
      } else if (field == "u16Value") {
        if (decoder.isNextNil()) {
          this.u16Value = null;
        } else {
          this.u16Value = new Value(decoder.readUInt16());
        }
      } else if (field == "u32Value") {
        if (decoder.isNextNil()) {
          this.u32Value = null;
        } else {
          this.u32Value = new Value(decoder.readUInt32());
        }
      } else if (field == "u64Value") {
        if (decoder.isNextNil()) {
          this.u64Value = null;
        } else {
          this.u64Value = new Value(decoder.readUInt64());
        }
      } else if (field == "s8Value") {
        if (decoder.isNextNil()) {
          this.s8Value = null;
        } else {
          this.s8Value = new Value(decoder.readInt8());
        }
      } else if (field == "s16Value") {
        if (decoder.isNextNil()) {
          this.s16Value = null;
        } else {
          this.s16Value = new Value(decoder.readInt16());
        }
      } else if (field == "s32Value") {
        if (decoder.isNextNil()) {
          this.s32Value = null;
        } else {
          this.s32Value = new Value(decoder.readInt32());
        }
      } else if (field == "s64Value") {
        if (decoder.isNextNil()) {
          this.s64Value = null;
        } else {
          this.s64Value = new Value(decoder.readInt64());
        }
      } else if (field == "f32Value") {
        if (decoder.isNextNil()) {
          this.f32Value = null;
        } else {
          this.f32Value = new Value(decoder.readFloat32());
        }
      } else if (field == "f64Value") {
        if (decoder.isNextNil()) {
          this.f64Value = null;
        } else {
          this.f64Value = new Value(decoder.readFloat64());
        }
      } else if (field == "stringValue") {
        if (decoder.isNextNil()) {
          this.stringValue = null;
        } else {
          this.stringValue = new Value(decoder.readString());
        }
      } else if (field == "bytesValue") {
        if (decoder.isNextNil()) {
          this.bytesValue = null;
        } else {
          this.bytesValue = decoder.readByteArray();
        }
      } else if (field == "objectValue") {
        if (decoder.isNextNil()) {
          this.objectValue = null;
        } else {
          this.objectValue = Thing.decode(decoder);
        }
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(14);
    encoder.writeString("boolValue");
    if (this.boolValue === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.boolValue!;
      encoder.writeBool(unboxed.value);
    }
    encoder.writeString("u8Value");
    if (this.u8Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.u8Value!;
      encoder.writeUInt8(unboxed.value);
    }
    encoder.writeString("u16Value");
    if (this.u16Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.u16Value!;
      encoder.writeUInt16(unboxed.value);
    }
    encoder.writeString("u32Value");
    if (this.u32Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.u32Value!;
      encoder.writeUInt32(unboxed.value);
    }
    encoder.writeString("u64Value");
    if (this.u64Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.u64Value!;
      encoder.writeUInt64(unboxed.value);
    }
    encoder.writeString("s8Value");
    if (this.s8Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.s8Value!;
      encoder.writeInt8(unboxed.value);
    }
    encoder.writeString("s16Value");
    if (this.s16Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.s16Value!;
      encoder.writeInt16(unboxed.value);
    }
    encoder.writeString("s32Value");
    if (this.s32Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.s32Value!;
      encoder.writeInt32(unboxed.value);
    }
    encoder.writeString("s64Value");
    if (this.s64Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.s64Value!;
      encoder.writeInt64(unboxed.value);
    }
    encoder.writeString("f32Value");
    if (this.f32Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.f32Value!;
      encoder.writeFloat32(unboxed.value);
    }
    encoder.writeString("f64Value");
    if (this.f64Value === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.f64Value!;
      encoder.writeFloat64(unboxed.value);
    }
    encoder.writeString("stringValue");
    if (this.stringValue === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.stringValue!;
      encoder.writeString(unboxed.value);
    }
    encoder.writeString("bytesValue");
    if (this.bytesValue === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.bytesValue!;
      encoder.writeByteArray(unboxed);
    }
    encoder.writeString("objectValue");
    if (this.objectValue === null) {
      encoder.writeNil();
    } else {
      const unboxed = this.objectValue!;
      unboxed.encode(encoder);
    }
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }

  static newBuilder(): OptionalBuilder {
    return new OptionalBuilder();
  }
}

export class OptionalBuilder {
  instance: Optional = new Optional();

  withBoolValue(boolValue: Value<bool> | null): OptionalBuilder {
    this.instance.boolValue = boolValue;
    return this;
  }

  withU8Value(u8Value: Value<u8> | null): OptionalBuilder {
    this.instance.u8Value = u8Value;
    return this;
  }

  withU16Value(u16Value: Value<u16> | null): OptionalBuilder {
    this.instance.u16Value = u16Value;
    return this;
  }

  withU32Value(u32Value: Value<u32> | null): OptionalBuilder {
    this.instance.u32Value = u32Value;
    return this;
  }

  withU64Value(u64Value: Value<u64> | null): OptionalBuilder {
    this.instance.u64Value = u64Value;
    return this;
  }

  withS8Value(s8Value: Value<i8> | null): OptionalBuilder {
    this.instance.s8Value = s8Value;
    return this;
  }

  withS16Value(s16Value: Value<i16> | null): OptionalBuilder {
    this.instance.s16Value = s16Value;
    return this;
  }

  withS32Value(s32Value: Value<i32> | null): OptionalBuilder {
    this.instance.s32Value = s32Value;
    return this;
  }

  withS64Value(s64Value: Value<i64> | null): OptionalBuilder {
    this.instance.s64Value = s64Value;
    return this;
  }

  withF32Value(f32Value: Value<f32> | null): OptionalBuilder {
    this.instance.f32Value = f32Value;
    return this;
  }

  withF64Value(f64Value: Value<f64> | null): OptionalBuilder {
    this.instance.f64Value = f64Value;
    return this;
  }

  withStringValue(stringValue: Value<string> | null): OptionalBuilder {
    this.instance.stringValue = stringValue;
    return this;
  }

  withBytesValue(bytesValue: ArrayBuffer | null): OptionalBuilder {
    this.instance.bytesValue = bytesValue;
    return this;
  }

  withObjectValue(objectValue: Thing | null): OptionalBuilder {
    this.instance.objectValue = objectValue;
    return this;
  }

  build(): Optional {
    return this.instance;
  }
}

export class Maps {
  mapStringPrimative: Map<u32, string> = new Map<u32, string>();
  mapU64Primative: Map<u32, u64> = new Map<u32, u64>();

  static decodeNullable(decoder: Decoder): Maps | null {
    if (decoder.isNextNil()) return null;
    return Maps.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): Maps {
    const o = new Maps();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "mapStringPrimative") {
        this.mapStringPrimative = decoder.readMap(
          (decoder: Decoder): u32 => {
            return decoder.readUInt32();
          },
          (decoder: Decoder): string => {
            return decoder.readString();
          }
        );
      } else if (field == "mapU64Primative") {
        this.mapU64Primative = decoder.readMap(
          (decoder: Decoder): u32 => {
            return decoder.readUInt32();
          },
          (decoder: Decoder): u64 => {
            return decoder.readUInt64();
          }
        );
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(2);
    encoder.writeString("mapStringPrimative");
    encoder.writeMap(
      this.mapStringPrimative,
      (encoder: Writer, key: u32): void => {
        encoder.writeUInt32(key);
      },
      (encoder: Writer, value: string): void => {
        encoder.writeString(value);
      }
    );
    encoder.writeString("mapU64Primative");
    encoder.writeMap(
      this.mapU64Primative,
      (encoder: Writer, key: u32): void => {
        encoder.writeUInt32(key);
      },
      (encoder: Writer, value: u64): void => {
        encoder.writeUInt64(value);
      }
    );
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }

  static newBuilder(): MapsBuilder {
    return new MapsBuilder();
  }
}

export class MapsBuilder {
  instance: Maps = new Maps();

  withMapStringPrimative(mapStringPrimative: Map<u32, string>): MapsBuilder {
    this.instance.mapStringPrimative = mapStringPrimative;
    return this;
  }

  withMapU64Primative(mapU64Primative: Map<u32, u64>): MapsBuilder {
    this.instance.mapU64Primative = mapU64Primative;
    return this;
  }

  build(): Maps {
    return this.instance;
  }
}

export class Lists {
  listStrings: Array<string> = new Array<string>();
  listU64s: Array<u64> = new Array<u64>();
  listObjects: Array<Thing> = new Array<Thing>();
  listObjectsOptional: Array<Thing | null> = new Array<Thing | null>();

  static decodeNullable(decoder: Decoder): Lists | null {
    if (decoder.isNextNil()) return null;
    return Lists.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): Lists {
    const o = new Lists();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "listStrings") {
        this.listStrings = decoder.readArray((decoder: Decoder): string => {
          return decoder.readString();
        });
      } else if (field == "listU64s") {
        this.listU64s = decoder.readArray(
          (decoder: Decoder): u64 => {
            return decoder.readUInt64();
          }
        );
      } else if (field == "listObjects") {
        this.listObjects = decoder.readArray(
          (decoder: Decoder): Thing => {
            return Thing.decode(decoder);
          }
        );
      } else if (field == "listObjectsOptional") {
        this.listObjectsOptional = decoder.readArray(
          (decoder: Decoder): Thing | null => {
            if (decoder.isNextNil()) {
              return null;
            } else {
              return Thing.decode(decoder);
            }
          }
        );
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(4);
    encoder.writeString("listStrings");
    encoder.writeArray(
      this.listStrings,
      (encoder: Writer, item: string): void => {
        encoder.writeString(item);
      }
    );
    encoder.writeString("listU64s");
    encoder.writeArray(this.listU64s, (encoder: Writer, item: u64): void => {
      encoder.writeUInt64(item);
    });
    encoder.writeString("listObjects");
    encoder.writeArray(
      this.listObjects,
      (encoder: Writer, item: Thing): void => {
        item.encode(encoder);
      }
    );
    encoder.writeString("listObjectsOptional");
    encoder.writeArray(
      this.listObjectsOptional,
      (encoder: Writer, item: Thing | null): void => {
        if (item === null) {
          encoder.writeNil();
        } else {
          const unboxed = item;
          unboxed.encode(encoder);
        }
      }
    );
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }

  static newBuilder(): ListsBuilder {
    return new ListsBuilder();
  }
}

export class ListsBuilder {
  instance: Lists = new Lists();

  withListStrings(listStrings: Array<string>): ListsBuilder {
    this.instance.listStrings = listStrings;
    return this;
  }

  withListU64s(listU64s: Array<u64>): ListsBuilder {
    this.instance.listU64s = listU64s;
    return this;
  }

  withListObjects(listObjects: Array<Thing>): ListsBuilder {
    this.instance.listObjects = listObjects;
    return this;
  }

  withListObjectsOptional(
    listObjectsOptional: Array<Thing | null>
  ): ListsBuilder {
    this.instance.listObjectsOptional = listObjectsOptional;
    return this;
  }

  build(): Lists {
    return this.instance;
  }
}

export class Thing {
  value: string = "";

  static decodeNullable(decoder: Decoder): Thing | null {
    if (decoder.isNextNil()) return null;
    return Thing.decode(decoder);
  }

  // decode
  static decode(decoder: Decoder): Thing {
    const o = new Thing();
    o.decode(decoder);
    return o;
  }

  decode(decoder: Decoder): void {
    var numFields = decoder.readMapSize();

    while (numFields > 0) {
      numFields--;
      const field = decoder.readString();

      if (field == "value") {
        this.value = decoder.readString();
      } else {
        decoder.skip();
      }
    }
  }

  encode(encoder: Writer): void {
    encoder.writeMapSize(1);
    encoder.writeString("value");
    encoder.writeString(this.value);
  }

  toBuffer(): ArrayBuffer {
    let sizer = new Sizer();
    this.encode(sizer);
    let buffer = new ArrayBuffer(sizer.length);
    let encoder = new Encoder(buffer);
    this.encode(encoder);
    return buffer;
  }

  static newBuilder(): ThingBuilder {
    return new ThingBuilder();
  }
}

export class ThingBuilder {
  instance: Thing = new Thing();

  withValue(value: string): ThingBuilder {
    this.instance.value = value;
    return this;
  }

  build(): Thing {
    return this.instance;
  }
}
