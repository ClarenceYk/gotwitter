package gotwitter

// FriendIDs retrieve friend ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-ids
func (app *Application) FriendIDs(param RequestParameters) (ids *FriendIDs, err error) {
	// NOTICE: stringify must be false
	param.(*FriendIDsParam).StringifyIDs = false

	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ids = &FriendIDs{}
	err = app.doRequest(req, ids)

	return
}

// FriendsList retrieve a friends list of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-list
func (app *Application) FriendsList(param RequestParameters) (list *FriendsList, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	list = &FriendsList{}
	err = app.doRequest(req, list)

	return
}
