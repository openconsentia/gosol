package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/openconsentia/gosol/pkg/infurautil"
	"github.com/openconsentia/gosol/pkg/tokens/trontoken"
)

func queryTxnByHash(txnHash string) {
	projectID, ok := os.LookupEnv("PID")
	if !ok {
		log.Fatal("Project ID enviroment variable not set")
	}

	mainNetURL := infurautil.MainNetURL
	result, err := infurautil.QueryTxnByHash(mainNetURL, projectID, txnHash, 1)
	if err != nil {
		log.Fatalf("Query transaction failed. Reason: %s", err.Error())
	}

	fmt.Println("ChainID: ", result.ChainID)
	fmt.Println("Nounce:  ", result.Nounce)
	fmt.Println("Gas limit", result.GasLimit)
	fmt.Println("Value", result.Value)
	fmt.Println("Payload", string(result.Payload))
}

func interactWithTronContract(infuralNetURL string, infuraPID string) {
	netURL := fmt.Sprintf("%s/%s", infuralNetURL, infuraPID)
	conn, err := ethclient.Dial(netURL)
	if err != nil {
		log.Fatalf("Failed to connect to the network: %s", err.Error())
	}

	contractAddr, ok := os.LookupEnv("CONTRACT_ADDR")
	if !ok {
		log.Fatal("Contract address not found from envar")
	}

	contract, err := trontoken.NewTronToken(common.HexToAddress(contractAddr), conn)
	if err != nil {
		log.Fatalf("Unable to obtain contract: %s", err.Error())
	}

	walletAddr, ok := os.LookupEnv("WALLET_ADDR")
	if !ok {
		log.Fatal("Wallet address not set")
	}

	amt, err := contract.BalanceOf(&bind.CallOpts{}, common.HexToAddress(walletAddr))
	if err != nil {
		log.Fatalf("Unable to get balance of contract for wallet %s: %s", walletAddr, err.Error())
	}

	fmt.Printf("Balance for wallet %s is %v", walletAddr, amt)
}

func main() {

	txnHash, ok := os.LookupEnv("TXN_HASH")
	if !ok {
		log.Fatalf("Missing transaction hash envvar")
	}
	queryTxnByHash(txnHash)

	projectID, ok := os.LookupEnv("PID")
	if !ok {
		log.Fatal("Project ID enviroment variable not set")
	}
	interactWithTronContract(infurautil.MainNetURL, projectID)
}
