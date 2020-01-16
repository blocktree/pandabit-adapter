package pandabit

import (
	"errors"
	"math/big"
	"strconv"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/shopspring/decimal"
)

type AddrBalance struct {
	Address    string
	Balance    *big.Int
	FeeBalance *big.Int
	index      int
}

func convertFlostStringToBigInt(amount string, decimals int) (*big.Int, error) {
	vDecimal, err := decimal.NewFromString(amount)
	if err != nil {
		log.Error("convert from string to decimal failed, err=", err)
		return nil, err
	}

	decimalInt := big.NewInt(1)
	for i := 0; i < decimals; i++ {
		decimalInt.Mul(decimalInt, big.NewInt(10))
	}
	d, _ := decimal.NewFromString(decimalInt.String())
	vDecimal = vDecimal.Mul(d)
	rst := new(big.Int)
	if _, valid := rst.SetString(vDecimal.String(), 10); !valid {
		log.Error("conver to big.int failed")
		return nil, errors.New("conver to big.int failed")
	}
	return rst, nil
}

func convertBigIntToFloatDecimal(amount string, decimals int) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(amount)
	if err != nil {
		log.Error("convert string to deciaml failed, err=", err)
		return d, err
	}

	decimalInt := big.NewInt(1)
	for i := 0; i < decimals; i++ {
		decimalInt.Mul(decimalInt, big.NewInt(10))
	}

	w, _ := decimal.NewFromString(decimalInt.String())
	d = d.Div(w)
	return d, nil
}

func convertIntStringToBigInt(amount string) (*big.Int, error) {
	vInt64, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		log.Error("convert from string to int failed, err=", err)
		return nil, err
	}

	return big.NewInt(vInt64), nil
}

type ContractDecoder struct {
	*openwallet.SmartContractDecoderBase
	wm *WalletManager
}

//NewContractDecoder 智能合约解析器
func NewContractDecoder(wm *WalletManager) *ContractDecoder {
	decoder := ContractDecoder{}
	decoder.wm = wm
	return &decoder
}

func (decoder *ContractDecoder) GetTokenBalanceByAddress(contract openwallet.SmartContract, address ...string) ([]*openwallet.TokenBalance, error) {

	var tokenBalanceList []*openwallet.TokenBalance

	for i := 0; i < len(address); i++ {

		balance, err := decoder.wm.RestClient.getBalance(address[i], contract.Address, "")
		if err != nil {
			return nil, err
		}

		amount, err := convertBigIntToFloatDecimal(balance.Balance.String(), int(contract.Decimals))
		if err != nil {
			return nil, err
		}

		tokenBalanceList = append(tokenBalanceList, &openwallet.TokenBalance{
			Contract: &contract,
			Balance:  &openwallet.Balance{
				Symbol:           contract.Symbol,
				Address:          address[i],
				ConfirmBalance:   amount.String(),
				UnconfirmBalance: "0",
				Balance:          amount.String(),
			},
		})
	}

	return tokenBalanceList, nil
}
