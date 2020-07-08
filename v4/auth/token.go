package auth

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const TypeTokenAccess = "access"
const TypeTokenRefresh = "refresh"

func GenerateTokenPair(userId uint64) (map[string]string, error) {
	tClaims := jwt.MapClaims{}
	tClaims["type"] = TypeTokenAccess
	tClaims["sub"] = userId
	tClaims["exp"] = time.Now().Add(time.Hour * 3).Unix() //Token expires after 3 hours
	tWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tClaims)
	token, err := tWithClaims.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["type"] = TypeTokenRefresh
	rtClaims["sub"] = userId
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() //Token expires after 24 hours
	rtWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rtWithClaims.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  token,
		"refresh_token": refreshToken,
	}, nil
}

func TokenValidFromRequest(r *http.Request) error {
	tokenString := ExtractToken(r)
	_, err := GetPayoutsFromToken(tokenString)

	return err
}

func GetPayoutsFromToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		err = errors.New("Token is invalid")
		return nil, err
	}
}

func ExtractToken(r *http.Request) string {
	//zachem yya ito napisal??
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func GetUserIdByRequest(r *http.Request) (uint64, error) {
	var err error
	tokenString := ExtractToken(r)
	payouts, err := GetPayoutsFromToken(tokenString)

	uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", payouts["sub"]), 10, 32)
	if err != nil {
		return 0, errors.New("Invalid token for getUserId")
	}

	return uid, nil
}
