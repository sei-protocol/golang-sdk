package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/nitro/types"
)

func (c *Client) SubmitFraudChallenge(
	startSlot uint64,
	endSlot uint64,
	fraudState string,
	merkleProof *types.MerkleProof,
	accountStates []*types.Account,
	programs []*types.Account,
	gasLimit uint64,
	gasWanted sdk.Coin,
) (*sdk.TxResponse, error) {
	senderAddr := sdk.AccAddress(c.privKey.PubKey().Address()).String()
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := types.MsgSubmitFraudChallenge{
		Sender:        senderAddr,
		StartSlot:     startSlot,
		EndSlot:       endSlot,
		FraudState:    fraudState,
		MerkleProof:   merkleProof,
		AccountStates: accountStates,
		Programs:      programs,
	}
	_ = txBuilder.SetMsgs(&msg)
	(txBuilder).SetGasLimit(gasLimit)
	(txBuilder).SetFeeAmount([]sdk.Coin{gasWanted})
	return c.signAndSendTx(&txBuilder)
}
