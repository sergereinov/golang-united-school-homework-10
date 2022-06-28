package headers

import (
	"net/http"
	"strconv"
)

func HeadersHandler(w http.ResponseWriter, r *http.Request) {

	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))

	w.Header().Add("a+b", strconv.Itoa(a+b))

	w.WriteHeader(http.StatusOK)
}
