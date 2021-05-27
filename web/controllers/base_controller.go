package controllers

import (
	// "go_eden/logs"
	"go_eden/model"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gopkg.in/go-playground/validator.v9"
)

type BaseController struct {
}
type formValue func(string) string

type bodyValue func(interface{})

var validate *validator.Validate

// BeforeActivation called once before the server start
// and before the controller's registration, here you can add
// dependencies, to this controller and only, that the main caller may skip.
func (c *BaseController) BeforeActivation(b mvc.BeforeActivation) {
	// bind the context's `FormValue` as well in order to be
	// acceptable on the controller or its methods' input arguments (NEW feature as well).
	b.Dependencies().Add(func(ctx iris.Context) formValue { return ctx.FormValue })

}

func baseNewValidate(ctx iris.Context, param interface{}, validate *validator.Validate) (model.Response, error) {
	//validate = validator.New()
	// validate.RegisterStructValidation(fn, types)
	//var param model.AuthParams
	response := model.Response{ResCode: 200, Message: ""}
	if err := ctx.ReadJSON(&param); err != nil {
		// Handle error.
		// logs.Logger().Log.Error(err.Error())

		response := model.Response{ResCode: 201, Message: err.Error()}
		//return mvc.Response{Code: iris.StatusOK, Object: response}
		return response, err
	}

	// logs.Logger().Log.Info("baseValidate", zap.Any("path", ctx.Path()), zap.Any("param", param))
	// Returns InvalidValidationError for bad validation input,
	// nil or ValidationErrors ( []FieldError )
	err := validate.Struct(param)
	if err != nil {
		// logs.Logger().Log.Info("baseValidate", zap.Any("error", err.Error()))
		// This check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			// ctx.StatusCode(iris.StatusInternalServerError)
			// ctx.WriteString(err.Error())
			// logs.Logger().Log.Error(err.Error())

			response := model.Response{ResCode: iris.StatusInternalServerError, Message: err.Error()}
			//return mvc.Response{Code: iris.StatusOK, Object: response}
			return response, err
		}

		// logs.Logger().Log.Error(err.Error())
		response := model.Response{ResCode: 201, Message: err.Error()}
		//return mvc.Response{Code: iris.StatusOK, Object: response}
		return response, err
	}

	return response, err
}
