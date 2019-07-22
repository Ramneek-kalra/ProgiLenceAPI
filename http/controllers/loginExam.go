package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/exam"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/common"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/register"
	"github.com/gin-gonic/gin"
)

func LoginExam(c *gin.Context) {
	var loginrequest *register.LoginDetails
	err := c.BindJSON(&loginrequest)
	if err != nil || loginrequest.Emailid == "" || loginrequest.Password == "" {
		c.JSON(http.StatusBadRequest, Reply{"status": false, "error_message": "email or password cant be empty"})
		return
	}
	user, err := register.GetUser(c, loginrequest.Emailid)
	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			c.JSON(http.StatusOK, Reply{"status": false, "error_message": "no user found"})
		} else {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "Internal server error"})
		}
		return
	}
	accesstable, err := register.Getaccess(c, user.Guid)
	if err != nil {
		c.JSON(http.StatusOK, Reply{"status": false, "error_message": "can't get claim "})
		return
	}
	passtable, err := register.GetPassword(c, user.Guid)
	if err != nil {
		c.JSON(http.StatusOK, Reply{"status": false, "error_message": "can't get password"})
		return
	}
	check := common.CheckPasswordHash(loginrequest.Password, passtable.Password)
	if check == false {
		c.JSON(http.StatusOK, Reply{"status": false, "error_message": "pass match fail"})
		return
	}
	if accessToken, err := common.GenerateToken(user.Guid, user.Emailid, user.Firstname, user.Lastname, accesstable.Access.Api, accesstable.Access.Claim, loginrequest.Exam); err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "Internal Server error in access token", "err": err, "accessToken": accessToken, "user": user})
		return
	} else if loginrequest.Exam > 1 {
		str := user.Guid + strconv.Itoa(loginrequest.Exam)
		examuser, err := exam.GetExamUser(c, str)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "Internal Server error in getting exam user", "err": err})
			return
		}
		//exam token ke saath dena hain
		c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "login": Reply{"access_token": accessToken, "user_details": user, "claim": accesstable}, "exam": examuser})
		return
	} else {
		c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "data": Reply{"access_token": accessToken, "user_details": user, "claim": accesstable}})
		return
	}

}
