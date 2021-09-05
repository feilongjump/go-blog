package article

import (
	"go-blog/app/models"
	"go-blog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title string `gorm:"type:varchar(255);" valid:"title"`
	Body  string `gorm:"type:varchar(255);" valid:"body"`
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
