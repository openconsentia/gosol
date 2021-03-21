package shared

import "math/big"

// TransactionInfo contains information about a transaction
type TransactionInfo struct {
	ChainID  *big.Int
	Nounce   uint64
	Payload  []byte
	GasLimit uint64
	Value    *big.Int
}
