# pandabit-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建PB.ini文件，编辑如下内容：

```ini
# transaction type
txType = "auth/StdTx"
# message type
msgSend = "cosmos-sdk/MsgSend"
msgVote = "cosmos-sdk/MsgVote"
msgDelegate = "cosmos-sdk/MsgDelegate"
# message choose 1-send  2-vote  3-delegate
msgType = 1

#rest api url
restAPI = "https://pandagram.caisinfo.co.kr"
# chain id
chainID = "pandagram"

# scan mempool or not
isScanMemPool = false
# minimum fee to pay
minFee = 50000
# standed gas
stdGas = 200000
# fee denom
feeDenom = "upanda
# Cache data file directory, default = "", current directory: ./data
dataDir = ""
```
