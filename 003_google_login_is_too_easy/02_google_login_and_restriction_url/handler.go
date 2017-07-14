package google_login_and_restriction_url


import (
    "net/http"
    "html/template"
    "app01"
    "app02"
    
	"google.golang.org/appengine"
    "google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/app01", app01.App01Handler)
	http.HandleFunc("/app02", app02.App02Handler)
    http.HandleFunc("/", handler)
}

var boostrapTemplate =  template.Must(template.ParseFiles("templates/index.html"))
func handler(w http.ResponseWriter, r *http.Request) {
	tc := make(map[string]interface{})
	ctx := appengine.NewContext(r)
    u := user.Current(ctx)
    if u == nil {
    	url, _ := user.LoginURL(ctx, "/")
    	tc["login_url"] = url
    }
    tc["logout_url"], _ = user.LogoutURL(ctx, "/")
    tc["user"] = u

    if err:=boostrapTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}