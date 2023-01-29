package main

import (
	authorhandler "LayeredArchitecture/handler/author"
	authorservice "LayeredArchitecture/service/author"
	authorstore "LayeredArchitecture/store/author"
	"net/http"
)

func main() {
	// get the mysql configs from env:
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
	authorStore := authorstore.MockStore{}
	authorService := authorservice.New(authorStore)
	authorHandler := authorhandler.New(authorService)
	http.HandleFunc("/author/{id}", authorHandler.GetByID)
	http.ListenAndServe(":8000", nil)
}
