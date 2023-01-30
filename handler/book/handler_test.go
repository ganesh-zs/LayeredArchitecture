package book

import (
	"LayeredArchitecture/models"
	bookservice "LayeredArchitecture/service/book"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

//func Test_ReadBooks(t *testing.T) {
//	testCases := []struct {
//		isbn               int
//		title              string
//		includeauthor      string
//		method             string
//		expectedStatusCode int
//		expectedOutput     []byte
//	}{
//		{isbn: 1, includeauthor: "false", method: http.MethodGet, expectedStatusCode: http.StatusOK, expectedOutput: []byte(`[{"id":1,"isbn":1,"title":"Hai","genre":"Comic","publication":"RELX","yearOfPublication":2000,"authorId":1}]`)},
//		{title: "Hai", includeauthor: "true", method: http.MethodGet, expectedStatusCode: http.StatusOK, expectedOutput: []byte(`[{"id":1,"isbn":1,"title":"Hai","genre":"Comic","publication":"RELX","yearOfPublication":2000,"authorId":1,"bookAuthor":{"id":1,"firstName":"Ganesh","lastName":"Manchi","penName":"Natraj","dateOfBirth":"2001-09-01","genre":"Comic"}}]`)},
//		{isbn: 1, title: "Hai", includeauthor: "true", method: http.MethodGet, expectedStatusCode: http.StatusOK, expectedOutput: []byte(`[{"id":1,"isbn":1,"title":"Hai","genre":"Comic","publication":"RELX","yearOfPublication":2000,"authorId":1,"bookAuthor":{"id":1,"firstName":"Ganesh","lastName":"Manchi","penName":"Natraj","dateOfBirth":"2001-09-01","genre":"Comic"}}]`)},
//		{isbn: 1000, title: "Hello", includeauthor: "true", method: http.MethodGet, expectedStatusCode: http.StatusNotFound, expectedOutput: []byte(``)},
//		{isbn: 2, title: "NoTitle", includeauthor: "true", method: http.MethodGet, expectedStatusCode: http.StatusNotFound, expectedOutput: []byte(``)},
//		{isbn: 1, title: "NoTitle", includeauthor: "true", method: http.MethodPost, expectedStatusCode: http.StatusMethodNotAllowed, expectedOutput: []byte(``)},
//	}
//	for i, tc := range testCases {
//		url := fmt.Sprintf("http://localhost:8000/books?isbn=%d", tc.isbn)
//		url = url + "&title="
//		url = url + tc.title
//		url = url + "&includeAuthor="
//		url = url + tc.includeauthor
//		req := httptest.NewRequest(tc.method, url, nil)
//		w := httptest.NewRecorder()
//		bookService := bookservice.Mockservice{}
//		bookHandler := New(bookService)
//		bookHandler.ReadBooks(w, req)
//		resp := w.Result()
//		body, _ := io.ReadAll(resp.Body)
//		if !reflect.DeepEqual(resp.StatusCode, tc.expectedStatusCode) {
//			t.Errorf("testcase %v failed, expected %v got %v", i, tc.expectedStatusCode, resp.StatusCode)
//		}
//		if !reflect.DeepEqual(body, tc.expectedOutput) {
//			t.Errorf("Task %v Failed", i)
//		}
//	}
//}

//func Test_ReadBook(t *testing.T) { // working
//	testCases := []struct {
//		id                 int
//		method             string
//		expectedStatusCode int
//		expectedOutput     []byte
//	}{
//		{16, http.MethodGet, http.StatusOK, []byte(`{"id":16,"isbn":778,"title":"Money","genre":"Fiction","publication":"RELX","yearOfPublication":2002,"authorId":1,"bookAuthor":{"id":1,"firstName":"ab","lastName":"Sukhla","penName":"Natraj","dateOfBirth":"2000-09-01","genre":"Comic"}}`)},
//		{1000, http.MethodGet, http.StatusNotFound, []byte(``)},
//		{1, http.MethodPost, http.StatusMethodNotAllowed, []byte(``)},
//	}
//	for i, tc := range testCases {
//		url := fmt.Sprintf("http://localhost:8000/books/%d", tc.id)
//		req := httptest.NewRequest(tc.method, url, nil)
//		w := httptest.NewRecorder()
//		bookService := bookservice.Mockservice{}
//		bookHandler := New(bookService)
//		bookHandler.ReadBook(w, req)
//		resp := w.Result()
//		body, err := io.ReadAll(resp.Body)
//		if err != nil {
//			t.Errorf("Error while reading the data")
//		}
//		if !reflect.DeepEqual(resp.StatusCode, tc.expectedStatusCode) {
//			t.Errorf("testcase %v failed, expected %v got %v", i, tc.expectedStatusCode, resp.StatusCode)
//		}
//		if !reflect.DeepEqual(body, tc.expectedOutput) {
//			expectedStructure := models.Book{}
//			actualStructure := models.Book{}
//			err1 := json.Unmarshal(body, expectedStructure)
//			err2 := json.Unmarshal(body, actualStructure)
//			if err1 != nil || err2 != nil {
//				t.Errorf("Task %v Failed, expected %v got %v", i, expectedStructure, actualStructure)
//			}
//		}
//	}
//}

func Test_CreateBook(t *testing.T) { // working
	testCases := []struct {
		expectedStatusCode int
		method             string
		input              []byte
		expOutput          []byte
	}{
		{http.StatusCreated, http.MethodPost, []byte(`{"isbn":3469,"title":"Money","genre":"Fiction","publication":"RELX","yearOfPublication":2002,"authorId":1}`), []byte(`{"id":1,"isbn":3469,"title":"Money Heist","genre":"Fiction","publication":"RELX","yearOfPublication":2002,"authorId":1}`)},
		{http.StatusMethodNotAllowed, http.MethodPut, []byte(`{"isbn"": 1456, "title"": "abcdefg", "genre"": "Classics ", "publication"": "Pearson", "yearOfPublication"": 2000}`), []byte(``)},
		{http.StatusInternalServerError, http.MethodPost, []byte(`{"isbn":3469,"title":"Money","genre":"Fiction","publication":"Harry","yearOfPublication":2002,"authorId":1}`), []byte(``)},
	}
	for i, tc := range testCases {
		req := httptest.NewRequest(tc.method, "http://localhost:8000/books", bytes.NewReader(tc.input))
		w := httptest.NewRecorder()
		bookService := bookservice.Mockservice{}
		bookHandler := New(bookService)
		bookHandler.CreateBook(w, req)
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		if !reflect.DeepEqual(resp.StatusCode, tc.expectedStatusCode) {
			t.Errorf("testcase %v failed expected %v got %v", i, tc.expectedStatusCode, resp.StatusCode)
		}
		expectedStructure := models.Book{}
		actualStructure := models.Book{}
		err1 := json.Unmarshal(body, expectedStructure)
		err2 := json.Unmarshal(body, actualStructure)
		if err1 != nil || err2 != nil {
			if !reflect.DeepEqual(actualStructure, expectedStructure) {
				t.Errorf("Task %v Failed, expected %v got %v", i, expectedStructure, actualStructure)

			}
		}
	}
}
