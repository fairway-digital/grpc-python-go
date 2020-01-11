proto-go:
	protoc --go_out=plugins=grpc:./go-client protos/math.proto

