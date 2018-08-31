package ethNotification

type EngineConfig struct {
	WSURL          string
	FCM_PUSH_KEY   string
	FCM_PUSH_TITLE string
	CHAIN_NAME     string
}

type TokenContract struct {
	Name            string `json:"name" bson:"name"`
	Symbol          string `json:"symbol" bson:"symbol"`
	TotalSupply     string `json:"total_supply" bson:"total_supply"`
	Decimals        int8   `json:"decimals" bson:"decimals"`
	ContractAddress string `json:"contract_address" bson:"contract_address"`
	ContractCreator string `json:"contract_creator" bson:"contract_creator"`
}

type InputData struct {
	MethodID          string `json:"method_id" bson:"method_id"`
	ToAddress         string `json:"to_address" bson:"to_address"`
	Value             string `json:"value" bson:"value"`
	ValueWithDecimals string `json:"value_with_dcms" bson:"value_with_dcms"`
}

type PushMessage struct {
	Title        string                 `json:"title" bson:"title"`
	Sound        string                 `json:"sound" bson:"sound"`
	Content      string                 `json:"content" bson:"content"`
	Badge        string                 `json:"badge" bson:"badge"`
	DeviceTokens []string               `json:"device_tokens" bson:"device_tokens"`
	Payload      map[string]interface{} `json:"payload" bson:"payload"`
	Responses    []PushResult           `json:"results,omitempty" bson:"results,omitempty"`
}

type PushResult struct {
	DeviceToken string            `json:"device_token" bson:"device_token"`
	Result      map[string]string `json:"result,omitempty" bson:"result,omitempty"`
}

type Block struct {
	Hash         string        `json:"hash"`
	Number       string        `json:"number"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	BlockHash         string   `json:"blockHash" bson:"block_hash"`
	BlockNumber       string   `json:"blockNumber" bson:"block_number"`
	From              string   `json:"from" bson:"from"`
	To                string   `json:"to" bson:"to"`
	Value             string   `json:"value" bson:"value"`
	Gas               string   `json:"gas" bson:"gas"`
	GasPrice          string   `json:"gasPrice" bson:"gas_price"`
	CumulativeGasUsed string   `json:"cumulativeGasUsed" bson:"cumulativeGasUsed"`
	Hash              string   `json:"hash" bson:"hash"`
	Input             string   `json:"input" bson:"input"`
	Logs              []Log    `json:"logs" bson:"logs"`
	LogsBloom         string   `json:"logsBloom" bson:"logsBloom"`
	TransactionIndex  string   `json:"transactionIndex" bson:"tx_index"`
	TimeStamp         string   `json:"timeStamp" bson:"time_stamp"`
	ContractAddress   string   `json:"contractAddress" bson:"contract_address"`
	TokenName         string   `json:"tokenName" bson:"token_name"`
	TokenSymbol       string   `json:"tokenSymbol" bson:"token_symbol"`
	TokenDecimal      int      `json:"tokenDecimal" bson:"token_decimal"`
	Status            TxStatus `json:"-" bson:"status"`
	StatusReceipt     string   `json:"status"`
	IsSeft            bool     `json:"isSeft"`
	ChainName         string   `json:"chainName"`
	// Nonce            string `json:"nonce" bson:"nonce"`
	// GasUsed          string `json:"gasUsed" bson:"gas_used"`
	// V                string `json:"v" bson:"v"`
	// R                string `json:"r" bson:"r"`
	// S                string `json:"s" bson:"s"`
}

type Log struct {
	Address          string   `json:"address" bson:"address"`
	Topics           []string `json:"topics" bson:"topics"`
	Data             string   `json:"data" bson:"data"`
	BlockNumber      string   `json:"blockNumber" bson:"blockNumber"`
	TransactionHash  string   `json:"transactionHash" bson:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex" bson:"transactionIndex"`
	BlockHash        string   `json:"blockHash" bson:"blockHash"`
	LogIndex         string   `json:"logIndex" bson:"logIndex"`
	Removed          bool     `json:"removed" bson:"removed"`
}

func (t *Transaction) IsNormalTx() bool {
	switch {
	case t.Value == "0x0":
		return false
	default:
		return true
	}
}

type Wallet struct {
	ID      string `json:"id" bson:"id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Type    string `json:"type" bson:"type"`
}

type WalletPush struct {
	Wallet
	DeviceUDID  string `json:"device_udid" bson:"_id"`
	DeviceToken string `json:"device_token" bson:"device_token"`
}

type WalletSubscriber struct {
	WalletName    string
	WalletAddress string
	DeviceToken   string
}

type WalletSubscriberResult struct {
	Transaction *Transaction
	Subscribers []WalletSubscriber
}

type CacheData struct {
	Transaction       Transaction
	WalletSubscribers []WalletSubscriber
}
