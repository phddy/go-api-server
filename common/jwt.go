package common

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/phddy/go-api-server/model"
	"strings"
	"time"
)

type AuthTokenClaims struct {
	Tid  string   `json:"tid"`
	Uid  int32    `json:"uid"`
	Name string   `json:"name"`
	Role []string `json:"role"`
	jwt.RegisteredClaims
}

var expireTime time.Duration

func init() {
	expireTime = time.Duration(config.Jwt.Expire) * time.Hour
}

func CreateToken(user *model.User) (string, error) {
	tid := uuid.NewString()
	at := AuthTokenClaims{
		Tid:  tid,
		Uid:  user.Id,
		Name: user.Name,
		Role: []string{"user", "admin"},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(expireTime).Truncate(time.Microsecond)},
		},
	}

	value, _ := json.Marshal(at)
	print(value)
	//rdb.Set(ctx, fmt.Sprintf("jwt:%s", tid), value, expireTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &at)
	return token.SignedString([]byte(config.Jwt.SecretKey))
}

func ValidateToken(authorization string) error {
	tokens := strings.Split(authorization, " ")
	claims := &AuthTokenClaims{}

	token, err := jwt.ParseWithClaims(tokens[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.SecretKey), nil
	})

	if token.Valid == false || err != nil {
		return &UnauthorizedError{
			Message: "token is invalid.",
		}
	}

	//if err := rdb.Get(ctx, fmt.Sprintf("jwt:%s", claims.Tid)).Err(); err != nil {
	//	return &UnauthorizedError{
	//		Message: "token does not exist.",
	//	}
	//}

	return nil
}

func RemoveToken(authorization string) error {
	tokens := strings.Split(authorization, " ")
	claims := &AuthTokenClaims{}
	jwt.ParseWithClaims(tokens[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Jwt.SecretKey), nil
	})

	//if err := rdb.Del(ctx, fmt.Sprintf("jwt:%s", claims.Tid)).Err(); err != nil {
	//	return &UnauthorizedError{
	//		Message: "token does not exist.",
	//	}
	//}

	print(claims)
	return nil

}
