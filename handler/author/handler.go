package author

import (
	service "LayeredArchitecture/service/author"
	"encoding/json"
	"fmt"
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

func (h handler) ReadAuthor(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	if reflect.DeepEqual(r.Method, http.MethodGet) {
		i := strings.Split(r.URL.Path, "/")
		id, err := strconv.Atoi(i[len(i)-1])
		data, err := h.ser.ReadAuthor(id)
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
