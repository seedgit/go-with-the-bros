package where_is_my_caption


import (
	"time"
    "net/http"
    "html/template"
    
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/blobstore"
	
	"models"
)

func init() {
	http.HandleFunc("/post", webHandler(upload_handler))
	http.HandleFunc("/photobook", webHandler(photobook_handler))
	http.HandleFunc("/images/", image_handler)
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
	ctx := appengine.NewContext(r)
	uploadOption := blobstore.UploadURLOptions{
		//MaxUploadBytes: 1024 * 256,
		MaxUploadBytesPerBlob: 1024 * 256,
	}
	uploadURL, err := blobstore.UploadURL(ctx, "/post", &uploadOption)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	tc["upload_url"] = uploadURL 
    if err:=boostrapTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
var blankTemplate =  template.Must(template.ParseFiles("templates/base.html", "templates/blank.html"))
func upload_handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
	if r.Method == "POST" {
		blobs, _, err := blobstore.ParseUpload(r)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		var photo = models.PhotoBook{
			Caption: r.FormValue("caption"),
			ImageKey: blobs["fileupload"][0].BlobKey,
			UploadDate: time.Now(),
		}
		if _, err := datastore.Put(ctx, datastore.NewIncompleteKey(ctx, "PhotoBook", nil), &photo); err!=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return			
		}
		tc["h1"] = "Upload complete"
		tc["p"] = "CONGLATURATION !!!"
	    if err:=blankTemplate.Execute(w, tc); err != nil {
	    	http.Error(w, err.Error(), http.StatusInternalServerError)
	    }
	}
}
var photoTemplate =  template.Must(template.ParseFiles("templates/base.html", "templates/photos.html"))
func photobook_handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
	var photoKeys []*datastore.Key
	var photos []models.PhotoBook
	var photoErr error
	
	if photoKeys, photoErr = datastore.NewQuery("PhotoBook").GetAll(ctx, &photos); photoErr!=nil{
		http.Error(w, photoErr.Error(), http.StatusInternalServerError)
		return
	}
	tc["photoKeys"] = photoKeys
	tc["photos"] = photos
	if err:=photoTemplate.Execute(w, tc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func image_handler(w http.ResponseWriter, r *http.Request) {
	blobstore.Send(w, appengine.BlobKey(r.FormValue("key")))
}