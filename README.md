# Introduction

Using gRPC to make a go client communicate with a python server.

A simple go programm (client) is requesting to a server (python) result of a sum (1 + 1).

Result (2) is returned to client.


# Protobuf code gen
## Generate go code

```
make proto-go
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

# TODO

[ ] Dockerifier
[ ] K8s-ifier -> gcp
[ ] Front
[ ] Generer les stubs protobuf python -> python-server/generated
[ ] Test de rpc mode stream (sera utile pour les calculs plus long)
[ ] Read input sum param from cli (vs hard coded)
