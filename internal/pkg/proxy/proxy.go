package proxy

import (
	"encoding/json"
	"net/http"
)

type handleFunc func(http.ResponseWriter, *http.Request) (interface{}, int, error)

type ResponseBody struct {
	IsOk   bool        `json:"isOk"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func New(handle handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		result, status, err := handle(w, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		if err != nil {
			json.NewEncoder(w).Encode(ResponseBody{
				IsOk:   false,
				Status: http.StatusText(status),
				Data:   err.Error()})
			return
		}

		json.NewEncoder(w).Encode(ResponseBody{
			IsOk:   true,
			Status: http.StatusText(status),
			Data:   result})
	}
}
