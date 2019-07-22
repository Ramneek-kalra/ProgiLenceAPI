package register

type User struct {
	Emailid      string
	Firstname    string
	Middlename   string
	Lastname     string
	Phonenumber  string
	IsRegistered bool
	Guid         string
}
type Accesstype struct {
	Api   string
	Claim int64
}
type RegisterUserDetails struct {
	Emailid     string `json:"email"`
	Firstname   string `json:"firstname"`
	Middlename  string `json:"middlename"`
	Lastname    string `json:"lastname"`
	Phonenumber string `json:"phonenumber"`
	AccessApi   string `json:"api"`
	AccessClaim int64  `json:"claim"`
}
type UserPassword struct {
	Guid     string
	Password string
}
type LoginDetails struct {
	Emailid  string `json:"email"`
	Password string `json:"password"`
	Exam     int    `json:"exam"`
}
type Accessstore struct {
	Guid   string
	Access Accesstype
}
type DataStruct struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type Bulk struct {
	Count    int           `json:"count"`
	Data     *[]DataStruct `json:"bulk"`
	Examcode int           `json:"exam"`
}
