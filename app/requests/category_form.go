package requests

import (
	"github.com/thedevsaddam/govalidator"
	"go-blog/app/models/category"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	// 定制认证规则
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	// 定制错误消息
	message := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
		},
	}

	// 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      message,
	}

	return govalidator.New(opts).ValidateStruct()
}
