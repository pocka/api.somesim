package flower

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFlowersHandlerFunc turns a function with the right signature into a get flowers handler
type GetFlowersHandlerFunc func(GetFlowersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFlowersHandlerFunc) Handle(params GetFlowersParams) middleware.Responder {
	return fn(params)
}

// GetFlowersHandler interface for that can handle valid get flowers params
type GetFlowersHandler interface {
	Handle(GetFlowersParams) middleware.Responder
}

// NewGetFlowers creates a new http.Handler for the get flowers operation
func NewGetFlowers(ctx *middleware.Context, handler GetFlowersHandler) *GetFlowers {
	return &GetFlowers{Context: ctx, Handler: handler}
}

/*GetFlowers swagger:route GET /flowers Flower getFlowers

花びら定義の取得

花びら定義を取得する

花びらのカテゴリによる絞り込みが可能(e.g. ラーファン, 課金花びら)


*/
type GetFlowers struct {
	Context *middleware.Context
	Handler GetFlowersHandler
}

func (o *GetFlowers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetFlowersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
