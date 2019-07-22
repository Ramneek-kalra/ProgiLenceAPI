package controllers

import (
	"net/http"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/common"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/user"
	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	var passwordDetails *user.ChangePasswordDetails
	var passtable *user.UserPassword
	// var ctx *registercontext.Context
	// var user1 *user.User
	// if user==nil{
	// 	c.JSON("xyz")
	// }
	ctx, err, token := registercontext.VerifyJwtToken(c)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "token failed", "err": err, "token": token})
		return
	}
	er := c.BindJSON(&passwordDetails)
	if er != nil || passwordDetails.OldPassword == "" || passwordDetails.NewPassword == "" {
		c.JSON(http.StatusBadRequest, Reply{"status": false, "error_message": "fields can't be empty", "err": err, "op": passwordDetails.OldPassword, "np": passwordDetails.NewPassword})
		return
	}

	// email := ctx.RegisterdUser.Email
	// guid := ctx.RegisterdUser.Guid
	passtable, err = user.GetPassword(ctx)
	check := common.CheckPasswordHash(passwordDetails.OldPassword, passtable.Password)
	if check == false {
		c.JSON(http.StatusOK, Reply{"status": false, "error_message": "pass match fail"})
		return
	}
	pass, err := common.HashPassword(passwordDetails.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in hash password"})

		return
	}

	passtable.Password = pass
	err = user.UpdatePassword(ctx, passtable)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in update password"})

		return
	}
	c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "msg": "pasword changed ", "value": passwordDetails.NewPassword})
}
