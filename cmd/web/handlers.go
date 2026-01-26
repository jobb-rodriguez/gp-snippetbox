package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Header Methods: Set(), Add(), Del(), Get(), Values()
	w.Header().Add("Server", "Go")
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	msg := fmt.Sprintf("Display a form for creating a new snippet")
	w.Write([]byte(msg))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	msg := fmt.Sprintf("Save a new snippet...")
	w.Write([]byte(msg))
}
