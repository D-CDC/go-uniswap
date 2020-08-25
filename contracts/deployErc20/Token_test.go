package token

import (
	"context"
	"ethereum/contract/contracts/backends"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlWarn, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)

	key2, _  = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	key3, _  = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	testAddr = crypto.PubkeyToAddress(key2.PublicKey)
	add3     = crypto.PubkeyToAddress(key3.PublicKey)
)

func TestErc20(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: big.NewInt(1000000000)},
		testAddr: {Balance: big.NewInt(10000000000)}},
		1000000)
	transactOpts := bind.NewKeyedTransactor(key)
	keyOpts := bind.NewKeyedTransactor(key2)

	// Deploy the ENS registry
	ensAddr, _, _, err := DeployToken(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	ens, err := NewToken(ensAddr, contractBackend)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

	contractBackend.Commit()

	// Set ourself as the owner of the name.
	name, err := ens.Name(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Token name:", name)

	// Set ourself as the owner of the name.
	symbol, err := ens.Symbol(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Token symbol:", symbol)

	totalSupply, err := ens.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("totalSupply ", totalSupply)

	balance, err := ens.BalanceOf(nil, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("addr balance BalanceOf", balance)

	tx, err := ens.Transfer(transactOpts, testAddr, big.NewInt(50000))
	if err != nil {
		log.Error("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	contractBackend.Commit()

	balance, err = ens.BalanceOf(nil, testAddr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("testAddr balance BalanceOf", balance)

	tx, err = ens.Approve(keyOpts, addr, big.NewInt(10000))
	if err != nil {
		log.Error("Failed to retrieve Approve ", "name: %v", err)
	}
	contractBackend.Commit()

	balance, err = ens.Allowance(nil, testAddr, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("balance Allowance", balance)

	tx, err = ens.TransferFrom(transactOpts, testAddr, add3, big.NewInt(5000))
	if err != nil {
		log.Error("Failed to request token transfer: %v", err)
	}
	fmt.Printf("Transfer pending: 0x%x\n", tx.Hash())
	contractBackend.Commit()

	balance, err = ens.Allowance(nil, testAddr, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("Allowance balance ", balance)

	balance, err = ens.BalanceOf(nil, testAddr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("balance BalanceOf", balance)

	balance, err = ens.BalanceOf(nil, add3)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("balance BalanceOf", balance)
	contractBackend.Commit()

	balance, _ = contractBackend.BalanceAt(context.Background(), testAddr, nil)
	fmt.Println("balance", balance)
	keyOpts.Value = new(big.Int).SetUint64(10000000)
	tx, err = ens.ApproveOne(keyOpts, addr)
	if err != nil {
		log.Error("Failed to retrieve ApproveOne ", "name: %v", err)
	}
	contractBackend.Commit()

	balance, _ = contractBackend.BalanceAt(context.Background(), testAddr, nil)
	fmt.Println("11balance", balance)
}
