package gotwitter

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

// FollowerIDsParam holds the paramters for retrieving follower'sid of a specified user.
type FollowerIDsParam struct {
	UserID       int64  `url:"user_id,omitempty"`
	ScreenName   string `url:"screen_name,omitempty"`
	Cursor       int64  `url:"cursor,omitempty"`
	StringifyIDs bool   `url:"stringify_ids,omitempty"`
	Count        int    `url:"count,omitempty"`
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

// FriendIDsParam holds the paramters for retrieving friend's id of a specified user.
type FriendIDsParam struct {
	UserID       int64  `url:"user_id,omitempty"`
	ScreenName   string `url:"screen_name,omitempty"`
	Cursor       int64  `url:"cursor,omitempty"`
	StringifyIDs bool   `url:"stringify_ids,omitempty"`
	Count        int    `url:"count,omitempty"`
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

// ShowFriendshipParam holds the paramters for showing a friendship bwtween two specified users.
type ShowFriendshipParam struct {
	SourceID         int64  `url:"source_id,omitempty"`
	SourceScreenName string `url:"source_screen_name,omitempty"`
	TargetID         int64  `url:"target_id,omitempty"`
	TargetScreenName string `url:"target_screen_name,omitempty"`
}

// UsersLookupParam holds the paramters for retrieving users.
type UsersLookupParam struct {
	UserID          []int64  `url:"user_id,comma,omitempty"`
	ScreenName      []string `url:"screen_name,comma,omitempty"`
	IncludeEntities bool     `url:"include_entities"`
}

// UserShowParam holds the paramters for retrieving informations of a specified user.
type UserShowParam struct {
	UserID          int64  `url:"user_id,omitempty"`
	ScreenName      string `url:"screen_name,omitempty"`
	IncludeEntities bool   `url:"include_entities"`
}

// UserSuggestionsParam holds the paramters for retrieving suggestions.
type UserSuggestionsParam struct {
	Lang    string `url:"lang,omitempty"`
	Slug    string `url:"slug,omitempty"`
	Members bool   `url:"-"`
}

// GetListsParam holds the paramters for retrieving lists.
type GetListsParam struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Reverse    bool   `url:"reverse,omitempty"`
}

// ListsCreateParam holds the parameters for creating a new list.
type ListsCreateParam struct {
	Name        string `url:"name"`
	Mode        string `url:"mode,omitempty"`
	Description string `url:"description,omitempty"`
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

//ListsMembershipsParam holds the parameters for retrieving lists the specified user has been added to.
type ListsMembershipsParam struct {
	UserID             int64  `url:"user_id,omitempty"`
	ScreenName         string `url:"screen_name,omitempty"`
	Count              int    `url:"count,omitempty"`
	Cursor             int    `url:"cursor,omitempty"`
	FilterToOwnedLists bool   `url:"filter_to_owned_lists,omitempty"`
}

// ListsOwnershipsParam holds the parameters for retrieving lists owned by a specified user.
type ListsOwnershipsParam struct {
	UserID     int64  `url:"user_id,omitempty"`
	ScreenName string `url:"screen_name,omitempty"`
	Count      int    `url:"count,omitempty"`
	Cursor     int    `url:"cursor,omitempty"`
}

// ListsShowParam holds the parameters for retrieving a specified list.
type ListsShowParam struct {
	ListID          int64  `url:"list_id,omitempty"`
	Slug            string `url:"slug,omitempty"`
	OwnerScreenName string `url:"owner_screen_name,omitempty"`
	OwnerID         int64  `url:"owner_id,omitempty"`
}
