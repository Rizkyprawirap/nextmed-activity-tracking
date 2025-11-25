package pkgjwt

import (
	"context"
	"encoding/json"
	"time"

	apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

	"github.com/golang-jwt/jwt/v5"
)

func (p *pkgJWT) Generate(ctx context.Context, payload any) (string, error) {
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	const defaultTTL = 7 * 24 * time.Hour
	now := time.Now()

	claims := &apidto.JWTClaims{
		Data: payloadJSON,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(defaultTTL)),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(p.Secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}
