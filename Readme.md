# suffiks-tinygo

## Installation

```bash
go get github.com/suffiks/suffiks-tinygo@latest
```

## Usage

The skeleton of a suffiks-tinygo program looks like this:

```go
package main

import (
	suffiks "github.com/suffiks/suffiks-tinygo"

	// Only needed if you need to use the generated types
	// "github.com/suffiks/suffiks-tinygo/protogen"
)


//export Sync
func Sync() {
	// Called whenever the controller is started or the resource is updated
}

//export Delete
func Delete() {
	// if you don't want to delete, you don't have to implement this function
}

//export Defaulting
func Defaulting() uint64 {
	// if you don't want to default, you can just return 0
	// otherwise, return the json patch with changes using suffiks.DefaultingResponse
	// return suffiks.DefaultingResponse(map[string]any{"spec": map[string]any{"foo": "bar"}})
	return 0
}

//export Validate
func Validate(typ protogen.ValidationType) {
	// if you don't want to validate on delete, you can just return
	// if typ == protogen.ValidationType_DELETE {
	//	return
	// }

	// Add errors with suffiks.ValidationError
}

// main has to be defined, but is not used
func main() {}
```

Build the extension using:

```bash
CGO_ENABLED=0 tinygo build -o build/release.wasm -scheduler=none -target=wasi .
```

## Development

To generate the protobuf files, ensure that Suffiks is cloned into the same directory as this repository.

```bash
make gen-proto
```
