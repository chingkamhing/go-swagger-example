// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"go-swagger-example/gen/models"
)

// MyselfOKCode is the HTTP code returned for type MyselfOK
const MyselfOKCode int = 200

/*
MyselfOK Success, return user's detail info

swagger:response myselfOK
*/
type MyselfOK struct {

	/*
	  In: Body
	*/
	Payload *models.UserAccount `json:"body,omitempty"`
}

// NewMyselfOK creates MyselfOK with default headers values
func NewMyselfOK() *MyselfOK {

	return &MyselfOK{}
}

// WithPayload adds the payload to the myself o k response
func (o *MyselfOK) WithPayload(payload *models.UserAccount) *MyselfOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the myself o k response
func (o *MyselfOK) SetPayload(payload *models.UserAccount) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MyselfOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
MyselfDefault Error

swagger:response myselfDefault
*/
type MyselfDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMyselfDefault creates MyselfDefault with default headers values
func NewMyselfDefault(code int) *MyselfDefault {
	if code <= 0 {
		code = 500
	}

	return &MyselfDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the myself default response
func (o *MyselfDefault) WithStatusCode(code int) *MyselfDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the myself default response
func (o *MyselfDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the myself default response
func (o *MyselfDefault) WithPayload(payload *models.Error) *MyselfDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the myself default response
func (o *MyselfDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MyselfDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
