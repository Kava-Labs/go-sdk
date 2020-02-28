package keys

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/go-bip39"

	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/kava-labs/kava/app"
)

const (
	defaultBIP39Passphrase = ""
	BIP44Prefix            = "44'/118'/"
	PartialPath            = "0'/0/0"
	FullPath               = BIP44Prefix + PartialPath
)

// KeyManager is an interface for common methods on KeyManagers
type KeyManager interface {
	GetPrivKey() crypto.PrivKey
	GetAddr() sdk.AccAddress
	GetCodec() *amino.Codec
	SetCodec(*amino.Codec)
	Sign(authtypes.StdSignMsg) ([]byte, error)
}

// NewMnemonicKeyManager creates a new KeyManager from a mnenomic
func NewMnemonicKeyManager(mnemonic string) (KeyManager, error) {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixes(config)
	config.Seal()

	k := keyManager{}
	err := k.recoveryFromMnemonic(mnemonic, FullPath)

	return &k, err
}

// NewPrivateKeyManager creates a new KeyManager from a private key
func NewPrivateKeyManager(priKey string) (KeyManager, error) {
	k := keyManager{}
	err := k.recoveryFromPrivateKey(priKey)
	return &k, err
}

type keyManager struct {
	cdc      *amino.Codec
	privKey  crypto.PrivKey
	addr     sdk.AccAddress
	mnemonic string
}

func (m *keyManager) GetPrivKey() crypto.PrivKey {
	return m.privKey
}

func (m *keyManager) GetAddr() sdk.AccAddress {
	return m.addr
}

func (m *keyManager) GetCodec() *amino.Codec {
	return m.cdc
}

func (m *keyManager) SetCodec(codec *amino.Codec) {
	m.cdc = codec
}

// Sign signs a standard msg and marshals the result to bytes
func (m *keyManager) Sign(stdMsg authtypes.StdSignMsg) ([]byte, error) {
	sig, err := m.makeSignature(stdMsg)
	if err != nil {
		return nil, err
	}

	newTx := authtypes.NewStdTx(stdMsg.Msgs, stdMsg.Fee, []authtypes.StdSignature{sig}, stdMsg.Memo)

	bz, err := m.cdc.MarshalBinaryLengthPrefixed(&newTx)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

func (m *keyManager) makeSignature(msg authtypes.StdSignMsg) (sig authtypes.StdSignature, err error) {
	if err != nil {
		return
	}

	sigBytes, err := m.privKey.Sign(msg.Bytes())
	if err != nil {
		return
	}

	return authtypes.StdSignature{
		PubKey:    m.privKey.PubKey(),
		Signature: sigBytes,
	}, nil
}

func (m *keyManager) recoveryFromMnemonic(mnemonic, keyPath string) error {

	words := strings.Split(mnemonic, " ")
	if len(words) != 12 && len(words) != 24 {
		return fmt.Errorf("mnemonic length should either be 12 or 24")
	}
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, defaultBIP39Passphrase)
	if err != nil {
		return err
	}
	// create master key and derive first key:
	masterPriv, ch := computeMastersFromSeed(seed)
	derivedPriv, err := derivePrivateKeyForPath(masterPriv, ch, keyPath)
	if err != nil {
		return err
	}
	priKey := secp256k1.PrivKeySecp256k1(derivedPriv)
	addr := sdk.AccAddress(priKey.PubKey().Address())
	if err != nil {
		return err
	}
	m.addr = addr
	m.privKey = priKey
	m.mnemonic = mnemonic
	return nil
}

// computeMastersFromSeed returns the master public key, master secret, and chain code in hex.
func computeMastersFromSeed(seed []byte) (secret [32]byte, chainCode [32]byte) {
	masterSecret := []byte("Bitcoin seed")
	secret, chainCode = i64(masterSecret, seed)

	return
}

// derivePrivateKeyForPath derives the private key by following the BIP 32/44 path from privKeyBytes,
// using the given chainCode.
func derivePrivateKeyForPath(privKeyBytes [32]byte, chainCode [32]byte, path string) ([32]byte, error) {
	data := privKeyBytes
	parts := strings.Split(path, "/")
	for _, part := range parts {
		// do we have an apostrophe?
		harden := part[len(part)-1:] == "'"
		// harden == private derivation, else public derivation:
		if harden {
			part = part[:len(part)-1]
		}
		idx, err := strconv.Atoi(part)
		if err != nil {
			return [32]byte{}, fmt.Errorf("invalid BIP 32 path: %s", err)
		}
		if idx < 0 {
			return [32]byte{}, errors.New("invalid BIP 32 path: index negative ot too large")
		}
		data, chainCode = derivePrivateKey(data, chainCode, uint32(idx), harden)
	}
	var derivedKey [32]byte
	n := copy(derivedKey[:], data[:])
	if n != 32 || len(data) != 32 {
		return [32]byte{}, fmt.Errorf("expected a (secp256k1) key of length 32, got length: %v", len(data))
	}

	return derivedKey, nil
}

// derivePrivateKey derives the private key with index and chainCode.
// If harden is true, the derivation is 'hardened'.
// It returns the new private key and new chain code.
// For more information on hardened keys see:
//  - https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki
func derivePrivateKey(privKeyBytes [32]byte, chainCode [32]byte, index uint32, harden bool) ([32]byte, [32]byte) {
	var data []byte
	if harden {
		index = index | 0x80000000
		data = append([]byte{byte(0)}, privKeyBytes[:]...)
	} else {
		// this can't return an error:
		_, ecPub := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes[:])
		pubkeyBytes := ecPub.SerializeCompressed()
		data = pubkeyBytes

		/* By using btcec, we can remove the dependency on tendermint/crypto/secp256k1
		pubkey := secp256k1.PrivKeySecp256k1(privKeyBytes).PubKey()
		public := pubkey.(secp256k1.PubKeySecp256k1)
		data = public[:]
		*/
	}
	data = append(data, uint32ToBytes(index)...)
	data2, chainCode2 := i64(chainCode[:], data)
	x := addScalars(privKeyBytes[:], data2[:])
	return x, chainCode2
}

func (m *keyManager) recoveryFromPrivateKey(privateKey string) error {
	priBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return err
	}

	if len(priBytes) != 32 {
		return fmt.Errorf("Len of Keybytes is not equal to 32 ")
	}
	var keyBytesArray [32]byte
	copy(keyBytesArray[:], priBytes[:32])
	priKey := secp256k1.PrivKeySecp256k1(keyBytesArray)
	addr := sdk.AccAddress(priKey.PubKey().Address())
	m.addr = addr
	m.privKey = priKey
	return nil
}

// modular big endian addition
func addScalars(a []byte, b []byte) [32]byte {
	aInt := new(big.Int).SetBytes(a)
	bInt := new(big.Int).SetBytes(b)
	sInt := new(big.Int).Add(aInt, bInt)
	x := sInt.Mod(sInt, btcec.S256().N).Bytes()
	x2 := [32]byte{}
	copy(x2[32-len(x):], x)
	return x2
}

func uint32ToBytes(i uint32) []byte {
	b := [4]byte{}
	binary.BigEndian.PutUint32(b[:], i)
	return b[:]
}

// i64 returns the two halfs of the SHA512 HMAC of key and data.
func i64(key []byte, data []byte) (IL [32]byte, IR [32]byte) {
	mac := hmac.New(sha512.New, key)
	// sha512 does not err
	_, _ = mac.Write(data)
	I := mac.Sum(nil)
	copy(IL[:], I[:32])
	copy(IR[:], I[32:])
	return
}
