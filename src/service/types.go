package service

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
)

type JsonAccount struct {
	Address string   `json:"address"`
	Balance *big.Int `json:"balance"`
	Nonce   uint64   `json:"nonce"`
}

type JsonAccountList struct {
	Accounts []JsonAccount `json:"accounts"`
}

// SendTxArgs represents the arguments to submit a new transaction into the transaction pool.
type SendTxArgs struct {
	From     common.Address  `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *big.Int        `json:"gas"`
	GasPrice *big.Int        `json:"gasPrice"`
	Value    *big.Int        `json:"value"`
	Data     string          `json:"data"`
	Nonce    *uint64         `json:"nonce"`
}

type JsonCallRes struct {
	Data string `json:"data"`
}

type JsonTxRes struct {
	TxHash string `json:"txHash"`
}

type JsonReceipt struct {
	Root              common.Hash     `json:"root"`
	TransactionHash   common.Hash     `json:"transactionHash"`
	From              common.Address  `json:"from"`
	To                *common.Address `json:"to"`
	Value             *big.Int        `json:"value"`
	Gas               *big.Int        `json:"gas"`
	GasUsed           *big.Int        `json:"gasUsed"`
	GasPrice          *big.Int        `json:"gasPrice"`
	CumulativeGasUsed *big.Int        `json:"cumulativeGasUsed"`
	ContractAddress   common.Address  `json:"contractAddress"`
	Logs              []*ethTypes.Log `json:"logs"`
	LogsBloom         ethTypes.Bloom  `json:"logsBloom"`
	Error             string          `json:"error"`
	Failed            bool            `json:"failed"`
	Status            uint64          `json:"status"`
}

type JsonBlock struct {
	Hash         string        `json:"hash"`
	Index        int64         `json:"index"`
	Round        int64         `json:"round"`
	StateHash    string        `json:"stateHash"`
	FrameHash    string        `json:"frameHash"`
	Transactions []JsonReceipt `json:"transactions"`
}
