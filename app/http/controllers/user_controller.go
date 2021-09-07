package controllers

import (
	"go-blog/app/models/article"
	"go-blog/app/models/user"
	"go-blog/pkg/route"
	"go-blog/pkg/view"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	BaseController
}

// Show 用户个人页面
func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 读取对应的文章数据
	_user, err := user.Get(id)

	// 如果出现错误
	if err != nil {
		uc.ResponseForSQLError(w, err)
	} else {
		// 读取成功，显示用户文章列表
		articles, pagerData, err := article.GetByUserID(_user.GetStringID(), r, 3)
		if err != nil {
			uc.ResponseForSQLError(w, err)
		} else {
			view.Render(w, view.D{
				"Articles":  articles,
				"PagerData": pagerData,
			}, "articles.index", "articles._article_meta")
		}
	}
}
