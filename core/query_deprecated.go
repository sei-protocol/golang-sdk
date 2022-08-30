package client

// import (
// 	"context"
// 	"fmt"
// 	"strconv"

// 	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
// 	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
// 	oracletypes "github.com/sei-protocol/sei-chain/x/oracle/types"
// )

// func (c *Client) QueryBalance(address string, denom string, contractAddr string) parser.ContractBalance {
// 	query := fmt.Sprintf("{\"get_balance\":{\"account\":\"%s\",\"symbol\":\"%s\"}}", address, denom)
// 	response := queryWasm(query, contractAddr)

// 	fmt.Println("QueryBalance res:\n", parser.ParseContractBalance(response))
// 	return parser.ParseContractBalance(response)
// }

// func (c *Client) queryPosition(address string, price_denom string, asset_denom string, contractAddr string) parser.ContractPosition {
// 	query := fmt.Sprintf("{\"get_position\":{\"account\":\"%s\",\"price_denom\":\"%s\", \"asset_denom\":\"%s\"}}", address, price_denom, asset_denom)
// 	response := queryWasm(query, contractAddr)

// 	// TODO: remove after finish integration tests
// 	fmt.Println("queryPosition res:\n", parser.ParseContractPosition(response))
// 	return parser.ParseContractPosition(response)
// }

// func (c *Client) queryOrder(address string, price_denom string, asset_denom string, contractAddr string) parser.ContractGetOrderResponse {
// 	query := fmt.Sprintf("{\"get_order\":{\"account\":\"%s\",\"price_denom\":\"%s\", \"asset_denom\":\"%s\"}}", address, price_denom, asset_denom)
// 	response := queryWasm(query, contractAddr)

// 	// TODO: remove after finish integration tests
// 	fmt.Println("queryOrder res:\n", parser.ParseContractGetOrderResponse(response))
// 	return parser.ParseContractGetOrderResponse(response)
// }

// func (c *Client) queryPortfolioSpecs(address string, contractAddr string) parser.PortfolioSpecs {
// 	query := fmt.Sprintf("{\"get_portfolio_specs\":{\"account\":\"%s\"}}", address)
// 	response := queryWasm(query, contractAddr)

// 	// TODO: remove after finish integration tests
// 	fmt.Println("queryPortfolioSpecs res:\n", parser.ParsePortfolioSpecs(response))
// 	return parser.ParsePortfolioSpecs(response)
// }
