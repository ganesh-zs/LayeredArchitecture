package author

import (
	"LayeredArchitecture/models"
	store "LayeredArchitecture/store/author"
	"LayeredArchitecture/store/book"
	"errors"
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

type Mockservice struct {
	store book.Store
}

func (s Mockservice) ReadAuthor(id int) (models.Author, error) {

	if id == 16 {
		return models.Author{Id: 16, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Fiction"}, nil
	}
	if id == 1000 {
		return models.Author{}, errors.New("entity not found")
	}
	return models.Author{}, errors.New("internal server error")
}
