package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lestrrat-go/jwx/jwk"
)

func ValidateAuthToken(tokenString string) (bool, error) {
	jwks, err := jwk.Fetch(context.Background(), os.Getenv("AUTHENTIK_JWKS_URL"))
	if err != nil {
		return false, fmt.Errorf("failed to fetch JWKs: %v", err)
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
		return false, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}
	return true, nil
}
