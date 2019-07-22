package controllers

import (
	"net/http"
	"time"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/exam"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"github.com/gin-gonic/gin"
)

func BuyExam(c *gin.Context) {
	var examDetails *exam.ExamData
	//var examUserDetails *exam.ExamUserDetails
	var examrequest *exam.RequestExamCode

	ctx, err, token := registercontext.VerifyJwtToken(c)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "token failed", "err": err, "token": token})
		return
	}
	err = c.BindJSON(&examrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, Reply{"status": false, "error_message": "exam code cant be empty"})
		return
	}
	examDetails, err = exam.GetExamData(ctx, examrequest.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in getting exam data"})

		return
	}
	exp := time.Now().Unix()
	if examDetails.ExpiryDate <= exp {
		c.JSON(http.StatusOK, Reply{"status": true, "error_message": "exam expers can't be buyed"})
		return
	}
	err = exam.SaveExamUser(ctx, examDetails.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in saving exam user data"})

		return
	}

	c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "exam type Buyed ": examDetails.EType, "code": examDetails.Code, "user": ctx.RegisterdUser})
}
