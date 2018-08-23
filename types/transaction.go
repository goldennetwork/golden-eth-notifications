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
	TransactionIndex string `json:"transactionIndex" bson:"tx_index"`
	TimeStamp        string `json:"timeStamp" bson:"time_stamp"`
	ContractAddress  string `json:"contractAddress" bson:"contract_address"`
	TokenName        string `json:"tokenName" bson:"token_name"`
	TokenSymbol      string `json:"tokenSymbol" bson:"token_symbol"`
	TokenDecimal     string `json:"tokenDecimal" bson:"token_decimal"`
	// Nonce            string `json:"nonce" bson:"nonce"`
	// GasUsed          string `json:"gasUsed" bson:"gas_used"`
	// V                string `json:"v" bson:"v"`
	// R                string `json:"r" bson:"r"`
	// S                string `json:"s" bson:"s"`
}
