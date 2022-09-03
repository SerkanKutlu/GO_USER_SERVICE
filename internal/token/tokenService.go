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
	Role           string
	UserId         string
	StandardClaims jwt.StandardClaims
}

func (c Claims) Valid() error {
	return nil
}

func GetTokenService(jwtConfig *config.Jwt) *TokenService {
	return &TokenService{JwtConfiguration: jwtConfig}
}

func (ts *TokenService) GenerateToken(user *model.User) (*string, error) {
	exp := time.Now().Add(10 * time.Minute)
	claims := &Claims{StandardClaims: jwt.StandardClaims{
		ExpiresAt: exp.Unix(),
		Issuer:    ts.JwtConfiguration.Issuer,
		NotBefore: time.Now().Unix(),
	}, UserId: user.Id, Role: user.Role}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(ts.JwtConfiguration.Secret))
	if err != nil {
		return nil, err
	}
	return &tokenStr, nil
}
