package main

import (
	"fmt"
	"github.com/mark3labs/mcp-go/server"
	"mcpTools/tools"
)

func main() {
	// Create MCP server
	s := server.NewMCPServer(
		"ip-mcp",
		"1.0.0",
	)

	// Register tools
	tools.RegisterTool(s,
		tools.NewQueryIp(),
		tools.NewQueryWeather())

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
