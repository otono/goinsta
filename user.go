package goinsta

// User is the representation of instagram's user profile
type User struct {
	//Feed *Feed
	//Followers *Followers
	//Following *Following
	//Status *Friendship
	//Story *Story
	//Messages *Messages

	// User values shared between user and account
	ID            int64  `json:"pk"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	Biography     string `json:"biography"`
	ProfilePicURL string `json:"profile_pic_url"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	IsBusiness    bool   `json:"is_business"`
	Gender        int    `json:"gender"`

	ProfilePicID               string       `json:"profile_pic_id"`
	HasAnonymousProfilePicture bool         `json:"has_anonymous_profile_picture"`
	IsPrivate                  bool         `json:"is_private"`
	IsVerified                 bool         `json:"is_verified"`
	MediaCount                 int          `json:"media_count"`
	FollowerCount              int          `json:"follower_count"`
	FollowingCount             int          `json:"following_count"`
	GeoMediaCount              int          `json:"geo_media_count"`
	ExternalURL                string       `json:"external_url"`
	HasBiographyTranslation    bool         `json:"has_biography_translation"`
	ExternalLynxURL            string       `json:"external_lynx_url"`
	HdProfilePicURLInfo        PicURLInfo   `json:"hd_profile_pic_url_info"`
	HdProfilePicVersions       []PicURLInfo `json:"hd_profile_pic_versions"`
	UsertagsCount              int          `json:"usertags_count"`
	HasChaining                bool         `json:"has_chaining"`
	IsFavorite                 bool         `json:"is_favorite"`
	ReelAutoArchive            string       `json:"reel_auto_archive"`
	School                     School       `json:"school"`
	PublicEmail                string       `json:"public_email"`
	PublicPhoneNumber          string       `json:"public_phone_number"`
	PublicPhoneCountryCode     string       `json:"public_phone_country_code"`
	ContactPhoneNumber         string       `json:"contact_phone_number"`
	CityID                     int64        `json:"city_id"`
	CityName                   string       `json:"city_name"`
	AddressStreet              string       `json:"address_street"`
	DirectMessaging            string       `json:"direct_messaging"`
	Latitude                   float64      `json:"latitude"`
	Longitude                  float64      `json:"longitude"`
	Category                   string       `json:"category"`
	BusinessContactMethod      string       `json:"business_contact_method"`
	IsCallToActionEnabled      bool         `json:"is_call_to_action_enabled"`
	FbPageCallToActionID       string       `json:"fb_page_call_to_action_id"`
	Zip                        string       `json:"zip"`
}
