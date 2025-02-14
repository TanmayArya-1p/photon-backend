package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"os"
	"photon-backend/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

func ValidateAuthToken(tokenString string) (bool, error, models.UnpackedAccessToken) {
	jwks, err := jwk.Fetch(context.Background(), os.Getenv("AUTHENTIK_JWKS_URL"))
	if err != nil {
		return false, fmt.Errorf("failed to fetch JWKs: %v", err), models.UnpackedAccessToken{}
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("token missing 'kid' field")
		}
		key, found := jwks.LookupKeyID(kid)
		if !found {
			return nil, fmt.Errorf("key with kid %s not found", kid)
		}
		var rsaKey rsa.PublicKey
		if err := key.Raw(&rsaKey); err != nil {
			return nil, fmt.Errorf("failed to parse RSA public key: %v", err)
		}
		return &rsaKey, nil
	})
	if err != nil {
		return false, fmt.Errorf("error parsing token: %v", err), models.UnpackedAccessToken{}
	}
	if !token.Valid {
		return false, fmt.Errorf("invalid token"), models.UnpackedAccessToken{}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false, fmt.Errorf("failed to extract claims"), models.UnpackedAccessToken{}
	}

	sub, ok1 := claims["sub"].(string)
	acr, ok2 := claims["acr"].(string)
	email, ok3 := claims["email"].(string)
	if !ok1 || !ok2 || !ok3 {
		return false, fmt.Errorf("claim not found or not a string"), models.UnpackedAccessToken{}
	}

	return true, nil, models.UnpackedAccessToken{
		UID:   sub,
		Acr:   acr,
		Email: email,
	}
}
