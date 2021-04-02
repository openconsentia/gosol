// This example demonstrates use of Geth generated tron token struct
// to obtain information about transactions from the main net.
// Please refer to:
//    https://etherscan.io/token/0xf230b790e05390fc8295f4d3f60332c93bed42e2
//

package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/openconsentia/gosol/pkg/ethutil"
	"github.com/openconsentia/gosol/pkg/tokens/trontoken"
)

// infuraPID obtain from here https://infura.io/dashboard/ethereum
func getConnection(infuraMainNetURL string, infuraPID string) (*ethclient.Client, error) {

	netURL := fmt.Sprintf("%s/%s", infuraMainNetURL, infuraPID)
	conn, err := ethclient.Dial(netURL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getBalanceOf(conn *ethclient.Client, contractAddr string, walletAddr string) (*big.Int, error) {

	contract, err := trontoken.NewTronToken(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	amt, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(walletAddr))
	if err != nil {
		return nil, err
	}

	return amt, nil
}

func main() {

	projectID, ok := os.LookupEnv("PID")
	if !ok {
		log.Fatal("Project ID enviroment variable not set")
	}

	conn, err := getConnection(ethutil.InfuraMainNetURL, projectID)
	if err != nil {
		log.Fatalf("Failed to connect. Reason: %v", err)
	}

	ctx := context.Background()

	// Get transaction by Hash
	// View expected result here https://etherscan.io/tx/0x5b3f91849a33bb423be1e4f6c6d90be4522e9a934a1798d61fb745c1e001583d
	txnHash := "0x033abf3212c5cb7b2172d0e8d1425cc43fc08e10c2bea25916597d047a8ac28a"
	tx, pending, err := conn.TransactionByHash(ctx, common.HexToHash(txnHash))
	if err != nil {
		log.Fatalf("Unable to get transaction: %v", err)
	}
	fmt.Printf("Transaction is pending: %t\n", pending)
	if tx != nil {
		fmt.Printf("Chain ID: %v\n", tx.ChainId())
		fmt.Printf("Nounce: %v\n", tx.Nonce())
		fmt.Printf("Payload: %v\n", tx.Data())
		fmt.Printf("Gas limit: %v\n", tx.Gas())
		fmt.Printf("Gas limit: %v\n", tx.Value())
	} else {
		fmt.Println("No transaction info")
	}

	// View here https://etherscan.io/address/0xf230b790e05390fc8295f4d3f60332c93bed42e2
	contractAddr := "0xf230b790e05390fc8295f4d3f60332c93bed42e2"
	// View expected amount here https://etherscan.io/token/0xf230b790e05390fc8295f4d3f60332c93bed42e2?a=0x2c1393d26bace7e209468b0b24a67c7f3a0c36fb
	walletAddr := "0x2c1393d26bace7e209468b0b24a67c7f3a0c36fb"

	amt, err := getBalanceOf(conn, contractAddr, walletAddr)
	if err != nil {
		log.Fatalf("Failed to get amount. Reason: %v", err)
	}
	fmt.Printf("The amount of trons in your wallet %s is %v", walletAddr, amt)

}
