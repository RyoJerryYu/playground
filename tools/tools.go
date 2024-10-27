//go:build tools

package tools

import (
	_ "github.com/RyoJerryYu/protoc-gen-plugins/cmd/protoc-gen-go-enumx"
	_ "github.com/RyoJerryYu/protoc-gen-plugins/cmd/protoc-gen-go-fieldmask"
	_ "github.com/RyoJerryYu/protoc-gen-plugins/cmd/protoc-gen-go-json"
	_ "github.com/akuity/grpc-gateway-client/protoc-gen-grpc-gateway-client"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/idodod/protoc-gen-fieldmask/cmd/protoc-gen-fieldmask"
)
