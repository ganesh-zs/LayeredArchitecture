package book

import "LayeredArchitecture/models"

type Store interface {
	Create(b models.Book) (models.Book, error)
	ReadBook(id int) (models.Book, error)
	GetAll(isbn int, title string, includeAuthor bool) ([]models.Book, error)
}
