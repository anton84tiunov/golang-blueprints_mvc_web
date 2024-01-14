package about

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET")
		templates := template.Must(template.ParseFiles(
			"internal/app/template/base.html",
			"internal/app/template/header.html",
			"internal/app/template/footer.html",
			"internal/blueprints/about/templates/about_page.html"))

		data := struct {
			Title string
		}{
			Title: "about Page",
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

func AboutRoutes() http.Handler {
	router := mux.NewRouter()

	router.PathPrefix("/about/static/").Handler(http.StripPrefix("/about/static/", http.FileServer(http.Dir("internal/blueprints/about/static"))))

	router.HandleFunc("/about/", AboutHandler)
	// router.HandleFunc("/", AboutHandler)
	return router
}
