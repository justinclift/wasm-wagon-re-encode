## Purpose

Super simple utility to re-encode a .wasm file using Wagon, so it can be used as Wagon test data.

This is so LLVM generated wasm files, which encode slightly differently to how Wagon does them,
can be used as a source of test data for Wagon.

```
$ go run main.go testdata/hello_world-conservative_gc.wasm
Re-encoded file saved to 'testdata/hello_world-conservative_gc.wasm-re-encoded'
```
