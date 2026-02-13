package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID		int
	Title	string
	Content	string
	Created	time.Time
	Expires	time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

func (m *SnippetMOdel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
}

// Return 10 recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
