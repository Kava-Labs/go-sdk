package keys

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
)

const (
	defaultBIP39Passphrase = ""
	PartialBIP44Prefix     = "44"
	PartialPath            = "0'/0/0"
)

// KeyManager is an interface for common methods on KeyManagers
type KeyManager interface {
	GetKeyBase() keyring.Keyring
	GetKeyRing() keyring.Info
	Sign(legacytx.StdSignMsg, *codec.LegacyAmino) ([]byte, error)
}

// NewMnemonicKeyManager creates a new KeyManager from a mnenomic
func NewMnemonicKeyManager(mnemonic string, coinID uint32) (KeyManager, error) {
	fullBIP44Prefix := fmt.Sprintf("%s'/%d'/", PartialBIP44Prefix, coinID)
	fullPath := fullBIP44Prefix + PartialPath

	k := keyManager{}
	err := k.recoveryFromMnemonic(mnemonic, fullPath)
	return &k, err
}

type keyManager struct {
	keybase keyring.Keyring
	keyring keyring.Info
}

func (m *keyManager) GetKeyBase() keyring.Keyring {
	return m.keybase
}

func (m *keyManager) GetKeyRing() keyring.Info {
	return m.keyring
}

// Sign signs a standard msg and marshals the result to bytes
func (m *keyManager) Sign(stdMsg legacytx.StdSignMsg, cdc *codec.LegacyAmino) ([]byte, error) {
	sig, err := m.makeSignature(stdMsg)
	if err != nil {
		return nil, err
	}

	newTx := legacytx.NewStdTx(stdMsg.Msgs, stdMsg.Fee, []legacytx.StdSignature{sig}, stdMsg.Memo)

	bz, err := cdc.MarshalLengthPrefixed(&newTx)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func (m *keyManager) makeSignature(msg legacytx.StdSignMsg) (sig legacytx.StdSignature, err error) {
	if err != nil {
		return
	}

	sigBytes, pubKey, err := m.keybase.Sign("kava-go-sdk", msg.Bytes())
	if err != nil {
		return
	}

	return legacytx.StdSignature{
		PubKey:    pubKey,
		Signature: sigBytes,
	}, nil
}

func (m *keyManager) recoveryFromMnemonic(mnemonic, keyPath string) error {

	words := strings.Split(mnemonic, " ")
	if len(words) != 12 && len(words) != 24 {
		return fmt.Errorf("mnemonic length should either be 12 or 24")
	}

	kb := keyring.NewInMemory()
	signAlgo := hd.Secp256k1

	krInfo, err := kb.NewAccount("kava-go-sdk", mnemonic, defaultBIP39Passphrase, keyPath, signAlgo)
	if err != nil {
		return err
	}

	m.keybase = kb
	m.keyring = krInfo
	return nil
}
