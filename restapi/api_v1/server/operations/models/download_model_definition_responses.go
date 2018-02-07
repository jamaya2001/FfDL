// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.ibm.com/ffdl/ffdl-core/restapi/api_v1/restmodels"
)

// DownloadModelDefinitionOKCode is the HTTP code returned for type DownloadModelDefinitionOK
const DownloadModelDefinitionOKCode int = 200

/*DownloadModelDefinitionOK Model definition

swagger:response downloadModelDefinitionOK
*/
type DownloadModelDefinitionOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadModelDefinitionOK creates DownloadModelDefinitionOK with default headers values
func NewDownloadModelDefinitionOK() *DownloadModelDefinitionOK {
	return &DownloadModelDefinitionOK{}
}

// WithPayload adds the payload to the download model definition o k response
func (o *DownloadModelDefinitionOK) WithPayload(payload io.ReadCloser) *DownloadModelDefinitionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download model definition o k response
func (o *DownloadModelDefinitionOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadModelDefinitionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// DownloadModelDefinitionUnauthorizedCode is the HTTP code returned for type DownloadModelDefinitionUnauthorized
const DownloadModelDefinitionUnauthorizedCode int = 401

/*DownloadModelDefinitionUnauthorized Unauthorized

swagger:response downloadModelDefinitionUnauthorized
*/
type DownloadModelDefinitionUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewDownloadModelDefinitionUnauthorized creates DownloadModelDefinitionUnauthorized with default headers values
func NewDownloadModelDefinitionUnauthorized() *DownloadModelDefinitionUnauthorized {
	return &DownloadModelDefinitionUnauthorized{}
}

// WithPayload adds the payload to the download model definition unauthorized response
func (o *DownloadModelDefinitionUnauthorized) WithPayload(payload *restmodels.Error) *DownloadModelDefinitionUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download model definition unauthorized response
func (o *DownloadModelDefinitionUnauthorized) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadModelDefinitionUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadModelDefinitionNotFoundCode is the HTTP code returned for type DownloadModelDefinitionNotFound
const DownloadModelDefinitionNotFoundCode int = 404

/*DownloadModelDefinitionNotFound The model cannot be found.

swagger:response downloadModelDefinitionNotFound
*/
type DownloadModelDefinitionNotFound struct {

	/*
	  In: Body
	*/
	Payload *restmodels.Error `json:"body,omitempty"`
}

// NewDownloadModelDefinitionNotFound creates DownloadModelDefinitionNotFound with default headers values
func NewDownloadModelDefinitionNotFound() *DownloadModelDefinitionNotFound {
	return &DownloadModelDefinitionNotFound{}
}

// WithPayload adds the payload to the download model definition not found response
func (o *DownloadModelDefinitionNotFound) WithPayload(payload *restmodels.Error) *DownloadModelDefinitionNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download model definition not found response
func (o *DownloadModelDefinitionNotFound) SetPayload(payload *restmodels.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadModelDefinitionNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
