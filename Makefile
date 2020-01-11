clean:
	rm -rf api/protos/*.go
	rm -rf computation/*_pb2.py
	rm -rf computation/*_pb2_grpc.py

proto: proto-go proto-py

proto-go:
	protoc --go_out=plugins=grpc:./api protos/math.proto

proto-py:
	python -m grpc_tools.protoc -I./protos --python_out=computation --grpc_python_out=computation protos/math.proto
