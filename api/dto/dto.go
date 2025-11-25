package apidto

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	Admin struct {
		ID        string    `db:"id"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		FullName  string    `db:"full_name"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}

	Client struct {
		ClientID uuid.UUID `json:"client_id"`
		Name     string    `json:"name"`
		Email    string    `json:"email"`
		ApiKey   string    `json:"api_key"`
	}

	Log struct {
		ClientID uuid.UUID `json:"client_id"`
		APIKey   string    `json:"api_key"`
		IP       string    `json:"ip"`
		Endpoint string    `json:"endpoint"`
	}

	Usage struct {
		Date  string
		Total int
	}

	TopClientUsage struct {
		ClientID string `json:"client_id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		APIKey   string `json:"api_key"`
		Total    int64  `json:"total"`
	}

	JWTClaims struct {
		Data json.RawMessage `json:"data"`
		jwt.RegisteredClaims
	}

	HTTPResponse struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    any    `json:"data,omitempty"`
		Error   string `json:"error,omitempty"`
	}
)
