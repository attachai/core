package setting

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	nuboJwt "github.com/nubo/jwt"
)

type jwtCtrl struct{}

/// แกะ token
func (j jwtCtrl) VerifyToken(rawToken string) string {
	var payload interface{}
	var careId string
	// rawToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJjYXJlSWQiOiJQb1RBWHRRQldnIn0.998AX3mvnYYfUQf6AZi12M1AmHnVUyQu5u9PlWOpHtM"

	token, ok := nuboJwt.ParseAndVerify(rawToken, "blueposh")
	if !ok {
		log.Fatal("Invalid token")
	} else {
		payload = token.ClaimSet["careId"]
		careId := fmt.Sprintf("%v", payload)
		// fmt.Println("Type", token.Header.Type)
		// fmt.Println("Algorithm", token.Header.Algorithm)
		// fmt.Println("Claim Set", token.ClaimSet)

		return careId
	}

	return careId
}

/// create token
func (j jwtCtrl) CreateToken(careId string) (string, error) {
	var err error

	secret := "blueposh"
	atClaims := jwt.MapClaims{}
	// atClaims["careId"] = "PoRcipekGZ"
	atClaims["careId"] = careId

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
