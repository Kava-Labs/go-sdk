package main

import (
	"fmt"

	binanceKeys "github.com/binance-chain/go-sdk/keys"

	"github.com/kava-labs/go-sdk/client"
	"github.com/kava-labs/go-sdk/types"
)

const (
	mnemonic       = "fragile flip puzzle adjust mushroom gas minimum maid love coach brush cattle match analyst oak spell blur thunder unfair inch mother park toilet toddler"
	mnemonicAddr   = "kava1l0xsq2z7gqd7yly0g40y5836g0appumark77ny"
	rpcAddr        = "tcp://localhost:26657"
	networkTestnet = 1
)

func main() {

	bip44Params := binanceKeys.NewBinanceBIP44Params(0, 0)
	fmt.Println("bip44Params:", bip44Params)

	// Set up Kava app and initialize codec
	config := types.GetConfig()
	types.SetBech32AddressPrefixes(config)
	cdc := types.MakeCodec()

	// Initialize new Kava client and set codec
	kavaClient := client.NewKavaClient(cdc, mnemonic, types.Bip44CoinType, rpcAddr, networkTestnet)
	kavaClient.Keybase.SetCodec(cdc)

	addr, err := types.AccAddressFromBech32(mnemonicAddr)
	if err != nil {
		panic(err)
	}

	// TODO: Won't work until we call RegisterCodec for kava app's ModuleBasics,
	//		 but kava/app will import from tendermint/tendermint and cosmos/cosmos-sdk
	//		 and must be updated to kava-labs/tendermint and kava-labs/cosmos-sdk first
	acc, err := kavaClient.GetAccount(addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account:", acc)

	// TODO: This is for testing kava's go-sdk keybase
	// keybase, err := keys.NewMnemonicKeyManager(mnemonic, app.Bip44CoinType)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// keybaseAddr := keybase.GetAddr().String()
	// if mnemonicAddr != keybaseAddr {
	// 	fmt.Println("Expect:", mnemonicAddr)
	// 	fmt.Println("Actual:", keybaseAddr)
	// } else {
	// 	fmt.Println("Success!")
	// }
}
