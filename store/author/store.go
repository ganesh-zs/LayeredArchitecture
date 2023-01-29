package author

import (
	"LayeredArchitecture/models"
	"database/sql"
	"errors"
	"log"
)

type store struct {
	db *sql.DB
}

type MockStore struct {
}

func New(db *sql.DB) store {
	return store{db: db}
}

func (d store) GetByID(id int) (models.Author, error) {
	var response models.Author
	row := d.db.QueryRow("SELECT * from Authors where id=?", id)
	err := row.Scan(&response.Id, &response.FirstName, &response.LastName, &response.PenName, &response.DateOfBirth, &response.Genre)
	if err != nil {
		log.Fatal(err)
		return response, errors.New("Entity doesnt exist")
	}
	return response, nil
}

func (d MockStore) GetByID(id int) (models.Author, error) {
	switch id {
	case 16:
		return models.Author{Id: 16, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Fiction"}, nil
	case 1000:
		return models.Author{}, errors.New("entity not found")
	default:
		return models.Author{}, errors.New("internal server error")
	}

}
