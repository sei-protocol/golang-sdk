package client

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

type Client struct {
	privKey        cryptotypes.PrivKey
	encodingConfig *EncodingConfig
	txConfig       *TxConfig
}

func NewClientWithDefaultConfig(key cryptotypes.PrivKey) *Client {
	// establish grpc connection
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}

	txConfig := NewTxConfig(
		"tcp://localhost:26657",
		"http://localhost:8088",
		"sei-chain",
		2000000,
		sdk.NewCoin("usei", sdk.NewInt(100000)),
		grpcConn,
	)

	return NewClient(key, txConfig, NewDefaultEncodingConfig())
}

func NewClient(
	key cryptotypes.PrivKey,
	txConfig *TxConfig,
	encodingConfig *EncodingConfig,
) *Client {
	return &Client{
		privKey:        key,
		txConfig:       txConfig,
		encodingConfig: encodingConfig,
	}
}

func (c *Client) GetTxConfig() *TxConfig {
	return c.txConfig
}
