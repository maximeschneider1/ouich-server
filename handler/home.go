package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/maximeschneider1/ouich-server/dao"
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
			log.Printf("Error encoding response: %v", err)
		}
	}
}
