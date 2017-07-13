package put_it_in_go_file


import (
    "net/http"
    "html/template"
)

func init() {
    http.HandleFunc("/", handler)
}

var myTemplate =  template.Must(template.New("my_template").Parse(`
<html>
  <head>
    <title>Put it in go file</title>
  </head>
  <body>
	  <h1>Put it in go file.</h1>
	  <p>Good or Bad idea?</p>
  </body>
</html>
`))

func handler(w http.ResponseWriter, r *http.Request) {
    if err:=myTemplate.Execute(w, nil); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
