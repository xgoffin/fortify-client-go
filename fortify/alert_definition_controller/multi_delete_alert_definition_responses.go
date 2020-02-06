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

// MultiDeleteAlertDefinitionReader is a Reader for the MultiDeleteAlertDefinition structure.
type MultiDeleteAlertDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiDeleteAlertDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiDeleteAlertDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMultiDeleteAlertDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMultiDeleteAlertDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMultiDeleteAlertDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMultiDeleteAlertDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMultiDeleteAlertDefinitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMultiDeleteAlertDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMultiDeleteAlertDefinitionOK creates a MultiDeleteAlertDefinitionOK with default headers values
func NewMultiDeleteAlertDefinitionOK() *MultiDeleteAlertDefinitionOK {
	return &MultiDeleteAlertDefinitionOK{}
}

/*MultiDeleteAlertDefinitionOK handles this case with default header values.

OK
*/
type MultiDeleteAlertDefinitionOK struct {
	Payload *models.APIResultVoid
}

func (o *MultiDeleteAlertDefinitionOK) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionOK  %+v", 200, o.Payload)
}

func (o *MultiDeleteAlertDefinitionOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAlertDefinitionBadRequest creates a MultiDeleteAlertDefinitionBadRequest with default headers values
func NewMultiDeleteAlertDefinitionBadRequest() *MultiDeleteAlertDefinitionBadRequest {
	return &MultiDeleteAlertDefinitionBadRequest{}
}

/*MultiDeleteAlertDefinitionBadRequest handles this case with default header values.

Bad Request
*/
type MultiDeleteAlertDefinitionBadRequest struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAlertDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *MultiDeleteAlertDefinitionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAlertDefinitionUnauthorized creates a MultiDeleteAlertDefinitionUnauthorized with default headers values
func NewMultiDeleteAlertDefinitionUnauthorized() *MultiDeleteAlertDefinitionUnauthorized {
	return &MultiDeleteAlertDefinitionUnauthorized{}
}

/*MultiDeleteAlertDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type MultiDeleteAlertDefinitionUnauthorized struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAlertDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *MultiDeleteAlertDefinitionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAlertDefinitionForbidden creates a MultiDeleteAlertDefinitionForbidden with default headers values
func NewMultiDeleteAlertDefinitionForbidden() *MultiDeleteAlertDefinitionForbidden {
	return &MultiDeleteAlertDefinitionForbidden{}
}

/*MultiDeleteAlertDefinitionForbidden handles this case with default header values.

Forbidden
*/
type MultiDeleteAlertDefinitionForbidden struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAlertDefinitionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *MultiDeleteAlertDefinitionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAlertDefinitionNotFound creates a MultiDeleteAlertDefinitionNotFound with default headers values
func NewMultiDeleteAlertDefinitionNotFound() *MultiDeleteAlertDefinitionNotFound {
	return &MultiDeleteAlertDefinitionNotFound{}
}

/*MultiDeleteAlertDefinitionNotFound handles this case with default header values.

Not Found
*/
type MultiDeleteAlertDefinitionNotFound struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAlertDefinitionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *MultiDeleteAlertDefinitionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAlertDefinitionConflict creates a MultiDeleteAlertDefinitionConflict with default headers values
func NewMultiDeleteAlertDefinitionConflict() *MultiDeleteAlertDefinitionConflict {
	return &MultiDeleteAlertDefinitionConflict{}
}

/*MultiDeleteAlertDefinitionConflict handles this case with default header values.

Conflict
*/
type MultiDeleteAlertDefinitionConflict struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAlertDefinitionConflict) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionConflict  %+v", 409, o.Payload)
}

func (o *MultiDeleteAlertDefinitionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAlertDefinitionInternalServerError creates a MultiDeleteAlertDefinitionInternalServerError with default headers values
func NewMultiDeleteAlertDefinitionInternalServerError() *MultiDeleteAlertDefinitionInternalServerError {
	return &MultiDeleteAlertDefinitionInternalServerError{}
}

/*MultiDeleteAlertDefinitionInternalServerError handles this case with default header values.

Internal Server Error
*/
type MultiDeleteAlertDefinitionInternalServerError struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAlertDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /alertDefinitions][%d] multiDeleteAlertDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *MultiDeleteAlertDefinitionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAlertDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
