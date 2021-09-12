package main

import (
	"fmt"
	"net/http"
)

func serve(pathsToUrls map[string]string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerIndex)

	handler := handlerPath(pathsToUrls, mux)

	fmt.Println("Starting the server on http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index page")
}

func handlerPath(pathsToUrls map[string]string, mux http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
		}
		mux.ServeHTTP(w, r)
	}
}
