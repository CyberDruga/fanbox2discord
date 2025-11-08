package fanbox

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

type Pages struct {
	Body []string `json:"body,omitempty"`
	ErrorResponse
}

type Posts struct {
	Body []Post `json:"body,omitempty"`
	ErrorResponse
}

type Post struct {
	Id                     string   `json:"id,omitempty"`
	Title                  string   `json:"title,omitempty"`
	FeeRequired            int      `json:"feeRequired"`
	PublishedTime          string   `json:"publishedTime"`
	UpdatedTime            string   `json:"updatedTime"`
	Tags                   []string `json:"tags"`
	LikeCount              int      `json:"likeCount"`
	IsCommentingRestricted bool     `json:"isCommentingRestricted"`
	CommentCount           int      `json:"commentCount"`
	IsRestricted           bool     `json:"isRestricted,omitempty"`
	User                   User     `json:"user"`
	CreatorId              string   `json:"creatorId,omitempty"`
	Cover                  Cover    `json:"cover"`
	Excerpt                string   `json:"excerpt,omitempty"`
	IsPinned               bool     `json:"IsPinned"`
}

type Cover struct {
	Type string `json:"type,omitempty"`
	Url  string `json:"url,omitempty"`
}

type User struct {
	UserId  string `json:"userId,omitempty"`
	Name    string `json:"name,omitempty"`
	IconUrl string `json:"iconUrl,omitempty"`
}
