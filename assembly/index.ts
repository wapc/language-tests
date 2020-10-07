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
