package client

import (
	"encoding/json"
	"os"
)

func ParseTestFile(filename string) Test {
	pwd, _ := os.Getwd()
	file, _ := os.ReadFile(pwd + "/tests/" + filename + ".json")
	test := Test{}
	if err := json.Unmarshal([]byte(file), &test); err != nil {
		panic(err)
	}
	return test
}

func ParseMultiCollateralAccounts(raw json.RawMessage) MultiCollateralAccounts {
	accounts := MultiCollateralAccounts{}
	if err := json.Unmarshal(raw, &accounts); err != nil {
		panic(err)
	}
	return accounts
}

func ParseFundedOrder(raw json.RawMessage) FundedOrder {
	order := FundedOrder{}
	if err := json.Unmarshal(raw, &order); err != nil {
		panic(err)
	}
	return order
}

func ParseCancel(raw json.RawMessage) Cancel {
	cancel := Cancel{}
	if err := json.Unmarshal(raw, &cancel); err != nil {
		panic(err)
	}
	return cancel
}

func ParseDeposit(raw json.RawMessage) Deposit {
	deposit := Deposit{}
	if err := json.Unmarshal(raw, &deposit); err != nil {
		panic(err)
	}
	return deposit
}

func ParseStartingBalance(raw json.RawMessage) StartingBalance {
	startingBalance := StartingBalance{}
	if err := json.Unmarshal(raw, &startingBalance); err != nil {
		panic(err)
	}
	return startingBalance
}

func ParseOracleUpdate(raw json.RawMessage) OracleUpdate {
	oracleUpdate := OracleUpdate{}
	if err := json.Unmarshal(raw, &oracleUpdate); err != nil {
		panic(err)
	}
	return oracleUpdate
}

func ParseLiquidation(raw json.RawMessage) LiquidationRequest {
	liquidationRequest := LiquidationRequest{}
	if err := json.Unmarshal(raw, &liquidationRequest); err != nil {
		panic(err)
	}
	return liquidationRequest
}

func ParseSleep(raw json.RawMessage) Sleep {
	sleep := Sleep{}
	if err := json.Unmarshal(raw, &sleep); err != nil {
		panic(err)
	}
	return sleep
}

func ParseContractBalance(balanceBytes []byte) ContractBalance {
	var res ContractBalance
	if err := json.Unmarshal(balanceBytes, &res); err != nil {
		panic(err)
	}
	return res
}

func ParseContractPosition(positionBytes []byte) ContractPosition {
	var res ContractPosition
	if err := json.Unmarshal(positionBytes, &res); err != nil {
		panic(err)
	}
	return res
}

func ParseContractGetOrderResponse(orderBytes []byte) ContractGetOrderResponse {
	var res ContractGetOrderResponse
	if err := json.Unmarshal(orderBytes, &res); err != nil {
		panic(err)
	}
	return res
}

func ParsePortfolioSpecs(specsBytes []byte) PortfolioSpecs {
	var res PortfolioSpecs
	if err := json.Unmarshal(specsBytes, &res); err != nil {
		panic(err)
	}
	return res
}
