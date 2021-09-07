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
