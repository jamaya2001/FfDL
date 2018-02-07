// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.ibm.com/ffdl/ffdl-core/restapi/api_v1/restmodels"
)

// NewCreateEventEndpointParams creates a new CreateEventEndpointParams object
// with the default values initialized.
func NewCreateEventEndpointParams() *CreateEventEndpointParams {
	var ()
	return &CreateEventEndpointParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateEventEndpointParamsWithTimeout creates a new CreateEventEndpointParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateEventEndpointParamsWithTimeout(timeout time.Duration) *CreateEventEndpointParams {
	var ()
	return &CreateEventEndpointParams{

		timeout: timeout,
	}
}

// NewCreateEventEndpointParamsWithContext creates a new CreateEventEndpointParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateEventEndpointParamsWithContext(ctx context.Context) *CreateEventEndpointParams {
	var ()
	return &CreateEventEndpointParams{

		Context: ctx,
	}
}

// NewCreateEventEndpointParamsWithHTTPClient creates a new CreateEventEndpointParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateEventEndpointParamsWithHTTPClient(client *http.Client) *CreateEventEndpointParams {
	var ()
	return &CreateEventEndpointParams{
		HTTPClient: client,
	}
}

/*CreateEventEndpointParams contains all the parameters to send to the API endpoint
for the create event endpoint operation typically these are written to a http.Request
*/
type CreateEventEndpointParams struct {

	/*CallbackURL
	  The URL that should be notified when an event triggers.

	*/
	CallbackURL *restmodels.EventEndpointRegistration
	/*EndpointID
	  The id of the endpoint.

	*/
	EndpointID string
	/*EventType
	  The type of event.

	*/
	EventType string
	/*ModelID
	  The id of the model.

	*/
	ModelID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create event endpoint params
func (o *CreateEventEndpointParams) WithTimeout(timeout time.Duration) *CreateEventEndpointParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create event endpoint params
func (o *CreateEventEndpointParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create event endpoint params
func (o *CreateEventEndpointParams) WithContext(ctx context.Context) *CreateEventEndpointParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create event endpoint params
func (o *CreateEventEndpointParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create event endpoint params
func (o *CreateEventEndpointParams) WithHTTPClient(client *http.Client) *CreateEventEndpointParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create event endpoint params
func (o *CreateEventEndpointParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallbackURL adds the callbackURL to the create event endpoint params
func (o *CreateEventEndpointParams) WithCallbackURL(callbackURL *restmodels.EventEndpointRegistration) *CreateEventEndpointParams {
	o.SetCallbackURL(callbackURL)
	return o
}

// SetCallbackURL adds the callbackUrl to the create event endpoint params
func (o *CreateEventEndpointParams) SetCallbackURL(callbackURL *restmodels.EventEndpointRegistration) {
	o.CallbackURL = callbackURL
}

// WithEndpointID adds the endpointID to the create event endpoint params
func (o *CreateEventEndpointParams) WithEndpointID(endpointID string) *CreateEventEndpointParams {
	o.SetEndpointID(endpointID)
	return o
}

// SetEndpointID adds the endpointId to the create event endpoint params
func (o *CreateEventEndpointParams) SetEndpointID(endpointID string) {
	o.EndpointID = endpointID
}

// WithEventType adds the eventType to the create event endpoint params
func (o *CreateEventEndpointParams) WithEventType(eventType string) *CreateEventEndpointParams {
	o.SetEventType(eventType)
	return o
}

// SetEventType adds the eventType to the create event endpoint params
func (o *CreateEventEndpointParams) SetEventType(eventType string) {
	o.EventType = eventType
}

// WithModelID adds the modelID to the create event endpoint params
func (o *CreateEventEndpointParams) WithModelID(modelID string) *CreateEventEndpointParams {
	o.SetModelID(modelID)
	return o
}

// SetModelID adds the modelId to the create event endpoint params
func (o *CreateEventEndpointParams) SetModelID(modelID string) {
	o.ModelID = modelID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateEventEndpointParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CallbackURL == nil {
		o.CallbackURL = new(restmodels.EventEndpointRegistration)
	}

	if err := r.SetBodyParam(o.CallbackURL); err != nil {
		return err
	}

	// path param endpoint_id
	if err := r.SetPathParam("endpoint_id", o.EndpointID); err != nil {
		return err
	}

	// path param event_type
	if err := r.SetPathParam("event_type", o.EventType); err != nil {
		return err
	}

	// path param model_id
	if err := r.SetPathParam("model_id", o.ModelID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
