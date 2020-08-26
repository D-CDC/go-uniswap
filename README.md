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
* 先部署`Weth`合约 在包`weth`
* 部署`CDC`合约 在包`cdc`
*  部署`工厂`合约 在包`factory`
*  部署`V2Router`合约

调用`V2Router`中`AddLiquidityETH`需要先`Approve`代币合约，然后调用可能成功。

### TestDeployUniswap
本地测试用例，无需申请测试币
### TestNode
测试节点是否可用，打印常规信息
### TestDialNode
部署合约到节点，查询合约执行状态，成功才进行下一步部署