package auth

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/cfg"
)

var secret_key []byte

// creates a JWT based on secret key
func CreateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	// creates a token time limit
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	if secret_key == nil {
		secret_key = []byte(cfg.GetConfig().SecretKey)
	}

	tokenStr, err := token.SignedString(secret_key)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// checks the JWT in the header to grant access
func ValidateJWT(c *gin.Context) {
	if c.Request.Header["Token"] != nil {
		token, err := jwt.Parse(c.Request.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
			_, ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "not authorized",
				})
				c.Abort()
			}
			if secret_key == nil {
				secret_key = []byte(cfg.GetConfig().SecretKey)
			}
			return secret_key, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"not authorized": err.Error(),
			})
			c.Abort()
			return
		}

		if token.Valid {
			c.Next()
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "not authorized",
		})
		c.Abort()
	}
}
