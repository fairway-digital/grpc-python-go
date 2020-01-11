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
make proto-py
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

* [ ] Dockerifier
* [ ] K8s-ifier
* [ ] CI/CD GCP
* [ ] Front
* [ ] Test de rpc mode stream (sera utile pour les calculs plus long)
* [ ] go CLI client
* [ ] front web client page
