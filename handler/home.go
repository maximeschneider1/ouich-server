package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// handleGetHome returns the list of quotes for home
func (s *server) handleGetHome() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		var resp response

		resp.Data = append(resp.Data, r)
		resp.StatusCode = http.StatusOK
		resp.Message = "OK"
		resp.Error = nil
		resp.Meta.Query = fmt.Sprintln("List of quotes")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			log.Printf("Error encoding response : %v", err)
		}
	}
}
