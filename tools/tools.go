//go:build tools

package tools

import (
	_ "github.com/RyoJerryYu/protoc-gen-plugins/cmd/protoc-gen-go-enumx"
	_ "github.com/RyoJerryYu/protoc-gen-plugins/cmd/protoc-gen-go-fieldmask"
	_ "github.com/RyoJerryYu/protoc-gen-plugins/cmd/protoc-gen-go-json"
	_ "github.com/akuity/grpc-gateway-client/protoc-gen-grpc-gateway-client"
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/idodod/protoc-gen-fieldmask/cmd/protoc-gen-fieldmask"
	_ "go.etcd.io/etcd/etcdctl/v3"
	_ "go.etcd.io/etcd/v3/tools/etcd-dump-db"
)
