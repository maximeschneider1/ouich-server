package handler

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

// server is the base structure of the API
type server struct {
	router *httprouter.Router

	// Here I would put a DB connection if there was one
}

// response contains all response infos at a glance
type response struct {
	StatusCode int  `json:"status_code"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	Meta       struct {
		Query       interface{} `json:"query,omitempty"`
		ResultCount int         `json:"result_count,omitempty"`} `json:"meta"`
	Data []interface{} `json:"data"`
}

// StartWebServer is the function responsible for launching the API
func StartWebServer() {
	s := server{
		router : httprouter.New(),
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