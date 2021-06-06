package jwts

///https://www.bookstack.cn/read/studyiris-examples/8194b1145f61e609.md
import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SecretKey = "ABCDEFGHI"
const Iss = "iris"

type JwtToken struct {
}

///new token
func NewToken(username string, userId string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"iss":      Iss,                                                      //issuer
		"iat":      time.Now().Unix(),                                        //Issued At
		"jti":      "9527",                                                   //JWT ID
		"exp":      time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //expiration time)
	})
	tokenString, _ := token.SignedString([]byte(SecretKey))
	return tokenString
}

/// parse token
func ParseToken(token interface{}) (username string, userId string) {
	jwtInfo := token.(*jwt.Token)
	usern := jwtInfo.Claims.(jwt.MapClaims)["username"].(string)
	uid := jwtInfo.Claims.(jwt.MapClaims)["userId"].(string)
	return usern, uid

}
