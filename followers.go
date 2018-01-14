package gotwitter

import (
	"encoding/json"
)

// FollowerIDs retrieve follower ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
func (app *Application) FollowerIDs(param *FollowerIDsParam) (*FollowerIDs, error) {
	// NOTICE: stringify must be false
	param.StringifyIDs = false

	req, err := app.followerIDsReq(param)
	if err != nil {
		return nil, err
	}

	grc, err := app.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer grc.Close()

	ids := FollowerIDs{}
	if err := json.NewDecoder(grc).Decode(&ids); err != nil {
		return nil, err
	}

	return &ids, nil
}

// FollowersList retrieve a cursored collection of user objects for users following the specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
func (app *Application) FollowersList(param *FollowersListParam) (*FollowersList, error) {
	req, err := app.followersListReq(param)
	if err != nil {
		return nil, err
	}

	grc, err := app.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer grc.Close()

	list := FollowersList{}
	if err := json.NewDecoder(grc).Decode(&list); err != nil {
		return nil, err
	}

	return &list, nil
}
