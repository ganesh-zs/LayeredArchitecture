package book

import (
	"LayeredArchitecture/models"
	"database/sql"
	"errors"
	"reflect"
)

type store struct {
	db *sql.DB
}

func New(d *sql.DB) store {
	return store{db: d}
}

func (d store) Create(b models.Book) (models.Book, error) {
	return models.Book{}, nil
}

func (d store) ReadBook(id int) (models.Book, error) {
	return models.Book{}, nil
}

func (d store) GetAll(isbn int, title string, includeAuthor bool) ([]models.Book, error) {
	return []models.Book{}, nil
}

type MockStore struct {
}

func (d MockStore) Create(book models.Book) (models.Book, error) {
	if reflect.DeepEqual(book, models.Book{0, 3469, "Money", "Fiction", "RELX", 2002, 1, nil}) {
		return models.Book{1, 3469, "Money", "Fiction", "RELX", 2002, 1, &models.Author{}}, nil
	}
	if reflect.DeepEqual(book, models.Book{0, 3469, "Money", "Fiction", "Harry", 2002, 1, nil}) {
		return models.Book{}, errors.New("bad request")
	}
	return models.Book{}, nil
}

func (d MockStore) ReadBook(id int) (models.Book, error) {
	if id == 16 {
		return models.Book{Id: 16, ISBN: 778, Title: "Money", Genre: "Fiction", Publication: "RELX", YearOfPublication: 2002, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "ab", LastName: "Sukhla", PenName: "Natraj", DateOfBirth: "2000-09-01", Genre: "Comic"}}, nil
	}
	if id == 1000 {
		return models.Book{}, errors.New("entity not found")
	}
	return models.Book{}, errors.New("internal server error")
}

func (d MockStore) GetAll(isbn int, title string, includeAuthor bool) ([]models.Book, error) {
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
