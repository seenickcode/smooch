package smooch

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func (s *Smooch) NewAppAuthToken() (tokenString string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"scope": "app",
	})
	token.Header["kid"] = s.appKeyID
	tokenString, err := token.SignedString([]byte(s.appSecret))
	if err != nil {
		fmt.Errorf("couldn't sign Smooch JWT token: %s", err.Error())
	}
	return
}
