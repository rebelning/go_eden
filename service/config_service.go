package service

import (
	"go_eden/dao"
	"go_eden/model"
)

type ConfigService interface {
	GetMenuList() model.Response
}

type configService struct {
	dao dao.ConfigDao
}

///
func NewConfigService(dao dao.ConfigDao) ConfigService {

	return &configService{dao: dao}
}

func (c *configService) GetMenuList() model.Response {
	response := model.Response{}
	menuList := c.dao.GetMenuList()

	response.Data = menuList

	return response

}
