package request

type Post struct {
	Title   string `json:"title" example:"标题"`
	Content string `json:"content" example:"内容"`
	UserId  uint   `json:"user_id" example:"1"`
	PostId  uint   `json:"post_id" example:"1"`
}
