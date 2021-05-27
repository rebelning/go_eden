package routes

import (
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12/v12/v12/mvc"
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
	// //wbd service
	// wbd := app.Party("/wbd")
	// ///

	// //casbin middleware
	// menu := wbd.Party("/cms/menu", casbinMiddleware.ServeHTTP)
	// configDao := dao.NewConfigDao(db)
	// configService := service.NewConfigService(configDao)
	// menu.Register(configService)
	// menu.Handle(new(controllers.MenuController))

	// 课程
	// /wbd/cms/course/

	//发布管理
	///wbd/cms/publish/

	///教师管理
	////wbd/cms/tec/
	///机构管理
	////wbd/cms/org/
	///推广管理
	////wbd/cms/recomand/
	///教务在线
	////wbd/cms/feature/
	///审核管理
	// /wbd/cms/check/
	///权限管理
	////wbd/cms/rule/
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
