package admin

import (
    "net/http"
    "html/template"
    
    "google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
    
    "models"
)


var adminTemplate =  template.Must(template.ParseFiles("templates/base.html", "admin/templates/index.html"))
func AdminHandler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)

	var messages []models.Guestbook
	if keys, err:=datastore.NewQuery("Guestbook").GetAll(ctx, &messages); err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		tc["keys"] = keys
		tc["messages"] = messages
	}
	
    
    if err:=adminTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
var adminEditTemplate =  template.Must(template.ParseFiles("templates/base.html", "admin/templates/edit.html"))
func AdminEditHandler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
    
    if r.Method == "GET" {
    	var message models.Guestbook
    	key, err := datastore.DecodeKey(r.FormValue("key"))
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
    	}
    	if err := datastore.Get(ctx, key, &message); err!=nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return    		
    	}
    	tc["key"] = key
    	tc["message"] = message
    }
    if r.Method == "POST" {
    	var message models.Guestbook
    	key, err := datastore.DecodeKey(r.FormValue("key"))
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
    	}
    	if err := datastore.Get(ctx, key, &message); err!=nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return    		
    	}
    	message.Name = r.FormValue("name")
    	message.Message = r.FormValue("message")
    	
    	if _, err := datastore.Put(ctx, key, &message); err!=nil{
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return    		
    	}
    	tc["key"] = key
    	tc["message"] = message
    	tc["edit_message"] = "Edit message complete"
    }
    
    if err:=adminEditTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
var adminDeleteTemplate =  template.Must(template.ParseFiles("templates/base.html", "admin/templates/delete.html"))
func AdminDeleteHandler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
    if r.Method == "GET"{
		key, err := datastore.DecodeKey(r.FormValue("key"))
    	if err != nil {
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
    	}
    	if err := datastore.Delete(ctx, key); err!=nil{
    		http.Error(w, err.Error(), http.StatusInternalServerError)
    		return
    	}
    }
    if err:=adminDeleteTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}