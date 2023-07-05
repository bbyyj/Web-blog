package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyCustomClaims struct {
	Username string
	UserId   uint32
	jwt.StandardClaims
}

var mySigningKey = []byte("web-homework")

func CreateToken(id uint32, userName string, expireDuration time.Duration) (string, error) {
	myClaims := &MyCustomClaims{
		Username: userName,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			Issuer:    "web-blog",
			Subject:   "myToken",
			ExpiresAt: time.Now().Add(expireDuration).Unix(),
		},
	}
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	myTokenStr, err := myToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return myTokenStr, nil
}

func VerifyToken(mytoken string) (string, uint32, bool) {
	if mytoken == "" {
		return "", 0, false
	}

	tok, err := jwt.ParseWithClaims(mytoken, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		Logger().Warning("ParseWithClaims error %v", err)
		return "", 0, false
	}

	if claims, ok := tok.Claims.(*MyCustomClaims); ok && tok.Valid {
		return claims.Username, claims.UserId, true
	} else {
		Logger().Warning("%v", err)
		return "", 0, false
	}

}
