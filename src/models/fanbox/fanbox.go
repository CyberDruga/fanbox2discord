package fanbox

import "encoding/json"

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
	PostId                 string     `json:"id,omitempty"`
	Title                  JsonString `json:"title,omitempty"`
	FeeRequired            int        `json:"feeRequired"`
	PublishedDatetime      string     `json:"publishedDatetime"`
	UpdatedDatetime        string     `json:"updatedDatetime"`
	Tags                   []string   `json:"tags"`
	LikeCount              int        `json:"likeCount"`
	IsCommentingRestricted bool       `json:"isCommentingRestricted"`
	CommentCount           int        `json:"commentCount"`
	IsRestricted           bool       `json:"isRestricted,omitempty"`
	User                   User       `json:"user"`
	CreatorId              string     `json:"creatorId,omitempty"`
	Cover                  Cover      `json:"cover"`
	Excerpt                JsonString `json:"excerpt,omitempty"`
	IsPinned               bool       `json:"IsPinned"`
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

type JsonString string

func (js *JsonString) UnmarshalText(text []byte) error {

	data, _ := json.Marshal(string(text))

	result := string(data)

	*js = JsonString(result[1 : len(result)-1])

	return nil
}
