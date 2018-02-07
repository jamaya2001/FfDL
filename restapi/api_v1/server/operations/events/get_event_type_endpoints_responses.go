// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.ibm.com/ffdl/ffdl-core/restapi/api_v1/restmodels"
)

// GetEventTypeEndpointsOKCode is the HTTP code returned for type GetEventTypeEndpointsOK
const GetEventTypeEndpointsOKCode int = 200

/*GetEventTypeEndpointsOK Event endpoints

swagger:response getEventTypeEndpointsOK
*/
type GetEventTypeEndpointsOK struct {

	/*
	  In: Body
	*/
	Payload *restmodels.EndpointList `json:"body,omitempty"`
}

// NewGetEventTypeEndpointsOK creates GetEventTypeEndpointsOK with default headers values
func NewGetEventTypeEndpointsOK() *GetEventTypeEndpointsOK {
	return &GetEventTypeEndpointsOK{}
}

// WithPayload adds the payload to the get event type endpoints o k response
func (o *GetEventTypeEndpointsOK) WithPayload(payload *restmodels.EndpointList) *GetEventTypeEndpointsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type endpoints o k response
func (o *GetEventTypeEndpointsOK) SetPayload(payload *restmodels.EndpointList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeEndpointsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeEndpointsUnauthorizedCode is the HTTP code returned for type GetEventTypeEndpointsUnauthorized
const GetEventTypeEndpointsUnauthorizedCode int = 401

/*GetEventTypeEndpointsUnauthorized Unauthorized

swagger:response getEventTypeEndpointsUnauthorized
*/
type GetEventTypeEndpointsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewGetEventTypeEndpointsUnauthorized creates GetEventTypeEndpointsUnauthorized with default headers values
func NewGetEventTypeEndpointsUnauthorized() *GetEventTypeEndpointsUnauthorized {
	return &GetEventTypeEndpointsUnauthorized{}
}

// WithPayload adds the payload to the get event type endpoints unauthorized response
func (o *GetEventTypeEndpointsUnauthorized) WithPayload(payload *restmodels.Error) *GetEventTypeEndpointsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type endpoints unauthorized response
func (o *GetEventTypeEndpointsUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeEndpointsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeEndpointsNotFoundCode is the HTTP code returned for type GetEventTypeEndpointsNotFound
const GetEventTypeEndpointsNotFoundCode int = 404

/*GetEventTypeEndpointsNotFound The model or event type cannot be found.

swagger:response getEventTypeEndpointsNotFound
*/
type GetEventTypeEndpointsNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewGetEventTypeEndpointsNotFound creates GetEventTypeEndpointsNotFound with default headers values
func NewGetEventTypeEndpointsNotFound() *GetEventTypeEndpointsNotFound {
	return &GetEventTypeEndpointsNotFound{}
}

// WithPayload adds the payload to the get event type endpoints not found response
func (o *GetEventTypeEndpointsNotFound) WithPayload(payload *restmodels.Error) *GetEventTypeEndpointsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type endpoints not found response
func (o *GetEventTypeEndpointsNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeEndpointsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventTypeEndpointsGoneCode is the HTTP code returned for type GetEventTypeEndpointsGone
const GetEventTypeEndpointsGoneCode int = 410

/*GetEventTypeEndpointsGone If the trained model storage time has expired and it has been deleted. It only gets deleted if it not stored on an external data store.

swagger:response getEventTypeEndpointsGone
*/
type GetEventTypeEndpointsGone struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewGetEventTypeEndpointsGone creates GetEventTypeEndpointsGone with default headers values
func NewGetEventTypeEndpointsGone() *GetEventTypeEndpointsGone {
	return &GetEventTypeEndpointsGone{}
}

// WithPayload adds the payload to the get event type endpoints gone response
func (o *GetEventTypeEndpointsGone) WithPayload(payload *restmodels.Error) *GetEventTypeEndpointsGone {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event type endpoints gone response
func (o *GetEventTypeEndpointsGone) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventTypeEndpointsGone) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(410)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
