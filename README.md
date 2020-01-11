# Introduction

Using gRPC to make a go client communicate with a python server.

A simple go programm (client) is requesting to a server (python) result of a sum (1 + 1).

Result (2) is returned to client.


# Protobuf code gen
## Generate go code

```
protoc --go_out=plugins=grpc:. protos/math.proto
```

## Generate python code

```
python -m grpc_tools.protoc -I./protos --python_out=python-server --grpc_python_out=python-server protos/math.proto
```

# Start python server

```
cd python-server

python server.py
```

# Start go client

```
cd go-client

go run main.go

```
