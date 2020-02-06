// Code generated by go-swagger; DO NOT EDIT.

package project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// PurgeProjectVersionReader is a Reader for the PurgeProjectVersion structure.
type PurgeProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurgeProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurgeProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPurgeProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPurgeProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPurgeProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPurgeProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPurgeProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPurgeProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPurgeProjectVersionOK creates a PurgeProjectVersionOK with default headers values
func NewPurgeProjectVersionOK() *PurgeProjectVersionOK {
	return &PurgeProjectVersionOK{}
}

/*PurgeProjectVersionOK handles this case with default header values.

OK
*/
type PurgeProjectVersionOK struct {
	Payload *models.APIResultVoid
}

func (o *PurgeProjectVersionOK) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionOK  %+v", 200, o.Payload)
}

func (o *PurgeProjectVersionOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *PurgeProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeProjectVersionBadRequest creates a PurgeProjectVersionBadRequest with default headers values
func NewPurgeProjectVersionBadRequest() *PurgeProjectVersionBadRequest {
	return &PurgeProjectVersionBadRequest{}
}

/*PurgeProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type PurgeProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *PurgeProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *PurgeProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeProjectVersionUnauthorized creates a PurgeProjectVersionUnauthorized with default headers values
func NewPurgeProjectVersionUnauthorized() *PurgeProjectVersionUnauthorized {
	return &PurgeProjectVersionUnauthorized{}
}

/*PurgeProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type PurgeProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *PurgeProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *PurgeProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeProjectVersionForbidden creates a PurgeProjectVersionForbidden with default headers values
func NewPurgeProjectVersionForbidden() *PurgeProjectVersionForbidden {
	return &PurgeProjectVersionForbidden{}
}

/*PurgeProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type PurgeProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *PurgeProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *PurgeProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeProjectVersionNotFound creates a PurgeProjectVersionNotFound with default headers values
func NewPurgeProjectVersionNotFound() *PurgeProjectVersionNotFound {
	return &PurgeProjectVersionNotFound{}
}

/*PurgeProjectVersionNotFound handles this case with default header values.

Not Found
*/
type PurgeProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *PurgeProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *PurgeProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeProjectVersionConflict creates a PurgeProjectVersionConflict with default headers values
func NewPurgeProjectVersionConflict() *PurgeProjectVersionConflict {
	return &PurgeProjectVersionConflict{}
}

/*PurgeProjectVersionConflict handles this case with default header values.

Conflict
*/
type PurgeProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *PurgeProjectVersionConflict) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *PurgeProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeProjectVersionInternalServerError creates a PurgeProjectVersionInternalServerError with default headers values
func NewPurgeProjectVersionInternalServerError() *PurgeProjectVersionInternalServerError {
	return &PurgeProjectVersionInternalServerError{}
}

/*PurgeProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type PurgeProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *PurgeProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projectVersions/action/purge][%d] purgeProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *PurgeProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
