package clientexample

import (
	"testing"

	seiSdk "github.com/sei-protocol/golang-sdk/core"
	"github.com/stretchr/testify/assert"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
)

func TestGetEpoch(t *testing.T) {	
	seiClient := seiSdk.NewClientWithDefaultConfig(secp256k1.GenPrivKey())

	currentEpoch := seiClient.GetCurrentEpoch()
	println("Current Epoch: {}", currentEpoch)
	assert.GreaterOrEqual(t, currentEpoch, uint64(0), "Epoch should be >= 0")
}
