package models
import (
	"time"
	"google.golang.org/appengine"
)
type PhotoBook struct{
	ImageKey		appengine.BlobKey
	UploadDate		time.Time
}