package couchdbgo

import "net/http"

// Connection information for couch db
type Client struct {
	c *http.Client

	host              string
	port              string
	scheme            string
	basicAuth         bool
	basicAuthPassword string
	basicAuthUsername string
}

// ClientOptionFunc is a function that configures a Client.
// It is used in NewClient.
type ClientOptionFunc func(*Client) error

func NewClient(options ...ClientOptionFunc) (*Client, error) {
	// Set up the client
	c := &Client{
		c:      http.DefaultClient,
		scheme: "http",
	}

	// Run the options on it
	for _, option := range options {
		if err := option(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}
