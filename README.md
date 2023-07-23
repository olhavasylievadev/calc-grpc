# Calculator gRPC
This is a simple implementation of CLI calculator using Go and gRPC.

## Functionality
Supports four basic math operations:
- '+' as Addition
- '-' as Subtraction
- '*' as Multiplication
- '/' as Division

## Required Setup
- Make sure you're using recent version of Go, i.e. Go 1.16+
- Install [protoc](https://github.com/golang/protobuf/), as the server implements a generated Protocol Buffer file and needs `protoc` compiler

You can generate go-code file from protobuf. To do this, go to the directory of the protofile, i.e.

`cd ./api/proto`

Then, use the following command

`
$ protoc \
--go_out=./../../pkg/calculator \
--go_opt=paths=source_relative \
--go-grpc_out=./../../pkg/calculator \
--go-grpc_opt=paths=source_relative \
calculator.proto
`

## Running the program
1. Start the server

`go run cmd/server/main.go`
2. Start the client and pass the arguments

`go run cmd/client/main.go 6 '*' 4`

Or you can build executable files with `go build` and run them with arguments

## More info
Even though you can calculate result running code like

`go run ./client/main.go 24 / 6`, it's recommended to enclose operators in single quotes like this
`4 '+' 6` as some of them might be system commands. For example `*` on Mac.

For more information about gRPC please visit [gRPC.io](https://grpc.io/)

You can also learn about Protocol Buffers here [Protocol Buffers v3](https://cloud.google.com/apis/design/proto3)
