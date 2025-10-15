package response

import "ginBlogsTask/server/model"

type PostResponse struct {
	Posts []model.Post `json:"posts"`
}
