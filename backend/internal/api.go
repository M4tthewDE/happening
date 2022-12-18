package internal

import "encoding/json"

func HandleNewSubscription(body string) (response string, status int) {
	var newSubscriptionBody NewSubscriptionBody
	err := json.Unmarshal([]byte(body), &newSubscriptionBody)
	if err != nil {
		return "", 400
	}

	resp, err := json.Marshal(newSubscriptionBody)
	if err != nil {
		return "", 400
	}

	return string(resp), 200
}

type NewSubscriptionBody struct {
	TargetUserID string `json:"target_id"`
	SubType      string `json:"subscription_type"`
}
