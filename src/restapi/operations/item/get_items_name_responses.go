package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pocka/api.somesim/models"
)

// GetItemsNameOKCode is the HTTP code returned for type GetItemsNameOK
const GetItemsNameOKCode int = 200

/*GetItemsNameOK 指定した名前の装備情報

swagger:response getItemsNameOK
*/
type GetItemsNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewGetItemsNameOK creates GetItemsNameOK with default headers values
func NewGetItemsNameOK() *GetItemsNameOK {
	return &GetItemsNameOK{}
}

// WithPayload adds the payload to the get items name o k response
func (o *GetItemsNameOK) WithPayload(payload *models.Item) *GetItemsNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get items name o k response
func (o *GetItemsNameOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetItemsNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetItemsNameNotFoundCode is the HTTP code returned for type GetItemsNameNotFound
const GetItemsNameNotFoundCode int = 404

/*GetItemsNameNotFound 指定された名前の装備情報が存在しない

swagger:response getItemsNameNotFound
*/
type GetItemsNameNotFound struct {
}

// NewGetItemsNameNotFound creates GetItemsNameNotFound with default headers values
func NewGetItemsNameNotFound() *GetItemsNameNotFound {
	return &GetItemsNameNotFound{}
}

// WriteResponse to the client
func (o *GetItemsNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
