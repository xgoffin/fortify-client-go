// Code generated by go-swagger; DO NOT EDIT.

package performance_indicator_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreatePerformanceIndicatorReader is a Reader for the CreatePerformanceIndicator structure.
type CreatePerformanceIndicatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePerformanceIndicatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePerformanceIndicatorCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePerformanceIndicatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreatePerformanceIndicatorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePerformanceIndicatorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreatePerformanceIndicatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePerformanceIndicatorConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePerformanceIndicatorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePerformanceIndicatorCreated creates a CreatePerformanceIndicatorCreated with default headers values
func NewCreatePerformanceIndicatorCreated() *CreatePerformanceIndicatorCreated {
	return &CreatePerformanceIndicatorCreated{}
}

/*CreatePerformanceIndicatorCreated handles this case with default header values.

Created
*/
type CreatePerformanceIndicatorCreated struct {
	Payload *models.APIResultPerformanceIndicator
}

func (o *CreatePerformanceIndicatorCreated) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorCreated  %+v", 201, o.Payload)
}

func (o *CreatePerformanceIndicatorCreated) GetPayload() *models.APIResultPerformanceIndicator {
	return o.Payload
}

func (o *CreatePerformanceIndicatorCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultPerformanceIndicator)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePerformanceIndicatorBadRequest creates a CreatePerformanceIndicatorBadRequest with default headers values
func NewCreatePerformanceIndicatorBadRequest() *CreatePerformanceIndicatorBadRequest {
	return &CreatePerformanceIndicatorBadRequest{}
}

/*CreatePerformanceIndicatorBadRequest handles this case with default header values.

Bad Request
*/
type CreatePerformanceIndicatorBadRequest struct {
	Payload *models.APIResult
}

func (o *CreatePerformanceIndicatorBadRequest) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePerformanceIndicatorBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreatePerformanceIndicatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePerformanceIndicatorUnauthorized creates a CreatePerformanceIndicatorUnauthorized with default headers values
func NewCreatePerformanceIndicatorUnauthorized() *CreatePerformanceIndicatorUnauthorized {
	return &CreatePerformanceIndicatorUnauthorized{}
}

/*CreatePerformanceIndicatorUnauthorized handles this case with default header values.

Unauthorized
*/
type CreatePerformanceIndicatorUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreatePerformanceIndicatorUnauthorized) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePerformanceIndicatorUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreatePerformanceIndicatorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePerformanceIndicatorForbidden creates a CreatePerformanceIndicatorForbidden with default headers values
func NewCreatePerformanceIndicatorForbidden() *CreatePerformanceIndicatorForbidden {
	return &CreatePerformanceIndicatorForbidden{}
}

/*CreatePerformanceIndicatorForbidden handles this case with default header values.

Forbidden
*/
type CreatePerformanceIndicatorForbidden struct {
	Payload *models.APIResult
}

func (o *CreatePerformanceIndicatorForbidden) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorForbidden  %+v", 403, o.Payload)
}

func (o *CreatePerformanceIndicatorForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreatePerformanceIndicatorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePerformanceIndicatorNotFound creates a CreatePerformanceIndicatorNotFound with default headers values
func NewCreatePerformanceIndicatorNotFound() *CreatePerformanceIndicatorNotFound {
	return &CreatePerformanceIndicatorNotFound{}
}

/*CreatePerformanceIndicatorNotFound handles this case with default header values.

Not Found
*/
type CreatePerformanceIndicatorNotFound struct {
	Payload *models.APIResult
}

func (o *CreatePerformanceIndicatorNotFound) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorNotFound  %+v", 404, o.Payload)
}

func (o *CreatePerformanceIndicatorNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreatePerformanceIndicatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePerformanceIndicatorConflict creates a CreatePerformanceIndicatorConflict with default headers values
func NewCreatePerformanceIndicatorConflict() *CreatePerformanceIndicatorConflict {
	return &CreatePerformanceIndicatorConflict{}
}

/*CreatePerformanceIndicatorConflict handles this case with default header values.

Conflict
*/
type CreatePerformanceIndicatorConflict struct {
	Payload *models.APIResult
}

func (o *CreatePerformanceIndicatorConflict) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorConflict  %+v", 409, o.Payload)
}

func (o *CreatePerformanceIndicatorConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreatePerformanceIndicatorConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePerformanceIndicatorInternalServerError creates a CreatePerformanceIndicatorInternalServerError with default headers values
func NewCreatePerformanceIndicatorInternalServerError() *CreatePerformanceIndicatorInternalServerError {
	return &CreatePerformanceIndicatorInternalServerError{}
}

/*CreatePerformanceIndicatorInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreatePerformanceIndicatorInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreatePerformanceIndicatorInternalServerError) Error() string {
	return fmt.Sprintf("[POST /performanceIndicators][%d] createPerformanceIndicatorInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePerformanceIndicatorInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreatePerformanceIndicatorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
