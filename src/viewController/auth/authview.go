package auth

import (
	"html/template"
	"net/http"
)

func LoginViewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/authView/index.html"))
	tmpl.Execute(w, nil)

}
func SignUpViewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/authView/sign_up.html"))
	tmpl.Execute(w, nil)

}
