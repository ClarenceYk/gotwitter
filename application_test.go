package gotwitter

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UserTimelineParam{
		UserID: 933516360365121536,
		// ScreenName: "MaYukkee",
		Count: 10,
	}

	ts, err := app.UserTimeline(param)
	if err != nil {
		t.Log(err)
	} else {
		for _, tweet := range ts {
			t.Log(tweet.Text)
			if tweet.QuotedStatus != nil {
				t.Log(tweet.QuotedStatus.Text)
			}
			if tweet.RetweetedStatus != nil {
				t.Log(tweet.RetweetedStatus.Text)
			}
		}
	}

	// param := &FollowerIDsParam{
	// 	ScreenName:   "MaYukkee",
	// 	Count:        3,
	// 	StringifyIDs: true,
	// }

	// ids, err := app.FollowerIDs(param)
	// if err != nil {
	// 	t.Error(err)
	// } else {
	// 	t.Log(*ids)
	// }
}
