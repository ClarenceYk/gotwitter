package gotwitter

// Tweet holds the information of a tweet.
// See more at https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/tweet-object
type Tweet struct {
	ID
	CreatedAt            string       `json:"created_at"`
	Text                 string       `json:"text"`
	Source               string       `json:"source"`
	Truncated            bool         `json:"truncated"`
	InReplyToStatusID    int64        `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string       `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64        `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string       `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string       `json:"in_reply_to_screen_name"`
	User                 *User        `json:"user"`
	Coordinates          *Coordinates `json:"coordinates"`
	Place                Places       `json:"place"`
	QuotedStatusID       int64        `json:"quoted_status_id"`
	QuotedStatusIDStr    string       `json:"quoted_status_id_str"`
	IsQuoteStatus        bool         `json:"is_quote_status"`
	QuotedStatus         *Tweet       `json:"quoted_status"`
	RetweetedStatus      *Tweet       `json:"retweeted_status"`
	QuoteCount           int          `json:"quote_count"`
	ReplyCount           int          `json:"reply_count"`
	RetweetCount         int          `json:"retweet_count"`
	FavoriteCount        int          `json:"favorite_count"`
	Entities             *Entities    `json:"entities"`
	ExtendedEntities     interface{}  `json:"extended_entities"`
	Favorited            bool         `json:"favorited"`
	Retweeted            bool         `json:"retweeted"`
	PossiblySensitive    bool         `json:"possibly_sensitive"`
	FilterLevel          string       `json:"filter_level"`
	Lang                 string       `json:"lang"`
	MatchingRules        interface{}  `json:"matching_rules"`
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
	Hashtags     []*Hashtag     `json:"hashtag"`
	Media        []*Media       `json:"media"`
	URLs         []*URL         `json:"urls"`
	UserMentions []*UserMention `json:"user_mentions"`
	Symbols      []*Symbol      `json:"symbols"`
	Polls        []*Poll        `json:"polls"`
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
	Sizes             *Sizes `json:"sizes"`
	SourceStatusID    int64  `json:"source_status_id"`
	SourceStatusIDStr string `json:"source_status_id_str"`
	Type              string `json:"type"`
	URL               string `json:"url"`
}

// Sizes holds the information of four type media sizes.
type Sizes struct {
	Thumb  *Size `json:"thumb"`
	Large  *Size `json:"large"`
	Medium *Size `json:"medium"`
	Small  *Size `json:"small"`
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
	ID          string       `json:"id"`
	URL         string       `json:"url"`
	PlaceType   string       `json:"place_type"`
	Name        string       `json:"name"`
	FullName    string       `json:"full_name"`
	CounryCode  string       `json:"country_code"`
	Country     string       `json:"country"`
	BoundingBox *BoundingBox `json:"bounding_box"`
	Attributes  interface{}  `json:"attributes"`
}

// BoundingBox holds the information of area.
type BoundingBox struct {
	Coordinates [][4][2]float64 `json:"coordinates"`
	Type        string          `json:"type"`
}

// Coordinates holds the information of exact point.
type Coordinates struct {
	Coordinates [2]float64 `json:"coordinates"`
	Type        string     `json:"type"`
}

// ID holds the information of resource's id.
type ID struct {
	IDInt int64  `json:"id"`
	IDStr string `json:"id_str"`
}

// FollowerIDs holds all the information of followers id of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
type FollowerIDs struct {
	IDs               []int64 `json:"ids"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
}

// FollowersList holds all the information of followers id of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
type FollowersList struct {
	Users             []*User `json:"users"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
}

// FriendIDs holds all the information of friends id of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-ids
type FriendIDs struct {
	IDs               []int64 `json:"ids"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
}

// FriendsList holds all the information of a friends list of a specified user.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friends-list:
type FriendsList struct {
	Users             []*User `json:"users"`
	NextCursor        int64   `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	PreviousCursor    int64   `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
}

// Relationship holds the information of the relationship between two users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-friendships-show
type Relationship struct {
	Relationship struct {
		Target Target `json:"target"`
		Source Source `json:"source"`
	} `json:"relationship"`
}

// Target holds the information of the target user.
type Target struct {
	ID
	ScreenName string `json:"screen_name"`
	Following  bool   `json:"following"`
	FollowedBy bool   `json:"followed_by"`
}

// Source holds the information of the source user.
type Source struct {
	ID
	CanDM                bool   `json:"can_dm"`
	Blocking             bool   `json:"blocking"`
	Muting               bool   `json:"muting"`
	AllReplies           bool   `json:"all_replies"`
	WantRetweets         bool   `json:"want_retweets"`
	MarkedSpam           bool   `json:"marked_spam"`
	ScreenName           string `json:"screen_name"`
	Following            bool   `json:"following"`
	FollowedBy           bool   `json:"followed_by"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
}

