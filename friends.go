package gotwitter

import "encoding/json"

// FriendIDs retrieve friend ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-ids
func (app *Application) FriendIDs(param *FriendIDsParam) (*FriendIDs, error) {
	// NOTICE: stringify must be false
	param.StringifyIDs = false

	req, err := app.friendIDsReq(param)
	if err != nil {
		return nil, err
	}

	grc, err := app.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer grc.Close()

	ids := FriendIDs{}
	if err := json.NewDecoder(grc).Decode(&ids); err != nil {
		return nil, err
	}

	return &ids, nil
}

// FriendsList retrieve a friends list of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-list
func (app *Application) FriendsList(param *FriendsListParam) (*FriendsList, error) {
	req, err := app.friendsListReq(param)
	if err != nil {
		return nil, err
	}

	grc, err := app.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer grc.Close()

	ids := FriendsList{}
	if err := json.NewDecoder(grc).Decode(&ids); err != nil {
		return nil, err
	}

	return &ids, nil
}
