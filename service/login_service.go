package service

import (
	"go_eden/dao"
	"go_eden/model"
)

type LoginService interface {
	GetLogin(username string, password string) model.Response
}

type loginService struct {
	dao dao.LoginDao
}

func NewLoginService(dao dao.LoginDao) LoginService {
	return &loginService{dao: dao}
}

///user login
func (ls *loginService) GetLogin(username string, password string) model.Response {
	user := ls.dao.GetLogin(username)

	if len(user.Username) == 0 {
		return model.Response{ResCode: 201, Message: "user is not exits"}
	}

	if user.Password == password {
		return model.Response{ResCode: 200, Message: "login success", Data: user}
	}

	return model.Response{ResCode: 201, Message: "password is Incorrect"}
}
