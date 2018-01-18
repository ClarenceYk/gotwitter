package gotwitter

import (
	"net/http"
)

// UsersLookup retrieve a list of users specified by their id or screen names.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
func (app *Application) UsersLookup(param *UsersLookupParam) (users []*UserWithStatus, err error) {
	req, err := app.usersLookupReq(param)
	if err != nil {
		return
	}

	users = make([]*UserWithStatus, 0)
	err = app.doRequest(req, &users)

	return
}

func (app *Application) usersLookupReq(param *UsersLookupParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
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
		"GET",
		"https://api.twitter.com/1.1/users/show.json",
		param,
	)
}

// UserSuggestions returns the list of suggested user categories.
// The category can be used in GET users / suggestions / :slug to get the users in that category.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-suggestions
func (app *Application) UserSuggestions(param *UserSuggestionsParam) (suggestions []*Suggestion, err error) {
	req, err := app.userSuggestionsParam(param)
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
func (app *Application) UserSuggestionsWithSlug(param *UserSuggestionsParam) (users *SuggestionUsers, err error) {
	req, err := app.userSuggestionsParam(param)
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
func (app *Application) UserSuggestionsWithSlugAndMembers(param *UserSuggestionsParam) (users []*UserWithStatus, err error) {
	req, err := app.userSuggestionsParam(param)
	if err != nil {
		return
	}

	users = make([]*UserWithStatus, 0)
	err = app.doRequest(req, &users)

	return
}

func (app *Application) userSuggestionsParam(param *UserSuggestionsParam) (*http.Request, error) {
	baseURL := "https://api.twitter.com/1.1/users/suggestions"
	if param.Slug != "" {
		if param.Members {
			baseURL = baseURL + "/" + param.Slug + "/members" + ".json"
		} else {
			baseURL = baseURL + "/" + param.Slug + ".json"
		}
		param.Slug = ""
	} else {
		baseURL = baseURL + ".json"
	}
	return app.getRequest(
		"GET",
		baseURL,
		param,
	)
}
