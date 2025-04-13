package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
	"net/http"
	"os"
)

// NewQueryWeather creates a new tool to query weather
// 目前比较粗糙，需要提供城市代码
// 后续只需告知城市名称即可
func NewQueryWeather() Tool {
	return Tool{
		Description: mcp.NewTool("query_weather",
			mcp.WithDescription("query the weather about given address"),
			mcp.WithString("address_code",
				mcp.Required(),
				mcp.Description("Address Code to be queried"),
			),
		),
		Handler: queryWeatherHandler,
	}
}

func queryWeatherHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	addressCode, ok := request.Params.Arguments["address_code"]
	if !ok {
		return nil, fmt.Errorf("param address_code not exitsted")
	}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%v", os.Getenv("KEY"), addressCode),
		nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
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
