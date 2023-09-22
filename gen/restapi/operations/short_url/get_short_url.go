// Code generated by go-swagger; DO NOT EDIT.

package short_url

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetShortURLHandlerFunc turns a function with the right signature into a get short Url handler
type GetShortURLHandlerFunc func(GetShortURLParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetShortURLHandlerFunc) Handle(params GetShortURLParams) middleware.Responder {
	return fn(params)
}

// GetShortURLHandler interface for that can handle valid get short Url params
type GetShortURLHandler interface {
	Handle(GetShortURLParams) middleware.Responder
}

// NewGetShortURL creates a new http.Handler for the get short Url operation
func NewGetShortURL(ctx *middleware.Context, handler GetShortURLHandler) *GetShortURL {
	return &GetShortURL{Context: ctx, Handler: handler}
}

/*
	GetShortURL swagger:route GET /{short_url} ShortUrl getShortUrl

# Get short URL

Get short URL with a redirect
*/
type GetShortURL struct {
	Context *middleware.Context
	Handler GetShortURLHandler
}

func (o *GetShortURL) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetShortURLParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}