package app02


import (
	"time"
    "net/http"
    "html/template"
)

var app02Template =  template.Must(template.ParseFiles("templates/app02.html"))
func App02Handler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["message"] = "Hello from app02"
	data["time"] = time.Now()
    if err:=app02Template.Execute(w, data); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}