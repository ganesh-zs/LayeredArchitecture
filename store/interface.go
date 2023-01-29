package store

import "LayeredArchitecture/models"

type Store interface {
	GetByID(id int) (models.Author, error)
}
