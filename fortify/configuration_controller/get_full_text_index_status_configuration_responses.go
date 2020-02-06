// Code generated by go-swagger; DO NOT EDIT.

package configuration_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// GetFullTextIndexStatusConfigurationReader is a Reader for the GetFullTextIndexStatusConfiguration structure.
type GetFullTextIndexStatusConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFullTextIndexStatusConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFullTextIndexStatusConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetFullTextIndexStatusConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetFullTextIndexStatusConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetFullTextIndexStatusConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFullTextIndexStatusConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetFullTextIndexStatusConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetFullTextIndexStatusConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFullTextIndexStatusConfigurationOK creates a GetFullTextIndexStatusConfigurationOK with default headers values
func NewGetFullTextIndexStatusConfigurationOK() *GetFullTextIndexStatusConfigurationOK {
	return &GetFullTextIndexStatusConfigurationOK{}
}

/*GetFullTextIndexStatusConfigurationOK handles this case with default header values.

OK
*/
type GetFullTextIndexStatusConfigurationOK struct {
	Payload *models.APIResultSearchIndexStatus
}

func (o *GetFullTextIndexStatusConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationOK) GetPayload() *models.APIResultSearchIndexStatus {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultSearchIndexStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFullTextIndexStatusConfigurationBadRequest creates a GetFullTextIndexStatusConfigurationBadRequest with default headers values
func NewGetFullTextIndexStatusConfigurationBadRequest() *GetFullTextIndexStatusConfigurationBadRequest {
	return &GetFullTextIndexStatusConfigurationBadRequest{}
}

/*GetFullTextIndexStatusConfigurationBadRequest handles this case with default header values.

Bad Request
*/
type GetFullTextIndexStatusConfigurationBadRequest struct {
	Payload *models.APIResult
}

func (o *GetFullTextIndexStatusConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFullTextIndexStatusConfigurationUnauthorized creates a GetFullTextIndexStatusConfigurationUnauthorized with default headers values
func NewGetFullTextIndexStatusConfigurationUnauthorized() *GetFullTextIndexStatusConfigurationUnauthorized {
	return &GetFullTextIndexStatusConfigurationUnauthorized{}
}

/*GetFullTextIndexStatusConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type GetFullTextIndexStatusConfigurationUnauthorized struct {
	Payload *models.APIResult
}

func (o *GetFullTextIndexStatusConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFullTextIndexStatusConfigurationForbidden creates a GetFullTextIndexStatusConfigurationForbidden with default headers values
func NewGetFullTextIndexStatusConfigurationForbidden() *GetFullTextIndexStatusConfigurationForbidden {
	return &GetFullTextIndexStatusConfigurationForbidden{}
}

/*GetFullTextIndexStatusConfigurationForbidden handles this case with default header values.

Forbidden
*/
type GetFullTextIndexStatusConfigurationForbidden struct {
	Payload *models.APIResult
}

func (o *GetFullTextIndexStatusConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFullTextIndexStatusConfigurationNotFound creates a GetFullTextIndexStatusConfigurationNotFound with default headers values
func NewGetFullTextIndexStatusConfigurationNotFound() *GetFullTextIndexStatusConfigurationNotFound {
	return &GetFullTextIndexStatusConfigurationNotFound{}
}

/*GetFullTextIndexStatusConfigurationNotFound handles this case with default header values.

Not Found
*/
type GetFullTextIndexStatusConfigurationNotFound struct {
	Payload *models.APIResult
}

func (o *GetFullTextIndexStatusConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFullTextIndexStatusConfigurationConflict creates a GetFullTextIndexStatusConfigurationConflict with default headers values
func NewGetFullTextIndexStatusConfigurationConflict() *GetFullTextIndexStatusConfigurationConflict {
	return &GetFullTextIndexStatusConfigurationConflict{}
}

/*GetFullTextIndexStatusConfigurationConflict handles this case with default header values.

Conflict
*/
type GetFullTextIndexStatusConfigurationConflict struct {
	Payload *models.APIResult
}

func (o *GetFullTextIndexStatusConfigurationConflict) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationConflict  %+v", 409, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFullTextIndexStatusConfigurationInternalServerError creates a GetFullTextIndexStatusConfigurationInternalServerError with default headers values
func NewGetFullTextIndexStatusConfigurationInternalServerError() *GetFullTextIndexStatusConfigurationInternalServerError {
	return &GetFullTextIndexStatusConfigurationInternalServerError{}
}

/*GetFullTextIndexStatusConfigurationInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetFullTextIndexStatusConfigurationInternalServerError struct {
	Payload *models.APIResult
}

func (o *GetFullTextIndexStatusConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /configuration/searchIndex][%d] getFullTextIndexStatusConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFullTextIndexStatusConfigurationInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetFullTextIndexStatusConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
