# Kava Go SDK

The Kava Go SDK provides unique types and functionality required by services that interact with Kava's core modules.

## Components 

Kava's Go SDK includes the following components:
- client: sends transactions and queries to Kava's blockchain
- keys: management of private keys and account recovery from mnenomic phrase

### Client

To initialize a new client we'll need to set up the codec and pass it into the constructor

```go
// Initialize codec with Kava's prefixes and coin type
config := sdk.GetConfig()
app.SetBech32AddressPrefixes(config)
app.SetBip44CoinType(config)
cdc := app.MakeCodec()

// Initialize new Kava client and set codec
kavaClient := client.NewKavaClient(cdc, mnemonic, app.Bip44CoinType, rpcAddr, networkTestnet)
kavaClient.Keybase.SetCodec(cdc)
```

Let's use our new client to query the Kava blockchain for information about an account

```go
kavaAddress := "kava1l0xsq2z7gqd7yly0g40y5836g0appumark77ny"
addr, err := sdk.AccAddressFromBech32(kavaAddress)
if err != nil {
    panic(err)
}

acc, err := kavaClient.GetAccount(addr)
if err != nil {
    panic(err)
}

fmt.Println("Account:", acc)
```

### Keys

Client uses the keys package for signing transactions, but keys can also be used standalone. The following example shows how to create a new key manager from a mnemonic phrase

```go
// Create a new mnemonic key manager
mnemonic := "secret words that unlock your address"
keybase, err := keys.NewMnemonicKeyManager(mnemonic, app.Bip44CoinType)
if err != nil {
    fmt.Println(err)
}
```

## Version compatibility

We recommend using the Go SDK with the following versions of relevant software:
- github.com/cosmos/cosmos-sdk v0.38.3
- github.com/tendermint/tendermint v0.33.3
- github.com/kava-labs/kava v0.7.1-0.20200424154444-e9a73b80ce91
