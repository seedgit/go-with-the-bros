package models

import (
	"time"
	"google.golang.org/appengine"
)

type Item struct{
	Name				string
	Price				int
	PriceDiscount		int
	ImageKey			appengine.BlobKey
	ImageUrl			string
	CreatedDate			time.Time
	CreatedBy			string
}

type Review struct{
	Message				string
	CreatedDate			time.Time
	CreatedBy			string
}