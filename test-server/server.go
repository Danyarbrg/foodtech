package server

import (
	"fmt"
	"net/http"
)

func Server() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", home)

	fmt.Println("Server runs on port: 3002")

	err := http.ListenAndServe(":3002", mux)
	if err != nil {
		panic(err)
	}
	return mux
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heeeeey!\n"))
}
