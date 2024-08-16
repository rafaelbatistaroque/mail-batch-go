package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafaelbatistaroque/mail-batch-go/internal/domain/campaign/main/endpoint"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API v1.0 on"))
	})

	campaignEndpoints := endpoint.GetCampaignEndpoints()
	r.Post("/campaign", campaignEndpoints.CreateCampaign)
	r.Post("/campaign/search", campaignEndpoints.SearchCampaign)
	r.Get("/campaign/{id}", campaignEndpoints.GetCampaignById)

	http.ListenAndServe(":3000", r)
}
