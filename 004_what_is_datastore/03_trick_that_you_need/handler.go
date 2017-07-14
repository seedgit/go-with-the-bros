package trick_that_you_need


import (
    "net/http"
    "html/template"
    
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/datastore"
	
	"admin"
	"models"
)

func init() {
    http.HandleFunc("/post", webHandler(guestbook_post_Handler))
    http.HandleFunc("/guest_messages", webHandler(guest_messages_Handler))
    http.HandleFunc("/admin/", webHandler(admin.AdminHandler))
    http.HandleFunc("/admin/edit/", webHandler(admin.AdminEditHandler))
    http.HandleFunc("/admin/delete/", webHandler(admin.AdminDeleteHandler))
    http.HandleFunc("/", webHandler(handler))
}

func webHandler(handlefunc func(http.ResponseWriter, *http.Request, map[string]interface{})) func(http.ResponseWriter, *http.Request) {
	outfunc := func(w http.ResponseWriter, r *http.Request) {
		tc := make(map[string]interface{})
		ctx := appengine.NewContext(r)
	    u := user.Current(ctx)
	    if u == nil {
	    	url, _ := user.LoginURL(ctx, "/")
	    	tc["login_url"] = url
	    }
	    tc["logout_url"], _ = user.LogoutURL(ctx, "/")
	    tc["user"] = u
	    tc["is_admin"] = user.IsAdmin(ctx)
		handlefunc(w, r, tc)
	}
	return outfunc
}


var boostrapTemplate =  template.Must(template.ParseFiles("templates/base.html", "templates/index.html"))
func handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
    if err:=boostrapTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var post_complete_Template =  template.Must(template.ParseFiles("templates/base.html", "templates/post_complete.html"))
func guestbook_post_Handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
	if r.Method == "POST" {
		guestbook_message := models.Guestbook{
			Name: r.FormValue("name"),
			Message: r.FormValue("message"),
		}
		incompleteKey := datastore.NewIncompleteKey(ctx, "Guestbook", nil)
		if key, err := datastore.Put(ctx, incompleteKey, &guestbook_message); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			tc := make(map[string]interface{})
			tc["incomplete_key"] = incompleteKey
			tc["inserted_key"] = key
		    if err:=post_complete_Template.Execute(w, tc); err != nil {
		    	http.Error(w, err.Error(), http.StatusInternalServerError)
		    }
		}
	}
}
var guest_messages_Template =  template.Must(template.ParseFiles("templates/base.html", "templates/messages.html"))
func guest_messages_Handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
	var messages []models.Guestbook
	if _, err:=datastore.NewQuery("Guestbook").GetAll(ctx, &messages); err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tc["messages"] = messages
	if err:=guest_messages_Template.Execute(w, tc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

