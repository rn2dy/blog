package main

import (
	"time"

  "golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
)

// Subscriber - describe subscribers
type Subscriber struct {
	Name  string
	Email string
	Since time.Time
}

func subscribersKey(c context.Context) *datastore.Key {
	return datastore.NewKey(c, "Subscribers", "blog_subscribers", 0, nil)
}
