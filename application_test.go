package gotwitter

import (
	"testing"
)

func TestUserTimeline(t *testing.T) {
	app, _ := NewApplication(0)

	param := &UserTimelineParam{
		UserID: 750713,
		// ScreenName: "MaYukkee",
		Count: 1,
	}

	ts, err := app.UserTimeline(param)
	if err != nil {
		t.Error(err)
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
	app, _ := NewApplication(0)

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
	app, _ := NewApplication(0)

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
	app, _ := NewApplication(0)

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
	app, _ := NewApplication(0)

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

func TestUsersLookup(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UsersLookupParam{
		UserID:          []int64{177547360, 248806976, 90183678},
		IncludeEntities: true,
	}

	users, err := app.UsersLookup(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, user := range users {
			t.Log(user.ScreenName, user.Status.CreatedAt, *user.Status.Entities)
		}
	}
}

func TestShowUser(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UserShowParam{
		UserID:          248806976,
		ScreenName:      "MortyCrypto",
		IncludeEntities: true,
	}

	user, err := app.ShowUser(param)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(user.ScreenName, user.Status.CreatedAt, *user.Status.Entities)
	}
}

func TestUserSuggestions(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UserSuggestionsParam{}

	sugs, err := app.UserSuggestions(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, sug := range sugs {
			t.Log(*sug)
		}
	}
}

func TestUserSuggestionsWithSlug(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UserSuggestionsParam{
		Slug: "news",
	}

	users, err := app.UserSuggestionsWithSlug(param)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(users.Name)
		for _, user := range users.Users {
			t.Log(user.ScreenName, user.CreatedAt, user.Entities.URL.Urls)
		}
	}
}

func TestUserSuggestionsWithSlugAndMembers(t *testing.T) {
	app, _ := NewApplication(2)

	param := &UserSuggestionsParam{
		Slug:    "funny",
		Members: true,
	}

	users, err := app.UserSuggestionsWithSlugAndMembers(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, user := range users {
			t.Log(user.ScreenName, user.IDStr, user.Status.Text)
		}
	}
}

func TestGetList(t *testing.T) {
	app, _ := NewApplication(2)

	param := &GetListsParam{
		UserID: 148781366,
	}

	lists, err := app.GetLists(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, list := range lists {
			t.Log(list.FullName, list.CreatedAt, list.MemberCount, list.SubscriberCount, *list.User)
		}
	}
}

func TestListsCreate(t *testing.T) {
	app, _ := NewApplication(2)

	param := &ListsCreateParam{
		Name: "TESTLIST",
	}

	list, err := app.ListsCreate(param)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(list.FullName, list.CreatedAt, list.MemberCount, list.SubscriberCount, *list.User)
	}
}

func TestListMembers(t *testing.T) {
	app, _ := NewApplication(2)

	param := &ListMembersParam{
		Slug:            "6",
		OwnerScreenName: "taofoufoufou",
	}

	list, err := app.ListMembers(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, user := range list.Users {
			t.Log(user.ScreenName, user.CreatedAt)
		}
	}
}

func TestListMemberships(t *testing.T) {
	app, _ := NewApplication(2)

	param := &ListsMembershipsParam{
		ScreenName: "MaYukkee",
	}

	list, err := app.ListsMemberships(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, l := range list.Lists {
			t.Log(l.Name, l.FullName, l.Description, l.Slug)
		}
	}
}

func TestListsOwnerships(t *testing.T) {
	app, _ := NewApplication(2)

	param := &ListsOwnershipsParam{
		ScreenName: "taofoufoufou",
	}

	list, err := app.ListsOwnerships(param)

	if err != nil {
		t.Error(err)
	} else {
		for _, l := range list.Lists {
			t.Log(l.Name, l.FullName, l.Description, l.Slug)
		}
	}
}
