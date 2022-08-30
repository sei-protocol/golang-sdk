package test

// import (
// 	"sync"
// 	"testing"

// 	"github.com/cosmos/cosmos-sdk/codec"
// 	"github.com/cosmos/cosmos-sdk/codec/types"
// 	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
// 	"github.com/cosmos/cosmos-sdk/std"
// 	"github.com/cosmos/cosmos-sdk/x/auth/tx"
// 	"github.com/sei-protocol/sei-chain/app"
// 	dextypes "github.com/sei-protocol/sei-chain/x/dex/types"
// )

// // each scenario now are independent
// var ParallelizableTests = []string{
// 	"scenario_single_order",
// 	// "scenario_0",
// 	"scenario_1",
// 	"simulation_1_open",
// 	"simulation_1_close",
// 	"simulation_2_open",
// 	"simulation_2_close",
// 	"simulation_3_close2positions",
// 	"simulation_3_close4positions",
// 	"simulation_4_open",
// 	"simulation_4_close",
// 	"simulation_5_open",
// 	"simulation_5_close",
// 	"simulation_6_open",
// 	"simulation_6_close",
// 	"simulation_7",
// 	"simulation_8",
// 	"simulation_9",
// 	"simulation_10",
// 	"simulation_12",
// 	"simulation_13",
// 	"simulation_14",
// }

// // Tests that check bank balances need to be run sequentially
// // Anything that needs to sleep till next epoch or load bank balance
// // should go here.
// var SequentialTests = []string{
// 	"simulation_11",
// 	"simulation_15",
// }

// var mu = sync.Mutex{}


// func TestAll(t *testing.T) {
// 	config := getEncodingConfig()
// 	adminKey := driver.GetKey(driver.ADMIN_KEY_NAME)

// 	var wg sync.WaitGroup

// 	// go over parallel tests first
// 	for _, filename := range ParallelizableTests {
// 		wg.Add(1)
// 		filename := filename

// 		go func() {
// 			defer wg.Done()
// 			runTest(t, filename, config, adminKey)
// 		}()
// 	}
// 	wg.Wait()

// 	// go over sequential tests
// 	for _, filename := range SequentialTests {
// 		runTest(t, filename, config, adminKey)
// 	}
// }

// func runTest(t *testing.T, filename string, config parser.EncodingConfig, adminKey cryptotypes.PrivKey) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	// instantiate  contract and deploy
// 	spotContractAddr := driver.InstantiateSpotContract(config, adminKey)
// 	t.Logf("Spot contract address %s\n", spotContractAddr)
// 	driver.SendRegisterContract(config, adminKey, spotContractAddr, 2, false)

// 	// instantiate  contract and deploy
// 	ContractAddr := driver.SendIntantiate(config, adminKey, spotContractAddr, filename)
// 	t.Logf(" contract address %s\n", ContractAddr)
// 	driver.SendRegisterContract(config, adminKey, ContractAddr, 1, true)

// 	t.Logf("Registering test-specific denoms with oracle for %s\n", filename)
// 	driver.RegisterOracleDenomAndWaitForApproval(config, adminKey, filename)

// 	t.Logf("Testing %s\n", filename)
// 	test := parser.ParseTestFile(filename)
// 	for _, pair := range test.Pairs {
// 		var contractAddr string
// 		if contractAddr = ContractAddr; pair.Contract == "spot" {
// 			contractAddr = spotContractAddr
// 		}

// 		driver.RegisterPairAndWaitForApproval(config, adminKey, filename, contractAddr, []*dextypes.Pair{
// 			{PriceDenom: pair.PriceDenom, AssetDenom: pair.AssetDenom, Ticksize: &pair.Ticksize},
// 		})
// 	}

// 	handler := driver.NewHandler(ContractAddr, spotContractAddr, filename)
// 	handler.HandleInputs(config, test.Inputs)
// 	handler.VerifyExpectations(t, test.Expectations, ContractAddr)
// }


// func SendConfigUpdate(config parser.EncodingConfig, key cryptotypes.PrivKey, accounts parser.MultiCollateralAccounts, contractAddr string) {
// 	whitelistedAccounts := []string{}
// 	for _, account := range accounts.Accounts {
// 		whitelistedAccounts = append(whitelistedAccounts, sdk.AccAddress(GetKey(account).PubKey().Address()).String())
// 	}

// 	updateMultiCollateralWhitelistMsg := parser.ToUpdateMultiCollateralWhitelistExecuteMsg(whitelistedAccounts)
// 	txBuilder := config.TxConfig.NewTxBuilder()
// 	msg := wasmdtypes.MsgExecuteContract{
// 		Sender:   sdk.AccAddress(key.PubKey().Address()).String(),
// 		Contract: contractAddr,
// 		Msg:      updateMultiCollateralWhitelistMsg,
// 	}

// 	fmt.Println(string(updateMultiCollateralWhitelistMsg))
// 	_ = txBuilder.SetMsgs(&msg)
// 	addGasFee(&txBuilder)
// 	signTx(config, key, &txBuilder)
// 	sendTx(config, key, &txBuilder)
// }

// func SendSwapOrder(config parser.EncodingConfig, key cryptotypes.PrivKey, order parser.FundedOrder, contractAddr string) {
// 	swapMsg := parser.ToOrderPlacementExecuteMsg(order)
// 	txBuilder := config.TxConfig.NewTxBuilder()
// 	msg := wasmdtypes.MsgExecuteContract{
// 		Sender:   sdk.AccAddress(key.PubKey().Address()).String(),
// 		Contract: contractAddr,
// 		Msg:      swapMsg,
// 	}

// 	fmt.Println(string(swapMsg))
// 	_ = txBuilder.SetMsgs(&msg)
// 	addGasFee(&txBuilder)
// 	signAndSendTx(config, key, &txBuilder)
// }
