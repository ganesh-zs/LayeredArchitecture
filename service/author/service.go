package author

import (
	"LayeredArchitecture/models"
	"LayeredArchitecture/store"
)

type service struct {
	store store.Store
}

func New(s store.Store) service {
	return service{store: s}
}

func (s service) GetByID(id int) (models.Author, error) {
	return s.store.GetByID(id)
}
