.PHONY: generate-user-v2-gateway-proto

generate-user-v2-gateway-proto:
	protoc -I . \
	    --grpc-gateway_out . \
	    --grpc-gateway_opt logtostderr=true \
	    --grpc-gateway_opt paths=source_relative \
	    protos/v2/users/users.proto


.PHONY: generate-user-v2-proto
generate-user-v2-proto:
	protoc -I=. \
	    --go_out . \
        --go_opt paths=source_relative \
        --go-grpc_out . \
        --go-grpc_opt paths=source_relative \
	    protos/v2/users/users.proto