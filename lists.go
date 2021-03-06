package gotwitter

// GetLists Returns all lists the authenticating or specified user subscribes to, including their own.
// The user is specified using the user_id or screen_name parameters.
// If no user is given, the authenticating user is used.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-list
func (app *Application) GetLists(param RequestParameters) (lists []*List, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	lists = make([]*List, 0)
	err = app.doRequest(req, &lists)

	return
}

// ListsCreate Creates a new list for the authenticated user.
// Note that you can create up to 1000 lists per account.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/post-lists-create
func (app *Application) ListsCreate(param RequestParameters) (list *List, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	list = &List{}
	err = app.doRequest(req, list)

	return
}

// ListMembers Returns the members of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-members
func (app *Application) ListMembers(param RequestParameters) (members *ListMembers, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	members = &ListMembers{}
	err = app.doRequest(req, members)

	return
}

// ListMembersShow check if the specified user is a member of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-members-show
func (app *Application) ListMembersShow(param RequestParameters) (show *ListMembersShow, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	show = &ListMembersShow{}
	err = app.doRequest(req, show)

	return
}

// ListsMemberships check if the specified user is a member of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-memberships
func (app *Application) ListsMemberships(param RequestParameters) (mships *ListsMemberships, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	mships = &ListsMemberships{}
	err = app.doRequest(req, mships)

	return
}

// ListsOwnerships returns the lists owned by the specified Twitter user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-ownerships
func (app *Application) ListsOwnerships(param RequestParameters) (oships *ListsOwnerships, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	oships = &ListsOwnerships{}
	err = app.doRequest(req, oships)

	return
}

// ListsShow returns the specified lists.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-show
func (app *Application) ListsShow(param RequestParameters) (list *List, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	list = &List{}
	err = app.doRequest(req, list)

	return
}

// ListsStatuses returns a timeline of a tweets authored by members of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-statuses
func (app *Application) ListsStatuses(param RequestParameters) (ts []*Tweet, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*Tweet, 0)
	err = app.doRequest(req, &ts)

	return
}

// ListsSubscribers returns the subscribers of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-subscribers
func (app *Application) ListsSubscribers(param RequestParameters) (subscribers *ListsSubscribers, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	subscribers = &ListsSubscribers{}
	err = app.doRequest(req, subscribers)

	return
}

// ListsSubscriberShow check if the specified suer is a subscriber of the specified list.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-subscribers-show
func (app *Application) ListsSubscriberShow(param RequestParameters) (subscriber *ListsSubscriber, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	subscriber = &ListsSubscriber{}
	err = app.doRequest(req, subscriber)

	return
}

// ListsSubscriptions Obtian a collection of the lists the specified user is subscribed to.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/create-manage-lists/api-reference/get-lists-subscriptions
func (app *Application) ListsSubscriptions(param RequestParameters) (subscriptions *ListsSubscriptions, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	subscriptions = &ListsSubscriptions{}
	err = app.doRequest(req, subscriptions)

	return
}
