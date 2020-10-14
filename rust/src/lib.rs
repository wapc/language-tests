pub mod generated;
extern crate wapc_guest as guest;
use generated::*;
use guest::prelude::*;

#[no_mangle]
pub fn wapc_init() {
    Handlers::register_test_function(test_function);
    Handlers::register_test_unary(test_unary);
    Handlers::register_test_decode(test_decode);
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

fn test_decode(tests: Tests) -> HandlerResult<String> {
    let ret = format!("
{}
{}
{}
{}
{}
{}
{}
{}
{}
{:e}
{:e}
{}
{}
",
tests.required.bool_value,
tests.required.u8_value,
tests.required.u16_value,
tests.required.u32_value,
tests.required.u64_value,
tests.required.s8_value,
tests.required.s16_value,
tests.required.s32_value,
tests.required.s64_value,
tests.required.f32_value as f64,
tests.required.f64_value,
tests.required.string_value,
tests.required.string_value,
);
    Ok("{".to_owned() + &ret.to_string() + "}")
}