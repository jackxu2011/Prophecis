// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// PushReportByReportVersionIDOKCode is the HTTP code returned for type PushReportByReportVersionIDOK
const PushReportByReportVersionIDOKCode int = 200

/*PushReportByReportVersionIDOK OK

swagger:response pushReportByReportVersionIdOK
*/
type PushReportByReportVersionIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Event `json:"body,omitempty"`
}

// NewPushReportByReportVersionIDOK creates PushReportByReportVersionIDOK with default headers values
func NewPushReportByReportVersionIDOK() *PushReportByReportVersionIDOK {

	return &PushReportByReportVersionIDOK{}
}

// WithPayload adds the payload to the push report by report version Id o k response
func (o *PushReportByReportVersionIDOK) WithPayload(payload *models.Event) *PushReportByReportVersionIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the push report by report version Id o k response
func (o *PushReportByReportVersionIDOK) SetPayload(payload *models.Event) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PushReportByReportVersionIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PushReportByReportVersionIDUnauthorizedCode is the HTTP code returned for type PushReportByReportVersionIDUnauthorized
const PushReportByReportVersionIDUnauthorizedCode int = 401

/*PushReportByReportVersionIDUnauthorized Unauthorized

swagger:response pushReportByReportVersionIdUnauthorized
*/
type PushReportByReportVersionIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPushReportByReportVersionIDUnauthorized creates PushReportByReportVersionIDUnauthorized with default headers values
func NewPushReportByReportVersionIDUnauthorized() *PushReportByReportVersionIDUnauthorized {

	return &PushReportByReportVersionIDUnauthorized{}
}

// WithPayload adds the payload to the push report by report version Id unauthorized response
func (o *PushReportByReportVersionIDUnauthorized) WithPayload(payload *models.Error) *PushReportByReportVersionIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the push report by report version Id unauthorized response
func (o *PushReportByReportVersionIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PushReportByReportVersionIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PushReportByReportVersionIDNotFoundCode is the HTTP code returned for type PushReportByReportVersionIDNotFound
const PushReportByReportVersionIDNotFoundCode int = 404

/*PushReportByReportVersionIDNotFound Report push fail

swagger:response pushReportByReportVersionIdNotFound
*/
type PushReportByReportVersionIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPushReportByReportVersionIDNotFound creates PushReportByReportVersionIDNotFound with default headers values
func NewPushReportByReportVersionIDNotFound() *PushReportByReportVersionIDNotFound {

	return &PushReportByReportVersionIDNotFound{}
}

// WithPayload adds the payload to the push report by report version Id not found response
func (o *PushReportByReportVersionIDNotFound) WithPayload(payload *models.Error) *PushReportByReportVersionIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the push report by report version Id not found response
func (o *PushReportByReportVersionIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PushReportByReportVersionIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