// Status holds the inforamtion of status.
type Status struct {
	Coordinates          *Coordinates `json:"coordinates"`
	CreatedAt            string       `json:"created_at"`
	Favorited            bool         `json:"favorited"`
	Truncated            bool         `json:"truncated"`
	Entities             *Entities    `json:"entities"`
	IDStr                string       `json:"id_str"`
	InReplyToUserIDStr   string       `json:"in_reply_to_user_id_str"`
	Text                 string       `json:"text"`
	Contributors         interface{}  `json:"contributors"`
	RetweetCount         int          `json:"retweet_count"`
	ID                   int64        `json:"id"`
	InReplyToStatusIDStr string       `json:"in_reply_to_status_id_str"`
	Geo                  interface{}  `json:"geo"`
	Retweeted            bool         `json:"retweeted"`
	InReplyToUserID      int          `json:"in_reply_to_user_id"`
	Place                *Places      `json:"place"`
	Source               string       `json:"source"`
	InReplyToScreenName  string       `json:"in_reply_to_screen_name"`
	InReplyToStatusID    int64        `json:"in_reply_to_status_id"`
}

// UserLookup hold the information of twitter users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
type UserLookup struct {
	Name                           string  `json:"name"`
	ProfileSidebarFillColor        string  `json:"profile_sidebar_fill_color"`
	ProfileBackgroundTile          bool    `json:"profile_background_tile"`
	ProfileSidebarBorderColor      string  `json:"profile_sidebar_border_color"`
	ProfileImageURL                string  `json:"profile_image_url"`
	Location                       string  `json:"location"`
	CreatedAt                      string  `json:"created_at"`
	FollowRequestSent              bool    `json:"follow_request_sent"`
	IDStr                          string  `json:"id_str"`
	ProfileLinkColor               string  `json:"profile_link_color"`
	IsTranslator                   bool    `json:"is_translator"`
	DefaultProfile                 bool    `json:"default_profile"`
	FavouritesCount                int     `json:"favourites_count"`
	ContributorsEnabled            bool    `json:"contributors_enabled"`
	URL                            string  `json:"url"`
	ProfileImageURLHTTPS           string  `json:"profile_image_url_https"`
	UtcOffset                      int     `json:"utc_offset"`
	ID                             int     `json:"id"`
	ProfileUseBackgroundImage      bool    `json:"profile_use_background_image"`
	ListedCount                    int     `json:"listed_count"`
	ProfileTextColor               string  `json:"profile_text_color"`
	Lang                           string  `json:"lang"`
	FollowersCount                 int     `json:"followers_count"`
	Protected                      bool    `json:"protected"`
	ProfileBackgroundImageURLHTTPS string  `json:"profile_background_image_url_https"`
	GeoEnabled                     bool    `json:"geo_enabled"`
	Description                    string  `json:"description"`
	ProfileBackgroundColor         string  `json:"profile_background_color"`
	Verified                       bool    `json:"verified"`
	Notifications                  bool    `json:"notifications"`
	TimeZone                       string  `json:"time_zone"`
	StatusesCount                  int     `json:"statuses_count"`
	Status                         *Status `json:"status"`
	ProfileBackgroundImageURL      string  `json:"profile_background_image_url"`
	DefaultProfileImage            bool    `json:"default_profile_image"`
	FriendsCount                   int     `json:"friends_count"`
	ScreenName                     string  `json:"screen_name"`
	Following                      bool    `json:"following"`
	ShowAllInlineMedia             bool    `json:"show_all_inline_media"`
}

// UserShow hold the information of twitter users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
type UserShow struct {
	ID
	Name                           string      `json:"name"`
	ScreenName                     string      `json:"screen_name"`
	Location                       string      `json:"location"`
	ProfileLocation                interface{} `json:"profile_location"`
	Description                    string      `json:"description"`
	URL                            string      `json:"url"`
	Entities                       *Entities   `json:"entities"`
	Protected                      bool        `json:"protected"`
	FollowersCount                 int         `json:"followers_count"`
	FriendsCount                   int         `json:"friends_count"`
	ListedCount                    int         `json:"listed_count"`
	CreatedAt                      string      `json:"created_at"`
	FavouritesCount                int         `json:"favourites_count"`
	UtcOffset                      int         `json:"utc_offset"`
	TimeZone                       string      `json:"time_zone"`
	GeoEnabled                     bool        `json:"geo_enabled"`
	Verified                       bool        `json:"verified"`
	StatusesCount                  int         `json:"statuses_count"`
	Lang                           string      `json:"lang"`
	Status                         *Status     `json:"status"`
	ContributorsEnabled            bool        `json:"contributors_enabled"`
	IsTranslator                   bool        `json:"is_translator"`
	IsTranslationEnabled           bool        `json:"is_translation_enabled"`
	ProfileBackgroundColor         string      `json:"profile_background_color"`
	ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool        `json:"profile_background_tile"`
	ProfileImageURL                string      `json:"profile_image_url"`
	ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
	ProfileBannerURL               string      `json:"profile_banner_url"`
	ProfileLinkColor               string      `json:"profile_link_color"`
	ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string      `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
	HasExtendedProfile             bool        `json:"has_extended_profile"`
	DefaultProfile                 bool        `json:"default_profile"`
	DefaultProfileImage            bool        `json:"default_profile_image"`
	Following                      bool        `json:"following"`
	FollowRequestSent              bool        `json:"follow_request_sent"`
	Notifications                  bool        `json:"notifications"`
	TranslatorType                 string      `json:"translator_type"`
}
