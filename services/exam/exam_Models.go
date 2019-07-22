package exam

type ExamData struct {
	Code       int
	EType      string
	ExpiryDate int64
}
type ExamCompleted struct {
	Code int
	Guid []string
}
type ExamUserDetails struct {
	Code int
	Guid string
}
type RequestExamData struct {
	EType string `json:"etype"`
}
type RequestExamCode struct {
	Code int `json:"code"`
}
