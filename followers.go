package gotwitter

import "net/http"

// FollowerIDs retrieve follower ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
func (app *Application) FollowerIDs(param *FollowerIDsParam) (ids *FollowerIDs, err error) {
	// NOTICE: stringify must be false
	param.StringifyIDs = false

	req, err := app.followerIDsReq(param)
	if err != nil {
		return
	}

	ids = &FollowerIDs{}
	err = app.doRequest(req, ids)

	return
}

func (app *Application) followerIDsReq(param *FollowerIDsParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/followers/ids.json",
		param,
	)
}

// FollowersList retrieve a cursored collection of user objects for users following the specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
func (app *Application) FollowersList(param *FollowersListParam) (list *FollowersList, err error) {
	req, err := app.followersListReq(param)
	if err != nil {
		return
	}

	list = &FollowersList{}
	err = app.doRequest(req, list)

	return
}

func (app *Application) followersListReq(param *FollowersListParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/followers/list.json",
		param,
	)
}
