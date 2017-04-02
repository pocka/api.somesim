package flower

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/pocka/api.somesim/models"
)

// GetFlowersOKCode is the HTTP code returned for type GetFlowersOK
const GetFlowersOKCode int = 200

/*GetFlowersOK 花びら定義

swagger:response getFlowersOK
*/
type GetFlowersOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Flower `json:"body,omitempty"`
}

// NewGetFlowersOK creates GetFlowersOK with default headers values
func NewGetFlowersOK() *GetFlowersOK {
	return &GetFlowersOK{}
}

// WithPayload adds the payload to the get flowers o k response
func (o *GetFlowersOK) WithPayload(payload []*models.Flower) *GetFlowersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get flowers o k response
func (o *GetFlowersOK) SetPayload(payload []*models.Flower) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFlowersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Flower, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}
