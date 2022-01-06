package routes

import (
	"fmt"
	"net/http"
)


func GetItems(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println( r.Body)
	fmt.Fprintf(rw, "test")
}