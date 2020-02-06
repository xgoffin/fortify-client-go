// Code generated by go-swagger; DO NOT EDIT.

package issue_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UpdateTagForIssueOfProjectVersionReader is a Reader for the UpdateTagForIssueOfProjectVersion structure.
type UpdateTagForIssueOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTagForIssueOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTagForIssueOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateTagForIssueOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateTagForIssueOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateTagForIssueOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateTagForIssueOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateTagForIssueOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateTagForIssueOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTagForIssueOfProjectVersionOK creates a UpdateTagForIssueOfProjectVersionOK with default headers values
func NewUpdateTagForIssueOfProjectVersionOK() *UpdateTagForIssueOfProjectVersionOK {
	return &UpdateTagForIssueOfProjectVersionOK{}
}

/*UpdateTagForIssueOfProjectVersionOK handles this case with default header values.

OK
*/
type UpdateTagForIssueOfProjectVersionOK struct {
	Payload *models.APIResultIssueActionResponse
}

func (o *UpdateTagForIssueOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionOK) GetPayload() *models.APIResultIssueActionResponse {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueActionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForIssueOfProjectVersionBadRequest creates a UpdateTagForIssueOfProjectVersionBadRequest with default headers values
func NewUpdateTagForIssueOfProjectVersionBadRequest() *UpdateTagForIssueOfProjectVersionBadRequest {
	return &UpdateTagForIssueOfProjectVersionBadRequest{}
}

/*UpdateTagForIssueOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateTagForIssueOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateTagForIssueOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForIssueOfProjectVersionUnauthorized creates a UpdateTagForIssueOfProjectVersionUnauthorized with default headers values
func NewUpdateTagForIssueOfProjectVersionUnauthorized() *UpdateTagForIssueOfProjectVersionUnauthorized {
	return &UpdateTagForIssueOfProjectVersionUnauthorized{}
}

/*UpdateTagForIssueOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateTagForIssueOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateTagForIssueOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForIssueOfProjectVersionForbidden creates a UpdateTagForIssueOfProjectVersionForbidden with default headers values
func NewUpdateTagForIssueOfProjectVersionForbidden() *UpdateTagForIssueOfProjectVersionForbidden {
	return &UpdateTagForIssueOfProjectVersionForbidden{}
}

/*UpdateTagForIssueOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type UpdateTagForIssueOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateTagForIssueOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForIssueOfProjectVersionNotFound creates a UpdateTagForIssueOfProjectVersionNotFound with default headers values
func NewUpdateTagForIssueOfProjectVersionNotFound() *UpdateTagForIssueOfProjectVersionNotFound {
	return &UpdateTagForIssueOfProjectVersionNotFound{}
}

/*UpdateTagForIssueOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type UpdateTagForIssueOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateTagForIssueOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForIssueOfProjectVersionConflict creates a UpdateTagForIssueOfProjectVersionConflict with default headers values
func NewUpdateTagForIssueOfProjectVersionConflict() *UpdateTagForIssueOfProjectVersionConflict {
	return &UpdateTagForIssueOfProjectVersionConflict{}
}

/*UpdateTagForIssueOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type UpdateTagForIssueOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *UpdateTagForIssueOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTagForIssueOfProjectVersionInternalServerError creates a UpdateTagForIssueOfProjectVersionInternalServerError with default headers values
func NewUpdateTagForIssueOfProjectVersionInternalServerError() *UpdateTagForIssueOfProjectVersionInternalServerError {
	return &UpdateTagForIssueOfProjectVersionInternalServerError{}
}

/*UpdateTagForIssueOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateTagForIssueOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateTagForIssueOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/updateTag][%d] updateTagForIssueOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateTagForIssueOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateTagForIssueOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
