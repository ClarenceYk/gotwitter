package gotwitter

import "encoding/json"

// FriendIDs retrieve friend ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
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
