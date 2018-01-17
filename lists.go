package gotwitter

import "net/http"

// GetLists Returns all lists the authenticating or specified user subscribes to, including their own.
// The user is specified using the user_id or screen_name parameters.
// If no user is given, the authenticating user is used.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
func (app *Application) GetLists(param *GetListsParam) (users []*List, err error) {
	req, err := app.getListsParam(param)
	if err != nil {
		return
	}

	users = make([]*List, 0)
	err = app.doRequest(req, &users)

	return
}

func (app *Application) getListsParam(param *GetListsParam) (*http.Request, error) {
	return app.getRequest(
		"https://api.twitter.com/1.1/lists/list.json",
		param,
	)
}
