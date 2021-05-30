package routes

import (
	"go_eden/dao"
	"go_eden/service"
	"go_eden/web/controllers"

	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/mvc"
)

///
func Route(mvc *mvc.Application) {

	golog.Info("amdin route register...")

	// //eden service
	eden := mvc.Party("/eden")

	// ///

	// //casbin middleware
	// menu := wbd.Party("/cms/menu", casbinMiddleware.ServeHTTP)
	login := eden.Party("/auth")

	///
	loginDao := dao.NewLoginDao()
	loginService := service.NewLoginService(loginDao)
	login.Register(loginService)
	login.Handle(new(controllers.LoginController))

	menu := eden.Party("/menu")
	menu.Handle(new(controllers.AppController))
	///jwt token
	jwt := eden.Party("/jwt")
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwtmiddleware.Token) (interface{}, error) {
			return []byte("Secret"), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwtmiddleware.SigningMethodHS256,
	})
	jwt.Router.Use(jwtHandler.Serve)

	///
	// rule := wbd.Party("/cms/rule")
	///rule
	// ruleDao := dao.NewRuleDao(enforcer, db)
	// ruleService := service.NewRuleService(ruleDao)

	// rule.Register(ruleService)
	// rule.Handle(new(controllers.RuleController))

	////个人信息
	////wbd/cms/account/

}
