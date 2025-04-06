package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"net"
	"net/http"
)

func NewIpQuery() Tool {
	return Tool{
		Description: mcp.NewTool("ip_query",
			mcp.WithDescription("query geo location of an IP address"),
			mcp.WithString("ip",
				mcp.Required(),
				mcp.Description("IP address to query"),
			),
		),
		Handler: ipQueryHandler,
	}
}

func ipQueryHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ip, ok := request.Params.Arguments["ip"].(string)
	if !ok {
		return nil, fmt.Errorf("ip must be string")
	}
	parsedIp := net.ParseIP(ip)
	if parsedIp == nil {
		return nil, fmt.Errorf("invalid ip")
	}
	resp, err := http.Get("https://ip.rpcx.io/api/ip?ip=" + ip)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return mcp.NewToolResultText(string(data)), nil
}
