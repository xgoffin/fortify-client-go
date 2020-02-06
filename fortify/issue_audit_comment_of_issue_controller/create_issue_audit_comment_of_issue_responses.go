// Code generated by go-swagger; DO NOT EDIT.

package issue_audit_comment_of_issue_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreateIssueAuditCommentOfIssueReader is a Reader for the CreateIssueAuditCommentOfIssue structure.
type CreateIssueAuditCommentOfIssueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIssueAuditCommentOfIssueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateIssueAuditCommentOfIssueCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIssueAuditCommentOfIssueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateIssueAuditCommentOfIssueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateIssueAuditCommentOfIssueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateIssueAuditCommentOfIssueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateIssueAuditCommentOfIssueConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateIssueAuditCommentOfIssueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateIssueAuditCommentOfIssueCreated creates a CreateIssueAuditCommentOfIssueCreated with default headers values
func NewCreateIssueAuditCommentOfIssueCreated() *CreateIssueAuditCommentOfIssueCreated {
	return &CreateIssueAuditCommentOfIssueCreated{}
}

/*CreateIssueAuditCommentOfIssueCreated handles this case with default header values.

Created
*/
type CreateIssueAuditCommentOfIssueCreated struct {
	Payload *models.APIResultIssueAuditComment
}

func (o *CreateIssueAuditCommentOfIssueCreated) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueCreated  %+v", 201, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueCreated) GetPayload() *models.APIResultIssueAuditComment {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueAuditComment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIssueAuditCommentOfIssueBadRequest creates a CreateIssueAuditCommentOfIssueBadRequest with default headers values
func NewCreateIssueAuditCommentOfIssueBadRequest() *CreateIssueAuditCommentOfIssueBadRequest {
	return &CreateIssueAuditCommentOfIssueBadRequest{}
}

/*CreateIssueAuditCommentOfIssueBadRequest handles this case with default header values.

Bad Request
*/
type CreateIssueAuditCommentOfIssueBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateIssueAuditCommentOfIssueBadRequest) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueBadRequest  %+v", 400, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIssueAuditCommentOfIssueUnauthorized creates a CreateIssueAuditCommentOfIssueUnauthorized with default headers values
func NewCreateIssueAuditCommentOfIssueUnauthorized() *CreateIssueAuditCommentOfIssueUnauthorized {
	return &CreateIssueAuditCommentOfIssueUnauthorized{}
}

/*CreateIssueAuditCommentOfIssueUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateIssueAuditCommentOfIssueUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateIssueAuditCommentOfIssueUnauthorized) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIssueAuditCommentOfIssueForbidden creates a CreateIssueAuditCommentOfIssueForbidden with default headers values
func NewCreateIssueAuditCommentOfIssueForbidden() *CreateIssueAuditCommentOfIssueForbidden {
	return &CreateIssueAuditCommentOfIssueForbidden{}
}

/*CreateIssueAuditCommentOfIssueForbidden handles this case with default header values.

Forbidden
*/
type CreateIssueAuditCommentOfIssueForbidden struct {
	Payload *models.APIResult
}

func (o *CreateIssueAuditCommentOfIssueForbidden) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueForbidden  %+v", 403, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIssueAuditCommentOfIssueNotFound creates a CreateIssueAuditCommentOfIssueNotFound with default headers values
func NewCreateIssueAuditCommentOfIssueNotFound() *CreateIssueAuditCommentOfIssueNotFound {
	return &CreateIssueAuditCommentOfIssueNotFound{}
}

/*CreateIssueAuditCommentOfIssueNotFound handles this case with default header values.

Not Found
*/
type CreateIssueAuditCommentOfIssueNotFound struct {
	Payload *models.APIResult
}

func (o *CreateIssueAuditCommentOfIssueNotFound) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueNotFound  %+v", 404, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIssueAuditCommentOfIssueConflict creates a CreateIssueAuditCommentOfIssueConflict with default headers values
func NewCreateIssueAuditCommentOfIssueConflict() *CreateIssueAuditCommentOfIssueConflict {
	return &CreateIssueAuditCommentOfIssueConflict{}
}

/*CreateIssueAuditCommentOfIssueConflict handles this case with default header values.

Conflict
*/
type CreateIssueAuditCommentOfIssueConflict struct {
	Payload *models.APIResult
}

func (o *CreateIssueAuditCommentOfIssueConflict) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueConflict  %+v", 409, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIssueAuditCommentOfIssueInternalServerError creates a CreateIssueAuditCommentOfIssueInternalServerError with default headers values
func NewCreateIssueAuditCommentOfIssueInternalServerError() *CreateIssueAuditCommentOfIssueInternalServerError {
	return &CreateIssueAuditCommentOfIssueInternalServerError{}
}

/*CreateIssueAuditCommentOfIssueInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateIssueAuditCommentOfIssueInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateIssueAuditCommentOfIssueInternalServerError) Error() string {
	return fmt.Sprintf("[POST /issues/{parentId}/comments][%d] createIssueAuditCommentOfIssueInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateIssueAuditCommentOfIssueInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateIssueAuditCommentOfIssueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
