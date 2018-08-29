package push

import (
	"os"

	"github.com/NaySoftware/go-fcm"
	"github.com/goldennetwork/golden-eth-notifications/types"
)

var gcm *fcm.FcmClient

func fcmClient() *fcm.FcmClient {

	if gcm == nil {
		gcm = fcm.NewFcmClient(os.Getenv("FCM_KEY_PUSH"))
	}
	return gcm
}

func SendMessage(message *types.PushMessage) error {
	notif := fcm.NotificationPayload{
		Title: message.Title,
		Icon:  "ic_launcher",
		Sound: message.Sound,
		Body:  message.Content,
		Badge: message.Badge,
	}

	gcm := fcmClient()

	var responses []types.PushResult
	for _, t := range message.DeviceTokens {
		pushRS := types.PushResult{
			DeviceToken: t,
			Result:      make(map[string]string),
		}
		responses = append(responses, pushRS)
	}
	message.Reponses = responses

	gcm.NewFcmRegIdsMsg(message.DeviceTokens, message.Payload)
	gcm.SetNotificationPayload(&notif)
	gcm.SetCollapseKey(message.Title)
	gcm.SetContentAvailable(true)

	status, err := gcm.Send()
	if err != nil {
		return err
	}

	for i, _ := range status.Results {
		message.Reponses[i].Result = status.Results[i]
	}
	return nil
}
