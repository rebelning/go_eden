package main

import (
	stdContext "context"
	"encoding/json"
	"fmt"
	"go_eden/model"
	"go_eden/routes"
	"time"

	"go_eden/web/controllers"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	fmt.Print("go main...")
	app := newApp()
	app.Run(
		iris.Addr(":8080"),
		iris.WithoutInterruptHandler,
		iris.WithoutBodyConsumptionOnUnmarshal,
		configure)
}

func newApp() *iris.Application {

	// e := casbins.GetEnforcer()
	// casbinMiddleware := cm.New(e)
	app := iris.New()
	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。
	app.Logger().SetLevel("debug")

	app.Use(recover.New())

	app.Use(logger.New())
	//session
	sess := sessions.New(sessions.Config{
		// Cookie string, the session's client cookie name, for example: "mysessionid"
		//
		// Defaults to "irissessionid"
		Cookie: "token",
		// it's time.Duration, from the time cookie is created, how long it can be alive?
		// 0 means no expire.
		// -1 means expire when browser closes
		// or set a value, like 2 hours:
		Expires: time.Hour * 2,
		// if you want to invalid cookies on different subdomains
		// of the same host, then enable it.
		// Defaults to false.
		DisableSubdomainPersistence: false,
	})

	app.Use(sess.Handler()) //// session is always non-nil inside handlers now.
	// app.Use(casbinMiddleware.ServeHTTP)
	//控制器根路由路径"/"
	mvc.Configure(app, configureMvc)

	//logger
	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		Columns: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(requestLogger)

	// app.Run(
	// 	// Start the web server at localhost:8080
	// 	iris.Addr("localhost:8080"),
	// 	// skip err server closed when CTRL/CMD+C pressed:
	// 	iris.WithoutServerError(iris.ErrServerClosed),
	// 	// enables faster json serialization and more:
	// 	iris.WithOptimizations,
	// )

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		//关闭所有主机
		app.Shutdown(ctx)
	})

	//或捕获所有http错误:
	app.OnAnyErrorCode(requestLogger, func(ctx iris.Context) {
		body, err := ctx.GetBody()
		if err != nil {
			// logs.Logger().Log.Error(err.Error())
		}
		// logs.Logger().Log.Error(fmt.Sprintf("path=%s", ctx.Path()),
		// 	zap.String("params", string(body)))
		golog.Error("OnAnyErrorCode=", string(body))
		response := model.Response{ResCode: ctx.GetStatusCode(), Message: ctx.Path()}
		ctx.ContentType("application/json;charset=UTF-8")
		ctx.Writef(toJSON(response))

	})

	return app
}

///app configure
func configure(app *iris.Application) {
	app.Configure(
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}

///config mvc
func configureMvc(app *mvc.Application) {
	//mvc controller
	//root /
	root := app.Party("/")
	root.Handle(new(controllers.RootController))
	//login
	routes.AdminRoute(app)

}
func toJSON(obj interface{}) string {
	jsonByte, err := json.Marshal(obj)
	if err != nil {
		//	fmt.Println("error:" + err.Error())
		return err.Error()
	}

	return string(jsonByte)
}
