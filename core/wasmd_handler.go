package client

import (
	"context"

	wasmdtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// This function takes custom instantiateMsg and instantiates cosmwasm contract
// it returns the instantiated contract address if succeeds
func (c *Client) InstantiateContract(code uint64, instantiateMsg string) (*sdk.TxResponse, error) {
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	adminAddr := sdk.AccAddress(c.privKey.PubKey().Address()).String()
	msg := wasmdtypes.MsgInstantiateContract{
		Sender: adminAddr,
		Admin:  adminAddr,
		CodeID: code,
		Label:  "dex",
		Msg:    asciiDecodeString(instantiateMsg),
		Funds: []sdk.Coin{
			sdk.NewCoin("usei", sdk.NewInt(100000)),
		},
	}
	_ = txBuilder.SetMsgs(&msg)
	addGasFee(&txBuilder, c.txConfig.gasLimit, c.txConfig.gasFee)
	return c.signAndSendTx(&txBuilder)
	// return getEventAttributeValue(txResp, "instantiate", "_contract_address")
}

// This function takes custom executeMsg and call designated cosmwasm contract execute endpoint
// it returns the instantiated contract address if succeeds
// Input fund example: "1000usei". Empty string can be passed if this execution doesn't intend to attach any fund.
func (c *Client) ExecuteContract(contractAddr string, code uint64, executeMsg string, fund string) (*sdk.TxResponse, error) {
	amount, _ := sdk.ParseCoinsNormalized(fund)
	txBuilder := c.encodingConfig.TxConfig.NewTxBuilder()
	msg := wasmdtypes.MsgExecuteContract{
		Sender:   sdk.AccAddress(c.privKey.PubKey().Address()).String(),
		Contract: contractAddr,
		Msg:      asciiDecodeString(executeMsg),
		Funds:    amount,
	}

	_ = txBuilder.SetMsgs(&msg)
	addGasFee(&txBuilder, c.txConfig.gasLimit, c.txConfig.gasFee)
	return c.signAndSendTx(&txBuilder)
}

// This function takes custom queryMsg and get the corresponding state from the contract
func (c *Client) QueryContract(queryMsg string, contractAddr string) (*wasmdtypes.QuerySmartContractStateResponse, error) {
	client := wasmdtypes.NewQueryClient(c.txConfig.grpcConn)
	res, err := client.SmartContractState(
		context.Background(),
		&wasmdtypes.QuerySmartContractStateRequest{
			Address:   contractAddr,
			QueryData: asciiDecodeString(queryMsg),
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
