package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snippetbox.jobbrodriguez.com/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Header Methods: Set(), Add(), Del(), Get(), Values()
	w.Header().Add("Server", "Go")
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	msg := fmt.Sprintf("Display a form for creating a new snippet")
	w.Write([]byte(msg))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// dummy data
	title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil{
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
