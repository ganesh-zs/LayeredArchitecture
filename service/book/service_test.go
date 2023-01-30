package book

import (
	"LayeredArchitecture/models"
	store "LayeredArchitecture/store/book"
	"errors"
	"reflect"
	"testing"
)

func Test_ReadBooks(t *testing.T) {
	testCases := []struct {
		isbn          int
		title         string
		includeauthor bool
		expOutput     []models.Book
		expError      error
	}{
		{isbn: 1, includeauthor: false, expOutput: []models.Book{{Id: 1, ISBN: 1, Title: "Hai", Genre: "Comic", Publication: "RELX", YearOfPublication: 2000, AuthorId: 1, BookAuthor: nil}}, expError: nil},
		{title: "Hai", includeauthor: true, expOutput: []models.Book{{Id: 1, ISBN: 1, Title: "Hai", Genre: "Comic", Publication: "RELX", YearOfPublication: 2000, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Comic"}}}, expError: nil},
		{isbn: 1, title: "Hai", includeauthor: true, expOutput: []models.Book{{Id: 1, ISBN: 1, Title: "Hai", Genre: "Comic", Publication: "RELX", YearOfPublication: 2000, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "Ganesh", LastName: "Manchi", PenName: "Natraj", DateOfBirth: "2001-09-01", Genre: "Comic"}}}, expError: nil},
		{isbn: 1000, title: "Hello", includeauthor: true, expOutput: []models.Book{}, expError: errors.New("no entity found")},
		{isbn: 2, title: "NoTitle", includeauthor: true, expOutput: []models.Book{}, expError: errors.New("no entity found")},
	}
	for i, tc := range testCases {
		bookStore := store.MockStore{}
		bookService := New(bookStore)
		actOutput, actError := bookService.GetAll(tc.isbn, tc.title, tc.includeauthor)
		if !reflect.DeepEqual(actOutput, tc.expOutput) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expOutput, actOutput)
		}
		if !reflect.DeepEqual(actError, tc.expError) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expError, actError)
		}
	}
}

func Test_CreateBook(t *testing.T) {
	testCases := []struct {
		input     models.Book
		expOutput models.Book
		expError  error
	}{
		{models.Book{0, 3469, "Money", "Fiction", "RELX", 2002, 1, nil}, models.Book{1, 3469, "Money", "Fiction", "RELX", 2002, 1, &models.Author{}}, nil},
		{models.Book{0, 3469, "Money", "Fiction", "Harry", 2002, 1, nil}, models.Book{}, errors.New("bad request")},
	}
	for i, tc := range testCases {
		bookStore := store.MockStore{}
		bookService := New(bookStore)
		actOutput, actError := bookService.Create(tc.input)
		if !reflect.DeepEqual(actOutput, tc.expOutput) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expOutput, actOutput)
		}
		if !reflect.DeepEqual(actError, tc.expError) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expError, actError)
		}
	}
}

func Test_ReadBook(t *testing.T) {
	testCases := []struct {
		id        int
		expOutput models.Book
		expError  error
	}{
		{16, models.Book{Id: 16, ISBN: 778, Title: "Money", Genre: "Fiction", Publication: "RELX", YearOfPublication: 2002, AuthorId: 1, BookAuthor: &models.Author{Id: 1, FirstName: "ab", LastName: "Sukhla", PenName: "Natraj", DateOfBirth: "2000-09-01", Genre: "Comic"}}, nil},
		{1000, models.Book{}, errors.New("entity not found")},
	}
	for i, tc := range testCases {
		bookStore := store.MockStore{}
		bookService := New(bookStore)
		actOutput, actError := bookService.ReadBook(tc.id)
		if !reflect.DeepEqual(actOutput, tc.expOutput) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expOutput, actOutput)
		}
		if !reflect.DeepEqual(actError, tc.expError) {
			t.Errorf("Task %v Failed, expected %v got %v", i, tc.expError, actError)
		}
	}
}
