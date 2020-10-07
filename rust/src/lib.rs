pub mod generated;
extern crate wapc_guest as guest;
use generated::*;
use guest::prelude::*;

#[no_mangle]
pub fn wapc_init() {
    Handlers::register_test_function(test_function);
    Handlers::register_test_unary(test_unary);
}

fn test_function(
    required: Required,
    optional: Optional,
    maps: Maps,
    lists: Lists,
) -> HandlerResult<Tests> {
    let mut tests = Tests::default();
    tests.required = required;
    tests.optional = optional;
    tests.maps = maps;
    tests.lists = lists;
    Ok(tests)
}

fn test_unary(tests: Tests) -> HandlerResult<Tests> {
    Ok(tests)
}
