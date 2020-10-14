import { handleCall, handleAbort } from "wapc-guest-as";
import {
  Tests,
  Required,
  Optional,
  Maps,
  Lists,
  Thing,
  Handlers,
} from "./module";

export function wapc_init(): void {
  Handlers.registerTestFunction(testFunction);
  Handlers.registerTestUnary(testUnary);
  Handlers.registerTestDecode(testDecode);
}

function testFunction(
  required: Required,
  optional: Optional,
  maps: Maps,
  lists: Lists
): Tests {
  // Echo arguments
  return Tests.newBuilder()
    .withRequired(required)
    .withOptional(optional)
    .withMaps(maps)
    .withLists(lists)
    .build();
}

function testUnary(tests: Tests): Tests {
  // Echo input
  return tests;
}

function testDecode(tests: Tests): string {
  let ret = "{\n";
  ret += tests.required.boolValue.toString() + "\n";
  ret += tests.required.u8Value.toString() + "\n";
  ret += tests.required.u16Value.toString() + "\n";
  ret += tests.required.u32Value.toString() + "\n";
  ret += tests.required.u64Value.toString() + "\n";
  ret += tests.required.s8Value.toString() + "\n";
  ret += tests.required.s16Value.toString() + "\n";
  ret += tests.required.s32Value.toString() + "\n";
  ret += tests.required.s64Value.toString() + "\n";
  ret += tests.required.f32Value.toString() + "\n";
  ret += tests.required.f64Value.toString() + "\n";
  ret += tests.required.stringValue + "\n";
  const bytesAsString = String.UTF8.decode(tests.required.bytesValue)
  ret += bytesAsString + "\n";
  ret += "}";
  return ret;
}

// Boilerplate code for waPC.  Do not remove.

export function __guest_call(operation_size: usize, payload_size: usize): bool {
  return handleCall(operation_size, payload_size);
}

// Abort function
function abort(
  message: string | null,
  fileName: string | null,
  lineNumber: u32,
  columnNumber: u32
): void {
  handleAbort(message, fileName, lineNumber, columnNumber);
}
