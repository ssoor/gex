// Code generated by github.com/izumin5210/gex. DO NOT EDIT.

// +build tools

package tools

// tool dependencies
import (
	_ "github.com/gogo/protobuf/protoc-gen-gogo"
	_ "github.com/gogo/protobuf/protoc-gen-gogofast"
	_ "github.com/golang/mock/mockgen"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	_ "golang.org/x/lint/golint"
)

// If you want to use tools, please run the following command:
//  go generate ./tools.go
//
//go:generate go build -v -o=./bin/protoc-gen-gogo ./vendor/github.com/gogo/protobuf/protoc-gen-gogo
//go:generate go build -v -o=./bin/protoc-gen-gogofast ./vendor/github.com/gogo/protobuf/protoc-gen-gogofast
//go:generate go build -v -o=./bin/mockgen ./vendor/github.com/golang/mock/mockgen
//go:generate go build -v -o=./bin/protoc-gen-grpc-gateway ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
//go:generate go build -v -o=./bin/protoc-gen-swagger ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
//go:generate go build -v -o=./bin/golint ./vendor/golang.org/x/lint/golint

