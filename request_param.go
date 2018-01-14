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
