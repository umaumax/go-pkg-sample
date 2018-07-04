package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func apiRouting() http.Handler {
	router := httprouter.New()
	router.GET("/api/hello/", helloHandler)
	//	NOTE not match /api/hello/
	router.GET("/api/hello/:name", helloHandler)
	return router
}

func helloHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")
	log.Println("helloHandler:", r, ps)

	if name == "" {
		err := fmt.Errorf("empty name error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	jsonData := struct {
		Name string
		Age  int
	}{
		Name: name,
		Age:  18,
	}
	if err := json.NewEncoder(w).Encode(jsonData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
