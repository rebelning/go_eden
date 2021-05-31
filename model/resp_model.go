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
