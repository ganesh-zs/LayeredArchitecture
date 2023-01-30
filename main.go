package main

import (
	bookhandler "LayeredArchitecture/handler/book"
	bookservice "LayeredArchitecture/service/book"
	"github.com/gorilla/mux"
	"log"
	//authorhandler "LayeredArchitecture/handler/author"
	//authorservice "LayeredArchitecture/service/author"
	//authorstore "LayeredArchitecture/store/author"
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

	r := mux.NewRouter()
	//r.HandleFunc("/books", handlers.ReadBooks).Methods("GET")
	//r.HandleFunc("/books/{id}", bookHandler.ReadBook).Methods("GET")
	r.HandleFunc("/books", bookHandler.CreateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))

}
