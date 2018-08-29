package types

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
