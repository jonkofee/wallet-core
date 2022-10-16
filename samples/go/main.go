package main

import (
	"encoding/hex"
	"fmt"
	"tw/core"
	"tw/protos/polkadot"
)

func main() {
	wallet, err := core.GenerateNewWallet(core.CoinTypeWestend)
	if err != nil {
		panic(err)
	}
	printWallet(wallet)

	txn := CreateSignedTx(wallet)
	fmt.Println("Signed tx:")
	fmt.Println("\t", txn)
}

func CreateSignedTx(ew *core.Wallet) string {
	priKeyByte, _ := hex.DecodeString(ew.PriKey)
	amountByte, _ := hex.DecodeString(`1`)

	input := polkadot.SigningInput{
		PrivateKey:         priKeyByte,
		MessageOneof: &polkadot.SigningInput_BalanceCall{BalanceCall: &polkadot.Balance{MessageOneof: &polkadot.Balance_Transfer_{Transfer: &polkadot.Balance_Transfer{
			ToAddress: "1gHcEZamRJim95LeCkG3YFWbft5aYaC3p1fKsjLcXtj3y7P",
			Value:     amountByte,
		}}}},
	}

	var output polkadot.SigningOutput
	err := core.CreateSignedTx(&input, ew.CoinType, &output)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(output.GetEncoded())
}

func printWallet(w *core.Wallet) {
	fmt.Printf("wallet: %s %s\n", w.CoinType.GetName(), w.CoinType.GetSymbol())
	fmt.Printf("\t address: %s \n", w.Address)
	fmt.Printf("\t pri key: %s \n", w.PriKey)
	fmt.Printf("\t pub key: %s \n", w.PubKey)
	fmt.Printf("\t mnemonic: %s \n", w.Mnemonic)
	fmt.Println("")
}
