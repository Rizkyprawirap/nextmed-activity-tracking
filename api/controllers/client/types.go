package apicontrollersclient

import apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"

type (
	CreateClientRequest struct {
		Name  string
		Email string
	}

	CreateClientResponse struct {
		Data apidto.Client `json:"data"`
	}

	InsertLogRequest struct {
		APIKey   string
		IP       string
		Endpoint string
	}
)
