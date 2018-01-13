package gotwitter

// UserTimelineParam hold the paramters for fetching user's time line
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
