# Generate gRPC code from specification proto-exchanger/exchanger/v1
.PHONY: gen-proto
gen-proto:
	@ rm -rf gen/exchange
	@ protoc --proto_path=proto --go_out=gen \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=gen --go-grpc_opt=paths=source_relative \
	  proto/exchanger/v1/exchanger.proto

.PHONY: gen-mock
gen-mock:
	@ mockgen -source gen/exchanger/v1/exchanger_grpc.pb.go -destination gen/mock/mock.go