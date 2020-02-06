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

// FileBugForIssueOfProjectVersionReader is a Reader for the FileBugForIssueOfProjectVersion structure.
type FileBugForIssueOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileBugForIssueOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileBugForIssueOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewFileBugForIssueOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewFileBugForIssueOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewFileBugForIssueOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewFileBugForIssueOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewFileBugForIssueOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewFileBugForIssueOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFileBugForIssueOfProjectVersionOK creates a FileBugForIssueOfProjectVersionOK with default headers values
func NewFileBugForIssueOfProjectVersionOK() *FileBugForIssueOfProjectVersionOK {
	return &FileBugForIssueOfProjectVersionOK{}
}

/*FileBugForIssueOfProjectVersionOK handles this case with default header values.

OK
*/
type FileBugForIssueOfProjectVersionOK struct {
	Payload *models.APIResultIssueFileBugResponse
}

func (o *FileBugForIssueOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionOK) GetPayload() *models.APIResultIssueFileBugResponse {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueFileBugResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileBugForIssueOfProjectVersionBadRequest creates a FileBugForIssueOfProjectVersionBadRequest with default headers values
func NewFileBugForIssueOfProjectVersionBadRequest() *FileBugForIssueOfProjectVersionBadRequest {
	return &FileBugForIssueOfProjectVersionBadRequest{}
}

/*FileBugForIssueOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type FileBugForIssueOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *FileBugForIssueOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileBugForIssueOfProjectVersionUnauthorized creates a FileBugForIssueOfProjectVersionUnauthorized with default headers values
func NewFileBugForIssueOfProjectVersionUnauthorized() *FileBugForIssueOfProjectVersionUnauthorized {
	return &FileBugForIssueOfProjectVersionUnauthorized{}
}

/*FileBugForIssueOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type FileBugForIssueOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *FileBugForIssueOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileBugForIssueOfProjectVersionForbidden creates a FileBugForIssueOfProjectVersionForbidden with default headers values
func NewFileBugForIssueOfProjectVersionForbidden() *FileBugForIssueOfProjectVersionForbidden {
	return &FileBugForIssueOfProjectVersionForbidden{}
}

/*FileBugForIssueOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type FileBugForIssueOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *FileBugForIssueOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileBugForIssueOfProjectVersionNotFound creates a FileBugForIssueOfProjectVersionNotFound with default headers values
func NewFileBugForIssueOfProjectVersionNotFound() *FileBugForIssueOfProjectVersionNotFound {
	return &FileBugForIssueOfProjectVersionNotFound{}
}

/*FileBugForIssueOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type FileBugForIssueOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *FileBugForIssueOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileBugForIssueOfProjectVersionConflict creates a FileBugForIssueOfProjectVersionConflict with default headers values
func NewFileBugForIssueOfProjectVersionConflict() *FileBugForIssueOfProjectVersionConflict {
	return &FileBugForIssueOfProjectVersionConflict{}
}

/*FileBugForIssueOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type FileBugForIssueOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *FileBugForIssueOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileBugForIssueOfProjectVersionInternalServerError creates a FileBugForIssueOfProjectVersionInternalServerError with default headers values
func NewFileBugForIssueOfProjectVersionInternalServerError() *FileBugForIssueOfProjectVersionInternalServerError {
	return &FileBugForIssueOfProjectVersionInternalServerError{}
}

/*FileBugForIssueOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type FileBugForIssueOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *FileBugForIssueOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/issues/action/fileBug][%d] fileBugForIssueOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *FileBugForIssueOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *FileBugForIssueOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
