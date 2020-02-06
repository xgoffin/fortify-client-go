// Code generated by go-swagger; DO NOT EDIT.

package project_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadProjectReader is a Reader for the ReadProject structure.
type ReadProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadProjectOK creates a ReadProjectOK with default headers values
func NewReadProjectOK() *ReadProjectOK {
	return &ReadProjectOK{}
}

/*ReadProjectOK handles this case with default header values.

OK
*/
type ReadProjectOK struct {
	Payload *models.APIResultProject
}

func (o *ReadProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectOK  %+v", 200, o.Payload)
}

func (o *ReadProjectOK) GetPayload() *models.APIResultProject {
	return o.Payload
}

func (o *ReadProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProjectBadRequest creates a ReadProjectBadRequest with default headers values
func NewReadProjectBadRequest() *ReadProjectBadRequest {
	return &ReadProjectBadRequest{}
}

/*ReadProjectBadRequest handles this case with default header values.

Bad Request
*/
type ReadProjectBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectBadRequest  %+v", 400, o.Payload)
}

func (o *ReadProjectBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProjectUnauthorized creates a ReadProjectUnauthorized with default headers values
func NewReadProjectUnauthorized() *ReadProjectUnauthorized {
	return &ReadProjectUnauthorized{}
}

/*ReadProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadProjectUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadProjectUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProjectForbidden creates a ReadProjectForbidden with default headers values
func NewReadProjectForbidden() *ReadProjectForbidden {
	return &ReadProjectForbidden{}
}

/*ReadProjectForbidden handles this case with default header values.

Forbidden
*/
type ReadProjectForbidden struct {
	Payload *models.APIResult
}

func (o *ReadProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectForbidden  %+v", 403, o.Payload)
}

func (o *ReadProjectForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProjectNotFound creates a ReadProjectNotFound with default headers values
func NewReadProjectNotFound() *ReadProjectNotFound {
	return &ReadProjectNotFound{}
}

/*ReadProjectNotFound handles this case with default header values.

Not Found
*/
type ReadProjectNotFound struct {
	Payload *models.APIResult
}

func (o *ReadProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectNotFound  %+v", 404, o.Payload)
}

func (o *ReadProjectNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProjectConflict creates a ReadProjectConflict with default headers values
func NewReadProjectConflict() *ReadProjectConflict {
	return &ReadProjectConflict{}
}

/*ReadProjectConflict handles this case with default header values.

Conflict
*/
type ReadProjectConflict struct {
	Payload *models.APIResult
}

func (o *ReadProjectConflict) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectConflict  %+v", 409, o.Payload)
}

func (o *ReadProjectConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadProjectInternalServerError creates a ReadProjectInternalServerError with default headers values
func NewReadProjectInternalServerError() *ReadProjectInternalServerError {
	return &ReadProjectInternalServerError{}
}

/*ReadProjectInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadProjectInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{id}][%d] readProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadProjectInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
