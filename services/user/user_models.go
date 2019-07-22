package user

type User struct {
	Emailid      string
	Firstname    string
	Middlename   string
	Lastname     string
	Phonenumber  string
	IsRegistered bool
	Guid         string
}
type UserPassword struct {
	Guid     string
	Password string
}
type LoginDetails struct {
	Emailid  string `json:"email"`
	Password string `json:"password"`
}
type ChangePasswordDetails struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}
