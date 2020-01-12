# Introduction

Simple fronend app -> go api -> python computation app.

The only operation is Sum(op1, op2)

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

# Install
## Generate protos

```
make proto
```

## npm

```
cd front

yarn install
```

## Start api and computation service

```
docker-compose -f docker-compose.yml up
```

## Start front

```
cd front

yarn start

```

# TODO

* [ ] Voir si on peut pas directement utiliser protoc a la place de grpcio-tools pout le python
* [ ] CI/CD GCP
* [ ] Test de rpc mode stream (sera utile pour les calculs plus long)
