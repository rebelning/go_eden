package routes

import (
	"go_eden/dao"
	"go_eden/model"
	"go_eden/service"
	"go_eden/web/controllers"
	"go_eden/web/middleware/jwts"

	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/v12/context"
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

	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		///error handler
		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			}

			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.JSON(model.Response{
				ResCode: ctx.GetStatusCode(),
				Message: err.Error(),
			})
		},
		ValidationKeyGetter: func(token *jwtmiddleware.Token) (interface{}, error) {
			return []byte(jwts.SecretKey), nil
		},
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwtmiddleware.SigningMethodHS256,
	})

	menu := eden.Party("/config")
	menu.Router.Use(jwtHandler.Serve)

	configDao := dao.NewConfigDao()
	configService := service.NewConfigService(configDao)
	menu.Register(configService)
	menu.Handle(new(controllers.ConfigController))

}
