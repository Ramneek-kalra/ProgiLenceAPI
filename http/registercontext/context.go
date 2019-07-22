package registercontext

import (
	"context"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
)

type RegisterdUser struct {
	Email     string
	Guid      string
	FirstName string
	LastName  string
	Claim     Accesstype
	Examcode  int
}
type Accesstype struct {
	Api   string
	Claim int64
}
type Claims struct {
	Email     string
	Guid      string
	FirstName string
	LastName  string
	Claim     Accesstype
	Examcode  int
	jwt.StandardClaims
}
type Context struct {
	GinCtx        *gin.Context
	RegisterdUser RegisterdUser
	AppEngineCtx  context.Context
}

func VerifyJwtToken(c *gin.Context) (Context, error, string) {
	authorizationHeader := c.GetHeader("Authorization")
	claims := &Claims{}
	bearerToken := strings.TrimSpace(strings.Split(authorizationHeader, "Bearer")[1])
	jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return "mysecrectkey", nil
	})

	ctx := appengine.NewContext(c.Request)

	return Context{
		GinCtx: c,
		RegisterdUser: RegisterdUser{
			Email:     claims.Email,
			Guid:      claims.Guid,
			FirstName: claims.FirstName,
			LastName:  claims.LastName,
			Claim:     claims.Claim,
			Examcode:  claims.Examcode,
		},
		AppEngineCtx: ctx,
	}, nil, authorizationHeader
}
