package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// note: this should be put in .env but to simplify pocess installation for this testing i use it directly
var jwtKey = []byte("n12u3hc1c38y389cy39n18cy231y29312y983cn19823cy218y3981cy2398vy1293y891v23")

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	return claims, nil
}

func GetUserIDFromJWT(tokenString string) (primitive.ObjectID, error) {
	// get user id based on the jwt

	claims, _ := VerifyToken(tokenString)
	userObjID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return userObjID, nil
}
