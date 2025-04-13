package houseSelection

import (
	"html/template"
	"net/http"
)

func HouseSelectionViewHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/houseSelection/house_selection.html"))
	tmpl.Execute(w, nil)

}
