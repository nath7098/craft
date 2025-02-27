// Code generated by go-swagger; DO NOT EDIT.

package monitoring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kilianpaquier/craft/examples/oas_v2/models"
)

// UnsecuredPingOKCode is the HTTP code returned for type UnsecuredPingOK
const UnsecuredPingOKCode int = 200

/*
UnsecuredPingOK success response returning ping results

swagger:response unsecuredPingOK
*/
type UnsecuredPingOK struct {

	/*
	  In: Body
	*/
	Payload *models.UnsecuredPing `json:"body,omitempty"`
}

// NewUnsecuredPingOK creates UnsecuredPingOK with default headers values
func NewUnsecuredPingOK() *UnsecuredPingOK {

	return &UnsecuredPingOK{}
}

// WithPayload adds the payload to the unsecured ping o k response
func (o *UnsecuredPingOK) WithPayload(payload *models.UnsecuredPing) *UnsecuredPingOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unsecured ping o k response
func (o *UnsecuredPingOK) SetPayload(payload *models.UnsecuredPing) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnsecuredPingOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UnsecuredPingDefault default error response

swagger:response unsecuredPingDefault
*/
type UnsecuredPingDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnsecuredPingDefault creates UnsecuredPingDefault with default headers values
func NewUnsecuredPingDefault(code int) *UnsecuredPingDefault {
	if code <= 0 {
		code = 500
	}

	return &UnsecuredPingDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the unsecured ping default response
func (o *UnsecuredPingDefault) WithStatusCode(code int) *UnsecuredPingDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the unsecured ping default response
func (o *UnsecuredPingDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the unsecured ping default response
func (o *UnsecuredPingDefault) WithPayload(payload *models.Error) *UnsecuredPingDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unsecured ping default response
func (o *UnsecuredPingDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnsecuredPingDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
