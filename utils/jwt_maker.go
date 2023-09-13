package utils

import (
	"os"
	"zsocket_service/model"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a new JWT token with the user's claims.
func GenerateToken(tenant model.TenantMaster) (string, error) {

	// Your JWT secret key should be stored securely.
	secretKey := []byte(os.Getenv("SECRETKEY"))

	// Create a new token with the user's claims
	claims := jwt.MapClaims{}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
