// Code generated by go-swagger; DO NOT EDIT.

package alert_definition_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// DeleteAlertDefinitionReader is a Reader for the DeleteAlertDefinition structure.
type DeleteAlertDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAlertDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAlertDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAlertDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteAlertDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAlertDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAlertDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteAlertDefinitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAlertDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAlertDefinitionOK creates a DeleteAlertDefinitionOK with default headers values
func NewDeleteAlertDefinitionOK() *DeleteAlertDefinitionOK {
	return &DeleteAlertDefinitionOK{}
}

/*DeleteAlertDefinitionOK handles this case with default header values.

OK
*/
type DeleteAlertDefinitionOK struct {
	Payload *models.APIResultVoid
}

func (o *DeleteAlertDefinitionOK) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionOK  %+v", 200, o.Payload)
}

func (o *DeleteAlertDefinitionOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *DeleteAlertDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefinitionBadRequest creates a DeleteAlertDefinitionBadRequest with default headers values
func NewDeleteAlertDefinitionBadRequest() *DeleteAlertDefinitionBadRequest {
	return &DeleteAlertDefinitionBadRequest{}
}

/*DeleteAlertDefinitionBadRequest handles this case with default header values.

Bad Request
*/
type DeleteAlertDefinitionBadRequest struct {
	Payload *models.APIResult
}

func (o *DeleteAlertDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAlertDefinitionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteAlertDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefinitionUnauthorized creates a DeleteAlertDefinitionUnauthorized with default headers values
func NewDeleteAlertDefinitionUnauthorized() *DeleteAlertDefinitionUnauthorized {
	return &DeleteAlertDefinitionUnauthorized{}
}

/*DeleteAlertDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteAlertDefinitionUnauthorized struct {
	Payload *models.APIResult
}

func (o *DeleteAlertDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteAlertDefinitionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteAlertDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefinitionForbidden creates a DeleteAlertDefinitionForbidden with default headers values
func NewDeleteAlertDefinitionForbidden() *DeleteAlertDefinitionForbidden {
	return &DeleteAlertDefinitionForbidden{}
}

/*DeleteAlertDefinitionForbidden handles this case with default header values.

Forbidden
*/
type DeleteAlertDefinitionForbidden struct {
	Payload *models.APIResult
}

func (o *DeleteAlertDefinitionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *DeleteAlertDefinitionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteAlertDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefinitionNotFound creates a DeleteAlertDefinitionNotFound with default headers values
func NewDeleteAlertDefinitionNotFound() *DeleteAlertDefinitionNotFound {
	return &DeleteAlertDefinitionNotFound{}
}

/*DeleteAlertDefinitionNotFound handles this case with default header values.

Not Found
*/
type DeleteAlertDefinitionNotFound struct {
	Payload *models.APIResult
}

func (o *DeleteAlertDefinitionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAlertDefinitionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteAlertDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefinitionConflict creates a DeleteAlertDefinitionConflict with default headers values
func NewDeleteAlertDefinitionConflict() *DeleteAlertDefinitionConflict {
	return &DeleteAlertDefinitionConflict{}
}

/*DeleteAlertDefinitionConflict handles this case with default header values.

Conflict
*/
type DeleteAlertDefinitionConflict struct {
	Payload *models.APIResult
}

func (o *DeleteAlertDefinitionConflict) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionConflict  %+v", 409, o.Payload)
}

func (o *DeleteAlertDefinitionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteAlertDefinitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAlertDefinitionInternalServerError creates a DeleteAlertDefinitionInternalServerError with default headers values
func NewDeleteAlertDefinitionInternalServerError() *DeleteAlertDefinitionInternalServerError {
	return &DeleteAlertDefinitionInternalServerError{}
}

/*DeleteAlertDefinitionInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteAlertDefinitionInternalServerError struct {
	Payload *models.APIResult
}

func (o *DeleteAlertDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions/{id}][%d] deleteAlertDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAlertDefinitionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteAlertDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
