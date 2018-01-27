package gotwitter

import (
	"net/http"
)

// RequestParameters is parameters sent to twitter when invorking api.
type RequestParameters interface {
	GetRequest(app *Application) (*http.Request, error)
}

// UserTimelineParam holds the paramters for retireving user's time line.
type UserTimelineParam struct {
	UserID         int64  `url:"user_id,omitempty"`
	ScreenName     string `url:"screen_name,omitempty"`
	SinceID        int64  `url:"since_id,omitempty"`
	Count          int    `url:"count,omitempty"`
	MaxID          int64  `url:"max_id,omitempty"`
	TrimUser       bool   `url:"trim_user,omitempty"`
	ExcludeReplies bool   `url:"exclude_replies,omitempty"`
	IncludeRTs     bool   `url:"include_rts,omitempty"`
}

// GetRequest return the request
func (param *UserTimelineParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/statuses/user_timeline.json",
		param,
	)
}

// FollowerIDsParam holds the paramters for retrieving follower'sid of a specified user.
type FollowerIDsParam struct {
	UserID       int64  `url:"user_id,omitempty"`
	ScreenName   string `url:"screen_name,omitempty"`
	Cursor       int64  `url:"cursor,omitempty"`
	StringifyIDs bool   `url:"stringify_ids,omitempty"`
	Count        int    `url:"count,omitempty"`
}

// GetRequest return the request
func (param *FollowerIDsParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/followers/ids.json",
		param,
	)
}

// FollowersListParam holds the paramters for retrieving followers list of a specified user.
type FollowersListParam struct {
	UserID              int64  `url:"user_id,omitempty"`
	ScreenName          string `url:"screen_name,omitempty"`
	Cursor              int64  `url:"cursor,omitempty"`
	Count               int    `url:"count,omitempty"`
	SkipStatus          bool   `url:"skip_status,omitempty"`
	IncludeUserEntities bool   `url:"include_user_entities"`
}

// GetRequest return the request
func (param *FollowersListParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/followers/list.json",
		param,
	)
}

// FriendIDsParam holds the paramters for retrieving friend's id of a specified user.
type FriendIDsParam struct {
	UserID       int64  `url:"user_id,omitempty"`
	ScreenName   string `url:"screen_name,omitempty"`
	Cursor       int64  `url:"cursor,omitempty"`
	StringifyIDs bool   `url:"stringify_ids,omitempty"`
	Count        int    `url:"count,omitempty"`
}

// GetRequest return the request
func (param *FriendIDsParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/friends/ids.json",
		param,
	)
}

// FriendsListParam holds the paramters for retrieving a friends list of a specified user.
type FriendsListParam struct {
	UserID              int64  `url:"user_id,omitempty"`
	ScreenName          string `url:"screen_name,omitempty"`
	Cursor              int64  `url:"cursor,omitempty"`
	Count               int    `url:"count,omitempty"`
	SkipStatus          bool   `url:"skip_status,omitempty"`
	IncludeUserEntities bool   `url:"include_user_entities"`
}

// GetRequest return the request
func (param *FriendsListParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/friends/list.json",
		param,
	)
}

