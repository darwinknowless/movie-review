package middleware

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var memToken *jwt.Token

func AuthMiddleware(c *gin.Context) {
	secret := os.Getenv("SECRET")
	tokenFromHeader := c.Request.Header.Get("Authorization")

	// proses verify token dari header authorization
	token, err := jwt.Parse(tokenFromHeader, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("method salah %e", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if token != nil && err == nil {
		payloads := token.Claims.(jwt.MapClaims)
		// logic apa yang akan dibuat
		// contoh: kalau role dari token yang akses api bukan admin
		// maka prosesnya di stop atau tidak mendapatkan akses untuk api tersebut
		if payloads["role"] != "admin" {
			c.Abort() // untuk menghentikan request
			c.JSON(401, gin.H{
				"message": "Unauth",
			})
		} else {
			c.Next() // untuk melanjutkan request
			memToken = token
		}
	} else {
		c.Abort()
		c.JSON(401, gin.H{
			"message": "Unauth",
		})
	}
}

// function pembantu untuk mengambil data dari payload
func GetPayloadData(data string) string {
	payloads := memToken.Claims.(jwt.MapClaims)
	return payloads[data].(string)
}
