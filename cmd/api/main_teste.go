package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// type product struct {
// 	Id   int
// 	Name string
// }

// func main() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)

// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("API running..."))
// 	})

// 	r.Get("/{parametroRota}", func(w http.ResponseWriter, r *http.Request) {
// 		parametroRota := chi.URLParam(r, "parametroRota")

// 		w.Write([]byte(parametroRota))
// 	})

// 	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
// 		parametroQuery := r.URL.Query().Get("parametroQuery")

// 		w.Write([]byte(parametroQuery))
// 	})

// 	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")

// 		jsonObject := map[string]string{"message": "sucess"}
// 		b, _ := json.Marshal(jsonObject)
// 		w.Write(b)
// 	})

// 	r.Post("/product", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		var product product

// 		json.NewDecoder(r.Body).Decode(&product)

// 		b, _ := json.Marshal(product)
// 		w.Write(b)
// 	})

// 	http.ListenAndServe(":3000", r)
// }
