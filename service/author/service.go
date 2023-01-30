package author

import (
	"LayeredArchitecture/models"
	store "LayeredArchitecture/store/author"
)

type service struct {
	store store.Store
}

func New(s store.Store) service {
	return service{store: s}
}

func (s service) ReadAuthor(id int) (models.Author, error) {
	return s.store.ReadAuthor(id)
}
