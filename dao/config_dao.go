package dao

import "go_eden/model"

type ConfigDao interface {
	GetMenuList() []model.MenuInfo
}

type configDao struct {
}

func NewConfigDao() ConfigDao {
	return &configDao{}

}

func (c *configDao) GetMenuList() []model.MenuInfo {

	var menuList []model.MenuInfo
	// /account/message
	var m model.MenuInfo

	m = model.MenuInfo{
		MenuId:  "2001",
		Section: "login",
		Action:  "/account/login",
	}
	menuList = append(menuList, m)
	///
	m = model.MenuInfo{
		MenuId:  "2002",
		Section: "message",
		Action:  "/account/message",
	}
	menuList = append(menuList, m)
	return menuList
}
