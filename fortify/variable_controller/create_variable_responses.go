// Code generated by go-swagger; DO NOT EDIT.

package variable_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreateVariableReader is a Reader for the CreateVariable structure.
type CreateVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVariableCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVariableBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateVariableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateVariableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateVariableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateVariableConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateVariableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVariableCreated creates a CreateVariableCreated with default headers values
func NewCreateVariableCreated() *CreateVariableCreated {
	return &CreateVariableCreated{}
}

/*CreateVariableCreated handles this case with default header values.

Created
*/
type CreateVariableCreated struct {
	Payload *models.APIResultVariable
}

func (o *CreateVariableCreated) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableCreated  %+v", 201, o.Payload)
}

func (o *CreateVariableCreated) GetPayload() *models.APIResultVariable {
	return o.Payload
}

func (o *CreateVariableCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableBadRequest creates a CreateVariableBadRequest with default headers values
func NewCreateVariableBadRequest() *CreateVariableBadRequest {
	return &CreateVariableBadRequest{}
}

/*CreateVariableBadRequest handles this case with default header values.

Bad Request
*/
type CreateVariableBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateVariableBadRequest) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableBadRequest  %+v", 400, o.Payload)
}

func (o *CreateVariableBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateVariableBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableUnauthorized creates a CreateVariableUnauthorized with default headers values
func NewCreateVariableUnauthorized() *CreateVariableUnauthorized {
	return &CreateVariableUnauthorized{}
}

/*CreateVariableUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateVariableUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateVariableUnauthorized) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateVariableUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateVariableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableForbidden creates a CreateVariableForbidden with default headers values
func NewCreateVariableForbidden() *CreateVariableForbidden {
	return &CreateVariableForbidden{}
}

/*CreateVariableForbidden handles this case with default header values.

Forbidden
*/
type CreateVariableForbidden struct {
	Payload *models.APIResult
}

func (o *CreateVariableForbidden) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableForbidden  %+v", 403, o.Payload)
}

func (o *CreateVariableForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateVariableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableNotFound creates a CreateVariableNotFound with default headers values
func NewCreateVariableNotFound() *CreateVariableNotFound {
	return &CreateVariableNotFound{}
}

/*CreateVariableNotFound handles this case with default header values.

Not Found
*/
type CreateVariableNotFound struct {
	Payload *models.APIResult
}

func (o *CreateVariableNotFound) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableNotFound  %+v", 404, o.Payload)
}

func (o *CreateVariableNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateVariableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableConflict creates a CreateVariableConflict with default headers values
func NewCreateVariableConflict() *CreateVariableConflict {
	return &CreateVariableConflict{}
}

/*CreateVariableConflict handles this case with default header values.

Conflict
*/
type CreateVariableConflict struct {
	Payload *models.APIResult
}

func (o *CreateVariableConflict) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableConflict  %+v", 409, o.Payload)
}

func (o *CreateVariableConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateVariableConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVariableInternalServerError creates a CreateVariableInternalServerError with default headers values
func NewCreateVariableInternalServerError() *CreateVariableInternalServerError {
	return &CreateVariableInternalServerError{}
}

/*CreateVariableInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateVariableInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateVariableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /variables][%d] createVariableInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateVariableInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateVariableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
