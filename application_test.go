package gotwitter

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, _ := NewApplication(0)
	err := app.Authorize()
	if err != nil {
		t.Log(err)
	}
	param := &UserTimelineParam{
		ScreenName: "MaYukkee",
		Count:      8,
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
}
