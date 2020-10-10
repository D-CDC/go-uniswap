package token

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"testing"
)

func TestTx(t *testing.T) {
	url := "http://138.201.255.176:8545"
	client, url := dialConn(url)
	printBaseInfo(client, url)
	addr := common.HexToAddress("0x854dc6f847472964fc7f2b3f7380d60f7284cae5")
	PrintBalance(client, addr)

	fmt.Println("addr ", addr.String())
	var err error
	nonce, err = client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		fmt.Println("PendingNonceAt", err)
	}
	fmt.Println("PendingNonceAt", nonce)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("PendingNonceAt", nonce, " chainID ", chainID)

	dataS := "2e1a7d4d0000000000000000000000000000000000000000000000000de0b6b3a7640000"
	data, err := hex.DecodeString(dataS)
	if err != nil {
		fmt.Println(err)
	}
	transaction := types.NewTransaction(nonce, common.HexToAddress("0x9a11fe0a8dc1d8fd7067eed03ff9415d6a3b75c2"), nil, 60000, new(big.Int).SetUint64(1000000000), data)

	msg := ethereum.CallMsg{From: addr, To: transaction.To(), GasPrice: transaction.GasPrice(), Value: transaction.Value(), Data: transaction.Data()}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		fmt.Println("Contract exec failed", err)
	}
	fmt.Println("gasLimit ", gasLimit)
}
