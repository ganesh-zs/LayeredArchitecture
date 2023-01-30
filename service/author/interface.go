package author

import "LayeredArchitecture/models"

type Service interface {
	ReadAuthor(id int) (models.Author, error)
}
