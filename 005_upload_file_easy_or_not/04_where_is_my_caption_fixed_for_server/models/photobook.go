package models
import (
	"time"
	"google.golang.org/appengine"
)
type PhotoBook struct{
	Caption			string
	ImageKey		appengine.BlobKey
	UploadDate		time.Time
}