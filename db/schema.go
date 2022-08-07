package db

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

type Token struct {
	ID              int64  `sql:"id,pk" json:"id"`
	Name            string `sql:"name" json:"name"`
	Description     string `sql:"description" json:"description"`
	ContractAddress string `pg:",unique" json:"contract_address"`
	LastReadBlock   uint64 `sql:"last_read_block" json:"last_read_block"`
	CreatedAt       int64  `sql:"created_at" json:"created_at"`
}

// Transaction is the model of the transaction table returned by parameter action=txlist in the API call
type Transaction struct {
	ID int64 `sql:"id,pk" json:"id"`

	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `pg:",unique" json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodID          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}

//TrxInternal is the model of the internal transaction table returned by parameter action=txlistinternal in the API call
type TrxInternal struct {
	ID              int64  `sql:"id,pk" json:"id"`
	BlockNumber     string `json:"blockNumber"`
	TimeStamp       string `json:"timeStamp"`
	Hash            string `pg:",unique" json:"hash"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	ContractAddress string `json:"contractAddress"`
	Input           string `json:"input"`
	Type            string `json:"type"`
	Gas             string `json:"gas"`
	GasUsed         string `json:"gasUsed"`
	TraceID         string `json:"traceId"`
	IsError         string `json:"isError"`
	ErrCode         string `json:"errCode"`
}

//TrxERC20 is the model of the ERC20 transaction table returned by parameter action=tokentx in the API call
type TrxERC20 struct {
	ID                int64  `sql:"id,pk" json:"id"`
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `pg:",unique" json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

// TrxERC721 is the model of the ERC721 transaction table returned by parameter action=tokennfttx in the API call
type TrxERC721 struct {
	ID                int64  `sql:"id,pk" json:"id"`
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `pg:",unique" json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}

// TrxERC1155 is the model of the ERC1155 transaction table returned by parameter action=token1155tx in the API call
type TrxERC1155 struct {
	ID                int64  `sql:"id,pk" json:"id"`
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `pg:",unique" json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	From              string `json:"from"`
	To                string `json:"to"`
	TokenID           string `json:"tokenID"`
	TokenValue        string `json:"tokenValue"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
}

// CreateSchema creates the schema in the database for the models and create token record
func CreateSchema() error {
	models := []interface{}{
		&Token{},
		&Transaction{},
		&TrxInternal{},
		&TrxERC20{},
		&TrxERC721{},
		&TrxERC1155{},
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}

	// Create token record
	_, err := db.Model(&Token{
		Name:            "BoredApeYachtClub",
		Description:     "BoredApeYachtClub token for the BoredApeYachtClub ICO",
		ContractAddress: "0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D",
		LastReadBlock:   0,
		CreatedAt:       time.Now().UnixNano(),
	}).OnConflict("DO NOTHING").Insert()

	if err != nil {
		return err
	}
	return nil
}
