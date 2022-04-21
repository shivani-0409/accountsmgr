// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-swagger/go-swagger/examples/accountmgr/gen/models"
)

// AddAccountCreatedCode is the HTTP code returned for type AddAccountCreated
const AddAccountCreatedCode int = 201

/*AddAccountCreated Created

swagger:response addAccountCreated
*/
type AddAccountCreated struct {
}

// NewAddAccountCreated creates AddAccountCreated with default headers values
func NewAddAccountCreated() *AddAccountCreated {

	return &AddAccountCreated{}
}

// WriteResponse to the client
func (o *AddAccountCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// AddAccountBadRequestCode is the HTTP code returned for type AddAccountBadRequest
const AddAccountBadRequestCode int = 400

/*AddAccountBadRequest Name should be atleast 3 characters long

swagger:response addAccountBadRequest
*/
type AddAccountBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAccountBadRequest creates AddAccountBadRequest with default headers values
func NewAddAccountBadRequest() *AddAccountBadRequest {

	return &AddAccountBadRequest{}
}

// WithPayload adds the payload to the add account bad request response
func (o *AddAccountBadRequest) WithPayload(payload *models.Error) *AddAccountBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add account bad request response
func (o *AddAccountBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAccountBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddAccountDefault error

swagger:response addAccountDefault
*/
type AddAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAccountDefault creates AddAccountDefault with default headers values
func NewAddAccountDefault(code int) *AddAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &AddAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add account default response
func (o *AddAccountDefault) WithStatusCode(code int) *AddAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add account default response
func (o *AddAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add account default response
func (o *AddAccountDefault) WithPayload(payload *models.Error) *AddAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add account default response
func (o *AddAccountDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
