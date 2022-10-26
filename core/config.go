package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/sei-protocol/sei-chain/app"
)

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	// NOTE: this field will be renamed to Codec
	Marshaler codec.Codec
	TxConfig  client.TxConfig
	Amino     *codec.LegacyAmino
}

type TxConfig struct {
	nodeURI   string
	keyServer string
	chainId   string
	gasLimit  uint64
	gasFee    sdk.Coin
	grpcConn  *grpc.ClientConn
}

func NewTxConfig(
	nodeURI string,
	keyServer string,
	chainId string,
	gasLimit uint64,
	gasFee sdk.Coin,
	grpcConn *grpc.ClientConn,
) *TxConfig {
	return &TxConfig{
		nodeURI:   nodeURI,
		keyServer: keyServer,
		chainId:   chainId,
		gasLimit:  gasLimit,
		gasFee:    gasFee,
		grpcConn:  grpcConn,
	}
}

func NewDefaultEncodingConfig() *EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	config := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}
	std.RegisterLegacyAminoCodec(config.Amino)
	std.RegisterInterfaces(config.InterfaceRegistry)
	app.ModuleBasics.RegisterLegacyAminoCodec(config.Amino)
	app.ModuleBasics.RegisterInterfaces(config.InterfaceRegistry)
	return &config
}

func (t *TxConfig) GetGasLimit() uint64 {
	return t.gasLimit
}

func (t *TxConfig) GetGasFee() sdk.Coin {
	return t.gasFee
}