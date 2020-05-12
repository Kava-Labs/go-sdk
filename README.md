# Kava Go SDK

The Kava Go SDK provides unique types and functionality required by services that interact with Kava's core modules.

## Components 

Kava's Go SDK includes the following components:
- client: sends transactions and queries to the Kava blockchain
- kava: msgs and types from the Kava blockchain required for complete codec registration
- keys: management of private keys and account recovery from mnenomic phrase

### Client

To initialize a new client we'll need to set up the codec and pass it into the constructor

```go
// Required imports
import (
	"github.com/kava-labs/go-sdk/client"
	"github.com/kava-labs/go-sdk/kava"
)
    
// Set up Kava prefixes and codec
config := sdk.GetConfig()
kava.SetBech32AddressPrefixes(config)
cdc := kava.MakeCodec()

// Initialize new Kava client and set codec
kavaClient := client.NewKavaClient(cdc, mnemonic, kava.Bip44CoinType, rpcAddr, networkTestnet)
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
// Required imports
import (
    "github.com/kava-labs/kava/app"
    "github.com/kava-labs/go-sdk/keys"
)

// Create a new mnemonic key manager
mnemonic := "secret words that unlock your address"
keybase, err := keys.NewMnemonicKeyManager(mnemonic, app.Bip44CoinType)
if err != nil {
    fmt.Println(err)
}
```

## Version compatibility

The go-sdk is compatible with other libraries that use different versions of Tendermint and the Cosmos SDK. To ensure compatibility, Kava's go-sdk uses stable forks of tendermint v0.33.3 and the cosmos-sdk v0.38.3. The go-sdk has the equivalent dependencies:
- github.com/kava-labs/cosmos-sdk v0.34.4-0.20200506043356-5d772797f9a3
- github.com/kava-labs/tendermint v0.33.4-0.20200506042050-c611c5308a53
