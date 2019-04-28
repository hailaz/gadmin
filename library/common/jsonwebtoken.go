package common

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/g/os/gtime"
)

type JsonWebToken struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const (
	TOKEN_SIGNING_KEY = "gadmin"
	TOKEN_TIMEOUT     = 60 * 60 * 24 //sec
)

// CreateJWT description
//
// createTime:2019年04月25日 10:28:51
// author:hailaz
func CreateJWT(username string) (string, error) {
	// Create the Claims
	claims := JsonWebToken{
		username,
		jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(gtime.Second() + TOKEN_TIMEOUT)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(TOKEN_SIGNING_KEY))

}

// PareseJWT description
//
// createTime:2019年04月25日 15:48:51
// author:hailaz
func PareseJWT(tokenString string) (*JsonWebToken, error) {
	// sample token is expired.  override time so it parses as valid
	token, err := jwt.ParseWithClaims(tokenString, &JsonWebToken{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(TOKEN_SIGNING_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JsonWebToken); ok && token.Valid {
		//fmt.Printf("%v %v", claims.Username, claims.StandardClaims.ExpiresAt)
		return claims, nil
	}
	return nil, errors.New("token parese fail")
}
