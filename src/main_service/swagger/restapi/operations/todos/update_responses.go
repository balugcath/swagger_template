// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"swagger_template/src/main_service/swagger/models_server"
)

// UpdateOKCode is the HTTP code returned for type UpdateOK
const UpdateOKCode int = 200

/*UpdateOK OK

swagger:response updateOK
*/
type UpdateOK struct {

	/*
	  In: Body
	*/
	Payload *models_server.Item `json:"body,omitempty"`
}

// NewUpdateOK creates UpdateOK with default headers values
func NewUpdateOK() *UpdateOK {

	return &UpdateOK{}
}

// WithPayload adds the payload to the update o k response
func (o *UpdateOK) WithPayload(payload *models_server.Item) *UpdateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update o k response
func (o *UpdateOK) SetPayload(payload *models_server.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateDefault error

swagger:response updateDefault
*/
type UpdateDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models_server.Error `json:"body,omitempty"`
}

// NewUpdateDefault creates UpdateDefault with default headers values
func NewUpdateDefault(code int) *UpdateDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update default response
func (o *UpdateDefault) WithStatusCode(code int) *UpdateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update default response
func (o *UpdateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update default response
func (o *UpdateDefault) WithPayload(payload *models_server.Error) *UpdateDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update default response
func (o *UpdateDefault) SetPayload(payload *models_server.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
