package controllers

import (
	"go_eden/model"
	"go_eden/service"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"github.com/kataras/iris/sessions"
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
	Ctx     iris.Context
	Session *sessions.Session
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
	token := cToken(loginInfo.Username, loginInfo.Password)
	//header["Authorization"] = "bears "+tokenString
	lc.Ctx.Header("Authorization", "bears "+token)
	// lc.Ctx.SetCookieKV(USERNAME, token)
	// lc.Ctx.Request().Cookie(USERNAME)
	golog.Debug("token=" + token)

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
func cToken(username string, password string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"iss":      "eden",                                                   //issuer
		"iat":      time.Now().Unix(),                                        //Issued At
		"jti":      "9527",                                                   //JWT ID
		"exp":      time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //expiration time)
	})
	tokenString, _ := token.SignedString([]byte("Secret"))

	return tokenString
}
