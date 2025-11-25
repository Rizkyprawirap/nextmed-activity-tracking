package apimodelsclient

import (
	"github.com/google/uuid"
)

type (
	InsertClientRequest struct {
		Name   string
		Email  string
		ApiKey string
	}

	InsertClientResponse struct {
		ClientID uuid.UUID
		Name     string
		Email    string
		ApiKey   string
	}

	InsertLogRequest struct {
		ClientID uuid.UUID
		APIKey   string
		IP       string
		Endpoint string
	}

	GetClientByAPIKeyRequest struct {
		APIKey string
	}

	GetClientByAPIKeyResponse struct {
		ClientID uuid.UUID
		Name     string
		Email    string
		ApiKey   string
	}
)
