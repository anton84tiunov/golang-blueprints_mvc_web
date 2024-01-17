package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	// auth "../../auth"
)

func RegHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("GET")
		templates := template.Must(template.ParseFiles(
			"internal/app/template/base.html",
			"internal/app/template/header.html",
			"internal/app/template/footer.html",
			"internal/blueprints/auth/templates/reg_page.html"))

		data := struct {
			Title string
		}{
			Title: "reg Page",
		}

		err := templates.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var formData map[string]interface{}

		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		// RegDataParse(formData)
		fmt.Printf("Received JSON: %+v\n", formData)
		response := map[string]string{"message": "JSON data received successfully!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)

	}
}
