package model

///user info
type UserInfo struct {
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	Mobile      string `json:"mobile"`
	Address     string `json:"address"`
	Password    string `json:"-"`
	AccessToken string `json:"accessToken"`
}

type MenuInfo struct {
	MenuId  string `json:"menuId"`
	Section string `json:"section"`
	Action  string `json:"action"`
}
