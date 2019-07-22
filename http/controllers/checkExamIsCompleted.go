package controllers

import (
	"net/http"
	"strconv"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/exam"
	"github.com/gin-gonic/gin"
)

func RowsShifting(c *gin.Context) {
	ctx, err, token := registercontext.VerifyJwtToken(c)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "token failed", "err": err, "token": token})
		return
	}
	examArray, err := exam.CheckExamCompleted(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in getting examarray", "err": err})
		return
	}
	for _, p := range examArray {

		examUserArrya, err := exam.CheckExamCode(ctx, p.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error or exam user array not found", "err": err, "examUserArrya": examUserArrya})
			return
		}
		examcom, err := exam.GetExamComplete(ctx, p.Code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error or data not found", "err": err})
			return
		}
		array := append(examcom.Guid)
		for _, a := range examUserArrya {
			array = append(array, a.Guid)
		}
		for _, a := range examUserArrya {
			str := a.Guid + strconv.Itoa(p.Code)
			err := exam.DeleteExamUser(ctx, str)
			if err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in deleting", "err": err})
				return
			}
		}

		c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "msg": array})
		err = exam.SaveExamComplete(ctx, p.Code, array)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error exam complete not save", "err": err})
			return
		}

	}
	c.JSON(http.StatusOK, Reply{"status": true, "error_message": "", "msg": "done ", "exam": examArray, "guid": ctx.RegisterdUser.Guid})

}
