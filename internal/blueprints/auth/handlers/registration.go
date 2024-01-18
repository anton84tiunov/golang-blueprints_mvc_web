package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	controllers "../controllers"
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

		fmt.Println("POST")
		err := json.NewDecoder(r.Body).Decode(&formData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		msg := controllers.AddUser(formData)
		// response := map[string]string{msg}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)

	}
}
