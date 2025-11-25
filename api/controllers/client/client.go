package apicontrollersclient

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	apidto "github.com/Rizkyprawirap/nextmed-activity-tracking/api/dto"
	apimodelsclient "github.com/Rizkyprawirap/nextmed-activity-tracking/api/models/client"
)

type controller struct {
	modelsClient apimodelsclient.IModelClient
}

func New(
	modelsClient apimodelsclient.IModelClient,
) IControllerClient {
	return &controller{
		modelsClient: modelsClient,
	}
}

func (c *controller) CreateClient(ctx context.Context, request CreateClientRequest) (response *CreateClientResponse, err error) {
	apiKey, err := generateAPIKey(32)
	if err != nil {
		return nil, fmt.Errorf("failed generating API key: %w", err)
	}

	client, err := c.modelsClient.InsertClient(ctx, apimodelsclient.InsertClientRequest{
		Name:   request.Name,
		Email:  request.Email,
		ApiKey: apiKey,
	})

	if err != nil {
		return nil, fmt.Errorf("failed inserting client: %w", err)
	}

	response = &CreateClientResponse{
		Data: apidto.Client{
			ClientID: client.ClientID,
			Name:     client.Name,
			Email:    client.Email,
			ApiKey:   client.ApiKey,
		},
	}

	return response, err
}

func (c *controller) InsertLog(ctx context.Context, request InsertLogRequest) error {

	client, err := c.modelsClient.GetClientByAPIKey(ctx, apimodelsclient.GetClientByAPIKeyRequest{
		APIKey: request.APIKey,
	})
	if err != nil {
		return fmt.Errorf("invalid api_key: %w", err)
	}

	logReq := apimodelsclient.InsertLogRequest{
		ClientID: client.ClientID,
		APIKey:   request.APIKey,
		IP:       request.IP,
		Endpoint: request.Endpoint,
	}

	if err := c.modelsClient.InsertLog(ctx, logReq); err != nil {
		return fmt.Errorf("failed inserting log: %w", err)
	}

	return nil
}

func generateAPIKey(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
