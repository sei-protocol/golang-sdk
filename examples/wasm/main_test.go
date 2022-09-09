package clientexample

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	seiSdk "github.com/sei-protocol/golang-sdk/core"
	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
)

func TestClient(t *testing.T) {	
	seiClient := seiSdk.NewClientWithDefaultConfig(secp256k1.GenPrivKey())	

	// after uploading the contract code to the blockchain, it will return a auto incrementing 
	// codeId that is then used to instantiate the contract. 
	contractCodeId := uint64(0)

	// Example of an instante message 
	exampleInstantiateMsg := `{
        "whitelist": ["sei15yfjprdcq8qk5f4rea2vqh8zt769c62j77l3n6"],
        "use_whitelist":false,
        "multicollateral_whitelist":["sei15yfjprdcq8qk5f4rea2vqh8zt769c62j77l3n6"],
        "multicollateral_whitelist_enable":true,
        "admin":"sei15yfjprdcq8qk5f4rea2vqh8zt769c62j77l3n6",
        "denoms": ["SEI","ATOM","USDC","SOL","ETH"],
        "full_denom_mapping": [["usei","SEI","0.000001"],["uatom","ATOM","0.000001"],["uusdc","USDC","0.000001"],["ueth","ETH","0.000001"]],
        "limit_order_fee":{"decimal":"0.0003","negative":false},
        "market_order_fee":{"decimal":"0.0005","negative":false},
        "liquidation_order_fee":{"decimal":"0.0001","negative":false},
        "margin_ratio":{"decimal":"0.0625","negative":false},
        "max_leverage":{"decimal":"4","negative":false},
        "funding_payment_lookback":3600,
        "default_base":"USDC",
        "native_token":"SEI",
        "spot_market_contract":"XXX",
        "oracle_denom_mapping": [["usei","SEI","1"],["uatom","ATOM","1"],["uusdc","USDC","1"],["ueth","ETH","1"]],
        "funding_payment_pairs": [["USDC","ATOM"],["USDC","SEI"],["SEI","ETH"]],
        "supported_multicollateral_denoms": ["ATOM"],
        "supported_collateral_denoms": ["USDC", "SEI"],
        "default_margin_ratios":{
          "initial":"0.0625",
          "partial":"0.0303",
          "maintenance":"0.02"
        }
    }`

	response, err := seiClient.InstantiateContract(
		contractCodeId,
		exampleInstantiateMsg,
	)
	if err != nil {
		panic(err)
	}

	contractAddr := seiSdk.GetEventAttributeValue(*response, "instantiate", "_contract_address")

  // Example Register Contract
  seiClient.SendRegisterContract(
    contractAddr,
    contractCodeId,
    true,
  )

  tikSize := sdk.NewDec(int64(1))
  err = seiClient.RegisterPairAndWaitForApproval("example", contractAddr, []*dextypes.Pair{
    {PriceDenom: "USDC", AssetDenom: "ATOM", Ticksize: &tikSize},
  })
  if err != nil {
    panic(err)
  }

  seiClient.RegisterOracleDenomAndWaitForApproval("example")

	// Example deposit message 
	exampleDepositExecuteMsg := `{"deposit": {}}`
	seiClient.ExecuteContract(
		contractAddr,
		contractCodeId,
		exampleDepositExecuteMsg,
		"1000usei",
	)


  // Sending Orders
  moniker := "example-1"
  account := "alice"
  exampleSendOrderMsgString := fmt.Sprintf(`{
    "type": "order_placement",
    "details": {
        "account": "%s",
        "order": {
            "position_direction": "LONG",
            "price": "10",
            "quantity": "2",
            "price_denom": "USDC",
            "asset_denom": "ATOM",
            "position_effect": "Open",
            "order_type": "LIMIT",
            "leverage": "1"
        },
        "fund": "20000000uusdc",
        "moniker": "%s"
    }
  }`, account, moniker)
  exampleSendOrderMsgJsonEncoded, err := json.Marshal(exampleSendOrderMsgString)
  fundedOrder := seiSdk.ParseFundedOrder(exampleSendOrderMsgJsonEncoded)
  sendOrderResponse, err := seiClient.SendOrder(
    fundedOrder,
    contractAddr,
  )
  if err != nil {
    panic(err)
  }

    // Cancelling Orders
    exampleCancelOrderMsgString := fmt.Sprintf(`{
        "type": "order_cancellation",
        "details": {
            "account": "%s",
            "%s": "example-1"
        }
    }`, account, moniker)
    exampleCancelOrderMsgJsonEncoded, err := json.Marshal(exampleCancelOrderMsgString)
    cancelOrder := seiSdk.ParseCancel(exampleCancelOrderMsgJsonEncoded)
    monikerToOrderIds := map[string][]uint64{
      moniker: sendOrderResponse.OrderIds,
    }
    err = seiClient.SendCancel(
      cancelOrder,
      contractAddr,
      monikerToOrderIds,
    )
    if err != nil {
      panic(err)
    }
  
  
}
