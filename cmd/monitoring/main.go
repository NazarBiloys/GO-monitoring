package main

import (
	"github.com/NazarBiloys/GO-monitoring/internal/service"
	"log"
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if err := service.MakeUser(); err != nil {
	    fmt.Fprint(w, err.Error())
	}

	if err := service.MakeStudent(); err != nil {
	    fmt.Fprint(w, err.Error())
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":90", nil))
}
