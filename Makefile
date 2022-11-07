PHONY: lint-full
lint-full:
	golangci-lint run --config=.golangci.pipeline.yaml ./...


PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p vendor.protogen
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi
		@if [ ! -d vendor.protogen/github.com/envoyproxy ]; then \
			mkdir -p vendor.protogen/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
			mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
			rm -rf vendor.protogen/protoc-gen-validate ;\
		fi
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
		go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
		go get -u github.com/golang/protobuf/proto
		go get -u github.com/golang/protobuf/protoc-gen-go
		go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
		go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
		go install github.com/envoyproxy/protoc-gen-validate
		go get github.com/fullstorydev/grpcui/...

.PHONY: load-test
load-test:
	wrk -t12 -c400 -d30s -s ./load_tests/test.lua --latency http://localhost:80/link-shortener/v1/long-link/get

PHONY: generate
generate:
		mkdir -p pkg/link_shortener/v1
		protoc --proto_path vendor.protogen --proto_path api/link_shortener/v1 \
				--go_out=pkg/link_shortener/v1 --go_opt=paths=import \
				--go-grpc_out=pkg/link_shortener/v1 --go-grpc_opt=paths=import \
				--grpc-gateway_out=pkg/link_shortener/v1 \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=import \
				--validate_out lang=go:pkg/link_shortener/v1 \
				--swagger_out=allow_merge=true,merge_file_name=api:pkg/link_shortener/v1 \
				api/link_shortener/v1/link_shortener.proto
		mv pkg/link_shortener/v1/github.com/olezhek28/link-shortener/pkg/* pkg/link_shortener/v1/
		rm -rf pkg/link_shortener/v1/github.com