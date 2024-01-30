package client

import (
	"context"
	"github.com/blocto/solana-go-sdk/rpc"

	"github.com/blocto/solana-go-sdk/program/token"
)

func (c *Client) GetTokenAccount(ctx context.Context, base58Addr string) (token.TokenAccount, error) {
	accountInfo, err := c.GetAccountInfo(ctx, base58Addr)
	if err != nil {
		return token.TokenAccount{}, err
	}
	return token.DeserializeTokenAccount(accountInfo.Data, accountInfo.Owner)
}

func (c *Client) GetTokenAccountConfirmed(ctx context.Context, base58Addr string) (token.TokenAccount, error) {
	accountInfo, err := c.GetAccountInfoWithConfig(ctx, base58Addr, GetAccountInfoConfig{Commitment: rpc.CommitmentConfirmed})
	if err != nil {
		return token.TokenAccount{}, err
	}
	return token.DeserializeTokenAccount(accountInfo.Data, accountInfo.Owner)
}
