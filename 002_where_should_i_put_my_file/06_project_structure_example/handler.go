package project_structure_example


import (
    "net/http"
    "html/template"
    "app01"
    "app02"
)

func init() {
	http.HandleFunc("/app01", app01.App01Handler)
	http.HandleFunc("/app02", app02.App02Handler)
    http.HandleFunc("/", handler)
}

var boostrapTemplate =  template.Must(template.ParseFiles("templates/index.html"))
func handler(w http.ResponseWriter, r *http.Request) {
    if err:=boostrapTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}