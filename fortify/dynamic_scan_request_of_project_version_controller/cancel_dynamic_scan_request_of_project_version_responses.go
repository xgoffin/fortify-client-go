// Code generated by go-swagger; DO NOT EDIT.

package dynamic_scan_request_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CancelDynamicScanRequestOfProjectVersionReader is a Reader for the CancelDynamicScanRequestOfProjectVersion structure.
type CancelDynamicScanRequestOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelDynamicScanRequestOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelDynamicScanRequestOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelDynamicScanRequestOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelDynamicScanRequestOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelDynamicScanRequestOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelDynamicScanRequestOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCancelDynamicScanRequestOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelDynamicScanRequestOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelDynamicScanRequestOfProjectVersionOK creates a CancelDynamicScanRequestOfProjectVersionOK with default headers values
func NewCancelDynamicScanRequestOfProjectVersionOK() *CancelDynamicScanRequestOfProjectVersionOK {
	return &CancelDynamicScanRequestOfProjectVersionOK{}
}

/*CancelDynamicScanRequestOfProjectVersionOK handles this case with default header values.

OK
*/
type CancelDynamicScanRequestOfProjectVersionOK struct {
	Payload *models.APIResultVoid
}

func (o *CancelDynamicScanRequestOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDynamicScanRequestOfProjectVersionBadRequest creates a CancelDynamicScanRequestOfProjectVersionBadRequest with default headers values
func NewCancelDynamicScanRequestOfProjectVersionBadRequest() *CancelDynamicScanRequestOfProjectVersionBadRequest {
	return &CancelDynamicScanRequestOfProjectVersionBadRequest{}
}

/*CancelDynamicScanRequestOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type CancelDynamicScanRequestOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *CancelDynamicScanRequestOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDynamicScanRequestOfProjectVersionUnauthorized creates a CancelDynamicScanRequestOfProjectVersionUnauthorized with default headers values
func NewCancelDynamicScanRequestOfProjectVersionUnauthorized() *CancelDynamicScanRequestOfProjectVersionUnauthorized {
	return &CancelDynamicScanRequestOfProjectVersionUnauthorized{}
}

/*CancelDynamicScanRequestOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type CancelDynamicScanRequestOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *CancelDynamicScanRequestOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDynamicScanRequestOfProjectVersionForbidden creates a CancelDynamicScanRequestOfProjectVersionForbidden with default headers values
func NewCancelDynamicScanRequestOfProjectVersionForbidden() *CancelDynamicScanRequestOfProjectVersionForbidden {
	return &CancelDynamicScanRequestOfProjectVersionForbidden{}
}

/*CancelDynamicScanRequestOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type CancelDynamicScanRequestOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *CancelDynamicScanRequestOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDynamicScanRequestOfProjectVersionNotFound creates a CancelDynamicScanRequestOfProjectVersionNotFound with default headers values
func NewCancelDynamicScanRequestOfProjectVersionNotFound() *CancelDynamicScanRequestOfProjectVersionNotFound {
	return &CancelDynamicScanRequestOfProjectVersionNotFound{}
}

/*CancelDynamicScanRequestOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type CancelDynamicScanRequestOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *CancelDynamicScanRequestOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDynamicScanRequestOfProjectVersionConflict creates a CancelDynamicScanRequestOfProjectVersionConflict with default headers values
func NewCancelDynamicScanRequestOfProjectVersionConflict() *CancelDynamicScanRequestOfProjectVersionConflict {
	return &CancelDynamicScanRequestOfProjectVersionConflict{}
}

/*CancelDynamicScanRequestOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type CancelDynamicScanRequestOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *CancelDynamicScanRequestOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelDynamicScanRequestOfProjectVersionInternalServerError creates a CancelDynamicScanRequestOfProjectVersionInternalServerError with default headers values
func NewCancelDynamicScanRequestOfProjectVersionInternalServerError() *CancelDynamicScanRequestOfProjectVersionInternalServerError {
	return &CancelDynamicScanRequestOfProjectVersionInternalServerError{}
}

/*CancelDynamicScanRequestOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CancelDynamicScanRequestOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *CancelDynamicScanRequestOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/dynamicScanRequests/action/cancel][%d] cancelDynamicScanRequestOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelDynamicScanRequestOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelDynamicScanRequestOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
