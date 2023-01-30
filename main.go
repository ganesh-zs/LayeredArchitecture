package main

import (
	authorhandler "LayeredArchitecture/handler/author"
	bookhandler "LayeredArchitecture/handler/book"
	authorservice "LayeredArchitecture/service/author"
	bookservice "LayeredArchitecture/service/book"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//conf := driver.MySQLConfig{
	//	Host:     os.Getenv("localhost"),
	//	User:     os.Getenv("root"),
	//	Password: os.Getenv("Ganesh@123"),
	//	Port:     os.Getenv("3306"),
	//	Db:       os.Getenv("library"),
	//}
	//db, err := driver.ConnectToMySQL(conf)
	//if err != nil {
	//	log.Println("could not connect to sql, err:", err)
	//	return
	//}
	//authorStore := authorstore.New(db)

	//Get author by id
	//authorStore := authorstore.MockStore{}
	//authorService := authorservice.New(authorStore)
	//authorHandler := authorhandler.New(authorService)

	//Get all books
	//bookService := bookservice.Mockservice{}
	//bookHandler := bookhandler.New(bookService)

	bookService := bookservice.Mockservice{}
	bookHandler := bookhandler.New(bookService)

	authorService := authorservice.Mockservice{}
	authorHandler := authorhandler.New(authorService)

	r := mux.NewRouter()
	r.HandleFunc("/books", bookHandler.ReadBooks).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.ReadBook).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	r.HandleFunc("/authors/{id}", authorHandler.ReadAuthor).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
