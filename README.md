# Introduction

Using gRPC to make a go client communicate with a python server.


# Generate go stub

```
protoc --go_out=plugins=grpc:. protos/sum.proto
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
