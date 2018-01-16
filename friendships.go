package gotwitter

import "net/http"

// ShowFriendship show a friendship between two special users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friendships-show
func (app *Application) ShowFriendship(param *ShowFriendshipParam) (relation *Relationship, err error) {
	req, err := app.showFriendshipReq(param)
	if err != nil {
		return
	}

	relation = &Relationship{}
	err = app.doRequest(req, relation)

	return
}

func (app *Application) showFriendshipReq(param *ShowFriendshipParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/friendships/show.json",
		param,
	)
}
