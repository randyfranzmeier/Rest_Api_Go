package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "thisisasecret"

func GenerateToken(email string, userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func ValidToken(token string) (error, int64) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		//check if token was signed the correct way
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method!!!")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("unable to parse token"), 0
	}
	isValTok := parseToken.Valid
	if !isValTok {
		return errors.New("Invalid token"), 0
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)

	if !ok {
		return errors.New("invalid claims"), 0
	}
	//claims is under-the-hood map, so square bracket are used
	//email := claims["email"].(string)
	userID := int64(claims["userID"].(float64))
	return nil, userID
}
