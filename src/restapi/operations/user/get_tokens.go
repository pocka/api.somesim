package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	validate "github.com/go-openapi/validate"

	"github.com/pocka/api.somesim/models"
)

// GetTokensHandlerFunc turns a function with the right signature into a get tokens handler
type GetTokensHandlerFunc func(GetTokensParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTokensHandlerFunc) Handle(params GetTokensParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetTokensHandler interface for that can handle valid get tokens params
type GetTokensHandler interface {
	Handle(GetTokensParams, *models.Principal) middleware.Responder
}

// NewGetTokens creates a new http.Handler for the get tokens operation
func NewGetTokens(ctx *middleware.Context, handler GetTokensHandler) *GetTokens {
	return &GetTokens{Context: ctx, Handler: handler}
}

/*GetTokens swagger:route GET /tokens User getTokens

アクセストークンとリフレッシュトークンを取得する

アクセストークンとリフレッシュトークンを取得する

ユーザ名/パスワードを渡すBASIC認証方式と
古いリフレッシュトークンを渡す更新方式がある


*/
type GetTokens struct {
	Context *middleware.Context
	Handler GetTokensHandler
}

func (o *GetTokens) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTokensParams()

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

// GetTokensOKBody get tokens o k body
// swagger:model GetTokensOKBody
type GetTokensOKBody struct {

	// アクセストークン(JWT)
	// Required: true
	AccessToken *string `json:"access_token"`

	// リフレッシュトークン(JWT)
	// Required: true
	RefreshToken *string `json:"refresh_token"`
}

// Validate validates this get tokens o k body
func (o *GetTokensOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateRefreshToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTokensOKBody) validateAccessToken(formats strfmt.Registry) error {

	if err := validate.Required("getTokensOK"+"."+"access_token", "body", o.AccessToken); err != nil {
		return err
	}

	return nil
}

func (o *GetTokensOKBody) validateRefreshToken(formats strfmt.Registry) error {

	if err := validate.Required("getTokensOK"+"."+"refresh_token", "body", o.RefreshToken); err != nil {
		return err
	}

	return nil
}
