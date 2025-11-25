package routesclient

type (
	CreateClientRequest struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}
	CreateLogRequest struct {
		APIKey   string `json:"api_key" binding:"required"`
		IP       string `json:"ip" binding:"required"`
		Endpoint string `json:"endpoint" binding:"required"`
	}
)
