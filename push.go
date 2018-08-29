package ethNotification

import (
	"github.com/NaySoftware/go-fcm"
)

var gcm *fcm.FcmClient

func fcmClient(key string) *fcm.FcmClient {

	if gcm == nil {
		gcm = fcm.NewFcmClient(key)
	}
	return gcm
}

func sendMessage(pushKey string, message *PushMessage) error {
	notif := fcm.NotificationPayload{
		Title: message.Title,
		Icon:  "ic_launcher",
		Sound: message.Sound,
		Body:  message.Content,
		Badge: message.Badge,
	}

	gcm := fcmClient(pushKey)

	var responses []PushResult
	for _, t := range message.DeviceTokens {
		pushRS := PushResult{
			DeviceToken: t,
			Result:      make(map[string]string),
		}
		responses = append(responses, pushRS)
	}
	message.Responses = responses

	gcm.NewFcmRegIdsMsg(message.DeviceTokens, message.Payload)
	gcm.SetNotificationPayload(&notif)
	gcm.SetCollapseKey(message.Title)
	gcm.SetContentAvailable(true)

	status, err := gcm.Send()
	if err != nil {
		return err
	}

	for i, _ := range status.Results {
		message.Responses[i].Result = status.Results[i]
	}
	return nil
}
