package token

import (
	"context"
	"ethereum/contract/contracts/backends"
	"ethereum/contract/contracts/deployUniswap/cdc"
	factory "ethereum/contract/contracts/deployUniswap/factory"
	"ethereum/contract/contracts/deployUniswap/weth"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	fatallog "log"

	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}

func TestDeployUniswapLocal(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: new(big.Int).SetUint64(10000000000000000000)},
		testAddr: {Balance: big.NewInt(100000000000000)}},
		100000000)
	transactOpts := bind.NewKeyedTransactor(key)

	weth, _, _, err := weth.DeployTokene(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	fac, _, _, err := factory.DeployTokenf(transactOpts, contractBackend, addr)
	if err != nil {
		t.Fatalf("can't DeployContract Factory: %v", err)
	}
	cdcT, _, _, err := cdc.DeployTokenc(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	cdcTran, err := cdc.NewTokenc(cdcT, contractBackend)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}

	contractBackend.Commit()

	balance, err := cdcTran.BalanceOf(nil, addr)
	if err != nil {
		log.Error("Failed to retrieve token ", "name: %v", err)
	}
	fmt.Println("addr balance BalanceOf", balance)

	// Deploy the ENS registry
	unisWapAddr, _, _, err := DeployToken(transactOpts, contractBackend, fac, weth)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	_, err = cdcTran.Approve(transactOpts, unisWapAddr, new(big.Int).SetUint64(1000000000000000000))
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}
	contractBackend.Commit()

	ens, err := NewToken(unisWapAddr, contractBackend)
	if err != nil {
		t.Fatalf("can't NewContract: %v", err)
	}
	tik := new(big.Int).SetUint64(10000000000000000)
	tik1 := new(big.Int).SetUint64(1000000000000)

	balance, err = cdcTran.Allowance(nil, addr, unisWapAddr)
	name, err := cdcTran.Name(nil)
	fmt.Println("balance ", balance, "name", name, " unisWapAddr ", unisWapAddr.String())
	transactOpts.Value = new(big.Int).SetUint64(1000000000000000000)
	fmt.Println(cdcT.String(), " ", addr.String(), " fac ", fac.String(), " eth ", weth.String())
	fmt.Println(hexutil.Encode(tik.Bytes()), " ", hexutil.Encode(tik1.Bytes()))
	backends.SimulateDebug = false
	_, err = ens.AddLiquidityETH(transactOpts, cdcT, tik, tik, tik1, addr, new(big.Int).SetUint64(1699658290))
	if err != nil {
		t.Fatalf("can't NewContract AddLiquidityETH : %v", err)
	}
	contractBackend.Commit()
}

func TestNodeConnect(t *testing.T) {
	url := "http://157.245.118.249:8545"
	client, url := dialConn(url)
	printBaseInfo(client, url)
	PrintBalance(client, addr)

	var err error
	nonce, err = client.PendingNonceAt(context.Background(), addr)
	if err != nil {
		fmt.Println("PendingNonceAt", err)
	}
	fmt.Println("PendingNonceAt", nonce)

	nonce, err = client.NonceAt(context.Background(), addr, nil)
	if err != nil {
		fmt.Println("PendingNonceAt", err)
	}
	fmt.Println(nonce)
}

func TestDeployUniswapWithSimulate(t *testing.T) {
	url := "http://157.245.118.249:8545"

	client, url := dialConn(url)
	printBaseInfo(client, url)
	PrintBalance(client, addr)

	// easy access to transactions get input
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: new(big.Int).SetUint64(10000000000000000000)},
		testAddr: {Balance: big.NewInt(100000000000000)}},
		100000000)
	transactOpts := bind.NewKeyedTransactor(key)

	basecontract, result := sendBaseContract(transactOpts, contractBackend, client)
	if !result {
		fmt.Println("sendBaseContract failed")
		return
	}
	routercontract, result := sendRouterContract(transactOpts, contractBackend, client, basecontract)
	if !result {
		fmt.Println("sendBaseContract failed")
		return
	}

	tik := new(big.Int).SetUint64(10000000000000000)
	tik1 := new(big.Int).SetUint64(1000000000000)
	transactOpts.Value = new(big.Int).SetUint64(1000000000000000000)
	input := packInput(routerAbi, "addLiquidityETH", "addLiquidityETH", basecontract.mapTR, tik, tik, tik1, addr, new(big.Int).SetUint64(1699658290))
	aHash := sendRouterTransaction(client, addr, routercontract.rethR, transactOpts.Value, key, input)
	result, _ = getResult(client, aHash)
	fmt.Println("over result", result)
}

func TestDeployUniswapRPC(t *testing.T) {
	url := "http://157.245.118.249:8545"

	client, url := dialConn(url)
	printBaseInfo(client, url)
	PrintBalance(client, addr)

	transactOpts := bind.NewKeyedTransactor(key)

	_, wtx, _, err := weth.DeployTokene(transactOpts, client)
	_, ftx, _, err := factory.DeployTokenf(transactOpts, client, addr)
	_, mtx, _, err := cdc.DeployTokenc(transactOpts, client)

	result, wethR := getResult(client, wtx.Hash())
	result1, facR := getResult(client, ftx.Hash())
	result2, mapTR := getResult(client, mtx.Hash())

	if !result || !result1 || !result2 {
		fatallog.Fatal("sendBaseContract", err)
		return
	}

	_, routerTx, _, err := DeployToken(transactOpts, client, facR, wethR)
	result, routerAddr := getResult(client, routerTx.Hash())
	if !result {
		fatallog.Fatal("sendBaseContract routerTx", err)
		return
	}

	mapTran, err := cdc.NewTokenc(mapTR, client)
	atx, err := mapTran.Approve(transactOpts, routerAddr, new(big.Int).SetUint64(1000000000000000000))
	result, _ = getResult(client, atx.Hash())
	if !result {
		fatallog.Fatal("sendBaseContract atx", err)
		return
	}

	tik := new(big.Int).SetUint64(10000000000000000)
	tik1 := new(big.Int).SetUint64(1000000000000)
	transactOpts.Value = new(big.Int).SetUint64(1000000000000000000)
	RTran, err := NewToken(routerAddr, client)
	aHash, _ := RTran.AddLiquidityETH(transactOpts, mapTR, tik, tik, tik1, addr, new(big.Int).SetUint64(1699658290))
	result, _ = getResult(client, aHash.Hash())
	fmt.Println("over result", result)
}
