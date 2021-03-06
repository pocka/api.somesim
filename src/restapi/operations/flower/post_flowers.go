package flower

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/pocka/api.somesim/models"
)

// PostFlowersHandlerFunc turns a function with the right signature into a post flowers handler
type PostFlowersHandlerFunc func(PostFlowersParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PostFlowersHandlerFunc) Handle(params PostFlowersParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PostFlowersHandler interface for that can handle valid post flowers params
type PostFlowersHandler interface {
	Handle(PostFlowersParams, *models.Principal) middleware.Responder
}

// NewPostFlowers creates a new http.Handler for the post flowers operation
func NewPostFlowers(ctx *middleware.Context, handler PostFlowersHandler) *PostFlowers {
	return &PostFlowers{Context: ctx, Handler: handler}
}

/*PostFlowers swagger:route POST /flowers Flower postFlowers

花びら定義の追加

花びらの定義を追加する
指定されたカテゴリが存在しなければカテゴリを新規作成する


*/
type PostFlowers struct {
	Context *middleware.Context
	Handler PostFlowersHandler
}

func (o *PostFlowers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostFlowersParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
