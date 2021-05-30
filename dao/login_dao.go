package dao

import "go_eden/model"

///Login dao interface
type LoginDao interface {
	GetLogin(username string) model.UserInfo
}

type loginDao struct {
}

func NewLoginDao() LoginDao {
	return &loginDao{}
}

func (login loginDao) GetLogin(username string) model.UserInfo {
	///query db
	// var u  model.UserInfo
	u := model.UserInfo{UserId: "1001", Username: "xiaoning", Password: "123456", Address: "BeiJin", Mobile: "260 300052"}

	return u
}
