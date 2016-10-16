package goinstaresponse

import (
	"strconv"
)

// Pagination every pagination have next_max_id
type Int64Pagination struct {
	NextMaxID int64 `json:"next_max_id"`
}

// Pagination every pagination have next_max_id
type StringPagination struct {
	NextMaxID string `json:"next_max_id"`
}

// UsersResponse
type UsersReponse struct {
	Status   string `json:"status"`
	BigList  bool   `json:"big_list"`
	Users    []User `json:"users"`
	PageSize int    `json:"page_size"`
	StringPagination
}

// User , Instagram user informations
type User struct {
	Username                   string `json:"username"`
	HasAnanymousProfilePicture bool   `json:"has_anonymouse_profile_picture"`
	ProfilePictureId           string `json:"profile_pic_id"`
	ProfilePictureURL          string `json:"profile_pic_url"`
	FullName                   string `json:"full_name"`
	PK                         int64  `json:"pk"`
	IsVerified                 bool   `json:"is_verified"`
	IsPrivate                  bool   `json:"is_private"`
	IsFavorite                 bool   `json:"is_favorite"`
}

// User.String return PK as string
func (user User) StringID() string {
	return strconv.FormatInt(user.PK, 10)
}

// FeedsResponse struct contains array of media and can pagination
type FeedsResponse struct {
	Status        string              `json:"status"`
	Items         []MediaItemResponse `json:"items"`
	NumResults    int                 `json:"num_results"`
	AutoLoadMore  bool                `json:"auto_load_more_enabled"`
	MoreAvailable bool                `json:"more_available"`
	StringPagination
}

// TagFeedsResponse struct contains array of MediaItemResponse
// and can pagination
// and array of MediaItemResponse for ranked_items
type TagFeedsResponse struct {
	FeedsResponse
	RankedItems []MediaItemResponse `json:"ranked_items"`
}

// MediaItemResponse struct for each media item
type MediaItemResponse struct {
	TakenAt                      int64             `json:"taken_at"`
	PK                           int64             `json:"pk"`
	ID                           string            `json:"id"`
	DeviceTimeStamp              int64             `json:"device_timestamp"`
	MediaType                    int               `json:"media_type"`
	Code                         string            `json:"code"`
	ClientCacheKey               string            `json:"client_cache_key"`
	FilterType                   int               `json:"filter_type"`
	ImageVersions                ImageVersions     `json:"image_versions2"`
	OriginalWidth                int               `json:"original_width"`
	OriginalHeight               int               `json:"original_height"`
	Location                     Location          `json:"location"`
	Lat                          float32           `json:"lat"`
	Lng                          float32           `json:"lng"`
	User                         User              `json:"user"`
	OrganicTrackingToken         string            `json:"organic_tracking_token"`
	LikeCount                    int               `json:"like_count"`
	TopLikers                    []string          `json:"top_likers"`
	HasLiked                     bool              `json:"has_liked"`
	HasMoreComments              bool              `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int               `json:"max_num_visible_preview_comments"`
	PreviewComments              []CommentResponse `json:"preview_comments"`
	Comments                     []CommentResponse `json:"comments"`
	CommentCount                 int               `json:"comment_count"`
	Caption                      string            `json:"caption",omitempty`
	CaptionIsEdited              bool              `json:"caption_is_edited"`
	PhotoOfYou                   bool              `json:"photo_of_you"`
	Int64Pagination
}

// ImageVersions struct for image information , urls and etc
type ImageVersions struct {
	Condidates []ImageCondidate `json:"condidates"`
}

// ImageCondidate have urls and image width , height
type ImageCondidate struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Location struct mean where photo or video taken
type Location struct {
	ExternalSource   string  `json:"external_source"`
	City             string  `json:"city",omitempty`
	Name             string  `json:"name"`
	FacebookPlacesID int64   `json:"facebook_places_id"`
	Address          string  `json:"address"`
	Lat              float32 `json:"lat"`
	Lng              float32 `json:"lng"`
	PK               int64   `json:"pk"`
}

// CommentResponse struct is a object for comment under media
type CommentResponse struct {
	Status       string `json:"status"`
	UserID       int64  `json:"user_id"`
	CreatedAtUTC int64  `json:"created_at_utc"`
	CreatedAt    int64  `json:"created_at"`
	BitFlags     int    `json:"bit_flags"`
	User         User   `json:"user"`
	ContentType  string `json:"content_type"`
	Text         string `json:"text"`
	MediaID      int64  `json:"media_id"`
	PK           int64  `json:"pk"`
	Type         int    `json:"type"`
}
