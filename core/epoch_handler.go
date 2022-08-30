package client

import (
	"context"

	epochtypes "github.com/sei-protocol/sei-chain/x/epoch/types"
)

func (c *Client) GetCurrentEpoch() uint64 {
	client := epochtypes.NewQueryClient(c.txConfig.grpcConn)
	res, err := client.Epoch(context.Background(), &epochtypes.QueryEpochRequest{})
	if err != nil {
		panic(err)
	}
	return res.Epoch.CurrentEpoch
}
