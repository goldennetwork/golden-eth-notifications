package types

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
