// Code generated by go-swagger; DO NOT EDIT.

package issue_attachment_of_issue_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UploadIssueAttachmentOfIssueReader is a Reader for the UploadIssueAttachmentOfIssue structure.
type UploadIssueAttachmentOfIssueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadIssueAttachmentOfIssueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUploadIssueAttachmentOfIssueCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadIssueAttachmentOfIssueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadIssueAttachmentOfIssueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadIssueAttachmentOfIssueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadIssueAttachmentOfIssueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUploadIssueAttachmentOfIssueConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadIssueAttachmentOfIssueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadIssueAttachmentOfIssueCreated creates a UploadIssueAttachmentOfIssueCreated with default headers values
func NewUploadIssueAttachmentOfIssueCreated() *UploadIssueAttachmentOfIssueCreated {
	return &UploadIssueAttachmentOfIssueCreated{}
}

/*UploadIssueAttachmentOfIssueCreated handles this case with default header values.

Created
*/
type UploadIssueAttachmentOfIssueCreated struct {
	Payload *models.APIResultIssueAttachment
}

func (o *UploadIssueAttachmentOfIssueCreated) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueCreated  %+v", 201, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueCreated) GetPayload() *models.APIResultIssueAttachment {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadIssueAttachmentOfIssueBadRequest creates a UploadIssueAttachmentOfIssueBadRequest with default headers values
func NewUploadIssueAttachmentOfIssueBadRequest() *UploadIssueAttachmentOfIssueBadRequest {
	return &UploadIssueAttachmentOfIssueBadRequest{}
}

/*UploadIssueAttachmentOfIssueBadRequest handles this case with default header values.

Bad Request
*/
type UploadIssueAttachmentOfIssueBadRequest struct {
	Payload *models.APIResult
}

func (o *UploadIssueAttachmentOfIssueBadRequest) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueBadRequest  %+v", 400, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadIssueAttachmentOfIssueUnauthorized creates a UploadIssueAttachmentOfIssueUnauthorized with default headers values
func NewUploadIssueAttachmentOfIssueUnauthorized() *UploadIssueAttachmentOfIssueUnauthorized {
	return &UploadIssueAttachmentOfIssueUnauthorized{}
}

/*UploadIssueAttachmentOfIssueUnauthorized handles this case with default header values.

Unauthorized
*/
type UploadIssueAttachmentOfIssueUnauthorized struct {
	Payload *models.APIResult
}

func (o *UploadIssueAttachmentOfIssueUnauthorized) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueUnauthorized  %+v", 401, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadIssueAttachmentOfIssueForbidden creates a UploadIssueAttachmentOfIssueForbidden with default headers values
func NewUploadIssueAttachmentOfIssueForbidden() *UploadIssueAttachmentOfIssueForbidden {
	return &UploadIssueAttachmentOfIssueForbidden{}
}

/*UploadIssueAttachmentOfIssueForbidden handles this case with default header values.

Forbidden
*/
type UploadIssueAttachmentOfIssueForbidden struct {
	Payload *models.APIResult
}

func (o *UploadIssueAttachmentOfIssueForbidden) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueForbidden  %+v", 403, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadIssueAttachmentOfIssueNotFound creates a UploadIssueAttachmentOfIssueNotFound with default headers values
func NewUploadIssueAttachmentOfIssueNotFound() *UploadIssueAttachmentOfIssueNotFound {
	return &UploadIssueAttachmentOfIssueNotFound{}
}

/*UploadIssueAttachmentOfIssueNotFound handles this case with default header values.

Not Found
*/
type UploadIssueAttachmentOfIssueNotFound struct {
	Payload *models.APIResult
}

func (o *UploadIssueAttachmentOfIssueNotFound) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueNotFound  %+v", 404, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadIssueAttachmentOfIssueConflict creates a UploadIssueAttachmentOfIssueConflict with default headers values
func NewUploadIssueAttachmentOfIssueConflict() *UploadIssueAttachmentOfIssueConflict {
	return &UploadIssueAttachmentOfIssueConflict{}
}

/*UploadIssueAttachmentOfIssueConflict handles this case with default header values.

Conflict
*/
type UploadIssueAttachmentOfIssueConflict struct {
	Payload *models.APIResult
}

func (o *UploadIssueAttachmentOfIssueConflict) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueConflict  %+v", 409, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadIssueAttachmentOfIssueInternalServerError creates a UploadIssueAttachmentOfIssueInternalServerError with default headers values
func NewUploadIssueAttachmentOfIssueInternalServerError() *UploadIssueAttachmentOfIssueInternalServerError {
	return &UploadIssueAttachmentOfIssueInternalServerError{}
}

/*UploadIssueAttachmentOfIssueInternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadIssueAttachmentOfIssueInternalServerError struct {
	Payload *models.APIResult
}

func (o *UploadIssueAttachmentOfIssueInternalServerError) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/attachments][%d] uploadIssueAttachmentOfIssueInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadIssueAttachmentOfIssueInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadIssueAttachmentOfIssueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
