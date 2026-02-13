package decoding

import (
	"data-explorer/utils"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type DecodedTx struct {
	MethodName string                 `json:"method_name"`
	From       common.Address         `json:"from"`
	To         common.Address         `json:"to"`
	Params     map[string]interface{} `json:"params"`
	Value      *big.Int               `json:"value"`
}

func DecodeTransaction(tx *types.Transaction, contractABI *abi.ABI) (*DecodedTx, error) {
	var to common.Address
	txTo := tx.To()
	if txTo == nil || *txTo != utils.GetAddress() {
		return nil, nil
	}
	to = *txTo

	txData := tx.Data()
	if len(txData) < 4 {
		return nil, fmt.Errorf("no method selector")
	}

	if tx.ChainId() == nil {
		return nil, fmt.Errorf("transaction has no chain ID")
	}

	methodId := txData[:4]
	method, err := contractABI.MethodById(methodId)
	if err != nil {
		return nil, fmt.Errorf("unknown method")
	}

	args := make(map[string]interface{})
	err = method.Inputs.UnpackIntoMap(args, txData[4:])
	if err != nil {
		return nil, err
	}

	var from common.Address
	signer := types.LatestSignerForChainID(tx.ChainId())
	from, err = types.Sender(signer, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to recover sender: %w", err)
	}

	return &DecodedTx{
		MethodName: method.Name,
		From:       from,
		To:         to,
		Params:     args,
		Value:      tx.Value(),
	}, nil
}
