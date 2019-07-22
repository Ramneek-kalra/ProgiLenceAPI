package exam

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const (
	examDataid        = "EXAMDATA"
	examUserDetailsid = "EXAMUSERDETAILS"
	examCompletedid   = "EXAMCOMPLETED"
)

func SaveExamData(c registercontext.Context, code int, etype string) error {
	ctx := c.AppEngineCtx
	exptime := time.Now().Add(5 * time.Hour)
	key := datastore.NewKey(ctx, examDataid, strconv.Itoa(code), 0, nil)
	exam := &ExamData{
		Code:       code,
		EType:      etype,
		ExpiryDate: exptime.Unix(),
	}
	_, err := datastore.Put(ctx, key, exam)
	return err
}

func GetExamData(c registercontext.Context, code int) (*ExamData, error) {
	key := datastore.NewKey(c.AppEngineCtx, examDataid, strconv.Itoa(code), 0, nil)
	var exam ExamData
	err := datastore.Get(c.AppEngineCtx, key, &exam)
	return &exam, err
}
func SaveExamUser(c registercontext.Context, code int) error {
	ctx := c.AppEngineCtx
	str := c.RegisterdUser.Guid + strconv.Itoa(code)
	key := datastore.NewKey(ctx, examUserDetailsid, str, 0, nil)
	exam := &ExamUserDetails{
		Code: code,
		Guid: c.RegisterdUser.Guid,
	}
	_, err := datastore.Put(ctx, key, exam)
	return err
}
func GetExamUser(c *gin.Context, str string) (*ExamUserDetails, error) {
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, examUserDetailsid, str, 0, nil)
	var exam ExamUserDetails
	err := datastore.Get(ctx, key, &exam)
	return &exam, err
}
func GetAllExam(c registercontext.Context) ([]ExamData, error) {
	ctx := c.AppEngineCtx
	var exam []ExamData
	q := datastore.NewQuery(examDataid)
	_, err := q.GetAll(ctx, &exam)
	return exam, err
}
func CheckExamCompleted(c registercontext.Context) ([]ExamData, error) {
	ctx := c.AppEngineCtx
	exp := time.Now().Unix()
	var exam []ExamData
	q := datastore.NewQuery(examDataid).Filter("ExpiryDate <=", exp)
	_, err := q.GetAll(ctx, &exam)
	return exam, err
}
func CheckExamCode(c registercontext.Context, code int) ([]ExamUserDetails, error) {
	ctx := c.AppEngineCtx
	var exam []ExamUserDetails
	q := datastore.NewQuery(examUserDetailsid).Filter("Code =", code)
	_, err := q.GetAll(ctx, &exam)
	return exam, err
}

func SaveExamComplete(c registercontext.Context, code int, id []string) error {
	ctx := c.AppEngineCtx
	key := datastore.NewKey(ctx, examCompletedid, strconv.Itoa(code), 0, nil)
	exam := &ExamCompleted{
		Code: code,
		Guid: id,
	}
	_, err := datastore.Put(ctx, key, exam)
	return err

}
func GetExamComplete(c registercontext.Context, code int) (*ExamCompleted, error) {
	ctx := c.AppEngineCtx
	key := datastore.NewKey(ctx, examCompletedid, strconv.Itoa(code), 0, nil)
	var exam ExamCompleted
	err := datastore.Get(ctx, key, &exam)
	return &exam, err
}
func DeleteExamUser(c registercontext.Context, str string) error {
	ctx := c.AppEngineCtx
	key := datastore.NewKey(ctx, examUserDetailsid, str, 0, nil)
	err := datastore.Delete(ctx, key)
	return err
}
