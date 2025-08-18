package utils

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/library/config"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUser struct {
	Aud       string   `json:"aud"`
	Email     string   `json:"email"`
	Exp       int64    `json:"exp"`
	GivenName string   `json:"givenName"`
	Id        string   `json:"sub"`
	Iat       int64    `json:"iat"`
	Iss       string   `json:"iss"`
	Roles     []string `json:"roles"`
	Surname   string   `json:"surname"`
}

func DecodeAuthUserFromBearerToken(bearerToken string) (*AuthUser, *APIError) {
	spaceRegex := regexp.MustCompile(`\s+`)
	cleaned := spaceRegex.ReplaceAllString(bearerToken, " ")

	parts := strings.Split(cleaned, " ")
	if len(parts) < 2 {
		return nil, &APIError{
			Message: "invalid Authorization header format",
			Code:    http.StatusUnauthorized,
		}
	}

	token := parts[1]

	if token == "" {
		return nil, &APIError{
			Message: "empty token",
			Code:    http.StatusUnauthorized,
		}
	}

	jwtParts := strings.Split(token, ".")
	if len(jwtParts) != 3 {
		return nil, &APIError{
			Message: "invalid JWT format",
			Code:    http.StatusUnauthorized,
		}
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(jwtParts[1])
	if err != nil {
		return nil, &APIError{
			Message: "failed to decode JWT payload",
			Code:    http.StatusUnauthorized,
		}
	}

	var authUser AuthUser
	if err := json.Unmarshal(payloadBytes, &authUser); err != nil {
		return nil, &APIError{
			Message: "invalid JWT payload structure",
			Code:    http.StatusUnauthorized,
		}
	}

	return &authUser, nil
}

func DecodeJWT(token string) (AuthUser, error) {
	token = strings.TrimPrefix(token, "Bearer ")
	authUser, err := parseJWT(token)
	if err != nil {
		return AuthUser{}, errors.New("unauthorized: invalid or expired token")
	}

	return authUser, nil
}

func EncodeJWT(authUser AuthUser, expiration time.Duration) (string, error) {
	cfg := config.GetConfig()

	now := time.Now()
	authUser.Iat = now.Unix()
	authUser.Exp = now.Add(expiration).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aud":       authUser.Aud,
		"email":     authUser.Email,
		"exp":       authUser.Exp,
		"givenName": authUser.GivenName,
		"iat":       authUser.Iat,
		"iss":       authUser.Iss,
		"roles":     authUser.Roles,
		"sub":       authUser.Id,
		"surname":   authUser.Surname,
	})

	return token.SignedString([]byte(cfg.JWTSecretKey))
}

func GetAuthUser(authorizer map[string]interface{}) (*AuthUser, error) {
	authorizerJSON, err := json.Marshal(authorizer)
	if err != nil {
		return nil, errors.New("failed to marshal authorizer")
	}

	var authUser AuthUser
	err = json.Unmarshal(authorizerJSON, &authUser)
	if err != nil {
		return nil, errors.New("failed to unmarshal authorizer into AuthUser")
	}

	return &authUser, nil
}

func HasPermission(authUser *AuthUser, resourceOwnerID string) bool {
	if ContainsRole(authUser.Roles, "admin") {
		return true
	}
	if resourceOwnerID == "" {
		return false
	}
	return authUser.Id == resourceOwnerID
}

func mapClaimsToStruct(claims jwt.MapClaims) (AuthUser, error) {
	var authUser AuthUser
	bytes, err := json.Marshal(claims)
	if err != nil {
		return authUser, err
	}
	err = json.Unmarshal(bytes, &authUser)
	return authUser, err
}

func parseJWT(tokenStr string) (AuthUser, error) {
	cfg := config.GetConfig()
	var authUser AuthUser

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.JWTSecretKey), nil
	})
	if err != nil {
		return authUser, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return authUser, errors.New("token expired")
			}
		}
		return mapClaimsToStruct(claims)
	}

	return authUser, errors.New("invalid token")
}
