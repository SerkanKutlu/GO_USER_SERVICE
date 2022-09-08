package tokenpkg

import (
	"github.com/golang-jwt/jwt"
	"time"
	"userService/config"
	"userService/model"
)

type TokenService struct {
	JwtConfiguration *config.Jwt
}
type Claims struct {
	Role   string
	UserId string
	Exp    int64 `json:"exp"`
	Nbf    int64 `json:"nbf"`
}

func (c Claims) Valid() error {
	return nil
}

func GetTokenService(jwtConfig *config.Jwt) *TokenService {
	return &TokenService{JwtConfiguration: jwtConfig}
}

func (ts *TokenService) GenerateToken(user *model.User) (*string, error) {
	exp := time.Now().Add(10 * time.Minute)
	claims := &Claims{
		Role:   user.Role,
		UserId: user.Id,
		Exp:    exp.Unix(),
		Nbf:    time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(ts.JwtConfiguration.Secret))
	if err != nil {
		return nil, err
	}
	return &tokenStr, nil
}
