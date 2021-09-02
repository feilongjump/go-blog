package user

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/model"
)

func (u *User) Create() (err error) {
	if err = model.DB.Create(&u).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}
