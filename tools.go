// +build tools

// This file used for attaching tools dependencies to the project
package tools

import (
	_ "github.com/golang/protobuf/protoc-gen-go"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
)
