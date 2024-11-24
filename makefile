
DECODER_PROTO_DIR = ./proto

build_decoder: ## Generate Pbs and build the Beef service
	@echo "Building decoder server and client..."
	@set -e  # Exit immediately if a command fails
	protoc -I${DECODER_PROTO_DIR} --go_opt=module=$(shell go list -m) --go_out=. --go-grpc_opt=module=$(shell go list -m) --go-grpc_out=. ${DECODER_PROTO_DIR}/*.proto


# Run the rest api
client: ## Run the rest api
	go run decoder/client
# Run the rest api
server: ## Run the rest api
	go run decoder/server

