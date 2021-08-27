package view

import (
	"go-blog/pkg/logger"
	"go-blog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

func Render(w io.Writer, data interface{}, tmplFiles ...string) {
	// --- 加载模板 ---

	// 设置模板相对路径
	viewDir := "resources/views/"

	// 遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tmplFiles {
		tmplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".tmpl"
	}

	// 所有布局模板文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "/layouts/*.tmpl")
	logger.LogError(err)

	// 在 Slice 里新增目标文件
	allFiles := append(layoutFiles, tmplFiles...)

	// 解析模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
		}).
		ParseFiles(allFiles...)
	logger.LogError(err)

	// 渲染模板，将所有文章的数据传输进去
	tmpl.ExecuteTemplate(w, "app", data)
}
