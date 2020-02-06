// Code generated by go-swagger; DO NOT EDIT.

package report_library_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadReportLibraryReader is a Reader for the ReadReportLibrary structure.
type ReadReportLibraryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadReportLibraryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadReportLibraryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadReportLibraryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadReportLibraryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadReportLibraryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadReportLibraryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadReportLibraryConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadReportLibraryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadReportLibraryOK creates a ReadReportLibraryOK with default headers values
func NewReadReportLibraryOK() *ReadReportLibraryOK {
	return &ReadReportLibraryOK{}
}

/*ReadReportLibraryOK handles this case with default header values.

OK
*/
type ReadReportLibraryOK struct {
	Payload *models.APIResultReportLibrary
}

func (o *ReadReportLibraryOK) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryOK  %+v", 200, o.Payload)
}

func (o *ReadReportLibraryOK) GetPayload() *models.APIResultReportLibrary {
	return o.Payload
}

func (o *ReadReportLibraryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultReportLibrary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadReportLibraryBadRequest creates a ReadReportLibraryBadRequest with default headers values
func NewReadReportLibraryBadRequest() *ReadReportLibraryBadRequest {
	return &ReadReportLibraryBadRequest{}
}

/*ReadReportLibraryBadRequest handles this case with default header values.

Bad Request
*/
type ReadReportLibraryBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadReportLibraryBadRequest) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryBadRequest  %+v", 400, o.Payload)
}

func (o *ReadReportLibraryBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadReportLibraryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadReportLibraryUnauthorized creates a ReadReportLibraryUnauthorized with default headers values
func NewReadReportLibraryUnauthorized() *ReadReportLibraryUnauthorized {
	return &ReadReportLibraryUnauthorized{}
}

/*ReadReportLibraryUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadReportLibraryUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadReportLibraryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadReportLibraryUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadReportLibraryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadReportLibraryForbidden creates a ReadReportLibraryForbidden with default headers values
func NewReadReportLibraryForbidden() *ReadReportLibraryForbidden {
	return &ReadReportLibraryForbidden{}
}

/*ReadReportLibraryForbidden handles this case with default header values.

Forbidden
*/
type ReadReportLibraryForbidden struct {
	Payload *models.APIResult
}

func (o *ReadReportLibraryForbidden) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryForbidden  %+v", 403, o.Payload)
}

func (o *ReadReportLibraryForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadReportLibraryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadReportLibraryNotFound creates a ReadReportLibraryNotFound with default headers values
func NewReadReportLibraryNotFound() *ReadReportLibraryNotFound {
	return &ReadReportLibraryNotFound{}
}

/*ReadReportLibraryNotFound handles this case with default header values.

Not Found
*/
type ReadReportLibraryNotFound struct {
	Payload *models.APIResult
}

func (o *ReadReportLibraryNotFound) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryNotFound  %+v", 404, o.Payload)
}

func (o *ReadReportLibraryNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadReportLibraryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadReportLibraryConflict creates a ReadReportLibraryConflict with default headers values
func NewReadReportLibraryConflict() *ReadReportLibraryConflict {
	return &ReadReportLibraryConflict{}
}

/*ReadReportLibraryConflict handles this case with default header values.

Conflict
*/
type ReadReportLibraryConflict struct {
	Payload *models.APIResult
}

func (o *ReadReportLibraryConflict) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryConflict  %+v", 409, o.Payload)
}

func (o *ReadReportLibraryConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadReportLibraryConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadReportLibraryInternalServerError creates a ReadReportLibraryInternalServerError with default headers values
func NewReadReportLibraryInternalServerError() *ReadReportLibraryInternalServerError {
	return &ReadReportLibraryInternalServerError{}
}

/*ReadReportLibraryInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadReportLibraryInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadReportLibraryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /reportLibraries/{id}][%d] readReportLibraryInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadReportLibraryInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadReportLibraryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
