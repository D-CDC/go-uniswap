package token

import (
	"encoding/hex"
	"ethereum/contract/contracts/backends"
	"ethereum/contract/contracts/deployUniswap/TokenE"
	tokenf "ethereum/contract/contracts/deployUniswap/TokenF"
	"ethereum/contract/contracts/deployUniswap/tokenc"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

func init() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(false))))
}

var (
	key, _ = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	addr   = crypto.PubkeyToAddress(key.PublicKey)

	key2, _  = crypto.HexToECDSA("8a1f9a8f95be41cd7ccb6168179afb4504aefe388d1e14474d32c45c72ce7b7a")
	key3, _  = crypto.HexToECDSA("49a7b37aa6f6645917e7b807e9d1c00d4fa71f18343b0d4122a4d2df64dd6fee")
	testAddr = crypto.PubkeyToAddress(key2.PublicKey)
	add3     = crypto.PubkeyToAddress(key3.PublicKey)
)

func TestDeployUniswap(t *testing.T) {
	contractBackend := backends.NewSimulatedBackend(core.GenesisAlloc{
		addr:     {Balance: new(big.Int).SetUint64(10000000000000000000)},
		testAddr: {Balance: big.NewInt(100000000000000)}},
		100000000)
	transactOpts := bind.NewKeyedTransactor(key)

	weth, _, _, err := TokenE.DeployTokene(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	fac, _, _, err := tokenf.DeployTokenf(transactOpts, contractBackend, addr)
	if err != nil {
		t.Fatalf("can't DeployContract Factory: %v", err)
	}
	cdcT, _, _, err := tokenc.DeployTokenc(transactOpts, contractBackend)
	if err != nil {
		t.Fatalf("can't DeployContract: %v", err)
	}
	cdcTran, err := tokenc.NewTokenc(cdcT, contractBackend)
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

func TestNewToken(t *testing.T) {
	sig := "_addLiquidity(address,address,uint,uint,uint,uint)"
	id := crypto.Keccak256([]byte(sig))[:4]
	fmt.Println(hexutil.Encode(id))

	sig = "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)"
	id = crypto.Keccak256([]byte(sig))[:4]
	fmt.Println(hexutil.Encode(id))
	sig = "pairFor(address,address,address)"
	id = crypto.Keccak256([]byte(sig))[:4]
	fmt.Println(hexutil.Encode(id))
}

func TestParseDepositInput(t *testing.T) {
	input := "fb3bdb41000000000000000000000000000000000000000000108b2a2c280290940000000000000000000000000000000000000000000000000000000000000000000080000000000000000000000000b634058bd3ac146b128824e4b9ab59561b0568b9000000000000000000000000000000000000000000000000000000005f3f6a000000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000ed1199093b1abd07a368dd1c0cdc77d8517ba2a0"
	inputData, _ := hex.DecodeString(input)
	abiStaking, _ := abi.JSON(strings.NewReader(Json))
	methodName, err := abiStaking.MethodById(inputData)
	data := inputData[4:]
	args := struct {
		AmountOut *big.Int
		Path      []common.Address
		To        common.Address
		Deadline  *big.Int
	}{}
	method, _ := abiStaking.Methods[methodName.Name]

	err = method.Inputs.Unpack(&args, data)
	if err != nil {
		log.Error("Unpack deposit pubkey error", "err", err)
	}
	fmt.Println("AmountOut ", hex.EncodeToString(args.AmountOut.Bytes()), " Deadline ", hex.EncodeToString(args.Deadline.Bytes()), " ", len(args.Path), " ", args.To.String())
}

var abiStaking, _ = abi.JSON(strings.NewReader(Json))

func TestUnpack(t *testing.T) {
	args := struct {
		Pubkey []byte
		Fee    *big.Int
		Value  *big.Int
	}{}
	// 5d322ae8
	inputstr := "5d322ae8000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000003e8000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000004050863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b23522cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6"
	input, e := hex.DecodeString(inputstr)
	if e != nil {
		fmt.Println(e)
	}
	method, ok := abiStaking.Methods["deposit"]
	if !ok {
		fmt.Println("cann't find")
	}
	input = input[4:]
	err := method.Inputs.Unpack(&args, input)
	if err != nil {
		fmt.Println("Unpack deposit pubkey error", err)
	}
	// vpk,e2 := hex.DecodeString("0450863ad64a87ae8a2fe83c1af1a8403cb53f53e486d8511dad8a04887e5b23522cd470243453a299fa9e77237716103abc11a1df38855ed6f2ee187e9c582ba6")
	// if e2 != nil {
	// 	fmt.Println("e2:",e2)
	// }
	if _, err := crypto.UnmarshalPubkey(args.Pubkey); err != nil {
		fmt.Println("invalid pk,err:", err)
	}
	fmt.Println("pk:", hex.EncodeToString(args.Pubkey))
	fmt.Println("fee", args.Fee.String())
	fmt.Println("Value", args.Value.String())
	fmt.Println("finish")
}

const Json = `
[
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "_factory",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "_WETH",
				"type": "address"
			}
		],
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"inputs": [],
		"name": "WETH",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "tokenA",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "tokenB",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amountADesired",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountBDesired",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountAMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountBMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "addLiquidity",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountA",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountB",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amountTokenDesired",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountTokenMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETHMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "addLiquidityETH",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountToken",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETH",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			}
		],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "factory",
		"outputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "reserveIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "reserveOut",
				"type": "uint256"
			}
		],
		"name": "getAmountIn",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "reserveIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "reserveOut",
				"type": "uint256"
			}
		],
		"name": "getAmountOut",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			}
		],
		"name": "getAmountsIn",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			}
		],
		"name": "getAmountsOut",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountA",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "reserveA",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "reserveB",
				"type": "uint256"
			}
		],
		"name": "quote",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountB",
				"type": "uint256"
			}
		],
		"stateMutability": "pure",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "tokenA",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "tokenB",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountAMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountBMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "removeLiquidity",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountA",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountB",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountTokenMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETHMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "removeLiquidityETH",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountToken",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETH",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountTokenMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETHMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "removeLiquidityETHSupportingFeeOnTransferTokens",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountETH",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountTokenMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETHMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "approveMax",
				"type": "bool"
			},
			{
				"internalType": "uint8",
				"name": "v",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "r",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "removeLiquidityETHWithPermit",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountToken",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETH",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "token",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountTokenMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountETHMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "approveMax",
				"type": "bool"
			},
			{
				"internalType": "uint8",
				"name": "v",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "r",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "removeLiquidityETHWithPermitSupportingFeeOnTransferTokens",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountETH",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "tokenA",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "tokenB",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "liquidity",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountAMin",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountBMin",
				"type": "uint256"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			},
			{
				"internalType": "bool",
				"name": "approveMax",
				"type": "bool"
			},
			{
				"internalType": "uint8",
				"name": "v",
				"type": "uint8"
			},
			{
				"internalType": "bytes32",
				"name": "r",
				"type": "bytes32"
			},
			{
				"internalType": "bytes32",
				"name": "s",
				"type": "bytes32"
			}
		],
		"name": "removeLiquidityWithPermit",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "amountA",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountB",
				"type": "uint256"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapETHForExactTokens",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOutMin",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapExactETHForTokens",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOutMin",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapExactETHForTokensSupportingFeeOnTransferTokens",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountOutMin",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapExactTokensForETH",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountOutMin",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapExactTokensForETHSupportingFeeOnTransferTokens",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountOutMin",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapExactTokensForTokens",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountIn",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountOutMin",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapExactTokensForTokensSupportingFeeOnTransferTokens",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountInMax",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapTokensForExactETH",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amountOut",
				"type": "uint256"
			},
			{
				"internalType": "uint256",
				"name": "amountInMax",
				"type": "uint256"
			},
			{
				"internalType": "address[]",
				"name": "path",
				"type": "address[]"
			},
			{
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "deadline",
				"type": "uint256"
			}
		],
		"name": "swapTokensForExactTokens",
		"outputs": [
			{
				"internalType": "uint256[]",
				"name": "amounts",
				"type": "uint256[]"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"stateMutability": "payable",
		"type": "receive"
	}
]
`
