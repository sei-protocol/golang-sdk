# Sei SDK (golang)

![Sei Logo](https://raw.githubusercontent.com/sei-protocol/sei-chain/master/assets/SeiLogo.png)

The Sei SDK is a framework for builidng [Sei](https://github.com/sei-protocol/sei-chain) applications in Golang. It provides helpers that interacts with the chain modules through Golang. 

See `examples/` for sample of how to use the SDK

## Cosmos Module
Path: `core/wasmd_handler.go`

Used to interact with the native Cosmos and interacting with smart contracts.

```[golang]
func (c *Client) InstantiateContract(code uint64, instantiateMsg string) (*sdk.TxResponse, error) {...}

func (c *Client) ExecuteContract(contractAddr string, code uint64, executeMsg string, fund string) (*sdk.TxResponse, error) {...}

func (c *Client) QueryContract(queryMsg string, contractAddr string) (*wasmdtypes.QuerySmartContractStateResponse, error) {...}
```

## Dex Module
Path: `core/dex_handler.go`

Allows smart contracts to leverage Sei's underlying orderbook infrastructure to spinup and customize new markets.

```[golang]
func (c *Client) SendRegisterContract(contractAddr string, codeId uint64, needHook bool) (*sdk.TxResponse, error) {...}
 
func (c *Client) RegisterPairAndWaitForApproval(title string, contractAddr string, pairs []*dextypes.Pair) error {...}

func (c *Client) RegisterPair(title string, contractAddr string, pairs []*dextypes.Pair) (*sdk.TxResponse, error) {...}
```

## Epoch Module
Path: `core/epoch_handler.go`

The epoch module gives modules the ability to execute code per constant period of time instead of based on block height.

```[golang]
func (c *Client) GetCurrentEpoch() uint64 {...}
```


## Oracle Module
Path: `core/oracle_handler.go`

Sei Network has an oracle module to support asset exchange rate pricing for use by other modules and contracts.

```[golang]
func (c *Client) RegisterOracleDenomAndWaitForApproval(filename string) {...}

func (c *Client) RegisterOracleDenom(title string, denoms []string) string {...}

func (c *Client) SendOracleDenomProposal(title string, denoms []string) string {...}

func (c *Client) GetOracleWhitelist() oracletypes.DenomList {...}

func (c *Client) SendOraclePrice(coins sdk.DecCoins) error {...}

func (c *Client) sendOracleVote(exchangeRatesStr string) (*sdk.TxResponse, error) {...}
```
