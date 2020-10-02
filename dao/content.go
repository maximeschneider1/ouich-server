package dao

import (
	"database/sql"

	"github.com/maximeschneider1/ouich-server/model"
)

func QueryAllQuotes(db *sql.DB) ([]model.Quote, error) {
	var ac []model.Quote
	rows, err := db.Query("SELECT id, title, content, file_path FROM film_quotes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		q := model.Quote{}
		err := rows.Scan(&q.ID, &q.Title, &q.Content, &q.FilePath)
		if err != nil {
			return nil, err
		}
		ac = append(ac, q)
	}
	return ac, nil
}


// 1. Prendre le payload
// 1. le transformer en Go struct
// 1.