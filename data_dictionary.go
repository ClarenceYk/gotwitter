package gotwitter

import "time"

// Tweet holds the information of a tweet.
// See more at https://developer.twitter.com/en/docs/tweets/data-dictionary/overview/tweet-object
type Tweet struct {
	ID
	CreatedAt            string            `json:"created_at"`
	Text                 string            `json:"text"`
	Source               string            `json:"source"`
	Truncated            bool              `json:"truncated"`
	InReplyToStatusID    int64             `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr string            `json:"in_reply_to_status_id_str"`
	InReplyToUserID      int64             `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   string            `json:"in_reply_to_user_id_str"`
	InReplyToScreenName  string            `json:"in_reply_to_screen_name"`
	User                 *UserWithEntities `json:"user"`
	Coordinates          *Coordinates      `json:"coordinates"`
	Place                Places            `json:"place"`
	QuotedStatusID       int64             `json:"quoted_status_id"`
	QuotedStatusIDStr    string            `json:"quoted_status_id_str"`
	IsQuoteStatus        bool              `json:"is_quote_status"`
	QuotedStatus         *Tweet            `json:"quoted_status"`
	RetweetedStatus      *Tweet            `json:"retweeted_status"`
	QuoteCount           int               `json:"quote_count"`
	ReplyCount           int               `json:"reply_count"`
	RetweetCount         int               `json:"retweet_count"`
	FavoriteCount        int               `json:"favorite_count"`
	Entities             *Entities         `json:"entities"`
	ExtendedEntities     interface{}       `json:"extended_entities"`
	Favorited            bool              `json:"favorited"`
	Retweeted            bool              `json:"retweeted"`
	PossiblySensitive    bool              `json:"possibly_sensitive"`
	FilterLevel          string            `json:"filter_level"`
	Lang                 string            `json:"lang"`
	MatchingRules        interface{}       `json:"matching_rules"`
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

// UserWithStatus hold the information of twitter users.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-users-lookup
type UserWithStatus struct {
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

// UserWithEntities for user with entities.
type UserWithEntities struct {
	User
	Entities struct {
		URL struct {
			Urls []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"url"`
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
}

// Suggestion for user.
type Suggestion struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Size int    `json:"size"`
}

// SuggestionUsers holds the informations of suggested users.
type SuggestionUsers struct {
	Suggestion
	Users []*UserWithEntities `json:"users"`
}

// List holds the informations of a twitter list.
type List struct {
	Slug            string `json:"slug"`
	Name            string `json:"name"`
	CreatedAt       string `json:"created_at"`
	URI             string `json:"uri"`
	SubscriberCount int    `json:"subscriber_count"`
	IDStr           string `json:"id_str"`
	MemberCount     int    `json:"member_count"`
	Mode            string `json:"mode"`
	ID              int    `json:"id"`
	FullName        string `json:"full_name"`
	Description     string `json:"description"`
	Following       bool   `json:"following"`
	User            *User  `json:"user"`
}

// ListMembers holds the users in a list.
type ListMembers struct {
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	NextCursor        int    `json:"next_cursor"`
	Users             []*struct {
		User
		Entities *Entities `json:"entities"`
		Status   *Status   `json:"status"`
	} `json:"users"`
	NextCursorStr string `json:"next_cursor_str"`
}

// ListMembersShow holds informations of a user.
type ListMembersShow struct {
	ID             string `json:"id"`
	IDStr          string `json:"id_str"`
	IsTranslator   bool   `json:"is_translator"`
	DefaultProfile bool   `json:"default_profile"`
	Entities       struct {
		URL struct {
			Urls []struct {
				URL         string `json:"url"`
				Indices     []int  `json:"indices"`
				DisplayURL  string `json:"display_url"`
				ExpandedURL string `json:"expanded_url"`
			} `json:"urls"`
		} `json:"url"`
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	ProfileTextColor               string `json:"profile_text_color"`
	UtcOffset                      int    `json:"utc_offset"`
	ListedCount                    int    `json:"listed_count"`
	Name                           string `json:"name"`
	Notifications                  bool   `json:"notifications"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	FollowRequestSent              bool   `json:"follow_request_sent"`
	URL                            string `json:"url"`
	Lang                           string `json:"lang"`
	ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
	CreatedAt                      string `json:"created_at"`
	Protected                      bool   `json:"protected"`
	FollowersCount                 int    `json:"followers_count"`
	ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
	ScreenName                     string `json:"screen_name"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	FriendsCount                   int    `json:"friends_count"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	Description                    string `json:"description"`
	TimeZone                       string `json:"time_zone"`
	DefaultProfileImage            bool   `json:"default_profile_image"`
	Location                       string `json:"location"`
	ProfileImageURL                string `json:"profile_image_url"`
	FavouritesCount                int    `json:"favourites_count"`
	Following                      bool   `json:"following"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	Verified                       bool   `json:"verified"`
	StatusesCount                  int    `json:"statuses_count"`
	Status                         struct {
		Entities struct {
			Urls []struct {
				URL         string `json:"url"`
				Indices     []int  `json:"indices"`
				DisplayURL  string `json:"display_url"`
				ExpandedURL string `json:"expanded_url"`
			} `json:"urls"`
			Hashtags     []interface{} `json:"hashtags"`
			UserMentions []struct {
				Name       string `json:"name"`
				Indices    []int  `json:"indices"`
				ScreenName string `json:"screen_name"`
				ID         string `json:"id"`
				IDStr      string `json:"id_str"`
			} `json:"user_mentions"`
		} `json:"entities"`
		Geo                  interface{} `json:"geo"`
		Place                interface{} `json:"place"`
		InReplyToScreenName  string      `json:"in_reply_to_screen_name"`
		InReplyToUserID      int         `json:"in_reply_to_user_id"`
		Retweeted            bool        `json:"retweeted"`
		InReplyToStatusID    int64       `json:"in_reply_to_status_id"`
		CreatedAt            string      `json:"created_at"`
		PossiblySensitive    bool        `json:"possibly_sensitive"`
		InReplyToStatusIDStr string      `json:"in_reply_to_status_id_str"`
		Contributors         interface{} `json:"contributors"`
		Favorited            bool        `json:"favorited"`
		Source               string      `json:"source"`
		InReplyToUserIDStr   string      `json:"in_reply_to_user_id_str"`
		RetweetCount         int         `json:"retweet_count"`
		ID                   string      `json:"id"`
		IDStr                string      `json:"id_str"`
		Coordinates          interface{} `json:"coordinates"`
		Truncated            bool        `json:"truncated"`
		Text                 string      `json:"text"`
	} `json:"status"`
	ContributorsEnabled       bool   `json:"contributors_enabled"`
	ProfileBackgroundImageURL string `json:"profile_background_image_url"`
	ProfileLinkColor          string `json:"profile_link_color"`
}

// ListsMemberships is the lists the specified user has been added to.
type ListsMemberships struct {
	PreviousCursor int `json:"previous_cursor"`
	Lists          []struct {
		Name            string `json:"name"`
		Slug            string `json:"slug"`
		URI             string `json:"uri"`
		IDStr           string `json:"id_str"`
		SubscriberCount int    `json:"subscriber_count"`
		MemberCount     int    `json:"member_count"`
		Mode            string `json:"mode"`
		ID              int    `json:"id"`
		FullName        string `json:"full_name"`
		Description     string `json:"description"`
		User            struct {
			ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
			ProfileBackgroundTile          bool        `json:"profile_background_tile"`
			ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
			Name                           string      `json:"name"`
			CreatedAt                      string      `json:"created_at"`
			Location                       string      `json:"location"`
			ProfileImageURL                string      `json:"profile_image_url"`
			FollowRequestSent              bool        `json:"follow_request_sent"`
			ProfileLinkColor               string      `json:"profile_link_color"`
			IsTranslator                   bool        `json:"is_translator"`
			IDStr                          string      `json:"id_str"`
			DefaultProfile                 bool        `json:"default_profile"`
			FavouritesCount                int         `json:"favourites_count"`
			ContributorsEnabled            bool        `json:"contributors_enabled"`
			URL                            interface{} `json:"url"`
			ID                             int         `json:"id"`
			ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
			UtcOffset                      int         `json:"utc_offset"`
			ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
			ListedCount                    int         `json:"listed_count"`
			Lang                           string      `json:"lang"`
			FollowersCount                 int         `json:"followers_count"`
			ProfileTextColor               string      `json:"profile_text_color"`
			Protected                      bool        `json:"protected"`
			ProfileBackgroundColor         string      `json:"profile_background_color"`
			Verified                       bool        `json:"verified"`
			TimeZone                       string      `json:"time_zone"`
			ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
			Description                    string      `json:"description"`
			Notifications                  bool        `json:"notifications"`
			GeoEnabled                     bool        `json:"geo_enabled"`
			StatusesCount                  int         `json:"statuses_count"`
			DefaultProfileImage            bool        `json:"default_profile_image"`
			FriendsCount                   int         `json:"friends_count"`
			ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
			Following                      bool        `json:"following"`
			ScreenName                     string      `json:"screen_name"`
			ShowAllInlineMedia             bool        `json:"show_all_inline_media"`
		} `json:"user"`
		Following bool `json:"following"`
	} `json:"lists"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	NextCursor        int64  `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
}

// ListsOwnerships lists owned by a specified user.
type ListsOwnerships struct {
	NextCursor        int    `json:"next_cursor"`
	NextCursorStr     string `json:"next_cursor_str"`
	PreviousCursor    int    `json:"previous_cursor"`
	PreviousCursorStr string `json:"previous_cursor_str"`
	Lists             []struct {
		ID              int    `json:"id"`
		IDStr           string `json:"id_str"`
		Name            string `json:"name"`
		URI             string `json:"uri"`
		SubscriberCount int    `json:"subscriber_count"`
		MemberCount     int    `json:"member_count"`
		Mode            string `json:"mode"`
		Description     string `json:"description"`
		Slug            string `json:"slug"`
		FullName        string `json:"full_name"`
		CreatedAt       string `json:"created_at"`
		Following       bool   `json:"following"`
		User            struct {
			ID          int    `json:"id"`
			IDStr       string `json:"id_str"`
			Name        string `json:"name"`
			ScreenName  string `json:"screen_name"`
			Location    string `json:"location"`
			Description string `json:"description"`
			URL         string `json:"url"`
			Entities    struct {
				URL struct {
					Urls []struct {
						URL         string      `json:"url"`
						ExpandedURL interface{} `json:"expanded_url"`
						Indices     []int       `json:"indices"`
					} `json:"urls"`
				} `json:"url"`
				Description struct {
					Urls []interface{} `json:"urls"`
				} `json:"description"`
			} `json:"entities"`
			Protected                      bool   `json:"protected"`
			FollowersCount                 int    `json:"followers_count"`
			FriendsCount                   int    `json:"friends_count"`
			ListedCount                    int    `json:"listed_count"`
			CreatedAt                      string `json:"created_at"`
			FavouritesCount                int    `json:"favourites_count"`
			UtcOffset                      int    `json:"utc_offset"`
			TimeZone                       string `json:"time_zone"`
			GeoEnabled                     bool   `json:"geo_enabled"`
			Verified                       bool   `json:"verified"`
			StatusesCount                  int    `json:"statuses_count"`
			Lang                           string `json:"lang"`
			ContributorsEnabled            bool   `json:"contributors_enabled"`
			IsTranslator                   bool   `json:"is_translator"`
			ProfileBackgroundColor         string `json:"profile_background_color"`
			ProfileBackgroundImageURL      string `json:"profile_background_image_url"`
			ProfileBackgroundImageURLHTTPS string `json:"profile_background_image_url_https"`
			ProfileBackgroundTile          bool   `json:"profile_background_tile"`
			ProfileImageURL                string `json:"profile_image_url"`
			ProfileImageURLHTTPS           string `json:"profile_image_url_https"`
			ProfileBannerURL               string `json:"profile_banner_url"`
			ProfileLinkColor               string `json:"profile_link_color"`
			ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
			ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
			ProfileTextColor               string `json:"profile_text_color"`
			ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
			DefaultProfile                 bool   `json:"default_profile"`
			DefaultProfileImage            bool   `json:"default_profile_image"`
			Following                      bool   `json:"following"`
			FollowRequestSent              bool   `json:"follow_request_sent"`
			Notifications                  bool   `json:"notifications"`
		} `json:"user"`
	} `json:"lists"`
}

// ListsSubscribers holds the users in a list.
type ListsSubscribers struct {
	PreviousCursor    int     `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
	NextCursor        int     `json:"next_cursor"`
	Users             []*User `json:"users"`
	NextCursorStr     string  `json:"next_cursor_str"`
}

// ListsSubscriber holds the info of a subscriber.
type ListsSubscriber struct {
	User
	Entities *Entities `json:"entities"`
	Status   *Status   `json:"status"`
}

// ListsSubscriptions holds the users in a list.
type ListsSubscriptions struct {
	PreviousCursor    int     `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
	NextCursor        int     `json:"next_cursor"`
	Lists             []*List `json:"lists"`
	NextCursorStr     string  `json:"next_cursor_str"`
}

// Retweeters holds the retweeter's ids.
type Retweeters struct {
	PreviousCursor    int     `json:"previous_cursor"`
	PreviousCursorStr string  `json:"previous_cursor_str"`
	NextCursor        int     `json:"next_cursor"`
	NextCursorStr     string  `json:"next_cursor_str"`
	IDs               []int64 `json:"ids"`
}

// SearchResult holds the users in a list.
type SearchResult struct {
	Statuses       []*Tweet `json:""`
	SearchMetadata struct {
		MaxID       int64   `json:"max_id"`
		SinceID     int64   `json:"since_id"`
		RefreshURL  string  `json:"refresh_url"`
		NextResults string  `json:"next_results"`
		Count       int     `json:"count"`
		CompletedIn float64 `json:"completed_in"`
		SinceIDStr  string  `json:"since_id_str"`
		Query       string  `json:"query"`
		MaxIDStr    string  `json:"max_id_str"`
	} `json:"search_metadata"`
}

// Trend holds the trends being retrived.
type Trend struct {
	Trends []struct {
		Name            string      `json:"name"`
		URL             string      `json:"url"`
		PromotedContent interface{} `json:"promoted_content"`
		Query           string      `json:"query"`
		TweetVolume     int         `json:"tweet_volume"`
	} `json:"trends"`
	AsOf      time.Time `json:"as_of"`
	CreatedAt time.Time `json:"created_at"`
	Locations []struct {
		Name  string `json:"name"`
		Woeid int    `json:"woeid"`
	} `json:"locations"`
}

// TrendAvailable is locations that Twitter has trending topic information for.
type TrendAvailable struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	Parentid    int    `json:"parentid"`
	PlaceType   struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"placeType"`
	URL   string `json:"url"`
	Woeid int    `json:"woeid"`
}

// TrendClosest is the locations that Twitter has trending topic information for, closest to a specified location.
type TrendClosest struct {
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	Parentid    int    `json:"parentid"`
	PlaceType   struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	} `json:"placeType"`
	URL   string `json:"url"`
	Woeid int    `json:"woeid"`
}
