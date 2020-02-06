// Code generated by go-swagger; DO NOT EDIT.

package plugin_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// DeletePluginReader is a Reader for the DeletePlugin structure.
type DeletePluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePluginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeletePluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeletePluginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePluginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeletePluginConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePluginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePluginOK creates a DeletePluginOK with default headers values
func NewDeletePluginOK() *DeletePluginOK {
	return &DeletePluginOK{}
}

/*DeletePluginOK handles this case with default header values.

OK
*/
type DeletePluginOK struct {
	Payload *models.APIResultVoid
}

func (o *DeletePluginOK) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginOK  %+v", 200, o.Payload)
}

func (o *DeletePluginOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *DeletePluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginBadRequest creates a DeletePluginBadRequest with default headers values
func NewDeletePluginBadRequest() *DeletePluginBadRequest {
	return &DeletePluginBadRequest{}
}

/*DeletePluginBadRequest handles this case with default header values.

Bad Request
*/
type DeletePluginBadRequest struct {
	Payload *models.APIResult
}

func (o *DeletePluginBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginBadRequest  %+v", 400, o.Payload)
}

func (o *DeletePluginBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeletePluginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginUnauthorized creates a DeletePluginUnauthorized with default headers values
func NewDeletePluginUnauthorized() *DeletePluginUnauthorized {
	return &DeletePluginUnauthorized{}
}

/*DeletePluginUnauthorized handles this case with default header values.

Unauthorized
*/
type DeletePluginUnauthorized struct {
	Payload *models.APIResult
}

func (o *DeletePluginUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginUnauthorized  %+v", 401, o.Payload)
}

func (o *DeletePluginUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeletePluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginForbidden creates a DeletePluginForbidden with default headers values
func NewDeletePluginForbidden() *DeletePluginForbidden {
	return &DeletePluginForbidden{}
}

/*DeletePluginForbidden handles this case with default header values.

Forbidden
*/
type DeletePluginForbidden struct {
	Payload *models.APIResult
}

func (o *DeletePluginForbidden) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginForbidden  %+v", 403, o.Payload)
}

func (o *DeletePluginForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeletePluginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginNotFound creates a DeletePluginNotFound with default headers values
func NewDeletePluginNotFound() *DeletePluginNotFound {
	return &DeletePluginNotFound{}
}

/*DeletePluginNotFound handles this case with default header values.

Not Found
*/
type DeletePluginNotFound struct {
	Payload *models.APIResult
}

func (o *DeletePluginNotFound) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginNotFound  %+v", 404, o.Payload)
}

func (o *DeletePluginNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeletePluginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginConflict creates a DeletePluginConflict with default headers values
func NewDeletePluginConflict() *DeletePluginConflict {
	return &DeletePluginConflict{}
}

/*DeletePluginConflict handles this case with default header values.

Conflict
*/
type DeletePluginConflict struct {
	Payload *models.APIResult
}

func (o *DeletePluginConflict) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginConflict  %+v", 409, o.Payload)
}

func (o *DeletePluginConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeletePluginConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePluginInternalServerError creates a DeletePluginInternalServerError with default headers values
func NewDeletePluginInternalServerError() *DeletePluginInternalServerError {
	return &DeletePluginInternalServerError{}
}

/*DeletePluginInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeletePluginInternalServerError struct {
	Payload *models.APIResult
}

func (o *DeletePluginInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /plugins/{id}][%d] deletePluginInternalServerError  %+v", 500, o.Payload)
}

func (o *DeletePluginInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeletePluginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
