package internal_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetDocsOKCode is the HTTP code returned for type GetDocsOK
const GetDocsOKCode int = 200

/*GetDocsOK APIドキュメント(ReDoc)


swagger:response getDocsOK
*/
type GetDocsOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewGetDocsOK creates GetDocsOK with default headers values
func NewGetDocsOK() *GetDocsOK {
	return &GetDocsOK{}
}

// WithPayload adds the payload to the get docs o k response
func (o *GetDocsOK) WithPayload(payload string) *GetDocsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get docs o k response
func (o *GetDocsOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDocsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetDocsNotFoundCode is the HTTP code returned for type GetDocsNotFound
const GetDocsNotFoundCode int = 404

/*GetDocsNotFound ファイルが存在しない


swagger:response getDocsNotFound
*/
type GetDocsNotFound struct {
}

// NewGetDocsNotFound creates GetDocsNotFound with default headers values
func NewGetDocsNotFound() *GetDocsNotFound {
	return &GetDocsNotFound{}
}

// WriteResponse to the client
func (o *GetDocsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}