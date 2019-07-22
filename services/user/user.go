package user

import (
	"github.com/ProgiLence/Backend/ProgiLenceAPI/http/registercontext"
	"google.golang.org/appengine/datastore"
)

const (
	userid = "USER"
	passid = "PASSWORD"
)

func GetUser(c registercontext.Context) (*User, error) {
	key := datastore.NewKey(c.AppEngineCtx, userid, c.RegisterdUser.Email, 0, nil)
	var usr User
	err := datastore.Get(c.AppEngineCtx, key, &usr)
	return &usr, err
}
func GetPassword(c registercontext.Context) (*UserPassword, error) {
	key := datastore.NewKey(c.AppEngineCtx, passid, c.RegisterdUser.Guid, 0, nil)
	var pass UserPassword
	err := datastore.Get(c.AppEngineCtx, key, &pass)
	return &pass, err
}

func UpdatePassword(c registercontext.Context, pass *UserPassword) error {
	key := datastore.NewKey(c.AppEngineCtx, passid, c.RegisterdUser.Guid, 0, nil)
	_, err := datastore.Put(c.AppEngineCtx, key, pass)
	return err
}
func UpdateUser(c registercontext.Context, user *User) error {
	key := datastore.NewKey(c.AppEngineCtx, userid, c.RegisterdUser.Email, 0, nil)
	_, err := datastore.Put(c.AppEngineCtx, key, user)
	return err
}
func GetAllUser(c registercontext.Context) ([]User, error) {
	var user []User
	q := datastore.NewQuery(userid)
	_, err := q.GetAll(c.AppEngineCtx, &user)
	return user, err
}
