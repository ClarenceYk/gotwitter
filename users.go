package gotwitter

// UsersLookup retrieve a list of users specified by their id or screen names.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
func (app *Application) UsersLookup(param RequestParameters) (users []*UserWithStatus, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	users = make([]*UserWithStatus, 0)
	err = app.doRequest(req, &users)

	return
}

// ShowUser retrieve a list of users specified by their id or screen names.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
func (app *Application) ShowUser(param RequestParameters) (user *UserShow, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	user = &UserShow{}
	err = app.doRequest(req, user)

	return
}

// UserSuggestions returns the list of suggested user categories.
// The category can be used in GET users / suggestions / :slug to get the users in that category.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-suggestions
func (app *Application) UserSuggestions(param RequestParameters) (suggestions []*Suggestion, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	suggestions = make([]*Suggestion, 0)
	err = app.doRequest(req, &suggestions)

	return
}

// UserSuggestionsWithSlug returns the list of suggested users.
// The category can be used in GET users / suggestions / :slug to get the users in that category.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-suggestions
func (app *Application) UserSuggestionsWithSlug(param RequestParameters) (users *SuggestionUsers, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	users = &SuggestionUsers{}
	err = app.doRequest(req, users)

	return
}

// UserSuggestionsWithSlugAndMembers returns the list of suggested users.
// The category can be used in GET users / suggestions / :slug to get the users in that category.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-suggestions
func (app *Application) UserSuggestionsWithSlugAndMembers(param RequestParameters) (users []*UserWithStatus, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	users = make([]*UserWithStatus, 0)
	err = app.doRequest(req, &users)

	return
}
