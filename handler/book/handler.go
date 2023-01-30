package book

import (
	"LayeredArchitecture/models"
	service "LayeredArchitecture/service/book"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type handler struct {
	ser service.Service
}

func New(s service.Service) handler {
	return handler{ser: s}
}

func (h handler) ReadBooks(w http.ResponseWriter, r *http.Request) {
	if reflect.DeepEqual(r.Method, http.MethodGet) {
		v := r.URL.Query()
		isbn := v.Get("isbn")
		title := v.Get("title")
		boolValue := v.Get("includeAuthor")
		var includeAuthor bool
		if boolValue == "true" {
			includeAuthor = true
		} else {
			includeAuthor = false
		}
		id, _ := strconv.Atoi(isbn)
		data, _ := h.ser.GetAll(id, title, includeAuthor)
		if len(data) == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(``))
		} else {
			w.WriteHeader(http.StatusOK)
			b, _ := json.Marshal(data)
			w.Write(b)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(``))
	}
}

func (h handler) ReadBook(w http.ResponseWriter, r *http.Request) {
	if reflect.DeepEqual(r.Method, http.MethodGet) {
		i := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(i[len(i)-1])
		data, err := h.ser.ReadBook(id)
		b, _ := json.Marshal(data)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(``))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(``))
	}
}

func (h handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if reflect.DeepEqual(r.Method, http.MethodPost) {
		body, _ := io.ReadAll(r.Body)
		var book models.Book
		_ = json.Unmarshal(body, &book)
		data, err := h.ser.Create(book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(``))
		} else {
			resp, _ := json.Marshal(data)
			w.WriteHeader(http.StatusCreated)
			w.Write(resp)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(``))
	}
}
