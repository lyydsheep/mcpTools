package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"net"
	"net/http"
)

func NewQueryIp() Tool {
	return Tool{
		Description: mcp.NewTool("query",
			mcp.WithDescription("query geo location of an IP address"),
			mcp.WithString("ip",
				mcp.Required(),
				mcp.Description("IP address to query"),
			),
		),
		Handler: queryIpHandler,
	}
}

func queryIpHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ip, ok := request.Params.Arguments["ip"]
	if !ok {
		return nil, fmt.Errorf("param ip not existed")
	}
	ipStr := fmt.Sprintf("%v", ip)
	parsedIp := net.ParseIP(ipStr)
	if parsedIp == nil {
		return nil, fmt.Errorf("invalid ip")
	}
	resp, err := http.Get("https://ip.rpcx.io/api/ip?ip=" + ipStr)
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
