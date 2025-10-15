package user

import (
	"errors"
	"ginBlogsTask/server/global"
	"ginBlogsTask/server/model"
	"ginBlogsTask/server/model/response"
	"ginBlogsTask/server/utils"

	"gorm.io/gorm"
)

type UserService struct{}

func (u *UserService) Register(userIn model.User) (userRes model.User, err error) {
	var user model.User
	if !errors.Is(global.GVA_DB.Where("username = ?", userIn.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userRes, errors.New("用户名已注册")
	}
	userIn.Password = utils.BcryptHash(userIn.Password)
	err = global.GVA_DB.Create(&userIn).Error
	return userIn, err
}

func (u *UserService) Login(userIn model.User) (userRes response.LoginResponse, err error) {
	var user model.User
	if err := global.GVA_DB.Where("username = ?", userIn.Username).First(&user).Error; err != nil {
		return userRes, err
	}
	check := utils.BcryptCheck(userIn.Password, user.Password)
	if !check {
		return userRes, errors.New("密码错误!")
	}
	userRes.User = userIn
	tokenString := utils.JwtCreate(user, err)
	userRes.Token = tokenString
	if err != nil {
		return userRes, errors.New("生成token失败!")
	}
	return userRes, nil
}
