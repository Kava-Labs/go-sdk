package main

import (
	"fmt"

	"github.com/kava-labs/go-sdk/client"
)

const (
	Mnenomic = "equip town gesture square tomorrow volume nephew minute witness beef rich gadget actress egg sing secret pole winter alarm law today check violin uncover"
	Provider = "tcp://localhost:26657"
)

func main() {

	// Start a tendermint node (and kvstore) in the background to test against
	// app := kvstore.NewApplication()
	// node := rpctest.StartTendermint(app, rpctest.SuppressStdout, rpctest.RecreateConfig)
	// defer rpctest.StopTendermint(node)

	// keyManager, _ := keys.NewMnemonicKeyManager(Mnenomic)
	// // if err != nil {
	// // 	fmt.Println(err)
	// // }

	// client := client.NewClient(Provider, keyManager)

	// acc, _ := client.GetAccount(sdk.AccAddress("kava15qdefkmwswysgg4qxgqpqr35k3m49pkx2jdfnw"))

	c, err := client.NewClient("tcp://127.0.0.1:26657", "/websocket")
	if err != nil {
		panic(err)
	}

	height, err := c.GetHeight()
	if err != nil {
		panic(err)
	}

	fmt.Println("height:", height)

	blockAndTxs, err := c.GetBlockAndTxs(height)
	if err != nil {
		panic(err)
	}

	fmt.Println("blockAndTxs:", blockAndTxs)

	// Create a transaction
	// k := []byte("price")
	// v := []byte("xrp:usd")
	// tx := append(k, append([]byte("="), v...)...)

	// fmt.Println("Sending tx     :", string(tx))

	// // Broadcast the transaction and wait for it to commit (rather use
	// // c.BroadcastTxSync though in production)
	// bres, err := c.BroadcastTxCommit(tx)
	// if err != nil {
	// 	panic(err)
	// }
	// if bres.CheckTx.IsErr() || bres.DeliverTx.IsErr() {
	// 	panic("BroadcastTxCommit transaction failed")
	// }

	// Now try to fetch the value for the key
	// qres, err := c.ABCIQuery("/key", k)
	// if err != nil {
	// 	panic(err)
	// }
	// if qres.Response.IsErr() {
	// 	panic("ABCIQuery failed")
	// }
	// if !bytes.Equal(qres.Response.Key, k) {
	// 	panic("returned key does not match queried key")
	// }
	// if !bytes.Equal(qres.Response.Value, v) {
	// 	panic("returned value does not match sent value")
	// }

	// fmt.Println("Sent tx     :", string(tx))
	// fmt.Println("Queried for :", string(qres.Response.Key))
	// fmt.Println("Got value   :", string(qres.Response.Value))

	// res := GetAssetPrice("xrp:usd")
	// fmt.Println(res)

}

// MakeTxKV returns a text transaction, allong with expected key, value pair
// func MakeTxKV() ([]byte, []byte, []byte) {
// 	k := []byte(tmrand.Str(8))
// 	v := []byte(tmrand.Str(8))
// 	return k, v, append(k, append([]byte("="), v...)...)
// }

// // GetAssetPrice gets an asset's current price on kava
// func GetAssetPrice(symbol string) string {
// 	baseURL := "http://localhost:1317"

// 	// Format URL and HTTP request
// 	requestURL := fmt.Sprintf("%s/%s/%s", baseURL, "pricefeed/price", symbol)
// 	resp, err := MakeReq(requestURL)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	return string(resp)
// }
