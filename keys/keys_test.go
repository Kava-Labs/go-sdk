package keys

import (
	"os"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kava-labs/kava/app"
	"github.com/stretchr/testify/require"
)

const (
	TestMnenomic     = "equip town gesture square tomorrow volume nephew minute witness beef rich gadget actress egg sing secret pole winter alarm law today check violin uncover"
	TestExpectedAddr = "kava1ffv7nhd3z6sych2qpqkk03ec6hzkmufy0r2s4c"
	TestKavaCoinID   = 459
)

func TestMain(m *testing.M) {
	kavaConfig := sdk.GetConfig()
	app.SetBech32AddressPrefixes(kavaConfig)
	kavaConfig.Seal()
	os.Exit(m.Run())
}
func TestNewMnemonicKeyManager(t *testing.T) {

	tests := []struct {
		name       string
		mnenomic   string
		coinID     uint32
		expectpass bool
	}{
		{"normal", TestMnenomic, TestKavaCoinID, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keyManager, err := NewMnemonicKeyManager(tc.mnenomic, tc.coinID)

			if tc.expectpass {
				require.Nil(t, err)

				// Confirm correct address
				addr := keyManager.GetAddr()
				require.Equal(t, TestExpectedAddr, addr.String())
			} else {
				require.NotNil(t, err)
			}
		})
	}
}
