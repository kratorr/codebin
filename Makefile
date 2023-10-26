BINARY_NAME=codebin
LOCAL_BIN:=$(CURDIR)/bin
SRC_PATH=cmd
LOCAL_MIGRATION_DIR=$(CURDIR)/migrations
LOCAL_MIGRATION_DSN="host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.17.0
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.32.0
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2


get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go@v1.3.0
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0


generate:
	make generate-user-api
	make generate-snippet-api

generate-proto:
	mkdir -p pkg/proto
	protoc --proto_path api/auth_v1 --proto_path api/snippet_v1 --proto_path vendor.protogen \
	--go_out=pkg/proto --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/proto --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--grpc-gateway_out=pkg/proto --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:swagger_ui \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
    api/auth_v1/auth.proto \
    api/snippet_v1/snippet.proto


#generate-snippet-api:
#	mkdir -p pkg/snippet_v1
#	protoc --proto_path api/snippet_v1  --proto_path vendor.protogen \
#	--go_out=pkg/snippet_v1 --go_opt=paths=source_relative \
#	--plugin=protoc-gen-go=bin/protoc-gen-go \
#	--go-grpc_out=pkg/snippet_v1 --go-grpc_opt=paths=source_relative \
#	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
#	--grpc-gateway_out=pkg/snippet_v1 --grpc-gateway_opt=paths=source_relative \
#	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
#	--openapiv2_out=allow_merge=true,merge_file_name=api:docs \
#	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
#	api/snippet_v1/snippet.proto

build:
	go build -o ./bin/$(BINARY_NAME) $(SRC_PATH)/main.go

run:
	go run $(SRC_PATH)/main.go

clean:
	if [ -f $(BINARY_NAME) ] ; then rm $(BINARY_NAME) ; fi

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

docker-compose-up:
	docker-compose up --build

docker-compose-down:
	docker-compose down

docker-compose-clear:
	docker-compose down -v

vendor-proto:
		@if [ ! -d vendor.protogen/validate ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
			mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
			git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
			mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
			rm -rf vendor.protogen/openapiv2 ;\
		fi