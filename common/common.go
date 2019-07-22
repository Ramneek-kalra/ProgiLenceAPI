package common

import (
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func GetGuid() (string, error) {
	id, err := uuid.NewUUID()

	guid := id.String()
	return guid, err
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GenerateTempPassword() (string, error) {
	pass, err := password.Generate(10, 3, 0, false, false)
	if err != nil {
		return "", err
	}
	return pass, nil
}

type Accesstype struct {
	Api   string
	Claim int64
}
type Claims struct {
	Email     string
	Guid      string
	FirstName string
	LastName  string
	Access    Accesstype
	Examcode  int
	jwt.StandardClaims
}

func GenerateToken(guid string, emailid string, firstname string, lastname string, api string, apicliam int64, examcode int) (string, error) {
	var access Accesstype
	access.Api = api
	access.Claim = apicliam
	exptime := time.Now().Add(2 * time.Hour)
	claims := &Claims{
		Guid:      guid,
		Email:     emailid,
		FirstName: firstname,
		LastName:  lastname,
		Access:    access,
		Examcode:  examcode,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exptime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString([]byte("mysecrectkey"))
	if err != nil {
		return "", err
	}
	return tokenstr, nil
}
func GenerateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	x := 1000 + rand.Intn(8999)
	y := rand.Intn(9)
	x = y*10000 + x
	return x
}
