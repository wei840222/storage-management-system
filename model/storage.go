package model

type Storage struct {
	ID       string            `json:"id"`
	Type     string            `json:"type" binding:"required"`
	Config   map[string]string `json:"config" binding:"required"`
	Endpoint string            `json:"endpoint"`
	Status   string            `json:"status"`
}
