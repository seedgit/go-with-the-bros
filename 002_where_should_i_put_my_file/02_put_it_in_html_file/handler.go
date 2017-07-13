package put_html_in_html_file


import (
    "net/http"
    "html/template"
)

func init() {
	http.HandleFunc("/hello", hello)
    http.HandleFunc("/", handler)
}

var myTemplate =  template.Must(template.ParseFiles("template.html"))
func handler(w http.ResponseWriter, r *http.Request) {
    if err:=myTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var helloTemplate =  template.Must(template.ParseFiles("hello.html"))
func hello(w http.ResponseWriter, r *http.Request) {
    if err:=helloTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
