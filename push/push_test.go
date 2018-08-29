package push

import (
	"os"
	"testing"

	"github.com/goldennetwork/golden-eth-notifications/types"
)

func Test_Push(t *testing.T) {
	os.Setenv("FCM_TITLE_PUSH", "APP_NAME")
	os.Setenv("FCM_KEY_PUSH", "FIREBASE_PUSH_KEY")

	message := types.PushMessage{
		Title:        "Golden",
		Sound:        "default",
		Content:      "Test push Rinkeby",
		Badge:        "1",
		DeviceTokens: []string{"DEVICE_TOKEN"},
	}

	SendMessage(&message)
}
