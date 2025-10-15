package post

import (
	"errors"
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/model"
	"ginBlogsTask/server/model/response"
)

type PostService struct{}

func (p *PostService) Create(post *model.Post) (err error) {
	err = global.GVA_DB.Create(&post).Error
	return err
}

func (p *PostService) SearchAll() (rep response.PostResponse, err error) {
	var postResult []model.Post
	err = global.GVA_DB.Order("updated_at desc").Find(&postResult).Error
	rep.Posts = postResult
	return rep, err
}

func (p *PostService) Search(post *model.Post) (postRes *model.Post, err error) {
	err = global.GVA_DB.First(post).Error
	return post, err
}

func (p *PostService) Update(post *model.Post) (err error) {
	var postResult model.Post
	err = global.GVA_DB.Where("id=?", post.ID).First(&postResult).Error
	if err != nil {
		return err
	}
	if post.UserID != postResult.UserID {
		return errors.New("用户权限不足")
	}
	err = global.GVA_DB.Model(&model.Post{}).Where("id = ?", post.ID).Updates(model.Post{Title: post.Title, Content: post.Content}).Error
	return err
}

func (p *PostService) Delete(post *model.Post) (err error) {
	var postResult model.Post
	err = global.GVA_DB.Where("id=?", post.ID).First(&postResult).Error
	if err != nil {
		return err
	}
	if post.UserID != postResult.UserID {
		return errors.New("用户权限不足")
	}
	err = global.GVA_DB.Where("id = ?", post.ID).Delete(&model.Post{}).Error
	return err
}
