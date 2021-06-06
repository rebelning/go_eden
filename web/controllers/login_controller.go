package controllers

import (
	"go_eden/model"
	"go_eden/service"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gopkg.in/go-playground/validator.v9"
)

var (
	USERNAME = "username"
	ISLOGIN  = "isLogin"
)

///login controller
type LoginController struct {
	BaseController
	Ctx iris.Context
	//Session *sessions.Session
	Service service.LoginService
}

func (lc LoginController) PostLogin() mvc.Result {
	var loginInfo model.LoginInfo
	validate = validator.New()
	validate.RegisterStructValidation(LoginStructLevelValidation, model.LoginInfo{})
	response, err := baseNewValidate(lc.Ctx, &loginInfo, validate)
	if err != nil {
		return mvc.Response{Code: iris.StatusCreated, Object: response}
	}
	resp := lc.Service.GetLogin(loginInfo.Username, loginInfo.Password)
	response.ResCode = resp.ResCode
	response.Message = resp.Message
	response.Data = resp.Data

	return mvc.Response{Code: iris.StatusOK, Object: response}
}
func LoginStructLevelValidation(sl validator.StructLevel) {
	user := sl.Current().Interface().(model.LoginInfo)
	if len(user.Username) == 0 {
		sl.ReportError(user.Username, "Username", "username", "username", "")
		// sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}
	if len(user.Password) == 0 {
		sl.ReportError(user.Password, "Password", "Password", "Password", "")
	}

}
