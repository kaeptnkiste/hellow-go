package main

import (
		"fmt"
		"net/http"
		"log"
)


func main(){
	http.HandleFunc("/",hellowGo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hellowGo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hellow Jan")
	}	