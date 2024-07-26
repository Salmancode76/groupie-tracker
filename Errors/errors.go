package Errors

import (
	"fmt"
	"net/http"
)

func Error500(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "<h1>500 Error</h1>")
}
func Error404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprint(w, "<h1>404 Error</h1> page not found :( )")
}
