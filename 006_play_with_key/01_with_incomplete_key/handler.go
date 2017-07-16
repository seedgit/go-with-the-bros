package with_incomplete_key


import (
	"time"
	"strconv"
    "net/http"
    "html/template"
    
	"google.golang.org/appengine"
	"google.golang.org/appengine/user"
	"google.golang.org/appengine/file"
	"google.golang.org/appengine/image"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/datastore"
	
	"models"
)

func init() {
    http.HandleFunc("/newitem", webHandler(newitem_Handler))
    http.HandleFunc("/newitem/post", webHandler(newitem_post_Handler))
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
	if r.Method == "GET"{
		
		var item_keys []*datastore.Key
		var items []models.Item
		var item_error error
		item_keys, item_error = datastore.NewQuery("Item").GetAll(ctx, &items)
		
		tc["item_keys"] = item_keys
		tc["items"] = items
		tc["item_error"] = item_error
		 
	}
	
    if err:=boostrapTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var newitemTemplate =  template.Must(template.ParseFiles("templates/base.html", "templates/newitem.html"))
func newitem_Handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
	bucket, err := file.DefaultBucketName(ctx)
	uploadOption := blobstore.UploadURLOptions{
		//MaxUploadBytes: 1024 * 256,
		MaxUploadBytesPerBlob: 1024 * 256,
		StorageBucket:  bucket + "/upload",
		//StorageBucket:  bucket + "/upload/" + time.Now().Format("2006/01/02"),
	}
	uploadURL, err := blobstore.UploadURL(ctx, "/newitem/post", &uploadOption)
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	tc["upload_url"] = uploadURL 
    if err:=newitemTemplate.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var newitem_post_Template =  template.Must(template.ParseFiles("templates/base.html", "templates/post_complete.html"))
func newitem_post_Handler(w http.ResponseWriter, r *http.Request, tc map[string]interface{}) {
	ctx := appengine.NewContext(r)
	if r.Method == "POST" {
		blobs, req, err := blobstore.ParseUpload(r)
		if err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var u = user.Current(ctx)
		var name = req["name"][0]
		price, _ := strconv.Atoi(req["price"][0])
		price_discount, _ := strconv.Atoi(req["price_discount"][0])

		var image_url = ""
		if serve_url, err := image.ServingURL(ctx,  blobs["fileupload"][0].BlobKey, &image.ServingURLOptions{Secure:true,}); err == nil {
			image_url = serve_url.String()
		}
		
		incompleteKey := datastore.NewIncompleteKey(ctx, "Item", nil)
		
		var item = models.Item{
			Name: name,
			Price: price,
			PriceDiscount: price_discount,
			ImageKey: blobs["fileupload"][0].BlobKey,
			ImageUrl: image_url,
			CreatedDate: time.Now(),
			CreatedBy: u.Email,
		}
		
		if key, err := datastore.Put(ctx, incompleteKey, &item); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			tc["h1"] = "New item complete"
			tc["key"] = key
		}
	}

    if err:=newitem_post_Template.Execute(w, tc); err != nil {
    	http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}