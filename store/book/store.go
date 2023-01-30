package book

import (
	"LayeredArchitecture/models"
	"database/sql"
	"errors"
)

type store struct {
	db *sql.DB
}

func New(d *sql.DB) store {
	return store{db: d}
}

func (d store) ReadBook(id int) (models.Book, error) {
	//var response models.Author
	//row := d.db.QueryRow("SELECT * from Books where id=?", id)
	//err := row.Scan(&response.Id, &response.FirstName, &response.LastName, &response.PenName, &response.DateOfBirth, &response.Genre)
	//if err != nil {
	//	log.Fatal(err)
	//	return response, errors.New("Entity doesnt exist")
	//}
	return models.Book{}, nil
}

type MockStore struct {
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
