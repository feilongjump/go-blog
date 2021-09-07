package controllers

import (
	"fmt"
	"go-blog/app/models/category"
	"go-blog/app/requests"
	"go-blog/pkg/flash"
	"go-blog/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	// 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// 表单验证
	errors := requests.ValidateCategoryForm(_category)

	// 检查错误
	if len(errors) == 0 {
		// 创建文章分类
		_category.Create()
		if _category.ID > 0 {
			flash.Success("创建分类成功")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

func (*CategoriesController) Show(w http.ResponseWriter, r *http.Request) {

}
