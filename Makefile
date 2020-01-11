proto-go:
	protoc --go_out=plugins=grpc:./go-client protos/math.proto

proto-py:
	python -m grpc_tools.protoc -I./protos --python_out=python-server/protos --grpc_python_out=python-server/protos protos/math.proto
