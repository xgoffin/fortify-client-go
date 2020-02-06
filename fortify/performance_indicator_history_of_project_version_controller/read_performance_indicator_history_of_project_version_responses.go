// Code generated by go-swagger; DO NOT EDIT.

package performance_indicator_history_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadPerformanceIndicatorHistoryOfProjectVersionReader is a Reader for the ReadPerformanceIndicatorHistoryOfProjectVersion structure.
type ReadPerformanceIndicatorHistoryOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionOK creates a ReadPerformanceIndicatorHistoryOfProjectVersionOK with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionOK() *ReadPerformanceIndicatorHistoryOfProjectVersionOK {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionOK{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionOK handles this case with default header values.

OK
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionOK struct {
	Payload *models.APIResultPerformanceIndicatorHistory
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionOK) GetPayload() *models.APIResultPerformanceIndicatorHistory {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultPerformanceIndicatorHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionBadRequest creates a ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionBadRequest() *ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized creates a ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized() *ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionForbidden creates a ReadPerformanceIndicatorHistoryOfProjectVersionForbidden with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionForbidden() *ReadPerformanceIndicatorHistoryOfProjectVersionForbidden {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionForbidden{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionNotFound creates a ReadPerformanceIndicatorHistoryOfProjectVersionNotFound with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionNotFound() *ReadPerformanceIndicatorHistoryOfProjectVersionNotFound {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionNotFound{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionConflict creates a ReadPerformanceIndicatorHistoryOfProjectVersionConflict with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionConflict() *ReadPerformanceIndicatorHistoryOfProjectVersionConflict {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionConflict{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError creates a ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError with default headers values
func NewReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError() *ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError {
	return &ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError{}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/performanceIndicatorHistories/{id}][%d] readPerformanceIndicatorHistoryOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPerformanceIndicatorHistoryOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
