package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateTokenPair(user_id uint64, passwordHash string) (map[string]string, error) {
	tClaims := jwt.MapClaims{}
	tClaims["authorized"] = true
	tClaims["user_id"] = user_id
	tClaims["exp"] = time.Now().Add(time.Second * 20).Unix() //Token expires after 3 hours
	tWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, tClaims)
	token, err := tWithClaims.SignedString([]byte(os.Getenv("API_SECRET")))

	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = user_id
	rtClaims["secret"] = GenerateHmacForHashPassword(passwordHash)
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

func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, nil
}

func GenerateHmacForHashPassword(hashPassword string) string {
	secret := os.Getenv("API_SECRET")
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(hashPassword))
	expectedMAC := mac.Sum(nil)

	return hex.EncodeToString(expectedMAC)
}

func EqualsHmacForHashPassword(encodedShaHashPassword, hashPassword string) error {
	hexDecoded, err := hex.DecodeString(encodedShaHashPassword)
	if err != nil {
		return err
	}

	secret := os.Getenv("API_SECRET")
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(hashPassword))
	expectedMAC := mac.Sum(nil)

	if ok := hmac.Equal(hexDecoded, expectedMAC); !ok {
		return errors.New("Password was changed")
	}

	return nil
}
