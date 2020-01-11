# Introduction

Using gRPC to make a go client communicate with a python server.

A simple go programm (client) is requesting to a server (python) result of a sum (1 + 1).

Result (2) is returned to client.

# Prequesite

1. install protoc

Mac
```
brew install protobuf
```

2. Install grpcio-tools (for proto -> python code generation)

```
pip install grpcio-tools
```


# Protobuf code gen
## Generate go code

```
make proto-go
```

## Generate python code

```
make proto-py
```

# Start python server

```
cd computation

python server.py
```

# Start go client

```
cd api

go run main.go

```

# TODO

* [ ] Docker-composifier
* [ ] Voir si on peut pas directement utiliser protoc a la place de grpcio-tools pout le python
* [ ] K8s-ifier
* [ ] CI/CD GCP
* [ ] Front
* [ ] Test de rpc mode stream (sera utile pour les calculs plus long)
* [ ] go CLI client
* [ ] front web client page
