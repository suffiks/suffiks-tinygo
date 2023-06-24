gen-proto:
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest

	protoc \
		-I ../suffiks/extension/proto/ \
		../suffiks/extension/proto/extension.proto \
		../suffiks/extension/proto/k8s.proto \
		--go_out=. \
		--go-vtproto_out=. \
    --go-vtproto_opt=features=marshal+unmarshal+size \

	rm -rf protogen
	mv ./github.com/suffiks/suffiks/extension/protogen ./protogen
	rm -rf ./github.com

gen-env:
	go run ./internal/gen-env/main.go
