// Code generated by go-swagger; DO NOT EDIT.

package auth_token_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// RevokeAuthTokenReader is a Reader for the RevokeAuthToken structure.
type RevokeAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeAuthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRevokeAuthTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeAuthTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRevokeAuthTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewRevokeAuthTokenConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRevokeAuthTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRevokeAuthTokenOK creates a RevokeAuthTokenOK with default headers values
func NewRevokeAuthTokenOK() *RevokeAuthTokenOK {
	return &RevokeAuthTokenOK{}
}

/*RevokeAuthTokenOK handles this case with default header values.

OK
*/
type RevokeAuthTokenOK struct {
	Payload *models.APIResultVoid
}

func (o *RevokeAuthTokenOK) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenOK  %+v", 200, o.Payload)
}

func (o *RevokeAuthTokenOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *RevokeAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthTokenBadRequest creates a RevokeAuthTokenBadRequest with default headers values
func NewRevokeAuthTokenBadRequest() *RevokeAuthTokenBadRequest {
	return &RevokeAuthTokenBadRequest{}
}

/*RevokeAuthTokenBadRequest handles this case with default header values.

Bad Request
*/
type RevokeAuthTokenBadRequest struct {
	Payload *models.APIResult
}

func (o *RevokeAuthTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenBadRequest  %+v", 400, o.Payload)
}

func (o *RevokeAuthTokenBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *RevokeAuthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthTokenUnauthorized creates a RevokeAuthTokenUnauthorized with default headers values
func NewRevokeAuthTokenUnauthorized() *RevokeAuthTokenUnauthorized {
	return &RevokeAuthTokenUnauthorized{}
}

/*RevokeAuthTokenUnauthorized handles this case with default header values.

Unauthorized
*/
type RevokeAuthTokenUnauthorized struct {
	Payload *models.APIResult
}

func (o *RevokeAuthTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *RevokeAuthTokenUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *RevokeAuthTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthTokenForbidden creates a RevokeAuthTokenForbidden with default headers values
func NewRevokeAuthTokenForbidden() *RevokeAuthTokenForbidden {
	return &RevokeAuthTokenForbidden{}
}

/*RevokeAuthTokenForbidden handles this case with default header values.

Forbidden
*/
type RevokeAuthTokenForbidden struct {
	Payload *models.APIResult
}

func (o *RevokeAuthTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenForbidden  %+v", 403, o.Payload)
}

func (o *RevokeAuthTokenForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *RevokeAuthTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthTokenNotFound creates a RevokeAuthTokenNotFound with default headers values
func NewRevokeAuthTokenNotFound() *RevokeAuthTokenNotFound {
	return &RevokeAuthTokenNotFound{}
}

/*RevokeAuthTokenNotFound handles this case with default header values.

Not Found
*/
type RevokeAuthTokenNotFound struct {
	Payload *models.APIResult
}

func (o *RevokeAuthTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenNotFound  %+v", 404, o.Payload)
}

func (o *RevokeAuthTokenNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *RevokeAuthTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthTokenConflict creates a RevokeAuthTokenConflict with default headers values
func NewRevokeAuthTokenConflict() *RevokeAuthTokenConflict {
	return &RevokeAuthTokenConflict{}
}

/*RevokeAuthTokenConflict handles this case with default header values.

Conflict
*/
type RevokeAuthTokenConflict struct {
	Payload *models.APIResult
}

func (o *RevokeAuthTokenConflict) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenConflict  %+v", 409, o.Payload)
}

func (o *RevokeAuthTokenConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *RevokeAuthTokenConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeAuthTokenInternalServerError creates a RevokeAuthTokenInternalServerError with default headers values
func NewRevokeAuthTokenInternalServerError() *RevokeAuthTokenInternalServerError {
	return &RevokeAuthTokenInternalServerError{}
}

/*RevokeAuthTokenInternalServerError handles this case with default header values.

Internal Server Error
*/
type RevokeAuthTokenInternalServerError struct {
	Payload *models.APIResult
}

func (o *RevokeAuthTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /tokens/action/revoke][%d] revokeAuthTokenInternalServerError  %+v", 500, o.Payload)
}

func (o *RevokeAuthTokenInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *RevokeAuthTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
