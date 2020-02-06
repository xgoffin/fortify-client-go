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

// ReadIssueAttachmentOfIssueReader is a Reader for the ReadIssueAttachmentOfIssue structure.
type ReadIssueAttachmentOfIssueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadIssueAttachmentOfIssueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadIssueAttachmentOfIssueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadIssueAttachmentOfIssueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadIssueAttachmentOfIssueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadIssueAttachmentOfIssueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadIssueAttachmentOfIssueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadIssueAttachmentOfIssueConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadIssueAttachmentOfIssueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadIssueAttachmentOfIssueOK creates a ReadIssueAttachmentOfIssueOK with default headers values
func NewReadIssueAttachmentOfIssueOK() *ReadIssueAttachmentOfIssueOK {
	return &ReadIssueAttachmentOfIssueOK{}
}

/*ReadIssueAttachmentOfIssueOK handles this case with default header values.

OK
*/
type ReadIssueAttachmentOfIssueOK struct {
	Payload *models.APIResultIssueAttachment
}

func (o *ReadIssueAttachmentOfIssueOK) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueOK  %+v", 200, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueOK) GetPayload() *models.APIResultIssueAttachment {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAttachmentOfIssueBadRequest creates a ReadIssueAttachmentOfIssueBadRequest with default headers values
func NewReadIssueAttachmentOfIssueBadRequest() *ReadIssueAttachmentOfIssueBadRequest {
	return &ReadIssueAttachmentOfIssueBadRequest{}
}

/*ReadIssueAttachmentOfIssueBadRequest handles this case with default header values.

Bad Request
*/
type ReadIssueAttachmentOfIssueBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadIssueAttachmentOfIssueBadRequest) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueBadRequest  %+v", 400, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAttachmentOfIssueUnauthorized creates a ReadIssueAttachmentOfIssueUnauthorized with default headers values
func NewReadIssueAttachmentOfIssueUnauthorized() *ReadIssueAttachmentOfIssueUnauthorized {
	return &ReadIssueAttachmentOfIssueUnauthorized{}
}

/*ReadIssueAttachmentOfIssueUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadIssueAttachmentOfIssueUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadIssueAttachmentOfIssueUnauthorized) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAttachmentOfIssueForbidden creates a ReadIssueAttachmentOfIssueForbidden with default headers values
func NewReadIssueAttachmentOfIssueForbidden() *ReadIssueAttachmentOfIssueForbidden {
	return &ReadIssueAttachmentOfIssueForbidden{}
}

/*ReadIssueAttachmentOfIssueForbidden handles this case with default header values.

Forbidden
*/
type ReadIssueAttachmentOfIssueForbidden struct {
	Payload *models.APIResult
}

func (o *ReadIssueAttachmentOfIssueForbidden) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueForbidden  %+v", 403, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAttachmentOfIssueNotFound creates a ReadIssueAttachmentOfIssueNotFound with default headers values
func NewReadIssueAttachmentOfIssueNotFound() *ReadIssueAttachmentOfIssueNotFound {
	return &ReadIssueAttachmentOfIssueNotFound{}
}

/*ReadIssueAttachmentOfIssueNotFound handles this case with default header values.

Not Found
*/
type ReadIssueAttachmentOfIssueNotFound struct {
	Payload *models.APIResult
}

func (o *ReadIssueAttachmentOfIssueNotFound) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueNotFound  %+v", 404, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAttachmentOfIssueConflict creates a ReadIssueAttachmentOfIssueConflict with default headers values
func NewReadIssueAttachmentOfIssueConflict() *ReadIssueAttachmentOfIssueConflict {
	return &ReadIssueAttachmentOfIssueConflict{}
}

/*ReadIssueAttachmentOfIssueConflict handles this case with default header values.

Conflict
*/
type ReadIssueAttachmentOfIssueConflict struct {
	Payload *models.APIResult
}

func (o *ReadIssueAttachmentOfIssueConflict) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueConflict  %+v", 409, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAttachmentOfIssueInternalServerError creates a ReadIssueAttachmentOfIssueInternalServerError with default headers values
func NewReadIssueAttachmentOfIssueInternalServerError() *ReadIssueAttachmentOfIssueInternalServerError {
	return &ReadIssueAttachmentOfIssueInternalServerError{}
}

/*ReadIssueAttachmentOfIssueInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadIssueAttachmentOfIssueInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadIssueAttachmentOfIssueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /issues/{parentId}/attachments/{id}][%d] readIssueAttachmentOfIssueInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadIssueAttachmentOfIssueInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAttachmentOfIssueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
