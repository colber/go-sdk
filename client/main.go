package client

import (
	"github.com/colber/go-sdk/models"
)

// Handler ...
type Client interface {
	Find(filter map[string][]string) ([]*models.File,error)
	Get(id string) (*models.File,error)
	Upload(file *models.File) (*models.File,error)
	Download(chunk *models.Chunk) (*models.Chunk,error)
	Delete(file *models.File) (*models.File,error)
}

// Create ...
func NewClient() (Client, error) {
	
	client, err:= NewHandler()
	if err!=nil{
		return nil,err
	}
	
	return client, nil
}