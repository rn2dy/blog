package main

import (
	"time"

	"appengine"
	"appengine/datastore"
)

// Subscriber - describe subscribers
type Subscriber struct {
	Name  string
	Email string
	Since time.Time
}

func subscribersKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Subscribers", "blog_subscribers", 0, nil)
}
