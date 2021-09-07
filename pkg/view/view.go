package view

import (
	"go-blog/app/models/category"
	"go-blog/app/models/user"
	"go-blog/pkg/auth"
	"go-blog/pkg/flash"
	"go-blog/pkg/logger"
	"go-blog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

// Render 渲染通用视图
func Render(w io.Writer, data D, tmplFiles ...string) {
	RenderTemplate(w, "app", data, tmplFiles...)
}

// RenderSimple 渲染简单的视图
func RenderSimple(w io.Writer, data D, tmplFiles ...string) {
	RenderTemplate(w, "simple", data, tmplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tmplFiles ...string) {

	// 通用模板数据
	data["isLogined"] = auth.Check()
	data["loginUser"] = auth.User
	data["flash"] = flash.All()
	// data["Categories"], err = category.All()
	data["Categories"], _ = category.All()
	data["Users"], _ = user.All()

	// 在 Slice 里新增目标文件
	allFiles := getTemplateFiles(tmplFiles...)

	// 解析模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).
		ParseFiles(allFiles...)
	logger.LogError(err)

	// 渲染模板，将所有文章的数据传输进去
	tmpl.ExecuteTemplate(w, name, data)
}

func getTemplateFiles(tmplFiles ...string) []string {

	// 设置模板相对路径
	viewDir := "resources/views/"

	// 遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tmplFiles {
		tmplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".tmpl"
	}

	// 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.tmpl")
	logger.LogError(err)

	return append(layoutFiles, tmplFiles...)
}
