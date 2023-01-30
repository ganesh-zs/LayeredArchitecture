package book

import (
	"LayeredArchitecture/models"
	"LayeredArchitecture/store/book"
	"errors"
	"reflect"
)

type service struct {
	store book.Store
}

func New(s book.Store) service {
	return service{store: s}
}

func (s service) ReadBook(id int) (models.Book, error) {
	return s.store.ReadBook(id)
}

func (s service) GetAll(isbn int, title string, includeAuthor bool) ([]models.Book, error) {
	return []models.Book{}, nil
}

func (s service) Create(b models.Book) (models.Book, error) {
	return models.Book{}, nil
}

type Mockservice struct {
	store book.Store
}

func (s Mockservice) Create(book models.Book) (models.Book, error) {
	if reflect.DeepEqual(book, models.Book{0, 3469, "Money", "Fiction", "RELX", 2002, 1, nil}) {
		return models.Book{1, 3469, "Money", "Fiction", "RELX", 2002, 1, &models.Author{}}, nil
	}
	if reflect.DeepEqual(book, models.Book{0, 3469, "Money", "Fiction", "Harry", 2002, 1, nil}) {
		return models.Book{}, errors.New("bad request")
	}
	return models.Book{}, nil
}

func (s Mockservice) ReadBook(id int) (models.Book, error) {
	if id == 16 {
		return models.Book{Id: 16, ISBN: 778, Title: "Money", Genre: "Fiction", Publication: "RELX", YearOfPublication: 2002, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "ab", LastName: "Sukhla", PenName: "Natraj", DateOfBirth: "2000-09-01", Genre: "Comic"}}, nil
	}
	if id == 1000 {
		return models.Book{}, errors.New("entity not found")
	}
	return models.Book{}, errors.New("internal server error")
}

func (s Mockservice) GetAll(isbn int, title string, includeAuthor bool) ([]models.Book, error) {
	if isbn == 1 && includeAuthor == false {
		return []models.Book{{Id: 1, ISBN: 1, Title: "Hai", Genre: "Comic", Publication: "RELX", YearOfPublication: 2000, AuthorId: 1}}, nil
	}
	if title == "Hai" && includeAuthor == true {
		return []models.Book{{Id: 1, ISBN: 1, Title: "Hai", Genre: "Comic", Publication: "RELX", YearOfPublication: 2000, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Comic"}}}, nil
	}
	if isbn == 1 && title == "Hai" && includeAuthor == true {
		return []models.Book{{Id: 1, ISBN: 1, Title: "Hai", Genre: "Comic", Publication: "RELX", YearOfPublication: 2000, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Comic"}}}, nil
	}
	if isbn == 1000 && title == "Hello" && includeAuthor == true {
		return []models.Book{}, errors.New("no entity found")
	}
	if isbn == 2 && title == "NoTitle" && includeAuthor == true {
		return []models.Book{}, errors.New("no entity found")
	}
	return []models.Book{}, nil
}
