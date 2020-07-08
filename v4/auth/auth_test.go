package auth

import (
	"net/http"
	"testing"
)

func TestGenerateTokenPair(t *testing.T) {
	var err error
	tokens, err := GenerateTokenPair(1)

	if err != nil {
		t.Error(err)
	}

	if _, ok := tokens["access_token"]; !ok {
		t.Error("Access token is not exists")
	}

	if _, ok := tokens["refresh_token"]; !ok {
		t.Error("Refresh token is not exists")
	}

	payouts, err := GetPayoutsFromToken(tokens["access_token"])

	if err != nil {
		t.Error("Access token is not exists")
	}

	if typeVal, ok := payouts["type"]; !ok || typeVal != TypeTokenAccess {
		t.Error("type in access token not exists or not equal access type")
	}
	if _, ok := payouts["sub"]; !ok {
		t.Error("sub in access token not exists")
	}
	if _, ok := payouts["exp"]; !ok {
		t.Error("exp in access token not exists")
	}

	payouts, err = GetPayoutsFromToken(tokens["refresh_token"])

	if err != nil {
		t.Error("Refresh token is not exists")
	}

	if typeVal, ok := payouts["type"]; !ok || typeVal != TypeTokenRefresh {
		t.Error("type in refresh token not exists or not equal refresh type")
	}
	if _, ok := payouts["sub"]; !ok {
		t.Error("sub in refresh token not exists")
	}
	if _, ok := payouts["exp"]; !ok {
		t.Error("exp in refresh token not exists")
	}
}

func TestTokenValidFromRequest(t *testing.T) {
	tokens, _ := GenerateTokenPair(1)
	req, _ := http.NewRequest("get", "/", nil)
	req.Header.Add("Authorization", "Bearer "+tokens["access_token"])

	err := TokenValidFromRequest(req)
	if err != nil {
		t.Error("Token is not valid", err)
	}
}

func TestGetPayoutsFromToken(t *testing.T) {
	tokens, _ := GenerateTokenPair(1)
	_, err := GetPayoutsFromToken(tokens["access_token"])

	if err != nil {
		t.Error("Error test get payouts from token", err)
	}
}

func TestExtractToken(t *testing.T) {
	req, _ := http.NewRequest("get", "/", nil)
	req.Header.Add("Authorization", "Bearer test_token")

	v := ExtractToken(req)
	if v != "test_token" {
		t.Error("Incorrect token value")
	}
}

func TestGetUserIdByRequest(t *testing.T) {
	testUserId := uint64(1)
	tokens, _ := GenerateTokenPair(testUserId)
	req, _ := http.NewRequest("get", "/", nil)
	req.Header.Add("Authorization", "Bearer "+tokens["access_token"])

	userId, err := GetUserIdByRequest(req)
	if err != nil {
		t.Error("Get user id by request error", err)
	}

	if userId != testUserId {
		t.Errorf("Expected id %v upon receipt. Received %v", testUserId, userId)
	}
}
