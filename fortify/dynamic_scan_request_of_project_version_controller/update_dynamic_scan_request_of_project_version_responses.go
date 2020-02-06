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

// UpdateDynamicScanRequestOfProjectVersionReader is a Reader for the UpdateDynamicScanRequestOfProjectVersion structure.
type UpdateDynamicScanRequestOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDynamicScanRequestOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDynamicScanRequestOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDynamicScanRequestOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateDynamicScanRequestOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDynamicScanRequestOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDynamicScanRequestOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateDynamicScanRequestOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDynamicScanRequestOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDynamicScanRequestOfProjectVersionOK creates a UpdateDynamicScanRequestOfProjectVersionOK with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionOK() *UpdateDynamicScanRequestOfProjectVersionOK {
	return &UpdateDynamicScanRequestOfProjectVersionOK{}
}

/*UpdateDynamicScanRequestOfProjectVersionOK handles this case with default header values.

OK
*/
type UpdateDynamicScanRequestOfProjectVersionOK struct {
	Payload *models.APIResultDynamicScanRequest
}

func (o *UpdateDynamicScanRequestOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionOK) GetPayload() *models.APIResultDynamicScanRequest {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultDynamicScanRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynamicScanRequestOfProjectVersionBadRequest creates a UpdateDynamicScanRequestOfProjectVersionBadRequest with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionBadRequest() *UpdateDynamicScanRequestOfProjectVersionBadRequest {
	return &UpdateDynamicScanRequestOfProjectVersionBadRequest{}
}

/*UpdateDynamicScanRequestOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateDynamicScanRequestOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateDynamicScanRequestOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynamicScanRequestOfProjectVersionUnauthorized creates a UpdateDynamicScanRequestOfProjectVersionUnauthorized with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionUnauthorized() *UpdateDynamicScanRequestOfProjectVersionUnauthorized {
	return &UpdateDynamicScanRequestOfProjectVersionUnauthorized{}
}

/*UpdateDynamicScanRequestOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateDynamicScanRequestOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateDynamicScanRequestOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynamicScanRequestOfProjectVersionForbidden creates a UpdateDynamicScanRequestOfProjectVersionForbidden with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionForbidden() *UpdateDynamicScanRequestOfProjectVersionForbidden {
	return &UpdateDynamicScanRequestOfProjectVersionForbidden{}
}

/*UpdateDynamicScanRequestOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type UpdateDynamicScanRequestOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateDynamicScanRequestOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynamicScanRequestOfProjectVersionNotFound creates a UpdateDynamicScanRequestOfProjectVersionNotFound with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionNotFound() *UpdateDynamicScanRequestOfProjectVersionNotFound {
	return &UpdateDynamicScanRequestOfProjectVersionNotFound{}
}

/*UpdateDynamicScanRequestOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type UpdateDynamicScanRequestOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateDynamicScanRequestOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynamicScanRequestOfProjectVersionConflict creates a UpdateDynamicScanRequestOfProjectVersionConflict with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionConflict() *UpdateDynamicScanRequestOfProjectVersionConflict {
	return &UpdateDynamicScanRequestOfProjectVersionConflict{}
}

/*UpdateDynamicScanRequestOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type UpdateDynamicScanRequestOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *UpdateDynamicScanRequestOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDynamicScanRequestOfProjectVersionInternalServerError creates a UpdateDynamicScanRequestOfProjectVersionInternalServerError with default headers values
func NewUpdateDynamicScanRequestOfProjectVersionInternalServerError() *UpdateDynamicScanRequestOfProjectVersionInternalServerError {
	return &UpdateDynamicScanRequestOfProjectVersionInternalServerError{}
}

/*UpdateDynamicScanRequestOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateDynamicScanRequestOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateDynamicScanRequestOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] updateDynamicScanRequestOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateDynamicScanRequestOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateDynamicScanRequestOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
