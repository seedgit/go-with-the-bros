package go_with_bootstrap_template_again


import (
    "net/http"
    "html/template"
)

func init() {
    http.HandleFunc("/", handler)
}

var boostrapTemplate =  template.Must(template.ParseFiles("templates/index.html"))
func handler(w http.ResponseWriter, r *http.Request) {
    if err:=boostrapTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}