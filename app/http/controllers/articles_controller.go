package controllers

import (
	"database/sql"
	"fmt"
	"go-blog/pkg/logger"
	"go-blog/pkg/route"
	"go-blog/pkg/types"
	"net/http"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

	// URL 参数
	id := route.GetRouteVariable("id", r)

	// 读取对应的文章数据
	article, err := getArticleByID(id)

	// 如果出现错误
	if err != nil {
		if err == sql.ErrNoRows {
			// 文章未找到
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			// 数据库错误
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		// 读取成功
		tmpl, err := template.New("show.tmpl").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.tmpl")
		logger.LogError(err)

		tmpl.Execute(w, article)
	}
}
