package gotwitter

// Tweet holds the information of a tweet.
// See more at https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/tweet-object
type Tweet struct {
	ID
	CreatedAt            string      `json:"created_at"`
	Text                 string      `json:"text"`
	Source               string      `json:"source"`
	Truncated            bool        `json:"truncated"`
	InReplyToStatusID    int64       `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string      `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64       `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string      `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
	User                 User        `json:"user"`
	Coordinates          Coordinates `json:"coordinates"`
	Place                Places      `json:"place"`
	QuotedStatusID       int64       `json:"quoted_status_id"`
	QuotedStatusIDStr    string      `json:"quoted_status_id_str"`
	IsQuoteStatus        bool        `json:"is_quote_status"`
	QuotedStatus         *Tweet      `json:"quoted_status"`
	RetweetedStatus      *Tweet      `json:"retweeted_status"`
	QuoteCount           int         `json:"quote_count"`
	ReplyCount           int         `json:"reply_count"`
	RetweetCount         int         `json:"retweet_count"`
	FavoriteCount        int         `json:"favorite_count"`
	Entities             Entities    `json:"entities"`
	ExtendedEntities     interface{} `json:"extended_entities"`
	Favorited            bool        `json:"favorited"`
	Retweeted            bool        `json:"retweeted"`
	PossiblySensitive    bool        `json:"possibly_sensitive"`
	FilterLevel          string      `json:"filter_level"`
	Lang                 string      `json:"lang"`
	MatchingRules        interface{} `json:"matching_rules"`
}

// User holds all the informations of a twitter user.
// See more at https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/user-object
type User struct {
	ID
	Name                           string      `json:"name"`
	ScreenName                     string      `json:"screen_name"`
	Location                       string      `json:"location"`
	URL                            string      `json:"url"`
	Description                    string      `json:"description"`
	Derived                        interface{} `json:"derived"`
	Protected                      bool        `json:"protected"`
	Verified                       bool        `json:"verified"`
	FollowersCount                 int         `json:"followers_count"`
	FriendsCount                   int         `json:"friends_count"`
	ListedCount                    int         `json:"listed_count"`
	FavouritesCount                int         `json:"favourites_count"`
	StatusesCount                  int         `json:"statuses_count"`
	CreatedAt                      string      `json:"created_at"`
	UTCOffset                      int         `json:"utc_offset"`
	TimeZone                       string      `json:"time_zone"`
	GEOEnabled                     bool        `json:"geo_enabled"`
	Lang                           string      `json:"lang"`
	ContributorsEnabled            bool        `json:"contributors_enabled"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHttps string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileBannerURL               string      `json:"profile_banner_url"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHttps           string      `json:"profile_image_url_https"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	DefaultProfile                 bool        `json:"default_profile"`
	DefaultProfileImage            bool        `json:"default_profile_image"`
	WithheldInCountries            string      `json:"withheld_in_countries"`
	WithheldScope                  string      `json:"withheld_scope"`
}

// Entities holds the information of entities in a tweet.
// See more at https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/entities-object
type Entities struct {
	Hashtags     []Hashtag     `json:"hashtag"`
	Media        []Media       `json:"media"`
	URLs         []URL         `json:"urls"`
	UserMentions []UserMention `json:"user_mentions"`
	Symbols      []Symbol      `json:"symbols"`
	Polls        []Poll        `json:"polls"`
}

// Hashtag holds the information of hashtag in a tweet.
type Hashtag struct {
	Indices [2]int `json:"indices"`
	Text    string `json:"text"`
}

// Media holds the information of a media object.
type Media struct {
	ID
	DisplayURL        string `json:"display_url"`
	ExpandedURL       string `json:"expanded_url"`
	Indices           [2]int `json:"indices"`
	MediaURL          string `json:"media_url"`
	MediaURLHttps     string `json:"media_url_https"`
	Sizes             Sizes  `json:"sizes"`
	SourceStatusID    int64  `json:"source_status_id"`
	SourceStatusIDStr string `json:"source_status_id_str"`
	Type              string `json:"type"`
	URL               string `json:"url"`
}

// Sizes holds the information of four type media sizes.
type Sizes struct {
	Thumb  Size `json:"thumb"`
	Large  Size `json:"large"`
	Medium Size `json:"medium"`
	Small  Size `json:"small"`
}

// Size hold the information of a media size.
type Size struct {
	W      int    `json:"w"`
	H      int    `json:"h"`
	Resize string `json:"resize"`
}

// URL holds the information of a url
type URL struct {
	DisplayURL  string `json:"display_url"`
	ExpandedURL string `json:"Expanded_url"`
	Indices     [2]int `json:"indices"`
	URL         string `json:"url"`
	Unwound     struct {
		URL         string `json:"url"`
		Status      int    `json:"status"`
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"unwound"`
}

// UserMention holds the information of metions in a tweet.
type UserMention struct {
	ID
	Indices    [2]int `json:"indices"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

// Symbol holds the information of symbol in a tweet.
type Symbol struct {
	Indices [2]int `json:"indices"`
	Text    string `json:"text"`
}

// Poll holds the information of poll in a tweet.
type Poll struct {
	Options []struct {
		Position int    `json:"position"`
		Text     string `json:"text"`
	} `json:"options"`
	EndDatetime     string `json:"end_datetime"`
	DurationMinutes string `json:"duration_minutes"`
}

// Places holds the information of a place.
// See more at https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/geo-objects#coordinates-dictionary
type Places struct {
	ID          string      `json:"id"`
	URL         string      `json:"url"`
	PlaceType   string      `json:""`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	CounryCode  string      `json:"country_code"`
	Country     string      `json:"country"`
	BoundingBox BoundingBox `json:"bounding_box"`
	Attributes  interface{} `json:"attributes"`
}

// BoundingBox holds the information of area.
type BoundingBox struct {
	Coordinates [4][2]float64 `json:"coordinates"`
	Type        string        `json:"type"`
}

// Coordinates holds the information of exact point.
type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

// FollowerIDs holds all the information of followers id of user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
type FollowerIDs struct {
	IDs               []int64 `json:"ids"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
}

// ID holds the information of resource's id.
type ID struct {
	ID    int64  `json:"id"`
	IDStr string `json:"id_str"`
}
