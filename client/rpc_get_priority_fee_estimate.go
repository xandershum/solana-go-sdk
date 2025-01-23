package client

import (
	"context"
	"github.com/blocto/solana-go-sdk/rpc"
)

// GetPriorityFeeEstimate returns the estimated priority fee
func (c *Client) GetPriorityFeeEstimate(ctx context.Context, request rpc.GetPriorityFeeEstimateRequest) (rpc.GetPriorityFeeEstimate, error) {
	return process(
		func() (rpc.JsonRpcResponse[rpc.GetPriorityFeeEstimate], error) {
			return c.RpcClient.GetPriorityFeeEstimate(ctx, request)
		},
		forward[rpc.GetPriorityFeeEstimate],
	)
}
