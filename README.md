# go-uniswap
Auto deploy Uniswap with golang

* Deploy Erc20

* Deploy Uniswap 

* Backend 
backends.SimulateDebug设置true，可以查看合约返回值和生成`trace log`

### 生成合约Go文件
```
abigen --bin=bin --abi=abi --pkg=token --out=Token.go
```
[详情参考](https://blog.csdn.net/JIYILANZHOU/article/details/104000285)

### Erc20
主要测试`Approve`和`TransferFrom`
### Uniswap
* 先部署`Weth`合约 在包`TokenE`
* 部署`CDC`合约 在包`Tokenc`
*  部署`工厂`合约 在包`TokenF`
*  部署`V2Router`合约

调用`V2Router`中`AddLiquidityETH`需要先`Approve`代币合约，然后调用可能成功。