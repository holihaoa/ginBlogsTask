package request

type Comment struct {
	Content string `json:"content" example:"内容"`
	PostID  uint   `json:"postId" example:"文章id"`
}
