package client

import (
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
)

func (c *Client) SendRegisterContract(contractAddr string, codeId uint64, needHook bool) (*sdk.TxResponse, error) {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := dextypes.MsgRegisterContract{
		Creator: sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Contract: &dextypes.ContractInfoV2{
			CodeId:            codeId,
			ContractAddr:      contractAddr,
			NeedOrderMatching: true,
			NeedHook:          needHook,
		},
	}
	_ = txBuilder.SetMsgs(&msg)
	(txBuilder).SetGasLimit(2000000)
	(txBuilder).SetFeeAmount([]sdk.Coin{
		sdk.NewCoin("usei", sdk.NewInt(100000)),
	})
	return c.signAndSendTx(&txBuilder)
}

func (c *Client) RegisterPairAndWaitForApproval(
	title string,
	contractAddr string,
	pairs []*dextypes.Pair,
) error {
	proposalResp, err := c.RegisterPair(title, contractAddr, pairs)
	if err != nil {
		return err
	}

	proposalId := GetEventAttributeValue(*proposalResp, "submit_proposal", "proposal_id")
	for {
		if c.IsProposalHandled(proposalId) {
			return nil
		}
		time.Sleep(time.Second * VOTE_WAIT_SECONDS)
	}
}

func (c *Client) RegisterPair(
	title string,
	contractAddr string,
	pairs []*dextypes.Pair,
) (*sdk.TxResponse, error) {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	from := sdk.AccAddress(c.privKey.PubKey().Address())

	msg := types.NewMsgRegisterPairs(
		from.String(),
		[]dextypes.BatchContractPair{
			{
				ContractAddr: contractAddr,
				Pairs:        pairs,
			},
		},
	)

	_ = txBuilder.SetMsgs(msg)
	(txBuilder).SetGasLimit(2000000)
	(txBuilder).SetFeeAmount([]sdk.Coin{
		sdk.NewCoin("usei", sdk.NewInt(10000000)),
	})

	return c.signAndSendTx(&txBuilder)
}

func (c *Client) SendOrder(order FundedOrder, contractAddr string) (dextypes.MsgPlaceOrdersResponse, error) {
	seiOrder := ToSeiOrderPlacement(order)
	orderPlacements := []*dextypes.Order{&seiOrder}
	amount, _ := sdk.ParseCoinsNormalized(order.Fund)
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := dextypes.MsgPlaceOrders{
		Creator:      sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Orders:       orderPlacements,
		ContractAddr: contractAddr,
		Funds:        amount,
	}
	_ = txBuilder.SetMsgs(&msg)
	resp, err := c.signAndSendTx(&txBuilder)
	if err != nil {
		return dextypes.MsgPlaceOrdersResponse{}, err
	}

	msgResp := sdk.TxMsgData{}
	respDataBytes, err := hex.DecodeString(resp.Data)
	if err != nil {
		return dextypes.MsgPlaceOrdersResponse{}, err
	}

	if err := msgResp.Unmarshal(respDataBytes); err != nil {
		return dextypes.MsgPlaceOrdersResponse{}, err
	}

	orderPlacementResponse := dextypes.MsgPlaceOrdersResponse{}
	orderMsgData := msgResp.Data[0].Data
	if err := orderPlacementResponse.Unmarshal([]byte(orderMsgData)); err != nil {
		return orderPlacementResponse, err
	}

	return orderPlacementResponse, nil
}

func (c *Client) SendCancel(
	order CancelOrder,
	contractAddr string,
) error {
	seiCancellation := ToSeiCancelOrderPlacement(order)
	orderCancellations := []*dextypes.Cancellation{&seiCancellation}
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := dextypes.MsgCancelOrders{
		Creator:       sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Cancellations: orderCancellations,
		ContractAddr:  contractAddr,
	}
	_ = txBuilder.SetMsgs(&msg)
	addGasFee(&txBuilder, c.txConfig.gasLimit, c.txConfig.gasFee)

	_, err := c.signAndSendTx(&txBuilder)
	return err
}
