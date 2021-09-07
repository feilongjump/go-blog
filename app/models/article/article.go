package article

import (
	"go-blog/app/models"
	"go-blog/app/models/user"
	"go-blog/pkg/route"
)

type Article struct {
	models.BaseModel
	Title      string `gorm:"type:varchar(255);" valid:"title"`
	Body       string `gorm:"type:varchar(255);" valid:"body"`
	UserID     uint64 `gorm:"not null;index"`
	User       user.User
	CategoryID uint64 `gorm:"not null;default:1;index;"`
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
