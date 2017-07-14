package app01


import (
    "net/http"
    "html/template"
)

var app01Template =  template.Must(template.ParseFiles("templates/app01.html"))
func App01Handler(w http.ResponseWriter, r *http.Request) {
    if err:=app01Template.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}