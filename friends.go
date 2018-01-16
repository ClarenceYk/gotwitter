package gotwitter

import "net/http"

// FriendIDs retrieve friend ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-ids
func (app *Application) FriendIDs(param *FriendIDsParam) (ids *FriendIDs, err error) {
	// NOTICE: stringify must be false
	param.StringifyIDs = false

	req, err := app.friendIDsReq(param)
	if err != nil {
		return
	}

	ids = &FriendIDs{}
	err = app.doRequest(req, ids)

	return
}

func (app *Application) friendIDsReq(param *FriendIDsParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/friends/ids.json",
		param,
	)
}

// FriendsList retrieve a friends list of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-list
func (app *Application) FriendsList(param *FriendsListParam) (list *FriendsList, err error) {
	req, err := app.friendsListReq(param)
	if err != nil {
		return
	}

	list = &FriendsList{}
	err = app.doRequest(req, list)

	return
}

func (app *Application) friendsListReq(param *FriendsListParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/friends/list.json",
		param,
	)
}
