package response

import "ginBlogsTask/server/model"

type CommentResponse struct {
	Comments []model.Comment `json:"comments"`
}
