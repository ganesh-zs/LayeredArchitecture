package author

import "LayeredArchitecture/models"

type Store interface {
	ReadAuthor(id int) (models.Author, error)
}
