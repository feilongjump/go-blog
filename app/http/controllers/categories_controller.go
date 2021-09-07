package controllers

import (
	"fmt"
	"go-blog/app/models/article"
	"go-blog/app/models/category"
	"go-blog/app/requests"
	"go-blog/pkg/flash"
	"go-blog/pkg/route"
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

func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 读取对应的数据
	_category, err := category.Get(id)

	// 获取结果集
	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 3)

	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		// 加载模板
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

func (cc *CategoriesController) Edit(w http.ResponseWriter, r *http.Request) {
	// 获取参数
	id := route.GetRouteVariable("id", r)

	// 获取数据
	_category, err := category.Get(id)

	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Category": _category,
		}, "categories.edit")
	}
}

func (cc *CategoriesController) Update(w http.ResponseWriter, r *http.Request) {
	// 获取参数
	id := route.GetRouteVariable("id", r)

	// 获取数据
	_category, err := category.Get(id)

	if err != nil {
		cc.ResponseForSQLError(w, err)
	} else {
		// 校验数据
		_category.Name = r.PostFormValue("name")

		errors := requests.ValidateCategoryForm(_category)

		if len(errors) == 0 {
			// 未出现错误
			rowsAffected, err := _category.Update()

			if err != nil {
				cc.ResponseForSQLError(w, err)
			}

			if rowsAffected > 0 {
				showURL := route.Name2URL("categories.show", "id", id)
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "您没有做任何更改！")
			}
		} else {
			// 校验不通过
			view.Render(w, view.D{
				"Category": _category,
				"Errors":   errors,
			}, "categories.edit")
		}
	}
}
