package client

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

func (c *Client) IsProposalHandled(proposalId string) bool {
	client := govtypes.NewQueryClient(c.txConfig.grpcConn)
	proposalID, err := strconv.ParseUint(proposalId, 10, 64)
	if err != nil {
		panic(err)
	}
	res, err := client.Proposal(context.Background(), &govtypes.QueryProposalRequest{ProposalId: proposalID})
	return err == nil && res.Proposal.Status == govtypes.StatusPassed
}

func (c *Client) Vote(proposalId string) error {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	from := sdk.AccAddress(c.privKey.PubKey().Address())
	proposalID, err := strconv.ParseUint(proposalId, 10, 64)
	if err != nil {
		panic(err)
	}

	byteVoteOption, err := types.VoteOptionFromString(govutils.NormalizeVoteOption("yes"))
	if err != nil {
		panic(err)
	}
	msg := govtypes.NewMsgVote(from, proposalID, byteVoteOption)
	_ = txBuilder.SetMsgs(msg)
	(txBuilder).SetGasLimit(2000000)
	(txBuilder).SetFeeAmount([]sdk.Coin{
		sdk.NewCoin("usei", sdk.NewInt(10000000)),
	})

	_, err = c.signAndSendTx(&txBuilder)
	return err
}

func (c *Client) GetBankBalance(account string, denom string) sdk.Coin {
	client := banktypes.NewQueryClient(c.txConfig.grpcConn)
	address, err := sdk.AccAddressFromBech32(account)
	if err != nil {
		panic(err)
	}
	res, err := client.Balance(context.Background(), banktypes.NewQueryBalanceRequest(address, denom))
	if err != nil {
		panic(err)
	}

	return *res.Balance
}
