package welcome_to_the_jungle


import (
    "fmt"
    "net/http"
)

func init() {
	http.HandleFunc("/welcome", welcome)
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to the jungle")
}

func welcome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "<h1>Welcome to the jungle</h1>")
}
