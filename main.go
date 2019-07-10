package main

// Trivial package to re-encode an LLVM generated wasm file with Wagon.  This
// allows the re-encoded file to be used as Wagon test data

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-interpreter/wagon/wasm"
)

func main() {
	// Ensure we've been given a command line argument
	var err error
	if len(os.Args) != 2 {
		log.Fatal("Needs a .wasm file name on the command line")
	}

	// Read in the wasm file
	fileName := os.Args[1]
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the wasm module
	r := bytes.NewReader(f)
	m, err := wasm.DecodeModule(r)
	if err != nil {
		log.Fatal(err)
	}

	// Re-encode the module file
	buf := new(bytes.Buffer)
	err = wasm.EncodeModule(buf, m)
	if err != nil {
		log.Fatalf("error writing module %v", err)
	}

	// Save out the new file
	outFile := fileName + "-re-encoded"
	err = ioutil.WriteFile(outFile, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Re-encoded file saved to '%s'\n", outFile)
}
