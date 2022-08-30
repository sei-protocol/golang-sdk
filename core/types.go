package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
)

type Test struct {
	Pairs        []Pair       `json:"pairs"`
	Inputs       []Input      `json:"inputs"`
	Expectations Expectations `json:"expectations"`
}

type Pair struct {
	PriceDenom string  `json:"price_denom"`
	AssetDenom string  `json:"asset_denom"`
	Ticksize   sdk.Dec `json:"ticksize"`
	Contract   string  `json:"contract"`
}

type Input struct {
	InputType string          `json:"type"`
	Details   json.RawMessage `json:"details"`
}

type FundedOrder struct {
	Moniker string         `json:"moniker"`
	Account string         `json:"account"`
	Order   OrderPlacement `json:"order"`
	Fund    string         `json:"fund"`
}

type OrderData struct {
	PositionEffect string `json:"position_effect"`
	Leverage       string `json:"leverage"`
}

type Cancel struct {
	Account string `json:"account"`
	Moniker string `json:"moniker"`
}

type Deposit struct {
	Account string `json:"account"`
	Fund    string `json:"fund"`
}

type SwapMulticollateralToBaseMsg struct {
	SwapMulticollateralToBase SwapMulticollateralToBase `json:"swap_multicollateral_to_base"`
}

type SwapMulticollateralToBase struct {
	Orders []dextypes.Order `json:"orders"`
}

type UpdateMultiCollateralWhitelistMsg struct {
	UpdateMultiCollateralWhitelist UpdateMultiCollateralWhitelist `json:"update_multi_collateral_whitelist"`
}

type UpdateMultiCollateralWhitelist struct {
	Whitelist       []string `json:"whitelist"`
	WhitelistEnable bool     `json:"whitelist_enable"`
}

type LiquidationRequest struct {
	AccountToBeLiquidated string `json:"account"`
	Fund                  string `json:"fund"`
}

type OrderPlacement struct {
	PositionDirection string `json:"position_direction"`
	Price             string `json:"price"`
	Quantity          string `json:"quantity"`
	PriceDenom        string `json:"price_denom"`
	AssetDenom        string `json:"asset_denom"`
	PositionEffect    string `json:"position_effect"`
	OrderType         string `json:"order_type"`
	Leverage          string `json:"leverage"`
}

type OracleUpdate struct {
	ExchangeRates sdk.DecCoins `json:"exchange_rates"`
}

type Sleep struct {
	TillNextEpoch bool `json:"till_next_epoch"`
}

type Expectations struct {
	Balances       []Balance                `json:"balances"`
	Positions      []Position               `json:"positions"`
	Orders         []Order                  `json:"orders"`
	PortfolioSpecs []PortfolioSpecs         `json:"portfolio_specs"`
	BankBalances   []BankBalanceExpectation `json:"bank_balances"`
}

type StartingBalance struct {
	Account string `json:"account"`
	Denom   string `json:"denom"`
}

type Balance struct {
	Account  string `json:"account"`
	Denom    string `json:"denom"`
	Amount   string `json:"amount"`
	Negative bool   `json:"negative"`
}

type BankBalanceExpectation struct {
	Account string  `json:"account"`
	Denom   string  `json:"denom"`
	Delta   sdk.Int `json:"delta"`
}

type Position struct {
	Account    string `json:"account"`
	PriceDenom string `json:"price_denom"`
	AssetDenom string `json:"asset_denom"`

	LongPosition                        ContractSignedDecimal `json:"long_position"`
	LongPositionMarginDebt              ContractSignedDecimal `json:"long_position_margin_debt"`
	LongPositioLastFundingPaymentEpoch  int64                 `json:"long_position_last_funding_payment_epoch"`
	ShortPosition                       ContractSignedDecimal `json:"short_position"`
	ShortPositionMarginDebt             ContractSignedDecimal `json:"short_position_margin_debt"`
	ShortPositioLastFundingPaymentEpoch int64                 `json:"short_position_last_funding_payment_epoch"`
}

type Order struct {
	Account    string `json:"account"`
	PriceDenom string `json:"price_denom"`
	AssetDenom string `json:"asset_denom"`

	ExpectedOrders []OrderDetails `json:"expected_orders"`
}

type OrderDetails struct {
	Price             ContractSignedDecimal `json:"price"`
	Quantity          ContractSignedDecimal `json:"quantity"`
	RemainingQuantity ContractSignedDecimal `json:"remaining_quantity"`
	Direction         string                `json:"direction"`
	Effect            string                `json:"effect"`
	Leverage          ContractSignedDecimal `json:"leverage"`
	OrderType         string                `json:"order_type"`
}

// contract types below
type ContractBalance struct {
	Amount ContractSignedDecimal `json:"amount"`
}

type ContractSignedDecimal struct {
	Decimal  string `json:"decimal"`
	Negative bool   `json:"negative"`
}

func (c *ContractSignedDecimal) Multiply(multiplier float64) ContractSignedDecimal {
	d, _ := strconv.ParseFloat(c.Decimal, 64)
	newDec := d * multiplier
	c.Decimal = fmt.Sprintf("%f", newDec)
	return *c
}

type ContractPosition struct {
	LongPosition                        ContractSignedDecimal `json:"long_position"`
	LongPositionMarginDebt              ContractSignedDecimal `json:"long_position_margin_debt"`
	LongPositioLastFundingPaymentEpoch  int64                 `json:"long_position_last_funding_payment_epoch"`
	ShortPosition                       ContractSignedDecimal `json:"short_position"`
	ShortPositionMarginDebt             ContractSignedDecimal `json:"short_position_margin_debt"`
	ShortPositioLastFundingPaymentEpoch int64                 `json:"short_position_last_funding_payment_epoch"`
}

type ContractOrder struct {
	Id                uint64                `json:"id"`
	Account           string                `json:"account"`
	PriceDenom        string                `json:"price_denom"`
	AssetDenom        string                `json:"asset_denom"`
	Price             ContractSignedDecimal `json:"price"`
	Quantity          ContractSignedDecimal `json:"quantity"`
	RemainingQuantity ContractSignedDecimal `json:"remaining_quantity"`
	Direction         string                `json:"direction"`
	Effect            string                `json:"effect"`
	Leverage          ContractSignedDecimal `json:"leverage"`
	OrderType         string                `json:"order_type"`
}

type ContractGetOrderResponse struct {
	Orders []ContractOrder `json:"orders"`
}

type PortfolioSpecs struct {
	Account            string                `json:"account"`
	Equity             ContractSignedDecimal `json:"equity"`
	TotalPositionValue ContractSignedDecimal `json:"total_position_value"`
	BuyingPower        ContractSignedDecimal `json:"buying_power"`
	UnrealizedPnl      ContractSignedDecimal `json:"unrealized_pnl"`
	Leverage           ContractSignedDecimal `json:"leverage"`
}

type MultiCollateralAccounts struct {
	Accounts []string `json:"accounts"`
}
