.PHONY: start-server
start-server:
	go run ./pkg1/main.go

.PHONY: start-client
start-client:
	go run ./pkg2/main.go

.PHONY: protoc
protoc:
	protoc --go_out=plugins=grpc:. ./pkg1/pb/rent.proto

.PHONY: cert
cert:
	./cert.sh
