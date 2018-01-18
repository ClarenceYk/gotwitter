package gotwitter

import "net/http"

// GetLists Returns all lists the authenticating or specified user subscribes to, including their own.
// The user is specified using the user_id or screen_name parameters.
// If no user is given, the authenticating user is used.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
func (app *Application) GetLists(param *GetListsParam) (lists []*List, err error) {
	req, err := app.getListsReq(param)
	if err != nil {
		return
	}

	lists = make([]*List, 0)
	err = app.doRequest(req, &lists)

	return
}

func (app *Application) getListsReq(param *GetListsParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/list.json",
		param,
	)
}

// ListsCreate Creates a new list for the authenticated user.
// Note that you can create up to 1000 lists per account.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/post-lists-create
func (app *Application) ListsCreate(param *ListsCreateParam) (list *List, err error) {
	req, err := app.getListsCreateReq(param)
	if err != nil {
		return
	}

	list = &List{}
	err = app.doRequest(req, list)

	return
}

func (app *Application) getListsCreateReq(param *ListsCreateParam) (*http.Request, error) {
	return app.getRequest(
		"POST",
		"https://api.twitter.com/1.1/lists/create.json",
		param,
	)
}

// ListMembers Returns the members of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-members
func (app *Application) ListMembers(param *ListMembersParam) (members *ListMembers, err error) {
	req, err := app.getListMembersReq(param)
	if err != nil {
		return
	}

	members = &ListMembers{}
	err = app.doRequest(req, members)

	return
}

func (app *Application) getListMembersReq(param *ListMembersParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/members.json",
		param,
	)
}

// ListMembersShow check if the specified user is a member of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-members-show
func (app *Application) ListMembersShow(param *ListMembersShowParam) (show *ListMembersShow, err error) {
	req, err := app.getListMembersShowReq(param)
	if err != nil {
		return
	}

	show = &ListMembersShow{}
	err = app.doRequest(req, show)

	return
}

func (app *Application) getListMembersShowReq(param *ListMembersShowParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/members/show.json",
		param,
	)
}

// ListsMemberships check if the specified user is a member of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-memberships
func (app *Application) ListsMemberships(param *ListsMembershipsParam) (mships *ListsMemberships, err error) {
	req, err := app.getListsMembershipsReq(param)
	if err != nil {
		return
	}

	mships = &ListsMemberships{}
	err = app.doRequest(req, mships)

	return
}

func (app *Application) getListsMembershipsReq(param *ListsMembershipsParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/memberships.json",
		param,
	)
}

// ListsOwnerships returns the lists owned by the specified Twitter user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-ownerships
func (app *Application) ListsOwnerships(param *ListsOwnershipsParam) (oships *ListsOwnerships, err error) {
	req, err := app.getListsOwnershipsReq(param)
	if err != nil {
		return
	}

	oships = &ListsOwnerships{}
	err = app.doRequest(req, oships)

	return
}

func (app *Application) getListsOwnershipsReq(param *ListsOwnershipsParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/ownerships.json",
		param,
	)
}
