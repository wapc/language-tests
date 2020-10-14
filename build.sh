#!/bin/sh
echo "Generating code"
wapc generate codegen.yaml

echo "Building AssemblyScript module"
npm run build

echo "Building TinyGo module"
tinygo build -o build/tinygo.wasm -target wasm -no-debug tinygo/main.go

echo "Building Rust module"
cargo build --target wasm32-unknown-unknown --release --manifest-path=rust/Cargo.toml && \
  cp rust/target/wasm32-unknown-unknown/release/rust_codegen_test.wasm build/rust.wasm
