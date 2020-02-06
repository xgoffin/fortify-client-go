// Code generated by go-swagger; DO NOT EDIT.

package project_version_of_auth_entity_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// AssignProjectVersionOfAuthEntityReader is a Reader for the AssignProjectVersionOfAuthEntity structure.
type AssignProjectVersionOfAuthEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssignProjectVersionOfAuthEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAssignProjectVersionOfAuthEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAssignProjectVersionOfAuthEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAssignProjectVersionOfAuthEntityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAssignProjectVersionOfAuthEntityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAssignProjectVersionOfAuthEntityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAssignProjectVersionOfAuthEntityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAssignProjectVersionOfAuthEntityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAssignProjectVersionOfAuthEntityOK creates a AssignProjectVersionOfAuthEntityOK with default headers values
func NewAssignProjectVersionOfAuthEntityOK() *AssignProjectVersionOfAuthEntityOK {
	return &AssignProjectVersionOfAuthEntityOK{}
}

/*AssignProjectVersionOfAuthEntityOK handles this case with default header values.

OK
*/
type AssignProjectVersionOfAuthEntityOK struct {
	Payload *models.APIResultVoid
}

func (o *AssignProjectVersionOfAuthEntityOK) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityOK  %+v", 200, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignProjectVersionOfAuthEntityBadRequest creates a AssignProjectVersionOfAuthEntityBadRequest with default headers values
func NewAssignProjectVersionOfAuthEntityBadRequest() *AssignProjectVersionOfAuthEntityBadRequest {
	return &AssignProjectVersionOfAuthEntityBadRequest{}
}

/*AssignProjectVersionOfAuthEntityBadRequest handles this case with default header values.

Bad Request
*/
type AssignProjectVersionOfAuthEntityBadRequest struct {
	Payload *models.APIResult
}

func (o *AssignProjectVersionOfAuthEntityBadRequest) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityBadRequest  %+v", 400, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignProjectVersionOfAuthEntityUnauthorized creates a AssignProjectVersionOfAuthEntityUnauthorized with default headers values
func NewAssignProjectVersionOfAuthEntityUnauthorized() *AssignProjectVersionOfAuthEntityUnauthorized {
	return &AssignProjectVersionOfAuthEntityUnauthorized{}
}

/*AssignProjectVersionOfAuthEntityUnauthorized handles this case with default header values.

Unauthorized
*/
type AssignProjectVersionOfAuthEntityUnauthorized struct {
	Payload *models.APIResult
}

func (o *AssignProjectVersionOfAuthEntityUnauthorized) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityUnauthorized  %+v", 401, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignProjectVersionOfAuthEntityForbidden creates a AssignProjectVersionOfAuthEntityForbidden with default headers values
func NewAssignProjectVersionOfAuthEntityForbidden() *AssignProjectVersionOfAuthEntityForbidden {
	return &AssignProjectVersionOfAuthEntityForbidden{}
}

/*AssignProjectVersionOfAuthEntityForbidden handles this case with default header values.

Forbidden
*/
type AssignProjectVersionOfAuthEntityForbidden struct {
	Payload *models.APIResult
}

func (o *AssignProjectVersionOfAuthEntityForbidden) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityForbidden  %+v", 403, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignProjectVersionOfAuthEntityNotFound creates a AssignProjectVersionOfAuthEntityNotFound with default headers values
func NewAssignProjectVersionOfAuthEntityNotFound() *AssignProjectVersionOfAuthEntityNotFound {
	return &AssignProjectVersionOfAuthEntityNotFound{}
}

/*AssignProjectVersionOfAuthEntityNotFound handles this case with default header values.

Not Found
*/
type AssignProjectVersionOfAuthEntityNotFound struct {
	Payload *models.APIResult
}

func (o *AssignProjectVersionOfAuthEntityNotFound) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityNotFound  %+v", 404, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignProjectVersionOfAuthEntityConflict creates a AssignProjectVersionOfAuthEntityConflict with default headers values
func NewAssignProjectVersionOfAuthEntityConflict() *AssignProjectVersionOfAuthEntityConflict {
	return &AssignProjectVersionOfAuthEntityConflict{}
}

/*AssignProjectVersionOfAuthEntityConflict handles this case with default header values.

Conflict
*/
type AssignProjectVersionOfAuthEntityConflict struct {
	Payload *models.APIResult
}

func (o *AssignProjectVersionOfAuthEntityConflict) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityConflict  %+v", 409, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAssignProjectVersionOfAuthEntityInternalServerError creates a AssignProjectVersionOfAuthEntityInternalServerError with default headers values
func NewAssignProjectVersionOfAuthEntityInternalServerError() *AssignProjectVersionOfAuthEntityInternalServerError {
	return &AssignProjectVersionOfAuthEntityInternalServerError{}
}

/*AssignProjectVersionOfAuthEntityInternalServerError handles this case with default header values.

Internal Server Error
*/
type AssignProjectVersionOfAuthEntityInternalServerError struct {
	Payload *models.APIResult
}

func (o *AssignProjectVersionOfAuthEntityInternalServerError) Error() string {
	return fmt.Sprintf("[POST /authEntities/{parentId}/projectVersions/action/assign][%d] assignProjectVersionOfAuthEntityInternalServerError  %+v", 500, o.Payload)
}

func (o *AssignProjectVersionOfAuthEntityInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *AssignProjectVersionOfAuthEntityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
