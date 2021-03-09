package endpoints

import (
	"fmt"
)

// Route is a basic struct containing Method and URL
type Route struct {
	Method Method
	URL    string
}

// NewRoute generates a new Route struct
func NewRoute(method Method, url string) Route {
	return Route{
		Method: method,
		URL:    url,
	}
}

// Compile builds a full request URL based on arguments
func (r Route) Compile(args ...interface{}) string {
	if len(args) == 0 {
		return API + r.URL
	}
	return API + fmt.Sprintf(r.URL, args...)
}
