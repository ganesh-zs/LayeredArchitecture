package author

import (
	"LayeredArchitecture/models"
	store "LayeredArchitecture/store/author"
	"errors"
	"reflect"
	"testing"
)

func Test_ReadBook(t *testing.T) {
	testCases := []struct {
		id        int
		expOutput models.Author
		expError  error
	}{
		{16, models.Author{Id: 16, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Fiction"}, nil},
		{1000, models.Author{}, errors.New("entity not found")},
	}
	for i, tc := range testCases {
		bookStore := store.MockStore{}
		bookService := New(bookStore)
		actOutput, actError := bookService.ReadAuthor(tc.id)
		if !reflect.DeepEqual(actOutput, tc.expOutput) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expOutput, actOutput)
		}
		if !reflect.DeepEqual(actError, tc.expError) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expError, actError)
		}
	}
}
