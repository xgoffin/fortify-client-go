// Code generated by go-swagger; DO NOT EDIT.

package artifact_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// PurgeArtifactReader is a Reader for the PurgeArtifact structure.
type PurgeArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PurgeArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPurgeArtifactOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPurgeArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPurgeArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPurgeArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPurgeArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPurgeArtifactConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPurgeArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPurgeArtifactOK creates a PurgeArtifactOK with default headers values
func NewPurgeArtifactOK() *PurgeArtifactOK {
	return &PurgeArtifactOK{}
}

/*PurgeArtifactOK handles this case with default header values.

OK
*/
type PurgeArtifactOK struct {
	Payload *models.APIResultVoid
}

func (o *PurgeArtifactOK) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactOK  %+v", 200, o.Payload)
}

func (o *PurgeArtifactOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *PurgeArtifactOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeArtifactBadRequest creates a PurgeArtifactBadRequest with default headers values
func NewPurgeArtifactBadRequest() *PurgeArtifactBadRequest {
	return &PurgeArtifactBadRequest{}
}

/*PurgeArtifactBadRequest handles this case with default header values.

Bad Request
*/
type PurgeArtifactBadRequest struct {
	Payload *models.APIResult
}

func (o *PurgeArtifactBadRequest) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *PurgeArtifactBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeArtifactUnauthorized creates a PurgeArtifactUnauthorized with default headers values
func NewPurgeArtifactUnauthorized() *PurgeArtifactUnauthorized {
	return &PurgeArtifactUnauthorized{}
}

/*PurgeArtifactUnauthorized handles this case with default header values.

Unauthorized
*/
type PurgeArtifactUnauthorized struct {
	Payload *models.APIResult
}

func (o *PurgeArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *PurgeArtifactUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeArtifactForbidden creates a PurgeArtifactForbidden with default headers values
func NewPurgeArtifactForbidden() *PurgeArtifactForbidden {
	return &PurgeArtifactForbidden{}
}

/*PurgeArtifactForbidden handles this case with default header values.

Forbidden
*/
type PurgeArtifactForbidden struct {
	Payload *models.APIResult
}

func (o *PurgeArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactForbidden  %+v", 403, o.Payload)
}

func (o *PurgeArtifactForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeArtifactNotFound creates a PurgeArtifactNotFound with default headers values
func NewPurgeArtifactNotFound() *PurgeArtifactNotFound {
	return &PurgeArtifactNotFound{}
}

/*PurgeArtifactNotFound handles this case with default header values.

Not Found
*/
type PurgeArtifactNotFound struct {
	Payload *models.APIResult
}

func (o *PurgeArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactNotFound  %+v", 404, o.Payload)
}

func (o *PurgeArtifactNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeArtifactConflict creates a PurgeArtifactConflict with default headers values
func NewPurgeArtifactConflict() *PurgeArtifactConflict {
	return &PurgeArtifactConflict{}
}

/*PurgeArtifactConflict handles this case with default header values.

Conflict
*/
type PurgeArtifactConflict struct {
	Payload *models.APIResult
}

func (o *PurgeArtifactConflict) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactConflict  %+v", 409, o.Payload)
}

func (o *PurgeArtifactConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeArtifactConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPurgeArtifactInternalServerError creates a PurgeArtifactInternalServerError with default headers values
func NewPurgeArtifactInternalServerError() *PurgeArtifactInternalServerError {
	return &PurgeArtifactInternalServerError{}
}

/*PurgeArtifactInternalServerError handles this case with default header values.

Internal Server Error
*/
type PurgeArtifactInternalServerError struct {
	Payload *models.APIResult
}

func (o *PurgeArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /artifacts/action/purge][%d] purgeArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *PurgeArtifactInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PurgeArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
