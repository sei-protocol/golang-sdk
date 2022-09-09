package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramsproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	oracletypes "github.com/sei-protocol/sei-chain/x/oracle/types"
)

func (c *Client) RegisterOracleDenomAndWaitForApproval(filename string) {
	filename = SanitizeFilename(filename)
	denoms := []string{
		fmt.Sprintf("usei%s", filename),
		fmt.Sprintf("uusdc%s", filename),
		fmt.Sprintf("uatom%s", filename),
		fmt.Sprintf("ueth%s", filename),
	}
	proposalId := c.RegisterOracleDenom(filename, denoms)
	for {
		if c.IsProposalHandled(proposalId) {
			break
		}
		time.Sleep(time.Second * VOTE_WAIT_SECONDS)
	}
}

func (c *Client) RegisterOracleDenom(title string, denoms []string) string {
	proposalId := c.SendOracleDenomProposal(title, denoms)
	err := c.Vote(proposalId)
	if err != nil {
		panic(err)
	}
	return proposalId
}

func (c *Client) SendOracleDenomProposal(title string, denoms []string) string {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	from := sdk.AccAddress(c.privKey.PubKey().Address())
	existingDenoms := c.GetOracleWhitelist()
	for _, newDenom := range denoms {
		existing := false
		for _, existingDenom := range existingDenoms {
			if newDenom == existingDenom.Name {
				existing = true
				break
			}
		}
		if existing {
			continue
		}
		existingDenoms = append(existingDenoms, oracletypes.Denom{
			Name: newDenom,
		})
	}
	serializedDenoms, err := json.Marshal(existingDenoms)
	if err != nil {
		panic(err)
	}
	content := paramsproposal.ParameterChangeProposal{
		Title:       title,
		Description: title,
		Changes: []paramsproposal.ParamChange{
			{
				Subspace: oracletypes.ModuleName,
				Key:      string(oracletypes.KeyWhitelist),
				Value:    string(serializedDenoms),
			},
		},
	}
	deposit := sdk.NewCoins(
		sdk.NewCoin("usei", govtypes.DefaultMinDepositTokens),
	)
	msg, err := govtypes.NewMsgSubmitProposal(&content, deposit, from)
	if err != nil {
		panic(err)
	}
	_ = txBuilder.SetMsgs(msg)
	(txBuilder).SetGasLimit(2000000)
	(txBuilder).SetFeeAmount([]sdk.Coin{
		sdk.NewCoin("usei", sdk.NewInt(10000000)),
	})
	txResp, err := c.signAndSendTx(&txBuilder)
	if err != nil {
		panic(err)
	}
	return GetEventAttributeValue(*txResp, "submit_proposal", "proposal_id")
}

func (c *Client) GetOracleWhitelist() oracletypes.DenomList {
	client := oracletypes.NewQueryClient(c.txConfig.grpcConn)
	res, err := client.Params(context.Background(), &oracletypes.QueryParamsRequest{})
	if err != nil {
		panic(err)
	}
	return res.Params.Whitelist
}

func (c *Client) SendOraclePrice(coins sdk.DecCoins) error {
	exchangeRatesStrs := []string{}
	for _, coin := range coins {
		exchangeRatesStrs = append(exchangeRatesStrs, fmt.Sprintf("%s%s", coin.Amount.String(), coin.Denom))
	}
	exchangeRatesStr := strings.Join(exchangeRatesStrs, ",")
	from := sdk.AccAddress(c.privKey.PubKey().Address())
	validator := sdk.ValAddress(from)
	hash := oracletypes.GetAggregateVoteHash(ORACLE_HASH, exchangeRatesStr, validator)

	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	prevoteMsg := oracletypes.NewMsgAggregateExchangeRatePrevote(hash, from, validator)
	_ = txBuilder.SetMsgs(prevoteMsg)
	(txBuilder).SetGasLimit(2000000)
	(txBuilder).SetFeeAmount([]sdk.Coin{
		sdk.NewCoin("usei", sdk.NewInt(100000)),
	})
	_, err := c.signAndSendTx(&txBuilder)
	if err != nil {
		return err
	}

	voteResp, err := c.sendOracleVote(exchangeRatesStr)
	if err != nil {
		return err
	}

	for voteResp.Code != 0 {
		if voteResp.Code == 9 {
			// Error code 11 means it's not yet time to vote for the previous prevote, so
			// we will retry until success
			voteResp, err = c.sendOracleVote(exchangeRatesStr)
			if err != nil {
				return err
			}
		} else {
			return errors.New("Failed to submit oracle price")
		}
	}

	return nil
}

func (c *Client) sendOracleVote(exchangeRatesStr string) (*sdk.TxResponse, error) {
	voter := sdk.AccAddress(c.privKey.PubKey().Address())
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	prevoteMsg := oracletypes.NewMsgAggregateExchangeRateVote(
		ORACLE_HASH,
		exchangeRatesStr,
		voter,
		sdk.ValAddress(voter),
	)
	_ = txBuilder.SetMsgs(prevoteMsg)
	(txBuilder).SetGasLimit(2000000)
	(txBuilder).SetFeeAmount([]sdk.Coin{
		sdk.NewCoin("usei", sdk.NewInt(100000)),
	})
	
	resp, err := c.signAndSendTx(&txBuilder)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
