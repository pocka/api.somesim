package flower

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutFlowersNameNoContentCode is the HTTP code returned for type PutFlowersNameNoContent
const PutFlowersNameNoContentCode int = 204

/*PutFlowersNameNoContent 更新に成功

swagger:response putFlowersNameNoContent
*/
type PutFlowersNameNoContent struct {
}

// NewPutFlowersNameNoContent creates PutFlowersNameNoContent with default headers values
func NewPutFlowersNameNoContent() *PutFlowersNameNoContent {
	return &PutFlowersNameNoContent{}
}

// WriteResponse to the client
func (o *PutFlowersNameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

// PutFlowersNameNotFoundCode is the HTTP code returned for type PutFlowersNameNotFound
const PutFlowersNameNotFoundCode int = 404

/*PutFlowersNameNotFound 指定された名前の花びらが存在しない

swagger:response putFlowersNameNotFound
*/
type PutFlowersNameNotFound struct {
}

// NewPutFlowersNameNotFound creates PutFlowersNameNotFound with default headers values
func NewPutFlowersNameNotFound() *PutFlowersNameNotFound {
	return &PutFlowersNameNotFound{}
}

// WriteResponse to the client
func (o *PutFlowersNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}