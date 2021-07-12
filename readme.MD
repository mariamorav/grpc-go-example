# gRPC and Protocol Buffers example

## What is it?

gRPC es un framework RPC open source de alto rendimiento.

- IDL = Interface Description Language
- Multilenguaje
- Serialización optimizada. (Más veloz que XML y JSON)
- Menos propenso a errores

## ¿Qué necesitamos?

- un archivo .proto
- un compilador para proto buffers (protoc)

Para este ejemplo usaré golang por lo tanto los Plugins necesarios para el protocol compiler son:

- go get -u google.golang.org/grpc
- go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
- go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

## Documentation
- https://developers.google.com/protocol-buffers/docs/overview
- https://grpc.io/docs/languages/go/quickstart/ - Golang Quick Start from gRPC docs.

## Commands:

```protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/<proto_file_name>.proto```