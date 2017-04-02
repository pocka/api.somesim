package internal_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetDocsHandlerFunc turns a function with the right signature into a get docs handler
type GetDocsHandlerFunc func(GetDocsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDocsHandlerFunc) Handle(params GetDocsParams) middleware.Responder {
	return fn(params)
}

// GetDocsHandler interface for that can handle valid get docs params
type GetDocsHandler interface {
	Handle(GetDocsParams) middleware.Responder
}

// NewGetDocs creates a new http.Handler for the get docs operation
func NewGetDocs(ctx *middleware.Context, handler GetDocsHandler) *GetDocs {
	return &GetDocs{Context: ctx, Handler: handler}
}

/*GetDocs swagger:route GET /docs InternalAPI getDocs

ReDocによるAPIドキュメントを取得する(HTML)


*/
type GetDocs struct {
	Context *middleware.Context
	Handler GetDocsHandler
}

func (o *GetDocs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetDocsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}