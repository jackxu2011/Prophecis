// Code generated by go-swagger; DO NOT EDIT.

package model_storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// ListModelsByGroupIDOKCode is the HTTP code returned for type ListModelsByGroupIDOK
const ListModelsByGroupIDOKCode int = 200

/*ListModelsByGroupIDOK OK

swagger:response listModelsByGroupIdOK
*/
type ListModelsByGroupIDOK struct {

	/*
	  In: Body
	*/
	Payload models.Models `json:"body,omitempty"`
}

// NewListModelsByGroupIDOK creates ListModelsByGroupIDOK with default headers values
func NewListModelsByGroupIDOK() *ListModelsByGroupIDOK {

	return &ListModelsByGroupIDOK{}
}

// WithPayload adds the payload to the list models by group Id o k response
func (o *ListModelsByGroupIDOK) WithPayload(payload models.Models) *ListModelsByGroupIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list models by group Id o k response
func (o *ListModelsByGroupIDOK) SetPayload(payload models.Models) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModelsByGroupIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.Models{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListModelsByGroupIDUnauthorizedCode is the HTTP code returned for type ListModelsByGroupIDUnauthorized
const ListModelsByGroupIDUnauthorizedCode int = 401

/*ListModelsByGroupIDUnauthorized Unauthorized

swagger:response listModelsByGroupIdUnauthorized
*/
type ListModelsByGroupIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListModelsByGroupIDUnauthorized creates ListModelsByGroupIDUnauthorized with default headers values
func NewListModelsByGroupIDUnauthorized() *ListModelsByGroupIDUnauthorized {

	return &ListModelsByGroupIDUnauthorized{}
}

// WithPayload adds the payload to the list models by group Id unauthorized response
func (o *ListModelsByGroupIDUnauthorized) WithPayload(payload *models.Error) *ListModelsByGroupIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list models by group Id unauthorized response
func (o *ListModelsByGroupIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModelsByGroupIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListModelsByGroupIDNotFoundCode is the HTTP code returned for type ListModelsByGroupIDNotFound
const ListModelsByGroupIDNotFoundCode int = 404

/*ListModelsByGroupIDNotFound Model create failed

swagger:response listModelsByGroupIdNotFound
*/
type ListModelsByGroupIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListModelsByGroupIDNotFound creates ListModelsByGroupIDNotFound with default headers values
func NewListModelsByGroupIDNotFound() *ListModelsByGroupIDNotFound {

	return &ListModelsByGroupIDNotFound{}
}

// WithPayload adds the payload to the list models by group Id not found response
func (o *ListModelsByGroupIDNotFound) WithPayload(payload *models.Error) *ListModelsByGroupIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list models by group Id not found response
func (o *ListModelsByGroupIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListModelsByGroupIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
