package book

import "LayeredArchitecture/models"

type Store interface {
	ReadBook(id int) (models.Book, error)
}
