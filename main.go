package main

import (
	"fmt"

	// binanceKeys "github.com/binance-chain/go-sdk/keys"
	sdk "github.com/kava-labs/cosmos-sdk/types"
	tmtime "github.com/kava-labs/tendermint/types/time"

	"github.com/kava-labs/go-sdk/client"
	"github.com/kava-labs/go-sdk/kava"
	"github.com/kava-labs/go-sdk/kava/msgs"
)

const (
	mnemonic       = "fragile flip puzzle adjust mushroom gas minimum maid love coach brush cattle match analyst oak spell blur thunder unfair inch mother park toilet toddler"
	mnemonicAddr   = "kava1l0xsq2z7gqd7yly0g40y5836g0appumark77ny"
	rpcAddr        = "tcp://localhost:26657"
	networkTestnet = 0
)

func main() {

	// bip44Params := binanceKeys.NewBinanceBIP44Params(0, 0)
	// fmt.Println("bip44Params:", bip44Params)

	// Set up Kava app and initialize codec
	config := sdk.GetConfig()
	kava.SetBech32AddressPrefixes(config)
	cdc := kava.MakeCodec()

	// Initialize new Kava client and set codec
	kavaClient := client.NewKavaClient(cdc, mnemonic, kava.Bip44CoinType, rpcAddr, networkTestnet)
	kavaClient.Keybase.SetCodec(cdc)

	fromAddr, err := sdk.AccAddressFromBech32(mnemonicAddr)
	if err != nil {
		panic(err)
	}

	// acc, err := kavaClient.GetAccount(fromAddr)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Account:", acc)

	toAddr, err := sdk.AccAddressFromBech32("kava1g0qywkx6mt5jmvefv6hs7c7h333qas5ks63a6t")
	if err != nil {
		panic(err)
	}

	recipientOtherChain := "bnb1urfermcg92dwq36572cx4xg84wpk3lfpksr5g7"
	senderOtherChain := "bnb1uky3me9ggqypmrsvxk7ur6hqkzq7zmv4ed4ng7"

	timestamp := tmtime.Now().Unix()
	randomNumber, _ := kava.GenerateSecureRandomNumber()
	randomNumberHash := kava.CalculateRandomHash(randomNumber.Bytes(), timestamp)

	amount := sdk.NewCoins(sdk.NewInt64Coin("bnb", int64(50000)))
	expectedIncome := amount.String()
	heightSpan := int64(360)
	crossChain := true

	msg := msgs.NewMsgCreateAtomicSwap(fromAddr, toAddr, recipientOtherChain, senderOtherChain,
		randomNumberHash, timestamp, amount, expectedIncome, heightSpan, crossChain)
	if err := msg.ValidateBasic(); err != nil {
		panic(fmt.Sprintf("msg basic validation failed: \n%v", msg))
	}

	res, err := kavaClient.Broadcast(msg, client.Sync)
	if err != nil {
		panic(err)
	}
	if res.Code != 0 {
		panic(fmt.Sprintf("\nres.Code: %d\nLog:%s", res.Code, res.Log))
	}

	fmt.Println("tx hash:", res.Hash.String())
}

func createAtomicSwap() {
	// Set up Kava app and initialize codec
	config := sdk.GetConfig()
	kava.SetBech32AddressPrefixes(config)
	cdc := kava.MakeCodec()

	// Initialize new Kava client and set codec
	kavaClient := client.NewKavaClient(cdc, mnemonic, kava.Bip44CoinType, rpcAddr, networkTestnet)
	kavaClient.Keybase.SetCodec(cdc)

	fromAddr, err := sdk.AccAddressFromBech32(mnemonicAddr)
	if err != nil {
		panic(err)
	}

	acc, err := kavaClient.GetAccount(fromAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Account:", acc)

	toAddr, err := sdk.AccAddressFromBech32("kava1g0qywkx6mt5jmvefv6hs7c7h333qas5ks63a6t")
	if err != nil {
		panic(err)
	}

	recipientOtherChain := "bnb1urfermcg92dwq36572cx4xg84wpk3lfpksr5g7"
	senderOtherChain := "bnb1uky3me9ggqypmrsvxk7ur6hqkzq7zmv4ed4ng7"

	timestamp := tmtime.Now().Unix()
	randomNumber, _ := kava.GenerateSecureRandomNumber()
	randomNumberHash := kava.CalculateRandomHash(randomNumber.Bytes(), timestamp)

	fmt.Println("Random number:", randomNumber)

	amount := sdk.NewCoins(sdk.NewInt64Coin("bnb", int64(50000)))
	expectedIncome := amount.String()
	heightSpan := int64(360)
	crossChain := true

	msg := msgs.NewMsgCreateAtomicSwap(fromAddr, toAddr, recipientOtherChain, senderOtherChain,
		randomNumberHash, timestamp, amount, expectedIncome, heightSpan, crossChain)

	fmt.Println("msg:", msg.String())
	if err := msg.ValidateBasic(); err != nil {
		panic(fmt.Sprintf("msg basic validation failed: \n%v", msg))
	}

	res, err := kavaClient.Broadcast(msg, client.Sync)
	if err != nil {
		panic(err)
	}
	if res.Code != 0 {
		panic(fmt.Sprintf("\nres.Code: %d\nLog:%s", res.Code, res.Log))
	}

	fmt.Println("tx hash:", res.Hash.String())
}

// TODO: This tests auth account query
// acc, err := kavaClient.GetAccount(addr)
// if err != nil {
// 	panic(err)
// }
// fmt.Println("Account:", acc)

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
