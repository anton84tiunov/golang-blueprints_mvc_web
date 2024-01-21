package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	// "../../../services"

	"github.com/steambap/captcha"

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

		img, err_c := captcha.NewMathExpr(150, 50)
		if err_c != nil {
			fmt.Fprint(w, nil)
			fmt.Println(err_c.Error())
			return
		}
		text := img.Text
		fmt.Println(text)
		buf := new(bytes.Buffer)
		// bufb := buf.([]byte)

		img.WriteImage(buf)
		butes := buf.Bytes()
		imgBase64Str := base64.StdEncoding.EncodeToString(butes)
		// msg := "<img src=\"data:image/png;base64," + imgBase64Str + "\" alt=\"Blank\" class=\"bg-white\"/>"
		// msg := "<h1>ppppppppp</h1>"
		// services.SendMessageMail("anton.tiunov.84.07@gmail.com", msg, "Подтверждение почты")

		data := struct {
			Title string
			Buf   string
		}{
			Title: "reg Page",
			Buf:   imgBase64Str,
		}

		err := templates.ExecuteTemplate(w, "base.html", data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var formData map[string]interface{}

		fmt.Println("POST")
		// contx := context.Context.Value "qqq", "ppppp")
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
