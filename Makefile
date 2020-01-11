clean:
	rm -rf go-client/protos/*.go
	rm -rf python-server/*_pb2.py
	rm -rf python-server/*_pb2_grpc.py

proto: proto-go proto-py

proto-go:
	protoc --go_out=plugins=grpc:./go-client protos/math.proto

proto-py:
	python -m grpc_tools.protoc -I./protos --python_out=python-server --grpc_python_out=python-server protos/math.proto
