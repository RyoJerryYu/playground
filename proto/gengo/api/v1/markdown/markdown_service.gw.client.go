// Code generated by protoc-gen-grpc-gateway-client. DO NOT EDIT.
// source: api/v1/markdown/markdown_service.proto

package markdown

import (
	context "context"
	fmt "fmt"
	gateway "github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
	url "net/url"
)

// MarkdownServiceGatewayClient is the interface for MarkdownService service client.
type MarkdownServiceGatewayClient interface {
	// ParseMarkdown parses the given markdown content and returns a list of nodes.
	ParseMarkdown(context.Context, *ParseMarkdownRequest) (*ParseMarkdownResponse, error)
	// RestoreMarkdownNodes restores the given nodes to markdown content.
	RestoreMarkdownNodes(context.Context, *RestoreMarkdownNodesRequest) (*RestoreMarkdownNodesResponse, error)
	// StringifyMarkdownNodes stringify the given nodes to plain text content.
	StringifyMarkdownNodes(context.Context, *StringifyMarkdownNodesRequest) (*StringifyMarkdownNodesResponse, error)
	// GetLinkMetadata returns metadata for a given link.
	GetLinkMetadata(context.Context, *GetLinkMetadataRequest) (*LinkMetadata, error)
}

func NewMarkdownServiceGatewayClient(c gateway.Client) MarkdownServiceGatewayClient {
	return &markdownServiceGatewayClient{
		gwc: c,
	}
}

type markdownServiceGatewayClient struct {
	gwc gateway.Client
}

func (c *markdownServiceGatewayClient) ParseMarkdown(ctx context.Context, req *ParseMarkdownRequest) (*ParseMarkdownResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/api/v1/markdown:parse")
	gwReq.SetBody(req)
	return gateway.DoRequest[ParseMarkdownResponse](ctx, gwReq)
}

func (c *markdownServiceGatewayClient) RestoreMarkdownNodes(ctx context.Context, req *RestoreMarkdownNodesRequest) (*RestoreMarkdownNodesResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/api/v1/markdown/node:restore")
	gwReq.SetBody(req)
	return gateway.DoRequest[RestoreMarkdownNodesResponse](ctx, gwReq)
}

func (c *markdownServiceGatewayClient) StringifyMarkdownNodes(ctx context.Context, req *StringifyMarkdownNodesRequest) (*StringifyMarkdownNodesResponse, error) {
	gwReq := c.gwc.NewRequest("POST", "/api/v1/markdown/node:stringify")
	gwReq.SetBody(req)
	return gateway.DoRequest[StringifyMarkdownNodesResponse](ctx, gwReq)
}

func (c *markdownServiceGatewayClient) GetLinkMetadata(ctx context.Context, req *GetLinkMetadataRequest) (*LinkMetadata, error) {
	gwReq := c.gwc.NewRequest("GET", "/api/v1/markdown/link:metadata")
	q := url.Values{}
	q.Add("link", fmt.Sprintf("%v", req.Link))
	gwReq.SetQueryParamsFromValues(q)
	return gateway.DoRequest[LinkMetadata](ctx, gwReq)
}
