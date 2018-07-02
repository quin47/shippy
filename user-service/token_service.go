package main

import "github.com/dgrijalva/jwt-go"
import (
	pb "github.com/quin47/shippy/user-service/proto/user"
	"time"
)

var(
	key =[]byte("this Is keY ")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}


type Authable interface {
	Decode(token string)(*CustomClaims,error)
	Encode(user *pb.User)(string,error)
}
type TokenService struct {
	repo Repository
}



func (s *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if claims,ok := token.Claims.(*CustomClaims);ok && token.Valid {
		return claims,nil
	}else {
		return  nil,err
	}
}

func (*TokenService) Encode(user *pb.User) (string, error) {
	expireAt := time.Now().Add(time.Hour * 72).Unix()

	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt:expireAt,
			Issuer:"go.micro.srv.user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(key)
}

