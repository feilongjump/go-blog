package category

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/model"
)

func (c *Category) Create() (err error) {
	if err = model.DB.Create(&c).Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}
