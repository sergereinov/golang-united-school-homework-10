package name

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello, %v!", vars["PARAM"])))
}
