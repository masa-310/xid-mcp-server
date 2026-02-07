package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/rs/xid"
)

type GenXIDParams struct {
	Count string `json:"count"`
}

func GenXID(ctx context.Context, req *mcp.CallToolRequest, args GenXIDParams) (*mcp.CallToolResult, any, error) {
	count := 1
	if args.Count != "" {
		var err error
		count, err = strconv.Atoi(args.Count)
		if err != nil || count < 1 {
			count = 1
		}
	}

	xids := make([]string, count)
	for i := 0; i < count; i++ {
		generatedXID := xid.New()
		xids[i] = generatedXID.String()
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: strings.Join(xids, ", "),
			},
		},
	}, nil, nil
}

func main() {
	ctx := context.Background()
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "XID Generator",
		Version: "1.0.0",
	}, nil)
	mcp.AddTool(server, &mcp.Tool{
		Name:        "XID Generator",
		Description: "Generates one or more unique XID identifiers.",
	}, GenXID)

	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		fmt.Println("Error running server:", err)
	}
}
