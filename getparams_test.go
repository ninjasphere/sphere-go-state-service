package main

//

import (
	"testing"
)

func TestUserIDRegex(t *testing.T) {
	topic := "5063777c-d609-4852-a604-c492e2e70248.$cloud.device.e43820b2f3.channel.1-6-in.event.state"

	params := getParams(topic)

	if params["user_id"] != "5063777c-d609-4852-a604-c492e2e70248" {
		t.Errorf("bad userID %v", params)
	}
}
