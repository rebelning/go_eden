package model

type UserInfo struct {
	UserId      string
	Username    string
	Mobile      string
	Address     string
	Password    string `json:"-"`
	AccessToken string
}
