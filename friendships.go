package gotwitter

// ShowFriendship show a friendship between two special users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friendships-show
func (app *Application) ShowFriendship(param RequestParameters) (relation *Relationship, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	relation = &Relationship{}
	err = app.doRequest(req, relation)

	return
}
