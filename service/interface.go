package service

import "LayeredArchitecture/models"

type Service interface {
	GetByID(id int) (models.Author, error)
}
