// Code generated by go-swagger; DO NOT EDIT.

package cloud_job_of_cloud_worker_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListCloudJobOfCloudWorkerReader is a Reader for the ListCloudJobOfCloudWorker structure.
type ListCloudJobOfCloudWorkerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudJobOfCloudWorkerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCloudJobOfCloudWorkerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCloudJobOfCloudWorkerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCloudJobOfCloudWorkerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCloudJobOfCloudWorkerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCloudJobOfCloudWorkerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListCloudJobOfCloudWorkerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCloudJobOfCloudWorkerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCloudJobOfCloudWorkerOK creates a ListCloudJobOfCloudWorkerOK with default headers values
func NewListCloudJobOfCloudWorkerOK() *ListCloudJobOfCloudWorkerOK {
	return &ListCloudJobOfCloudWorkerOK{}
}

/*ListCloudJobOfCloudWorkerOK handles this case with default header values.

OK
*/
type ListCloudJobOfCloudWorkerOK struct {
	Payload *models.APIResultListCloudJob
}

func (o *ListCloudJobOfCloudWorkerOK) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerOK  %+v", 200, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerOK) GetPayload() *models.APIResultListCloudJob {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListCloudJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobOfCloudWorkerBadRequest creates a ListCloudJobOfCloudWorkerBadRequest with default headers values
func NewListCloudJobOfCloudWorkerBadRequest() *ListCloudJobOfCloudWorkerBadRequest {
	return &ListCloudJobOfCloudWorkerBadRequest{}
}

/*ListCloudJobOfCloudWorkerBadRequest handles this case with default header values.

Bad Request
*/
type ListCloudJobOfCloudWorkerBadRequest struct {
	Payload *models.APIResult
}

func (o *ListCloudJobOfCloudWorkerBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerBadRequest  %+v", 400, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobOfCloudWorkerUnauthorized creates a ListCloudJobOfCloudWorkerUnauthorized with default headers values
func NewListCloudJobOfCloudWorkerUnauthorized() *ListCloudJobOfCloudWorkerUnauthorized {
	return &ListCloudJobOfCloudWorkerUnauthorized{}
}

/*ListCloudJobOfCloudWorkerUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCloudJobOfCloudWorkerUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListCloudJobOfCloudWorkerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobOfCloudWorkerForbidden creates a ListCloudJobOfCloudWorkerForbidden with default headers values
func NewListCloudJobOfCloudWorkerForbidden() *ListCloudJobOfCloudWorkerForbidden {
	return &ListCloudJobOfCloudWorkerForbidden{}
}

/*ListCloudJobOfCloudWorkerForbidden handles this case with default header values.

Forbidden
*/
type ListCloudJobOfCloudWorkerForbidden struct {
	Payload *models.APIResult
}

func (o *ListCloudJobOfCloudWorkerForbidden) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerForbidden  %+v", 403, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobOfCloudWorkerNotFound creates a ListCloudJobOfCloudWorkerNotFound with default headers values
func NewListCloudJobOfCloudWorkerNotFound() *ListCloudJobOfCloudWorkerNotFound {
	return &ListCloudJobOfCloudWorkerNotFound{}
}

/*ListCloudJobOfCloudWorkerNotFound handles this case with default header values.

Not Found
*/
type ListCloudJobOfCloudWorkerNotFound struct {
	Payload *models.APIResult
}

func (o *ListCloudJobOfCloudWorkerNotFound) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerNotFound  %+v", 404, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobOfCloudWorkerConflict creates a ListCloudJobOfCloudWorkerConflict with default headers values
func NewListCloudJobOfCloudWorkerConflict() *ListCloudJobOfCloudWorkerConflict {
	return &ListCloudJobOfCloudWorkerConflict{}
}

/*ListCloudJobOfCloudWorkerConflict handles this case with default header values.

Conflict
*/
type ListCloudJobOfCloudWorkerConflict struct {
	Payload *models.APIResult
}

func (o *ListCloudJobOfCloudWorkerConflict) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerConflict  %+v", 409, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudJobOfCloudWorkerInternalServerError creates a ListCloudJobOfCloudWorkerInternalServerError with default headers values
func NewListCloudJobOfCloudWorkerInternalServerError() *ListCloudJobOfCloudWorkerInternalServerError {
	return &ListCloudJobOfCloudWorkerInternalServerError{}
}

/*ListCloudJobOfCloudWorkerInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCloudJobOfCloudWorkerInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListCloudJobOfCloudWorkerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudworkers/{parentId}/cloudjobs][%d] listCloudJobOfCloudWorkerInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCloudJobOfCloudWorkerInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudJobOfCloudWorkerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
