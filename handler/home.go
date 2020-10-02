package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/maximeschneider1/ouich-server/dao"
	"github.com/maximeschneider1/ouich-server/model"
	"io/ioutil"
	"log"
	"net/http"
)

// handleGetHome returns the list of quotes for home
func (s *server) handleGetHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		resp := response{}
		results, err := dao.QueryAllQuotes(s.database)
		if err != nil {
			log.Println(err)
		}
		resp.Data = append(resp.Data, results)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = nil
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// handlePostQuote posts new quote to db
func (s *server) handlePostQuote() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		body, err := ioutil.ReadAll(r.Body); if err != nil {
			log.Println("Error reading request body, error :", err.Error())
			resp.Error = "Error reading request body"
			resp.Message = "Internal Server Error"
			resp.StatusCode = http.StatusInternalServerError
			w.WriteHeader(http.StatusInternalServerError)
			err = json.NewEncoder(w).Encode(resp); if err!= nil {
				log.Printf("Error encoding response : %v", err)
			}
			return
		}
		newQuote := model.Quote{}
		json.Unmarshal(body, &newQuote)

		AddQuote(s.database, newQuote)

		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = "No error"
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(resp); if err!= nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}

// AddQuote inserts in the database
func AddQuote(db *sql.DB, newQuote model.Quote) error {
	//Get last tag id
	var lastID int
	err := db.QueryRow("SELECT id FROM film_quotes ORDER BY id DESC LIMIT 1;").Scan(&lastID)
	if err != nil {
		log.Println("Error querying last tag id, error :", err.Error())
		return err
	}

	// Post tag in the DB
	var ctx = context.Background()
	tx, err := db.BeginTx(ctx, nil); if err != nil {
		log.Println("Error begining transaction :", err.Error())
		return err
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO film_quotes (id, title, content, file_path) VALUES ($1, $2, $3, $4)", lastID + 1, newQuote.Title, newQuote.Title, newQuote.FilePath); if err != nil {
		// In case we find any error in the query execution, rollback the transaction
		log.Println("Error executing transaction :", err.Error())
		err = tx.Rollback(); if err != nil {
			log.Println("Error during rollback on transaction :", err.Error())
			return err
		}
		return err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("Error committing the transaction :", err.Error())
	}

	return nil
}