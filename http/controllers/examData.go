package controllers

import (
	"net/http"
	"strings"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/common"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/exam"
	"github.com/gin-gonic/gin"
)

func ExamDataSet(c *gin.Context) {
	var examrequest *exam.RequestExamData
	ctx, err, token := registercontext.VerifyJwtToken(c)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "token failed", "err": err, "token": token})
		return
	}
	err = c.BindJSON(&examrequest)
	if err != nil || examrequest.EType == "" {
		c.JSON(http.StatusBadRequest, Reply{"status": false, "error_message": "exam type cant be empty"})
		return
	}
	code := common.GenerateRandomNumber()
	if _, err := exam.GetExamData(ctx, code); err != nil {

		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			if err := exam.SaveExamData(ctx, code, examrequest.EType); err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in saveing claim", "err": err})
				return
			}
			var a []string
			if err := exam.SaveExamComplete(ctx, code, a); err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in saveing claim", "err": err})
				return
			}
		} else {

			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "internal server error"})
			return
		}

	}

	c.JSON(http.StatusOK, Reply{"status": true, "error_msg": "", "data": Reply{"code": code, "examtye": examrequest.EType}})
}

// func GetAllExam(c *gin.Context) {
// 	exm, err := exam.GetAllExam(c)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error cant get examdata"})

// 		return
// 	}

// 	c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "user": exm})

// }
