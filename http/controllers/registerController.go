package controllers

import (
	"net/http"
	"strings"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/common"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/register"
	"github.com/gin-gonic/gin"
)

func GetRegisterDetails(c *gin.Context) {
	var registeruserRequest *register.RegisterUserDetails
	err := c.BindJSON(&registeruserRequest)
	if err != nil || registeruserRequest.Emailid == "" {
		c.JSON(http.StatusBadRequest, Reply{"status": false, "error_message": "email cant be empty"})
		return
	}
	guid, err := common.GetGuid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "internal server error in guid"})
		return
	}
	tempPassword, err := common.GenerateTempPassword()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "internal server Error in temp"})
		return
	}
	hashpassword, err := common.HashPassword(tempPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "internal server Error in hasspassword"})
		return
	}

	if _, err := register.GetUser(c, registeruserRequest.Emailid); err != nil {

		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			if err := register.Saveaccess(c, registeruserRequest.AccessApi, registeruserRequest.AccessClaim, guid); err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in saveing claim", "err": err})
				return
			}
			if err := register.SaveUser(c, registeruserRequest.Emailid, guid, registeruserRequest.Firstname, registeruserRequest.Middlename, registeruserRequest.Lastname, registeruserRequest.Phonenumber); err != nil {

				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in save user"})
				return
			}
			if err := register.SavePassword(c, hashpassword, guid); err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in save password"})

				return
			}

		} else {

			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "internal server error"})
			return
		}

	}

	c.JSON(http.StatusOK, Reply{"status": true, "error_msg": "", "data": Reply{"guid": guid, "password": tempPassword}})
}
