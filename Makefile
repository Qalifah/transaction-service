proto:
	protoc --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            proto/wallet.proto
tests:
	go test ./handler
run-service:
	./run-service.sh