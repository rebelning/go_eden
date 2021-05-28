package routes

import (
	"go_eden/dao"
	"go_eden/service"
	"go_eden/web/controllers"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/mvc"
)

///
func AdminRoute(app *mvc.Application) {

	golog.Info("amdin route register...")

	// j := jwt.New(jwt.Config{
	// 	// Extract by "token" url parameter.
	// 	Extractor: jwt.FromParameter("token"),

	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
	// 		return []byte("My Secret"), nil
	// 	},
	// 	SigningMethod: jwt.SigningMethodHS256,
	// })
	///
	// sess := sessions.New(sessions.Config{Cookie: "session_cookie_name"})
	// enforcer := casbins.GetEnforcer()
	// casbinMiddleware := casbins.New(enforcer)

	// //
	// db := conn.MasterEngine()
	// adminDao := dao.NewAdminUserDao(nil)
	// adminService := service.NewAdminUserService(adminDao)
	// //eden service
	eden := app.Party("/eden")
	// ///

	// //casbin middleware
	// menu := wbd.Party("/cms/menu", casbinMiddleware.ServeHTTP)
	login := eden.Party("/auth")
	loginDao := dao.NewLoginDao()
	loginService := service.NewLoginService(loginDao)
	login.Register(loginService)
	login.Handle(new(controllers.LoginController))
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
