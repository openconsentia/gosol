// This example demonstrates use of Geth generated tron token struct
// to obtain information about transactions from the main net.
// Please refer to:
//    https://etherscan.io/token/0xf230b790e05390fc8295f4d3f60332c93bed42e2
//

package main

import (
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/openconsentia/gosol/pkg/trontoken"
)

const (
	infuraMainNetURL = "https://mainnet.infura.io/v3"
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

	// Instantiate a handler to a contract based on tron struct
	contract, err := trontoken.NewTronToken(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	// Invoke a method from tron struct
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

	conn, err := getConnection(infuraMainNetURL, projectID)
	if err != nil {
		log.Fatalf("Failed to connect. Reason: %v", err)
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
