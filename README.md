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


# mainnet rest api url
mainnetRestAPI = "https://pandagram.caisinfo.co.kr"
# mainnet node api url
mainnetNodeAPI = "http://ip:port"
# chain id
mainnetChainID = "pandagram"
# mainnet denom
mainnetDenom = "upanda"

# testnet rest api url
testnetRestAPI = "http://ip:port"
# testnet node api url
testnetNodeAPI = "http://ip:port"
# chain id
testnetChainID = "gaia-13003"
# testnet denom
testnetDenom = "muon"

# Is network test?
isTestNet = false

# scan mempool or not
isScanMemPool = false

# pay fee or not
payFee = true
# minimum fee to pay in upanda(1 PB = 1000000upanda)
minFee = 5000
# standed gas
stdGas = 200000

# Cache data file directory, default = "", current directory: ./data
dataDir = ""
```