package tools

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type handler = func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error)

type Tool struct {
	Description mcp.Tool
	Handler     handler
}

func RegisterTool(s *server.MCPServer, tools ...Tool) {
	for _, tool := range tools {
		s.AddTool(tool.Description, tool.Handler)
	}
}
