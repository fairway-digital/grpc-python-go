# Introduction

Using gRPC to make a go client communicate with a python server.

Service is just a dumb 1 + 1 = 2 service.


# Protobuf code gen
## Generate go code

```
protoc --go_out=plugins=grpc:. protos/sum.proto
```

## Generate python code

```
python -m grpc_tools.protoc -I./protos --python_out=python-server --grpc_python_out=python-server protos/sum.proto
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
