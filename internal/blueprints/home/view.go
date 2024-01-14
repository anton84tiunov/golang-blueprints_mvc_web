package home

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET")

		templates := template.Must(template.ParseFiles(
			"internal/app/template/base.html",
			"internal/app/template/header.html",
			"internal/app/template/footer.html",
			"internal/blueprints/home/templates/home_page.html"))

		data := struct {
			Title string
		}{
			Title: "home Page",
		}

		err := templates.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var formData map[string]interface{}

		if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		fmt.Printf("Received JSON: %+v\n", formData)
		response := map[string]string{"message": "JSON data received successfully!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}

func HomeRoutes() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/home/static/").Handler(http.StripPrefix("/home/static/", http.FileServer(http.Dir("internal/blueprints/home/static"))))

	router.HandleFunc("/home/", HomeHandler)
	router.HandleFunc("/", HomeHandler)

	return router
}
