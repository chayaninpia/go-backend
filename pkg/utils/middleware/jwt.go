package middleware

import (
	"apps/pkg/utils/resx"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func AuthorizationValidate(c *gin.Context) {

	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	if err := validateToken(token); err != nil {
		resx.ErrorUnauthorize(c, err.Error(), nil)
		return
	}
}

func validateToken(token string) error {
	if token == "" {
		return fmt.Errorf("token is required")
	}
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(viper.GetString("jwt.signature")), nil
	})

	return err
}

func GenerateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
	})

	ss, err := token.SignedString([]byte(viper.GetString("jwt.signature")))
	if err != nil {
		return "", err
	}
	return ss, nil
}
