package controllers

import (
	"net/http"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/common"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"

	"github.com/gin-gonic/gin"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/exam"
	"github.com/ProgiLence/Backend/ProgiLenceAPI/services/register"
)

type retunArray struct {
	Username string
	Password string
	Examcode int
	Error    bool
}

func BulkRegister(c *gin.Context) {

	var bulkRequest *register.Bulk
	var rearry *retunArray
	var retunarr []*retunArray
	ctx, err, token := registercontext.VerifyJwtToken(c)
	if err != nil {
		c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "token failed", "err": err, "token": token})
		return
	}
	err = c.BindJSON(&bulkRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, Reply{"status": false, "error_message": "fields can't be empty", "err": err})
		return
	}
	//c.JSON(http.StatusOK, Reply{"status": true, "error_msg":"", "data": bulkRequest})
	// for i := 0; i < bulkRequest.Count; i++ {
	// 	name := (*bulkRequest.Data)[i].Name
	// 	user, err := register.GetUser(c, name)
	// 	c.JSON(http.StatusOK, Reply{"status": true, "error_msg": "", "data": user, "erroe": err, "name": name, "emaild": user.Emailid})
	// }
	// return
	// access := ctx.RegisterdUser.Claim
	// if access.Claim == 5 { //code can be changed
	// 	c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "no right user for this api", "cliam": ctx.RegisterdUser})
	// 	return
	// }
	// if ctx.RegisterdUser.Examcode != bulkRequest.Examcode {
	// 	c.JSON(http.StatusNonAuthoritativeInfo, Reply{"status": false, "error_message": "no right user for this exam"})
	// 	return
	// }

	for i := 0; i < bulkRequest.Count; i++ {
		erro := false
		guid, err := common.GetGuid()
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "internal server error in guid"})
			return
		}
		name := (*bulkRequest.Data)[i].Name
		password := (*bulkRequest.Data)[i].Password
		if name == "" {
			name, err = common.GenerateTempPassword()
			if err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "internal server Error in genrating name "})
				return
			}
		}

		if password == "" {
			password, err = common.GenerateTempPassword()
			if err != nil {
				c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_message": "internal server Error in genrating password"})
				return
			}
		}
		hashpassword, err := common.HashPassword(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "internal server Error in hasspassword"})
			return
		}
		user, err := register.GetUser(c, name)

		if user.Emailid == name {
			//array of error
			erro = true
			goto end
		}

		if err := register.SaveUser(c, name, guid, "", "", "", ""); err != nil {

			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in save user"})
			return
		}

		if err := register.SavePassword(c, hashpassword, guid); err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in save password"})

			return
		}
		if err := register.Saveaccess(c, "", 1, guid); err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in saveing claim", "err": err})
			return
		}
		err = exam.SaveExamUser(ctx, bulkRequest.Examcode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{"status": false, "error_msg": "Internal Server Error in saving exam user data"})

			return
		}

	end:
		rearry = &retunArray{
			Username: name,
			Password: password,
			Examcode: bulkRequest.Examcode,
			Error:    erro,
		}
		retunarr = append(retunarr, rearry)

	}
	c.JSON(http.StatusOK, Reply{"status": true, "error_msg": "", "data": retunarr})
}
