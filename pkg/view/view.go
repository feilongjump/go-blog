package view

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, name string, data interface{}) {
	// --- 加载模板 ---

	// 设置模板相对路径
	viewDir := "resources/views/"

	name = strings.Replace(name, ".", "/", -1)

	// 所有布局模板文件 Slice
	files, err := filepath.Glob(viewDir + "/layouts/*.tmpl")
	logger.LogError(err)

	// 在 Slice 里新增目标文件
	newFiles := append(files, viewDir+name+".tmpl")

	// 解析模板文件
	tmpl, err := template.New(name + ".tmpl").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).
		ParseFiles(newFiles...)
	logger.LogError(err)

	// 渲染模板，将所有文章的数据传输进去
	tmpl.ExecuteTemplate(w, "app", data)
}
