#!/bin/sh
wapc generate codegen-go.yaml

echo "Building AssemblyScript module"
wapc generate codegen-as.yaml
npm run build

echo "Building TinyGo module"
wapc generate codegen-tinygo.yaml
tinygo build -o build/tinygo.wasm -target wasm -no-debug tinygo/main.go

echo "Building Rust module"
wapc generate codegen-rust.yaml
cargo build --target wasm32-unknown-unknown --release --manifest-path=rust/Cargo.toml && \
  cp rust/target/wasm32-unknown-unknown/release/rust_codegen_test.wasm build/rust.wasm
