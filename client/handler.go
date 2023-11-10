package client

import (
	"log"
	"github.com/colber/go-sdk/models"
)


type Handler struct {
	
}

func NewHandler() (*Handler, error) {
}

Find(filter map[string][]string) ([]*models.File,error)
Get(id string) (*models.File,error)
Upload(file *models.File) (*models.File,error)
Download(chunk *models.Chunk) (*models.Chunk,error)
Delete(file *models.File) (*models.File,error)


func (srv *Handler) Find(filter map[string][]string) ([]*models.File,error){
	log.Println("Find")
	out :=[]*models.File{}
	return out,nil
}

func (srv *Handler) Get(id string) (*models.File,error){
	out :=&models.File{}
	return out,nil
}

func (srv *Handler) Upload(file *models.File) (*models.File,error){
	out :=&models.File{}
	return out,nil
}

func (srv *Handler) Download(chunk *models.Chunk) (*models.Chunk,error){
	out :=&models.Chunk{}
	return out,nil
}

func (srv *Handler) Delete(file *models.File) (*models.File,error){
	out :=&models.File{}
	return out,nil
}