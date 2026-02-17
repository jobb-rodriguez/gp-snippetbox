package main

import "snippetbox.jobbrodriguez.com/internal/models"

type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}
