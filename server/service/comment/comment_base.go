package comment

import (
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/model"
	"ginBlogsTask/server/model/response"
)

type CommentService struct{}

func (c *CommentService) Create(commengIn *model.Comment) (err error) {
	err = global.GVA_DB.Create(&commengIn).Error
	return err
}

func (c *CommentService) Search(comment *model.Comment) (commentRep response.CommentResponse, err error) {
	var commentResult []model.Comment
	err = global.GVA_DB.Where("post_id = ?", comment.PostID).Order("updated_at desc").Find(&commentResult).Error
	commentRep.Comments = commentResult
	return commentRep, err
}
