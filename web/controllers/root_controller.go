package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

///
type RootController struct {
}

///
func (root RootController) Get(ctx iris.Context) mvc.Result {

	return mvc.Response{
		ContentType: "text/html",
		Text:        "<h1>welcome</h1>"}
}
