package flower

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFlowersNameParams creates a new GetFlowersNameParams object
// with the default values initialized.
func NewGetFlowersNameParams() GetFlowersNameParams {
	var ()
	return GetFlowersNameParams{}
}

// GetFlowersNameParams contains all the bound params for the get flowers name operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFlowersName
type GetFlowersNameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*花びらの名前
	  Required: true
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetFlowersNameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetFlowersNameParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Name = raw

	return nil
}
