package client

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
	dextypesutils "github.com/sei-protocol/sei-chain/x/dex/types/utils"
)

func SanitizeFilename(filename string) string {
	return strings.Replace(filename, "_", "", -1)
}

func asciiDecodeString(s string) []byte {
	return []byte(s)
}

func GetEventAttributeValue(response sdk.TxResponse, eventType string, attributeKey string) string {
	for _, log := range response.Logs {
		for _, event := range log.Events {
			if event.Type != eventType {
				continue
			}
			for _, attribute := range event.Attributes {
				if attribute.Key != attributeKey {
					continue
				}
				return attribute.Value
			}
		}
	}
	panic(fmt.Sprintf("Event %s attribute %s not found", eventType, attributeKey))
}

func ToSeiOrderPlacement(fundedOrder FundedOrder) dextypes.Order {
	order := fundedOrder.Order
	positionDirection, err := dextypesutils.GetPositionDirectionFromStr(order.PositionDirection)
	if err != nil {
		panic(err)
	}
	orderType, err := dextypesutils.GetOrderTypeFromStr(order.OrderType)
	if err != nil {
		panic(err)
	}
	price := sdk.MustNewDecFromStr(order.Price)
	quantity := sdk.MustNewDecFromStr(order.Quantity)
	orderData := OrderData{
		PositionEffect: order.PositionEffect,
		Leverage:       order.Leverage,
	}
	orderDataBz, err := json.Marshal(orderData)
	if err != nil {
		panic(err)
	}
	return dextypes.Order{
		Account:           fundedOrder.Account,
		PositionDirection: positionDirection,
		Price:             price,
		Quantity:          quantity,
		PriceDenom:        order.PriceDenom,
		AssetDenom:        order.AssetDenom,
		Data:              string(orderDataBz),
		OrderType:         orderType,
	}
}

func ToOrderPlacementExecuteMsg(fundedOrder FundedOrder) []byte {
	order := ToSeiOrderPlacement(fundedOrder)
	swapMsg := SwapMulticollateralToBaseMsg{
		SwapMulticollateralToBase{Orders: []dextypes.Order{order}},
	}
	msgString, err := json.Marshal(swapMsg)
	if err != nil {
		panic(err)
	}
	return []byte(msgString)
}

func ToUpdateMultiCollateralWhitelistExecuteMsg(whitelistedAccounts []string) []byte {
	whitelistedAccountsMsg := UpdateMultiCollateralWhitelistMsg{
		UpdateMultiCollateralWhitelist{
			Whitelist:       whitelistedAccounts,
			WhitelistEnable: true,
		},
	}
	msgString, err := json.Marshal(whitelistedAccountsMsg)
	if err != nil {
		panic(err)
	}
	return []byte(msgString)
}
