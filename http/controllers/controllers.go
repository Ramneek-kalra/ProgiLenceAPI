package controllers

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Accesstype struct {
	Api   string
	Claim int64
}
type Claims struct {
	Email     string
	Guid      string
	FirstName string
	LastName  string
	Access    Accesstype
	Examcode  int
	jwt.StandardClaims
}
type Reply map[string]interface{}

func Verifytoken(flag bool, check int64) gin.HandlerFunc {

	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		claims := &Claims{}
		bearerToken := strings.TrimSpace(strings.Split(authorizationHeader, "Bearer")[1])
		tkn, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("mysecrectkey"), nil
		})
		if !tkn.Valid {
			c.AbortWithError(50, err)
			//c.AbortWithStatus(50)
		}
		if err != nil {
			c.AbortWithError(40, err)
		}

		if flag {
			Getscope(c, claims.Access.Claim, claims.Access.Api, check)
		}
	}
}
func Getscope(c *gin.Context, a int64, api string, check int64) {
	apiarrayadmin := []string{"GetAllUsers", "xyz"}
	apiarrayuser := []string{"ChangePassword", "xyz"}
	switch a {
	case 0:
		if contains(apiarrayadmin, api) && check == 0 {
			return
		}
		c.AbortWithStatus(601)
	case 1:
		if contains(apiarrayuser, api) && check == 1 {
			return
		}
		c.AbortWithStatus(701)
	default:
		c.AbortWithStatus(801)
	}
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
