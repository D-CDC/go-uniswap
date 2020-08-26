package token

import (
	"context"
	"ethereum/contract/contracts/backends"
	"ethereum/contract/contracts/deployUniswap/cdc"
	factory "ethereum/contract/contracts/deployUniswap/factory"
	"ethereum/contract/contracts/deployUniswap/weth"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"

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

func TestDeployUniswap(t *testing.T) {
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

func TestNode(t *testing.T) {
	url := "http://165.227.99.131:8545"
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

func TestDialNode(t *testing.T) {
	url := "http://165.227.99.131:8545"

	client, url := dialConn(url)
	printBaseInfo(client, url)
	PrintBalance(client, addr)

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

	fmt.Println(" routerAddr ", routercontract.rethR.String())
	tik := new(big.Int).SetUint64(10000000000000000)
	tik1 := new(big.Int).SetUint64(1000000000000)
	balance, err := basecontract.mapTran.Allowance(nil, addr, routercontract.fAddr)
	name, err := basecontract.mapTran.Name(nil)
	fmt.Println("balance ", balance, "name", name, " ", routercontract.fAddr.String(), "err", err)
	transactOpts.Value = new(big.Int).SetUint64(1000000000000000000)
	fmt.Println(basecontract.mapT.String(), " ", addr.String())
	_, err = routercontract.RTran.AddLiquidityETH(transactOpts, basecontract.mapT, tik, tik, tik1, addr, new(big.Int).SetUint64(1699658290))
	fmt.Println("simulate result", err)
	contractBackend.Commit()

	input := packInput(routerAbi, "addLiquidityETH", "addLiquidityETH", basecontract.mapTR, tik, tik, tik1, addr, new(big.Int).SetUint64(1699658290))
	aHash := sendRouterTransaction(client, addr, routercontract.rethR, transactOpts.Value, key, input)
	result, _ = getResult(client, aHash)
	fmt.Println("over result", result)
}
