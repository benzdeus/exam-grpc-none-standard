grpc-auth:
	protoc --go_out=./../gate-way/pkg/ --go-grpc_out=./../gate-way/pkg/ pb-auth-gen/*.proto
	protoc --go_out=./../auth-service/pkg/ --go-grpc_out=./../auth-service/pkg/ pb-auth-gen/*.proto

grpc-payment:
	protoc --go_out=./../gate-way/pkg/ --go-grpc_out=./../gate-way/pkg/ pb-payment-gen/*.proto
	protoc --go_out=./../payment-service/pkg/ --go-grpc_out=./../payment-service/pkg/ pb-auth-gen/*.proto