// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.ibm.com/ffdl/ffdl-core/restapi/api_v1/restmodels"
)

// DeleteModelReader is a Reader for the DeleteModel structure.
type DeleteModelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteModelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteModelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteModelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteModelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteModelOK creates a DeleteModelOK with default headers values
func NewDeleteModelOK() *DeleteModelOK {
	return &DeleteModelOK{}
}

/*DeleteModelOK handles this case with default header values.

Model deleted successfully.
*/
type DeleteModelOK struct {
	Payload *restmodels.BasicModel
}

func (o *DeleteModelOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}][%d] deleteModelOK  %+v", 200, o.Payload)
}

func (o *DeleteModelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.BasicModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModelUnauthorized creates a DeleteModelUnauthorized with default headers values
func NewDeleteModelUnauthorized() *DeleteModelUnauthorized {
	return &DeleteModelUnauthorized{}
}

/*DeleteModelUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteModelUnauthorized struct {
	Payload *restmodels.Error
}

func (o *DeleteModelUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}][%d] deleteModelUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteModelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteModelNotFound creates a DeleteModelNotFound with default headers values
func NewDeleteModelNotFound() *DeleteModelNotFound {
	return &DeleteModelNotFound{}
}

/*DeleteModelNotFound handles this case with default header values.

The model cannot be found.
*/
type DeleteModelNotFound struct {
	Payload *restmodels.Error
}

func (o *DeleteModelNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}][%d] deleteModelNotFound  %+v", 404, o.Payload)
}

func (o *DeleteModelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
