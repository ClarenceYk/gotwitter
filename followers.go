package gotwitter

// FollowerIDs retrieve follower ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
func (app *Application) FollowerIDs(param RequestParameters) (ids *FollowerIDs, err error) {
	// NOTICE: stringify must be false
	param.(*FollowerIDsParam).StringifyIDs = false

	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ids = &FollowerIDs{}
	err = app.doRequest(req, ids)

	return
}

// FollowersList retrieve a cursored collection of user objects for users following the specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
func (app *Application) FollowersList(param RequestParameters) (list *FollowersList, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	list = &FollowersList{}
	err = app.doRequest(req, list)

	return
}
