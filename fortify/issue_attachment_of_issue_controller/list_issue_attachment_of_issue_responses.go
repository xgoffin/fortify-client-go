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

// ListIssueAttachmentOfIssueReader is a Reader for the ListIssueAttachmentOfIssue structure.
type ListIssueAttachmentOfIssueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIssueAttachmentOfIssueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIssueAttachmentOfIssueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListIssueAttachmentOfIssueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListIssueAttachmentOfIssueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListIssueAttachmentOfIssueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListIssueAttachmentOfIssueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListIssueAttachmentOfIssueConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListIssueAttachmentOfIssueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIssueAttachmentOfIssueOK creates a ListIssueAttachmentOfIssueOK with default headers values
func NewListIssueAttachmentOfIssueOK() *ListIssueAttachmentOfIssueOK {
	return &ListIssueAttachmentOfIssueOK{}
}

/*ListIssueAttachmentOfIssueOK handles this case with default header values.

OK
*/
type ListIssueAttachmentOfIssueOK struct {
	Payload *models.APIResultListIssueAttachment
}

func (o *ListIssueAttachmentOfIssueOK) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueOK  %+v", 200, o.Payload)
}

func (o *ListIssueAttachmentOfIssueOK) GetPayload() *models.APIResultListIssueAttachment {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListIssueAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAttachmentOfIssueBadRequest creates a ListIssueAttachmentOfIssueBadRequest with default headers values
func NewListIssueAttachmentOfIssueBadRequest() *ListIssueAttachmentOfIssueBadRequest {
	return &ListIssueAttachmentOfIssueBadRequest{}
}

/*ListIssueAttachmentOfIssueBadRequest handles this case with default header values.

Bad Request
*/
type ListIssueAttachmentOfIssueBadRequest struct {
	Payload *models.APIResult
}

func (o *ListIssueAttachmentOfIssueBadRequest) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueBadRequest  %+v", 400, o.Payload)
}

func (o *ListIssueAttachmentOfIssueBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAttachmentOfIssueUnauthorized creates a ListIssueAttachmentOfIssueUnauthorized with default headers values
func NewListIssueAttachmentOfIssueUnauthorized() *ListIssueAttachmentOfIssueUnauthorized {
	return &ListIssueAttachmentOfIssueUnauthorized{}
}

/*ListIssueAttachmentOfIssueUnauthorized handles this case with default header values.

Unauthorized
*/
type ListIssueAttachmentOfIssueUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListIssueAttachmentOfIssueUnauthorized) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueUnauthorized  %+v", 401, o.Payload)
}

func (o *ListIssueAttachmentOfIssueUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAttachmentOfIssueForbidden creates a ListIssueAttachmentOfIssueForbidden with default headers values
func NewListIssueAttachmentOfIssueForbidden() *ListIssueAttachmentOfIssueForbidden {
	return &ListIssueAttachmentOfIssueForbidden{}
}

/*ListIssueAttachmentOfIssueForbidden handles this case with default header values.

Forbidden
*/
type ListIssueAttachmentOfIssueForbidden struct {
	Payload *models.APIResult
}

func (o *ListIssueAttachmentOfIssueForbidden) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueForbidden  %+v", 403, o.Payload)
}

func (o *ListIssueAttachmentOfIssueForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAttachmentOfIssueNotFound creates a ListIssueAttachmentOfIssueNotFound with default headers values
func NewListIssueAttachmentOfIssueNotFound() *ListIssueAttachmentOfIssueNotFound {
	return &ListIssueAttachmentOfIssueNotFound{}
}

/*ListIssueAttachmentOfIssueNotFound handles this case with default header values.

Not Found
*/
type ListIssueAttachmentOfIssueNotFound struct {
	Payload *models.APIResult
}

func (o *ListIssueAttachmentOfIssueNotFound) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueNotFound  %+v", 404, o.Payload)
}

func (o *ListIssueAttachmentOfIssueNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAttachmentOfIssueConflict creates a ListIssueAttachmentOfIssueConflict with default headers values
func NewListIssueAttachmentOfIssueConflict() *ListIssueAttachmentOfIssueConflict {
	return &ListIssueAttachmentOfIssueConflict{}
}

/*ListIssueAttachmentOfIssueConflict handles this case with default header values.

Conflict
*/
type ListIssueAttachmentOfIssueConflict struct {
	Payload *models.APIResult
}

func (o *ListIssueAttachmentOfIssueConflict) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueConflict  %+v", 409, o.Payload)
}

func (o *ListIssueAttachmentOfIssueConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAttachmentOfIssueInternalServerError creates a ListIssueAttachmentOfIssueInternalServerError with default headers values
func NewListIssueAttachmentOfIssueInternalServerError() *ListIssueAttachmentOfIssueInternalServerError {
	return &ListIssueAttachmentOfIssueInternalServerError{}
}

/*ListIssueAttachmentOfIssueInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListIssueAttachmentOfIssueInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListIssueAttachmentOfIssueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments][%d] listIssueAttachmentOfIssueInternalServerError  %+v", 500, o.Payload)
}

func (o *ListIssueAttachmentOfIssueInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAttachmentOfIssueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
