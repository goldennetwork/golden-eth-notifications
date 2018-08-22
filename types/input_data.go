package types

type InputData struct {
	MethodID  string `json:"method_id" bson:"method_id"`
	ToAddress string `json:"to_address" bson:"to_address"`
	Value     string `json:"value" bson:"value"`
}
