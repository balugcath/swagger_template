// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swagger_template/src/main_service/swagger/models_server"
)

// DestroyNoContentCode is the HTTP code returned for type DestroyNoContent
const DestroyNoContentCode int = 204

/*DestroyNoContent Deleted

swagger:response destroyNoContent
*/
type DestroyNoContent struct {
}

// NewDestroyNoContent creates DestroyNoContent with default headers values
func NewDestroyNoContent() *DestroyNoContent {

	return &DestroyNoContent{}
}

// WriteResponse to the client
func (o *DestroyNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DestroyDefault error

swagger:response destroyDefault
*/
type DestroyDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models_server.Error `json:"body,omitempty"`
}

// NewDestroyDefault creates DestroyDefault with default headers values
func NewDestroyDefault(code int) *DestroyDefault {
	if code <= 0 {
		code = 500
	}

	return &DestroyDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the destroy default response
func (o *DestroyDefault) WithStatusCode(code int) *DestroyDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the destroy default response
func (o *DestroyDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the destroy default response
func (o *DestroyDefault) WithPayload(payload *models_server.Error) *DestroyDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the destroy default response
func (o *DestroyDefault) SetPayload(payload *models_server.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DestroyDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
