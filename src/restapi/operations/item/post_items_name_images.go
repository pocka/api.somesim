package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostItemsNameImagesHandlerFunc turns a function with the right signature into a post items name images handler
type PostItemsNameImagesHandlerFunc func(PostItemsNameImagesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostItemsNameImagesHandlerFunc) Handle(params PostItemsNameImagesParams) middleware.Responder {
	return fn(params)
}

// PostItemsNameImagesHandler interface for that can handle valid post items name images params
type PostItemsNameImagesHandler interface {
	Handle(PostItemsNameImagesParams) middleware.Responder
}

// NewPostItemsNameImages creates a new http.Handler for the post items name images operation
func NewPostItemsNameImages(ctx *middleware.Context, handler PostItemsNameImagesHandler) *PostItemsNameImages {
	return &PostItemsNameImages{Context: ctx, Handler: handler}
}

/*PostItemsNameImages swagger:route POST /items/{name}/images Item postItemsNameImages

染色用画像を追加

指定された名前の装備に新しい装備画像セットを追加

e.g. ハイキャスター装備にNFの画像セットを追加


*/
type PostItemsNameImages struct {
	Context *middleware.Context
	Handler PostItemsNameImagesHandler
}

func (o *PostItemsNameImages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostItemsNameImagesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
