package register

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const (
	userid   = "USER"
	passid   = "PASSWORD"
	accessid = "ACCESS"
)

func GetUser(c *gin.Context, emailId string) (*User, error) {
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, userid, emailId, 0, nil)
	var user User
	err := datastore.Get(ctx, key, &user)
	return &user, err
}

func SaveUser(c *gin.Context, emailId string, guid string, firstName string, middleName string, lastName string, phone string) error {
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, userid, emailId, 0, nil)
	user := &User{
		Firstname:    firstName,
		Middlename:   middleName,
		Lastname:     lastName,
		IsRegistered: true,
		Guid:         guid,
		Phonenumber:  phone,
		Emailid:      emailId,
	}
	_, err := datastore.Put(ctx, key, user)
	return err
}
func SavePassword(c *gin.Context, password string, guid string) error {
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, passid, guid, 0, nil)
	pass := &UserPassword{
		Guid:     guid,
		Password: password,
	}
	_, err := datastore.Put(ctx, key, pass)
	return err

}
func GetPassword(c *gin.Context, guid string) (*UserPassword, error) {
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, passid, guid, 0, nil)
	var passtable UserPassword
	err := datastore.Get(ctx, key, &passtable)
	return &passtable, err
}
func Getaccess(c *gin.Context, guid string) (*Accessstore, error) {
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, accessid, guid, 0, nil)
	var claimstore Accessstore
	err := datastore.Get(ctx, key, &claimstore)
	return &claimstore, err
}
func Saveaccess(c *gin.Context, api string, apicliam int64, guid string) error {
	var access Accesstype
	access.Api = api
	access.Claim = apicliam
	ctx := appengine.NewContext(c.Request)
	key := datastore.NewKey(ctx, accessid, guid, 0, nil)
	acs := &Accessstore{
		Guid:   guid,
		Access: access,
	}
	_, err := datastore.Put(ctx, key, acs)
	return err

}
