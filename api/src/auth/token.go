package auth

import (
	"api/src/config"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func CreateToken(userId uint) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["user_id"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(c echo.Context) error {
	encodedToken, err := extractToken(c)
	if err != nil {
		return err
	}

	token, err := parseToken(encodedToken)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return echo.NewHTTPError(401, "Token invalid")
	}

	return nil
}

func ExtractUserId(c echo.Context) (uint, error) {
	encodedToken, err := extractToken(c)
	if err != nil {
		return 0, err
	}

	token, err := parseToken(encodedToken)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return uint(permissions["user_id"].(float64)), nil
	}

	return 0, echo.NewHTTPError(401, "Token invalid")
}

func extractToken(c echo.Context) (string, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return "", echo.NewHTTPError(401, "Token not found")
	}

	if len(strings.Split(token, " ")) != 2 {
		return "nil", echo.NewHTTPError(401, "Token invalid")
	}

	return strings.Split(token, " ")[1], nil
}

func parseToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(401, "Token invalid")
		}

		return []byte(config.SecretKey), nil
	})
}
