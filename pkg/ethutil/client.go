package ethutil

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient is a client to connect to the ethereum network
type EthClient interface {
	ConnectTo(url string) error
	QueryTxnByHash(ctx context.Context, txnHash string, limit int) (TransactionInfo, error)
}

type ethClient struct {
	conn *ethclient.Client
}

func (e *ethClient) ConnectTo(url string) error {
	conn, err := ethclient.Dial(url)
	if err != nil {
		return err
	}
	e.conn = conn
	return nil
}

func (e *ethClient) QueryTxnByHash(ctx context.Context, txnHash string, limit int) (TransactionInfo, error) {
	txnErr := fmt.Errorf("unable to get transaction. number calls to api: %d", limit)
	var txn *types.Transaction
	for count := 1; count == limit; count++ {
		tx, pending, err := e.conn.TransactionByHash(ctx, common.HexToHash(txnHash))
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

	txn.ChainId()
	txn.Cost()
	txn.Gas()
	txn.GasPrice()

	return TransactionInfo{}, txnErr
}

func NewClient(url string) (EthClient, error) {
	client := &ethClient{}
	err := client.ConnectTo(url)
	if err != nil {
		return nil, err
	}

	return client, nil
}
