package gotwitter

import (
	"testing"
)

func TestUserTimeline(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UserTimelineParam{
		UserID: 750713,
		// ScreenName: "MaYukkee",
		Count: 1,
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

func TestFollowerIDs(t *testing.T) {
	app, _ := NewApplication(2)

	param := &FollowerIDsParam{
		ScreenName:   "MaYukkee",
		Count:        3,
		StringifyIDs: true,
	}

	ids, err := app.FollowerIDs(param)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(*ids)
	}
}

func TestFollowersList(t *testing.T) {
	app, _ := NewApplication(2)

	param := &FollowersListParam{
		ScreenName:          "MaYukkee",
		Count:               3,
		IncludeUserEntities: true,
	}

	list, err := app.FollowersList(param)
	if err != nil {
		t.Error(err)
	} else {
		for _, user := range list.Users {
			t.Log(*user)
			t.Log(user.IDInt)
			t.Log(user.IDStr)
		}
	}
}

func TestFriendIDs(t *testing.T) {
	app, _ := NewApplication(0)

	param := &FriendIDsParam{
		ScreenName: "MaYukkee",
		Count:      3,
	}

	ids, err := app.FriendIDs(param)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(*ids)
	}
}

func TestFriendsList(t *testing.T) {
	app, _ := NewApplication(2)

	param := &FriendsListParam{
		ScreenName:          "MaYukkee",
		Count:               2,
		IncludeUserEntities: true,
	}

	fs, err := app.FriendsList(param)
	if err != nil {
		t.Error(err)
	} else {
		for _, user := range fs.Users {
			t.Log(user.ScreenName, user.Name)
		}
	}
}

func TestShowFriendship(t *testing.T) {
	app, _ := NewApplication(2)

	param := &ShowFriendshipParam{
		SourceScreenName: "MaYukkee",
		TargetID:         750713,
	}

	fs, err := app.ShowFriendship(param)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("source: ", fs.Relationship.Source.Following, fs.Relationship.Source.FollowedBy)
		t.Log("target: ", fs.Relationship.Target.Following, fs.Relationship.Target.FollowedBy)
	}
}
