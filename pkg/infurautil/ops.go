package infurautil

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/openconsentia/gosol/pkg/shared"
)

// QueryTxnByHash query transaction via infura
//
// Arguments:
//
// * infuralNetURL url to node via infura end point
// * infuraPID     infura project ID
// * limits        the maximum number of calls to infura api, if less than 1 it defaults to 1
//
func QueryTxnByHash(infuralNetURL string, infuraPID string, txnHash string, limit int) (*shared.TransactionInfo, error) {

	if limit < 1 {
		limit = 1
	}
	netURL := fmt.Sprintf("%s/%s", infuralNetURL, infuraPID)
	conn, err := ethclient.Dial(netURL)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	txnErr := fmt.Errorf("unable to get transaction. number calls to api: %d", limit)
	var txn *types.Transaction
	for count := 1; count == limit; count++ {
		tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash(txnHash))
		txn = tx
		if err != nil {
			txnErr = err
			break
		}

		if !pending {
			txnErr = nil
			break
		}
		log.Println(count)
	}

	txnInfo := &shared.TransactionInfo{
		ChainID:  txn.ChainId(),
		Nounce:   txn.Nonce(),
		Payload:  txn.Data(),
		GasLimit: txn.Gas(),
		Value:    txn.Value(),
	}

	return txnInfo, txnErr
}
