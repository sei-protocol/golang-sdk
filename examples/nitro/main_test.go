package main

import (
	seiSdk "github.com/sei-protocol/golang-sdk/core"
	"github.com/sei-protocol/sei-chain/x/nitro/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func main() {
	seiClient := seiSdk.NewClientWithDefaultConfig(secp256k1.GenPrivKey())

	// prepare fraud proof
	startSlot := 0
	endSlot := 5
	fraudStatePubKey := "state_pubkey"
	proof := &types.MerkleProof{}
	accountStates := []*types.Account{}
	programs := []*types.Account{}
	sysvarAccounts := []*types.Account{}

	_, err := seiClient.SubmitFraudChallenge(
		uint64(startSlot),
		uint64(endSlot),
		fraudStatePubKey,
		proof,
		accountStates,
		programs,
		sysvarAccounts,
		seiClient.GetTxConfig().GetGasLimit(),
		seiClient.GetTxConfig().GetGasFee(),
	)
	if err != nil {
		panic(err)
	}
}
