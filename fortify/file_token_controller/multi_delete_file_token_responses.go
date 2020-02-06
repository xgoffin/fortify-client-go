// Code generated by go-swagger; DO NOT EDIT.

package file_token_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// MultiDeleteFileTokenReader is a Reader for the MultiDeleteFileToken structure.
type MultiDeleteFileTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiDeleteFileTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiDeleteFileTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMultiDeleteFileTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMultiDeleteFileTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMultiDeleteFileTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMultiDeleteFileTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMultiDeleteFileTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMultiDeleteFileTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMultiDeleteFileTokenOK creates a MultiDeleteFileTokenOK with default headers values
func NewMultiDeleteFileTokenOK() *MultiDeleteFileTokenOK {
	return &MultiDeleteFileTokenOK{}
}

/*MultiDeleteFileTokenOK handles this case with default header values.

OK
*/
type MultiDeleteFileTokenOK struct {
	Payload *models.APIResultVoid
}

func (o *MultiDeleteFileTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenOK  %+v", 200, o.Payload)
}

func (o *MultiDeleteFileTokenOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *MultiDeleteFileTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteFileTokenBadRequest creates a MultiDeleteFileTokenBadRequest with default headers values
func NewMultiDeleteFileTokenBadRequest() *MultiDeleteFileTokenBadRequest {
	return &MultiDeleteFileTokenBadRequest{}
}

/*MultiDeleteFileTokenBadRequest handles this case with default header values.

Bad Request
*/
type MultiDeleteFileTokenBadRequest struct {
	Payload *models.APIResult
}

func (o *MultiDeleteFileTokenBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenBadRequest  %+v", 400, o.Payload)
}

func (o *MultiDeleteFileTokenBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteFileTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteFileTokenUnauthorized creates a MultiDeleteFileTokenUnauthorized with default headers values
func NewMultiDeleteFileTokenUnauthorized() *MultiDeleteFileTokenUnauthorized {
	return &MultiDeleteFileTokenUnauthorized{}
}

/*MultiDeleteFileTokenUnauthorized handles this case with default header values.

Unauthorized
*/
type MultiDeleteFileTokenUnauthorized struct {
	Payload *models.APIResult
}

func (o *MultiDeleteFileTokenUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *MultiDeleteFileTokenUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteFileTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteFileTokenForbidden creates a MultiDeleteFileTokenForbidden with default headers values
func NewMultiDeleteFileTokenForbidden() *MultiDeleteFileTokenForbidden {
	return &MultiDeleteFileTokenForbidden{}
}

/*MultiDeleteFileTokenForbidden handles this case with default header values.

Forbidden
*/
type MultiDeleteFileTokenForbidden struct {
	Payload *models.APIResult
}

func (o *MultiDeleteFileTokenForbidden) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenForbidden  %+v", 403, o.Payload)
}

func (o *MultiDeleteFileTokenForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteFileTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteFileTokenNotFound creates a MultiDeleteFileTokenNotFound with default headers values
func NewMultiDeleteFileTokenNotFound() *MultiDeleteFileTokenNotFound {
	return &MultiDeleteFileTokenNotFound{}
}

/*MultiDeleteFileTokenNotFound handles this case with default header values.

Not Found
*/
type MultiDeleteFileTokenNotFound struct {
	Payload *models.APIResult
}

func (o *MultiDeleteFileTokenNotFound) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenNotFound  %+v", 404, o.Payload)
}

func (o *MultiDeleteFileTokenNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteFileTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteFileTokenConflict creates a MultiDeleteFileTokenConflict with default headers values
func NewMultiDeleteFileTokenConflict() *MultiDeleteFileTokenConflict {
	return &MultiDeleteFileTokenConflict{}
}

/*MultiDeleteFileTokenConflict handles this case with default header values.

Conflict
*/
type MultiDeleteFileTokenConflict struct {
	Payload *models.APIResult
}

func (o *MultiDeleteFileTokenConflict) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenConflict  %+v", 409, o.Payload)
}

func (o *MultiDeleteFileTokenConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteFileTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteFileTokenInternalServerError creates a MultiDeleteFileTokenInternalServerError with default headers values
func NewMultiDeleteFileTokenInternalServerError() *MultiDeleteFileTokenInternalServerError {
	return &MultiDeleteFileTokenInternalServerError{}
}

/*MultiDeleteFileTokenInternalServerError handles this case with default header values.

Internal Server Error
*/
type MultiDeleteFileTokenInternalServerError struct {
	Payload *models.APIResult
}

func (o *MultiDeleteFileTokenInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /fileTokens][%d] multiDeleteFileTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *MultiDeleteFileTokenInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteFileTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
