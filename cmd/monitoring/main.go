package main

import (
	"fmt"
	"github.com/NazarBiloys/GO-monitoring/internal/service"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := service.MakeUser(); err != nil {
			panic(err)
		}

		if err := service.MakeStudent(); err != nil {
			panic(err)
		}
	}

	if r.Method == "GET" {
		students, err := service.FetchFirstFiveStudent()
		if err != nil {
			panic(err)
		}
		users, err := service.FetchFirstFiveUser()
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "Students : %s, Users : %s", students, users)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":90", nil))
}
