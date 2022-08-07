package etherscanapi

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"nft/db"
	"nft/graph/model"
	"strconv"
)

var (
	// apiKey is the api key of  eitherscan.io
	apiKey = "H6ZUH18UBVQBCA6IQEZM7U5IWUIBWUITMF"

	// contractAddress is the given contract address of the NFT
	contractAddress = "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d"

	// transactionUrl is the url of eitherscan.io for "normal" transaction
	transactionUrl = `https://api.etherscan.io/api?module=account&action=txlist&address=%s&startblock=%d&endblock=99999999&sort=desc&apikey=%s`

	// trxInternalUrl is the url of eitherscan.io for "internal" transaction
	trxInternalUrl = `https://api.etherscan.io/api?module=account&action=txlistinternal&address=%s&startblock=%d&endblock=99999999&sort=desc&apikey=%s`

	// trxERC20Url is the url of eitherscan.io for "ERC20" transaction
	trxERC20Url = `https://api.etherscan.io/api?module=account&action=tokentx&contractaddress=%s&startblock=%d&endblock=99999999&sort=desc&apikey=%s`

	// trxERC721Url is the url of eitherscan.io for "ERC721" transaction
	trxERC721Url = `https://api.etherscan.io/api?module=account&action=tokennfttx&contractaddress=%s&startblock=%d&endblock=99999999&sort=desc&apikey=%s`

	// trxERC1155Url is the url of eitherscan.io for "ERC1155" transaction
	trxERC1155Url = `https://api.etherscan.io/api?module=account&action=token1155tx&contractaddress=%s&startblock=%d&endblock=99999999&sort=desc&apikey=%s`

	// currentBlockUrl is the url of etherscan.io for current block number
	currentBlockUrl = `https://api.etherscan.io/api?module=proxy&action=eth_blockNumber&apikey=%s`
)

// GetEvents fetches the all 5 type of transactions from eitherscan.io
func GetEvents(startBlock uint64) model.Event {
	// get current block number
	resp, err := http.Get(fmt.Sprintf(currentBlockUrl, apiKey))
	if err != nil {
		log.Fatal("Error in getting current block number", err)
	}
	defer resp.Body.Close()

	data := struct {
		JsonRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  string `json:"result"` // block number
	}{}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatal("Error in decoding current block number", err)
	}

	// get the block number
	block := new(big.Int)
	block.SetString(data.Result, 0)
	currentBlockNumber := block.Uint64()

	// get the all 5 type of transactions from etherscan.io by last block number
	// fetch normal transaction
	resp, err = http.Get(fmt.Sprintf(transactionUrl, contractAddress, startBlock, apiKey))
	if err != nil {
		log.Fatal(err)
	}

	r1 := struct {
		Status  string               `json:"status"`
		Message string               `json:"message"`
		Result  []*model.Transaction `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&r1)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// fetch internal transaction
	resp, err = http.Get(fmt.Sprintf(trxInternalUrl, contractAddress, startBlock, apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	r2 := struct {
		Status  string               `json:"status"`
		Message string               `json:"message"`
		Result  []*model.TrxInternal `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&r2)
	if err != nil {
		log.Fatal(err)
	}

	// fetch ERC20 transaction
	resp, err = http.Get(fmt.Sprintf(trxERC20Url, contractAddress, startBlock, apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	r3 := struct {
		Status  string            `json:"status"`
		Message string            `json:"message"`
		Result  []*model.TrxErc20 `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&r3)
	if err != nil {
		log.Fatal(err)
	}

	// fetch ERC721 transaction
	resp, err = http.Get(fmt.Sprintf(trxERC721Url, contractAddress, startBlock, apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	r4 := struct {
		Status  string             `json:"status"`
		Message string             `json:"message"`
		Result  []*model.TrxErc721 `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&r4)
	if err != nil {
		log.Fatal(err)
	}

	// fetch ERC1155 transaction
	resp, err = http.Get(fmt.Sprintf(trxERC1155Url, contractAddress, startBlock, apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	r5 := struct {
		Status  string              `json:"status"`
		Message string              `json:"message"`
		Result  []*model.TrxErc1155 `json:"result"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&r5)
	if err != nil {
		log.Fatal(err)
	}

	strBlockNumber := strconv.FormatUint(currentBlockNumber, 10)

	// persist the all 5 type of transactions to database
	db.PushTransactionInfo(r1.Result, r2.Result, r3.Result, r4.Result, r5.Result, currentBlockNumber)

	return model.Event{
		CurrentBlock: &strBlockNumber,
		Transactions: r1.Result,
		TrxInternal:  r2.Result,
		TrxErc20:     r3.Result,
		TrxErc721:    r4.Result,
		TrxErc1155:   r5.Result,
	}

}

// func cronFunction() {
// 	tokenInfo, err := db.GetTokenInfo()
// 	if err != nil {
// 		log.Fatal("Token info is not found", err)
// 	}

// 	// get the all 5 type of transactions from etherscan.io by last block number
// 	event := fetchTransactions(tokenInfo.LastReadBlock)

// 	db.PushTransactionInfo(r1, r2, r3, r4, r5, blockNumber)

// }
