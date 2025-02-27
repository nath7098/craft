// Code generated by go-swagger; DO NOT EDIT.

package monitoring

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kilianpaquier/craft/examples/oas_v2/models"
)

// UnsecuredPingReader is a Reader for the UnsecuredPing structure.
type UnsecuredPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnsecuredPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnsecuredPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnsecuredPingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnsecuredPingOK creates a UnsecuredPingOK with default headers values
func NewUnsecuredPingOK() *UnsecuredPingOK {
	return &UnsecuredPingOK{}
}

/*
UnsecuredPingOK describes a response with status code 200, with default header values.

success response returning ping results
*/
type UnsecuredPingOK struct {
	Payload *models.UnsecuredPing
}

// IsSuccess returns true when this unsecured ping o k response has a 2xx status code
func (o *UnsecuredPingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unsecured ping o k response has a 3xx status code
func (o *UnsecuredPingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unsecured ping o k response has a 4xx status code
func (o *UnsecuredPingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unsecured ping o k response has a 5xx status code
func (o *UnsecuredPingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unsecured ping o k response a status code equal to that given
func (o *UnsecuredPingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the unsecured ping o k response
func (o *UnsecuredPingOK) Code() int {
	return 200
}

func (o *UnsecuredPingOK) Error() string {
	return fmt.Sprintf("[GET /unsecured/ping][%d] unsecuredPingOK  %+v", 200, o.Payload)
}

func (o *UnsecuredPingOK) String() string {
	return fmt.Sprintf("[GET /unsecured/ping][%d] unsecuredPingOK  %+v", 200, o.Payload)
}

func (o *UnsecuredPingOK) GetPayload() *models.UnsecuredPing {
	return o.Payload
}

func (o *UnsecuredPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnsecuredPing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnsecuredPingDefault creates a UnsecuredPingDefault with default headers values
func NewUnsecuredPingDefault(code int) *UnsecuredPingDefault {
	return &UnsecuredPingDefault{
		_statusCode: code,
	}
}

/*
UnsecuredPingDefault describes a response with status code -1, with default header values.

default error response
*/
type UnsecuredPingDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this unsecured ping default response has a 2xx status code
func (o *UnsecuredPingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unsecured ping default response has a 3xx status code
func (o *UnsecuredPingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unsecured ping default response has a 4xx status code
func (o *UnsecuredPingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unsecured ping default response has a 5xx status code
func (o *UnsecuredPingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unsecured ping default response a status code equal to that given
func (o *UnsecuredPingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unsecured ping default response
func (o *UnsecuredPingDefault) Code() int {
	return o._statusCode
}

func (o *UnsecuredPingDefault) Error() string {
	return fmt.Sprintf("[GET /unsecured/ping][%d] unsecuredPing default  %+v", o._statusCode, o.Payload)
}

func (o *UnsecuredPingDefault) String() string {
	return fmt.Sprintf("[GET /unsecured/ping][%d] unsecuredPing default  %+v", o._statusCode, o.Payload)
}

func (o *UnsecuredPingDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UnsecuredPingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
