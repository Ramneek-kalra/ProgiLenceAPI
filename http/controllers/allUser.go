package controllers

import (
	"net/http"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/user"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	ctx, err, token := registercontext.VerifyJwtToken(c)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "token failed", "err": err, "token": token})
		return
	}
	usr, err := user.GetAllUser(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error cant get users"})

		return
	}

	c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "user": usr})
	// for _, p := range usr {
	// 	//log.Infof(ctx, "%s %s, %d inches tall", p.FirstName, p.LastName, p.Height)
	// 	c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "user": p})
	// }
}
