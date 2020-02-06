// Code generated by go-swagger; DO NOT EDIT.

package responsibility_assignment_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListResponsibilityAssignmentOfProjectVersionReader is a Reader for the ListResponsibilityAssignmentOfProjectVersion structure.
type ListResponsibilityAssignmentOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListResponsibilityAssignmentOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListResponsibilityAssignmentOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListResponsibilityAssignmentOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListResponsibilityAssignmentOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListResponsibilityAssignmentOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListResponsibilityAssignmentOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListResponsibilityAssignmentOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListResponsibilityAssignmentOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListResponsibilityAssignmentOfProjectVersionOK creates a ListResponsibilityAssignmentOfProjectVersionOK with default headers values
func NewListResponsibilityAssignmentOfProjectVersionOK() *ListResponsibilityAssignmentOfProjectVersionOK {
	return &ListResponsibilityAssignmentOfProjectVersionOK{}
}

/*ListResponsibilityAssignmentOfProjectVersionOK handles this case with default header values.

OK
*/
type ListResponsibilityAssignmentOfProjectVersionOK struct {
	Payload *models.APIResultListResponsibilityAssignment
}

func (o *ListResponsibilityAssignmentOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionOK) GetPayload() *models.APIResultListResponsibilityAssignment {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListResponsibilityAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResponsibilityAssignmentOfProjectVersionBadRequest creates a ListResponsibilityAssignmentOfProjectVersionBadRequest with default headers values
func NewListResponsibilityAssignmentOfProjectVersionBadRequest() *ListResponsibilityAssignmentOfProjectVersionBadRequest {
	return &ListResponsibilityAssignmentOfProjectVersionBadRequest{}
}

/*ListResponsibilityAssignmentOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ListResponsibilityAssignmentOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ListResponsibilityAssignmentOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResponsibilityAssignmentOfProjectVersionUnauthorized creates a ListResponsibilityAssignmentOfProjectVersionUnauthorized with default headers values
func NewListResponsibilityAssignmentOfProjectVersionUnauthorized() *ListResponsibilityAssignmentOfProjectVersionUnauthorized {
	return &ListResponsibilityAssignmentOfProjectVersionUnauthorized{}
}

/*ListResponsibilityAssignmentOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListResponsibilityAssignmentOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListResponsibilityAssignmentOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResponsibilityAssignmentOfProjectVersionForbidden creates a ListResponsibilityAssignmentOfProjectVersionForbidden with default headers values
func NewListResponsibilityAssignmentOfProjectVersionForbidden() *ListResponsibilityAssignmentOfProjectVersionForbidden {
	return &ListResponsibilityAssignmentOfProjectVersionForbidden{}
}

/*ListResponsibilityAssignmentOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ListResponsibilityAssignmentOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ListResponsibilityAssignmentOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResponsibilityAssignmentOfProjectVersionNotFound creates a ListResponsibilityAssignmentOfProjectVersionNotFound with default headers values
func NewListResponsibilityAssignmentOfProjectVersionNotFound() *ListResponsibilityAssignmentOfProjectVersionNotFound {
	return &ListResponsibilityAssignmentOfProjectVersionNotFound{}
}

/*ListResponsibilityAssignmentOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ListResponsibilityAssignmentOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ListResponsibilityAssignmentOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResponsibilityAssignmentOfProjectVersionConflict creates a ListResponsibilityAssignmentOfProjectVersionConflict with default headers values
func NewListResponsibilityAssignmentOfProjectVersionConflict() *ListResponsibilityAssignmentOfProjectVersionConflict {
	return &ListResponsibilityAssignmentOfProjectVersionConflict{}
}

/*ListResponsibilityAssignmentOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ListResponsibilityAssignmentOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ListResponsibilityAssignmentOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListResponsibilityAssignmentOfProjectVersionInternalServerError creates a ListResponsibilityAssignmentOfProjectVersionInternalServerError with default headers values
func NewListResponsibilityAssignmentOfProjectVersionInternalServerError() *ListResponsibilityAssignmentOfProjectVersionInternalServerError {
	return &ListResponsibilityAssignmentOfProjectVersionInternalServerError{}
}

/*ListResponsibilityAssignmentOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListResponsibilityAssignmentOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListResponsibilityAssignmentOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/responsibilities][%d] listResponsibilityAssignmentOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListResponsibilityAssignmentOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListResponsibilityAssignmentOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
