package user

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/model"
	"go-blog/pkg/types"
)

func Get(idStr string) (User, error) {
	var user User
	id := types.StringToInt(idStr)
	if err := model.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func GetByEmail(email string) (User, error) {
	var user User
	if err := model.DB.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u *User) Create() (err error) {
	if err = model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
