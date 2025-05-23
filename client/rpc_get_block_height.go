package client

import (
	"context"
	"github.com/blocto/solana-go-sdk/rpc"
)

type GetBlockHeightConfig struct {
	Commitment rpc.Commitment
}

func (c GetBlockHeightConfig) toRpc() rpc.GetBlockHeightConfig {
	return rpc.GetBlockHeightConfig{
		Commitment: c.Commitment,
	}
}

// GetBlockHeight fetch block height
func (c *Client) GetBlockHeight(ctx context.Context) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetBlockHeight(ctx)
		},
		forward[uint64],
	)
}

// GetBlockHeightWithConfig fetch block height with specific commitment
func (c *Client) GetBlockHeightWithConfig(ctx context.Context, cfg GetBlockHeightConfig) (uint64, error) {
	return process(
		func() (rpc.JsonRpcResponse[uint64], error) {
			return c.RpcClient.GetBlockHeightWithConfig(ctx, cfg.toRpc())
		},
		forward[uint64],
	)
}
