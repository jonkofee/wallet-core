package core

// #cgo CFLAGS: -I../../../include
// #cgo LDFLAGS: -L../../../build -L../../../build/trezor-crypto -lTrustWalletCore -lprotobuf -lTrezorCrypto -lstdc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
import "C"

import (
	"errors"
	"tw/types"
)

type Wallet struct {
	Address string
	PriKey  string
	PubKey  string
	Mnemonic string
	CoinType
}

func CreateWalletWithMnemonic(mn string, ct CoinType) (*Wallet, error) {
	if !IsMnemonicValid(mn) {
		return nil, errors.New("mnemonic is not valid")
	}

	str := types.TWStringCreateWithGoString(mn)
	empty := types.TWStringCreateWithGoString("")
	defer C.TWStringDelete(str)
	defer C.TWStringDelete(empty)

	tw := C.TWHDWalletCreateWithMnemonic(str, empty)
	defer C.TWHDWalletDelete(tw)

	priKey := C.TWHDWalletGetKeyForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWPrivateKeyDelete(priKey)
	priKeyData := C.TWPrivateKeyData(priKey)
	defer C.TWDataDelete(priKeyData)

	pubKey := C.TWPrivateKeyGetPublicKeySecp256k1(priKey, true)
	defer C.TWPublicKeyDelete(pubKey)
	pubKeyData := C.TWPublicKeyData(pubKey)
	defer C.TWDataDelete(pubKeyData)

	address := C.TWHDWalletGetAddressForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWStringDelete(address)

	mnemonic := C.TWHDWalletMnemonic(tw)
	defer C.TWStringDelete(mnemonic)

	return &Wallet{
		Address:  types.TWStringGoString(address),
		PriKey:   types.TWDataHexString(priKeyData),
		PubKey:   types.TWDataHexString(pubKeyData),
		Mnemonic: types.TWStringGoString(mnemonic),
		CoinType: ct,
	}, nil
}

func GenerateNewWallet(ct CoinType) (*Wallet, error) {
	tw := C.TWHDWalletCreate(256, types.TWStringCreateWithGoString(""))
	defer C.TWHDWalletDelete(tw)

	priKey := C.TWHDWalletGetKeyForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWPrivateKeyDelete(priKey)
	priKeyData := C.TWPrivateKeyData(priKey)
	defer C.TWDataDelete(priKeyData)

	pubKey := C.TWPrivateKeyGetPublicKeySecp256k1(priKey, true)
	defer C.TWPublicKeyDelete(pubKey)
	pubKeyData := C.TWPublicKeyData(pubKey)
	defer C.TWDataDelete(pubKeyData)

	address := C.TWHDWalletGetAddressForCoin(tw, C.enum_TWCoinType(ct))
	defer C.TWStringDelete(address)

	mnemonic := C.TWHDWalletMnemonic(tw)
	defer C.TWStringDelete(mnemonic)

	return &Wallet{
		Address:  types.TWStringGoString(address),
		PriKey:   types.TWDataHexString(priKeyData),
		PubKey:   types.TWDataHexString(pubKeyData),
		Mnemonic: types.TWStringGoString(mnemonic),
		CoinType: ct,
	}, nil
}