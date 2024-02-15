package handler

import "net/http"

type PathParam interface {
	// Read reads the path params from the request
	Read(r *http.Request, key string) string
}

// ----------------------------------------------------------------------------
// PathParamStub is a stub for PathParam
type PathParamStub struct {
	// Value is the value to be returned by Read
	Value string
}

// Read reads the path params from the request
func (p *PathParamStub) Read(r *http.Request, key string) string {
	return p.Value
}
