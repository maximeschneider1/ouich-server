package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

// server is the base structure of the API
type server struct {
	router   *httprouter.Router
	database *sql.DB
}

// response contains all response infos at a glance
type response struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
	Message    string      `json:"message"`
	Meta       struct {
		Query       interface{} `json:"query,omitempty"`
		ResultCount int         `json:"result_count,omitempty"`
	} `json:"meta"`
	Data []interface{} `json:"data"`
}

// StartWebServer is the function responsible for launching the API
func StartWebServer() {
	psqlInfo := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	s := server{
		router:   httprouter.New(),
		database: db,
	}
	s.router.PanicHandler = handlePanic
	s.routes()
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), s.router))
}

// routes function launches all application's routes
func (s *server) routes() {
	//home
	s.router.GET("/", s.handleGetHome())
	s.router.POST("/new", s.handlePostQuote())
	//replique
	// random
}

// Gracefully handle panic without crashing the server
func handlePanic(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	w.WriteHeader(http.StatusInternalServerError)
}
