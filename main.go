package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/promethiumchain/go-promethium/ethclient"
)

var nodeIP = "95.179.178.196"
var nodePort = "9988"

func connect() (blockNumber, blockTime, diff, hashrate, nrOfTxs, info string) {
	client, err := ethclient.Dial("http://" + nodeIP + ":" + nodePort)
	if err != nil {
		log.Fatal(err)
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockByNumber(context.Background(), header.Number)
	blockNumber = "Current block number    : " + header.Number.String()
	blockTime = "Current block time      : " + strconv.Itoa(int(header.Time))
	diff = "Difficulty              : " + header.Difficulty.String()
	hsrate := header.Difficulty.Uint64() / uint64(14)
	hashrate = "Hashrate                : " + strconv.Itoa(int(hsrate))
	nrOfTxs = "Number of txs in block  : " + strconv.Itoa(block.Transactions().Len())
	info = "Refresh page for new block info ..."
	return blockNumber, blockTime, diff, hashrate, nrOfTxs, info
}

func main() {
	printStartMessage()
	r := mux.NewRouter()
	r.HandleFunc("/", printPage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func printStartMessage() {
	fmt.Println("<-------------------------------------------------------------------------------->")
	fmt.Println("<                    Welcome to Promethium chain monitor v0.1                    >")
	fmt.Println("<-------------------------------------------------------------------------------->")
	fmt.Println("<                       By HexAviator for Promethium 2019                        >")
	fmt.Println("<-------------------------------------------------------------------------------->")
	fmt.Println("< Connecting to node with IP : ", nodeIP, " and port : ", nodePort, "                >")
	fmt.Println("<-------------------------------------------------------------------------------->")
}

func printPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<-------------------------------------------------------------------------------->")
	fmt.Fprintln(w, "<                    Welcome to Promethium chain monitor v0.1                    >")
	fmt.Fprintln(w, "<-------------------------------------------------------------------------------->")
	fmt.Fprintln(w, "<                       By HexAviator for Promethium 2019                        >")
	fmt.Fprintln(w, "<-------------------------------------------------------------------------------->")
	fmt.Fprintln(w, "<         Connecting to node with IP : ", nodeIP, " and port : ", nodePort, "        >")
	fmt.Fprintln(w, "<-------------------------------------------------------------------------------->")
	bln, blt, diff, hsrate, nrOfTxs, info := connect()
	fmt.Fprintln(w, bln)
	fmt.Fprintln(w, blt)
	fmt.Fprintln(w, diff)
	fmt.Fprintln(w, hsrate)
	fmt.Fprintln(w, "Price                   : Not available yet")
	fmt.Fprintln(w, nrOfTxs)
	fmt.Fprintln(w, info)
}