// ShowFriendshipParam holds the paramters for showing a friendship bwtween two specified users.
type ShowFriendshipParam struct {
	SourceID         int64  `url:"source_id,omitempty"`
	SourceScreenName string `url:"source_screen_name,omitempty"`
	TargetID         int64  `url:"target_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
}

// GetRequest return the request
func (param *ShowFriendshipParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/friendships/show.json",
		param,
	)
}

// UsersLookupParam holds the paramters for retrieving users.
type UsersLookupParam struct {
	UserID          []int64  `url:"user_id,comma,omitempty"`
	ScreenName      []string `url:"screen_name,comma,omitempty"`
	IncludeEntities bool     `url:"include_entities"`
}

// GetRequest return the request
func (param *UsersLookupParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/users/lookup.json",
		param,
	)
}

// UserShowParam holds the paramters for retrieving informations of a specified user.
type UserShowParam struct {
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	IncludeEntities bool   `url:"include_entities"`
}

// GetRequest return the request
func (param *UserShowParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/users/show.json",
		param,
	)
}

// UserSuggestionsParam holds the paramters for retrieving suggestions.
type UserSuggestionsParam struct {
	Lang    string `url:"lang,omitempty"`
	Slug    string `url:"slug,omitempty"`
	Members bool   `url:"-"`
}

// GetRequest return the request
func (param *UserSuggestionsParam) GetRequest(app *Application) (*http.Request, error) {
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

// GetListsParam holds the paramters for retrieving lists.
type GetListsParam struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Reverse    bool   `url:"reverse,omitempty"`
}

// GetRequest return the request
func (param *GetListsParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/list.json",
		param,
	)
}

// ListsCreateParam holds the parameters for creating a new list.
type ListsCreateParam struct {
	Name        string `url:"name"`
	Mode        string `url:"mode,omitempty"`
	Description string `url:"description,omitempty"`
}

// GetRequest return the request
func (param *ListsCreateParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"POST",
		"https://api.twitter.com/1.1/lists/create.json",
		param,
	)
}

// ListMembersParam holds the parameters for retrieving members in a list.
type ListMembersParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
	Count           int    `url:"count,omitempty"`
	Cursor          int    `url:"cursor,omitempty"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	SkipStatus      bool   `url:"skip_status,omitempty"`
}

// GetRequest return the request
func (param *ListMembersParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/members.json",
		param,
	)
}

//ListMembersShowParam holds the parameters for showing whether a user is in a specified list.
type ListMembersShowParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	SkipStatus      bool   `url:"skip_status,omitempty"`
}

// GetRequest return the request
func (param *ListMembersShowParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/members/show.json",
		param,
	)
}

//ListsMembershipsParam holds the parameters for retrieving lists the specified user has been added to.
type ListsMembershipsParam struct {
	UserID             int64  `url:"user_id,omitempty"`
	ScreenName         string `url:"screen_name,omitempty"`
	Count              int    `url:"count,omitempty"`
	Cursor             int    `url:"cursor,omitempty"`
	FilterToOwnedLists bool   `url:"filter_to_owned_lists,omitempty"`
}

// GetRequest return the request
func (param *ListsMembershipsParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/memberships.json",
		param,
	)
}

// ListsOwnershipsParam holds the parameters for retrieving lists owned by a specified user.
type ListsOwnershipsParam struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Count      int    `url:"count,omitempty"`
	Cursor     int    `url:"cursor,omitempty"`
}

// GetRequest return the request
func (param *ListsOwnershipsParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/ownerships.json",
		param,
	)
}

// ListsShowParam holds the parameters for retrieving a specified list.
type ListsShowParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
}

// GetRequest return the request
func (param *ListsShowParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/show.json",
		param,
	)
}

// ListsStatusesParam holds the parameters for retrieving timeline of a specified list.
type ListsStatusesParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
	SinceID         int64  `url:"since_id,omitempty"`
	MaxID           int64  `url:"max_id,omitempty"`
	Count           int    `url:"count,omitempty"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	IncludeRTs      bool   `url:"include_rts,omitempty"`
}

// GetRequest return the request
func (param *ListsStatusesParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/statuses.json",
		param,
	)
}

// ListsSubscribersParam holds the parameters for retrieving users who suberscribe the specified list.
type ListsSubscribersParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
	Count           int    `url:"count,omitempty"`
	Cursor          int    `url:"cursor,omitempty"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	SkipStatus      bool   `url:"skip_status,omitempty"`
}

// GetRequest return the request
func (param *ListsSubscribersParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/subscribers.json",
		param,
	)
}

// ListsSubscriberShowParam holds the parameters for checking if the specified user is a subscriber of the specified list.
type ListsSubscriberShowParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	IncludeEntities bool   `url:"include_entities,omitempty"`
	SkipStatus      bool   `url:"skip_status,omitempty"`
}

// GetRequest return the request
func (param *ListsSubscriberShowParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/subscribers/show.json",
		param,
	)
}

// ListsSubscriptionsParam holds the parameters for obtaining a collection of the lists the specified user is subscribed to.
type ListsSubscriptionsParam struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Count      int    `url:"count,omitempty"`
	Cursor     int    `url:"cursor,omitempty"`
}

// GetRequest return the request
func (param *ListsSubscriptionsParam) GetRequest(app *Application) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/lists/subscriptions.json",
		param,
	)
}
