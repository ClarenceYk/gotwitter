package gotwitter

import "net/http"

// UsersLookup retrieve a list of users specified by their id or screen names.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
func (app *Application) UsersLookup(param *UsersLookupParam) (users []*UserLookup, err error) {
	req, err := app.usersLookupReq(param)
	if err != nil {
		return
	}

	users = make([]*UserLookup, 0)
	err = app.doRequest(req, &users)

	return
}

func (app *Application) usersLookupReq(param *UsersLookupParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/users/lookup.json",
		param,
	)
}

// ShowUser retrieve a list of users specified by their id or screen names.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
func (app *Application) ShowUser(param *UserShowParam) (user *UserShow, err error) {
	req, err := app.userShowParam(param)
	if err != nil {
		return
	}

	user = &UserShow{}
	err = app.doRequest(req, user)

	return
}

func (app *Application) userShowParam(param *UserShowParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/users/show.json",
		param,
	)
}
