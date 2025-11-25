package pkgjwt

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"
	"github.com/golang-jwt/jwt/v5"
)

func (p *pkgJWT) Validate(ctx context.Context, request ValidateRequest) (response *ValidateResponse, err error) {

	parsed, err := jwt.ParseWithClaims(
		request.Token,
		&apidto.JWTClaims{},
		func(t *jwt.Token) (any, error) {
			return []byte(p.Secret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(*apidto.JWTClaims)
	if !ok {
		return nil, errors.New("failed to parse claims format")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("claims data is expired")
	}

	var adminData AdminData

	if err := json.Unmarshal(claims.Data, &adminData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal claims data: %w", err)
	}

	return &ValidateResponse{
		Claims: *claims,
	}, nil
}
