package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

//jwt: 1. generate token, 2. melakukan validasi token(toke valid / tidak)

//1. set interface Service
type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

//2. set cetakan jwtService
type jwtService struct {
}

//5. set fungsi, return objek cetakan jwtService
func NewService() *jwtService {
	return &jwtService{}
}

//4. set secret key
var SECRET_KEY = []byte("BWASTARUP_s3cr3T_k3Y")

//3. implement kontrak GenerateToken milik jwtService
func (s *jwtService) GenerateToken(userID int) (string, error) {
	//-- payload :data yg akan disisipkan : userID
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	//set generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY) //assign signature verified to token
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

//6. implement ValidateToken()
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	//parsing token
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		//-- validasi apakah token method matching
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalite token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
