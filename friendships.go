package gotwitter

import "encoding/json"

// ShowFriendship show a friendship between two special users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friendships-show
func (app *Application) ShowFriendship(param *ShowFriendshipParam) (*Relationship, error) {
	req, err := app.showFriendshipReq(param)
	if err != nil {
		return nil, err
	}

	grc, err := app.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer grc.Close()

	relation := Relationship{}
	if err := json.NewDecoder(grc).Decode(&relation); err != nil {
		return nil, err
	}

	return &relation, nil
}
