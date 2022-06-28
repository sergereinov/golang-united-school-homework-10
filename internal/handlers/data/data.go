package data

import (
	"io/ioutil"
	"net/http"
)

func DataHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I got message:\n"))
	w.Write(body)
}
