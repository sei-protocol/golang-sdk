package clientexample

import (
	"testing"

	seiSdk "github.com/sei-protocol/golang-sdk/core"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	// establish grpc connection
	grpcConn, err := grpc.Dial(
		"127.0.0.1:9090",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	txConfig := seiSdk.NewTxConfig(
		"tcp://localhost:26657",
		"http://localhost:8088",
		"sei-chain",
		2000000,
		sdk.NewCoin("usei", sdk.NewInt(100000)),
		grpcConn,
	)

	seiSdk.NewClient(
		secp256k1.GenPrivKey(),
		txConfig,
		seiSdk.NewDefaultEncodingConfig(),
	)
}
