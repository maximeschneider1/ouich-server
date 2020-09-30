package handler

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// server is the base structure of the API
type server struct {
	router   *httprouter.Router
	database *sql.DB
}

// response contains all response infos at a glance
type response struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	Meta       struct {
		Query       interface{} `json:"query,omitempty"`
		ResultCount int         `json:"result_count,omitempty"`
	} `json:"meta"`
	Data []interface{} `json:"data"`
}

// StartWebServer is the function responsible for launching the API
func StartWebServer() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	db, _ := sql.Open("postgres", "postgres://postgres:"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":5432/"+os.Getenv("DB_NAME")+"?sslmode=disable")
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	s := server{
		router:   httprouter.New(),
		database: db,
	}
	s.router.PanicHandler = handlePanic

	s.routes()

	log.Fatal(http.ListenAndServe(":8085", s.router))
}

// routes function launches all application's routes
func (s *server) routes() {
	//home
	s.router.GET("/home", s.handleGetHome())
	//replique
	// random
}

// Gracefully handle panic without crashing the server
func handlePanic(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Println(r.URL.Path, err)
	w.WriteHeader(http.StatusInternalServerError)
}
