package types

type TokenContract struct {
	Name            string `json:"name" bson:"name"`
	Symbol          string `json:"symbol" bson:"symbol"`
	TotalSupply     string `json:"total_supply" bson:"total_supply"`
	Decimals        int8   `json:"decimals" bson:"decimals"`
	ContractAddress string `json:"contract_address" bson:"contract_address"`
	ContractCreator string `json:"contract_creator" bson:"contract_creator"`
}
