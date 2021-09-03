package controllers

import (
	"go-blog/app/models/user"
	"go-blog/app/requests"
	"go-blog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 初始化变量
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 表单验证
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 表单不通过 - 重新显示表单
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		// 验证通过 - 入库 - 跳转到首页
		//_user.Create()
		//
		//if _user.ID > 0 {
		//	fmt.Fprint(w, "插入成功，ID 为"+_user.GetStringID())
		//} else {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	fmt.Fprint(w, "创建用户失败，请联系管理员")
		//}
	}
}
