package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// RunMCPServer starts the MCP server for the backtest functionality.
// It configures the server with the necessary tools and handlers for running backtests.
func RunMCPServer() *server.MCPServer {
	_ = "STUB: not implemented"
	// Create a new MCP server
	return nil
}

// Add backtest tool with schema

// Add tool handler using the typed handler

// handleBacktest processes a backtest request by executing the specified strategy
// with the provided OHLCV data. It returns the transaction actions and the
// outcome of the backtest as a JSON string.
//
// The function takes a context and a tool call request, along with the parsed
// strategy request containing the strategy type and data. It runs the backtest,
// converts the resulting actions to a numeric format (1 for Buy, -1 for Sell, 0 for Hold),
// and marshals the response into a JSON object.
//
// If the backtest fails or no transactions are generated, it returns a tool result error.
func handleBacktest(_ context.Context, _ mcp.CallToolRequest, args StrategyRequest) (*mcp.CallToolResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert strategy.Action to numeric values (1=BUY, 0=HOLD, -1=SELL)

// strategy.Hold is 0 by default

// Create the response JSON

// GetAllStrategyTypes returns a slice of all available strategy types as strings.
// This list includes base, trend, momentum, and volume strategies, providing a
// comprehensive set of options for backtesting.
func GetAllStrategyTypes() []string {
	_ = "STUB: not implemented"

	// Base strategies
	return nil
}

// Trend strategies

// Momentum strategies

// Volume strategies
