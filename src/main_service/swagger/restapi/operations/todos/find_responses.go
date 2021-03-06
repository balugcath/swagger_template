// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swagger_template/src/main_service/swagger/models_server"
)

// FindOKCode is the HTTP code returned for type FindOK
const FindOKCode int = 200

/*FindOK Ok

swagger:response findOK
*/
type FindOK struct {

	/*
	  In: Body
	*/
	Payload *models_server.Item `json:"body,omitempty"`
}

// NewFindOK creates FindOK with default headers values
func NewFindOK() *FindOK {

	return &FindOK{}
}

// WithPayload adds the payload to the find o k response
func (o *FindOK) WithPayload(payload *models_server.Item) *FindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find o k response
func (o *FindOK) SetPayload(payload *models_server.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindDefault error

swagger:response findDefault
*/
type FindDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models_server.Error `json:"body,omitempty"`
}

// NewFindDefault creates FindDefault with default headers values
func NewFindDefault(code int) *FindDefault {
	if code <= 0 {
		code = 500
	}

	return &FindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find default response
func (o *FindDefault) WithStatusCode(code int) *FindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find default response
func (o *FindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find default response
func (o *FindDefault) WithPayload(payload *models_server.Error) *FindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find default response
func (o *FindDefault) SetPayload(payload *models_server.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
