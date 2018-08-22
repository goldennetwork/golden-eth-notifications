package types

type Transactions []Transaction

type Transaction struct {
	BlockHash        string `json:"blockHash" bson:"block_hash"`
	BlockNumber      string `json:"blockNumber" bson:"block_number"`
	From             string `json:"from" bson:"from"`
	To               string `json:"to" bson:"to"`
	Value            string `json:"value" bson:"value"`
	Gas              string `json:"gas" bson:"gas"`
	GasPrice         string `json:"gasPrice" bson:"gas_price"`
	Hash             string `json:"hash" bson:"hash"`
	Input            string `json:"input" bson:"input"`
	Nonce            string `json:"nonce" bson:"nonce"`
	TransactionIndex string `json:"transactionIndex" bson:"tx_index"`
	V                string `json:"v" bson:"v"`
	R                string `json:"r" bson:"r"`
	S                string `json:"s" bson:"s"`
}
