// Code generated by go-swagger; DO NOT EDIT.

package ldap_object_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// DeleteLdapObjectReader is a Reader for the DeleteLdapObject structure.
type DeleteLdapObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLdapObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLdapObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteLdapObjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteLdapObjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteLdapObjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteLdapObjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteLdapObjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteLdapObjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLdapObjectOK creates a DeleteLdapObjectOK with default headers values
func NewDeleteLdapObjectOK() *DeleteLdapObjectOK {
	return &DeleteLdapObjectOK{}
}

/*DeleteLdapObjectOK handles this case with default header values.

OK
*/
type DeleteLdapObjectOK struct {
	Payload *models.APIResultVoid
}

func (o *DeleteLdapObjectOK) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectOK  %+v", 200, o.Payload)
}

func (o *DeleteLdapObjectOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *DeleteLdapObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapObjectBadRequest creates a DeleteLdapObjectBadRequest with default headers values
func NewDeleteLdapObjectBadRequest() *DeleteLdapObjectBadRequest {
	return &DeleteLdapObjectBadRequest{}
}

/*DeleteLdapObjectBadRequest handles this case with default header values.

Bad Request
*/
type DeleteLdapObjectBadRequest struct {
	Payload *models.APIResult
}

func (o *DeleteLdapObjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteLdapObjectBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteLdapObjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapObjectUnauthorized creates a DeleteLdapObjectUnauthorized with default headers values
func NewDeleteLdapObjectUnauthorized() *DeleteLdapObjectUnauthorized {
	return &DeleteLdapObjectUnauthorized{}
}

/*DeleteLdapObjectUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteLdapObjectUnauthorized struct {
	Payload *models.APIResult
}

func (o *DeleteLdapObjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteLdapObjectUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteLdapObjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapObjectForbidden creates a DeleteLdapObjectForbidden with default headers values
func NewDeleteLdapObjectForbidden() *DeleteLdapObjectForbidden {
	return &DeleteLdapObjectForbidden{}
}

/*DeleteLdapObjectForbidden handles this case with default header values.

Forbidden
*/
type DeleteLdapObjectForbidden struct {
	Payload *models.APIResult
}

func (o *DeleteLdapObjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectForbidden  %+v", 403, o.Payload)
}

func (o *DeleteLdapObjectForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteLdapObjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapObjectNotFound creates a DeleteLdapObjectNotFound with default headers values
func NewDeleteLdapObjectNotFound() *DeleteLdapObjectNotFound {
	return &DeleteLdapObjectNotFound{}
}

/*DeleteLdapObjectNotFound handles this case with default header values.

Not Found
*/
type DeleteLdapObjectNotFound struct {
	Payload *models.APIResult
}

func (o *DeleteLdapObjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectNotFound  %+v", 404, o.Payload)
}

func (o *DeleteLdapObjectNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteLdapObjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapObjectConflict creates a DeleteLdapObjectConflict with default headers values
func NewDeleteLdapObjectConflict() *DeleteLdapObjectConflict {
	return &DeleteLdapObjectConflict{}
}

/*DeleteLdapObjectConflict handles this case with default header values.

Conflict
*/
type DeleteLdapObjectConflict struct {
	Payload *models.APIResult
}

func (o *DeleteLdapObjectConflict) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectConflict  %+v", 409, o.Payload)
}

func (o *DeleteLdapObjectConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteLdapObjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLdapObjectInternalServerError creates a DeleteLdapObjectInternalServerError with default headers values
func NewDeleteLdapObjectInternalServerError() *DeleteLdapObjectInternalServerError {
	return &DeleteLdapObjectInternalServerError{}
}

/*DeleteLdapObjectInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteLdapObjectInternalServerError struct {
	Payload *models.APIResult
}

func (o *DeleteLdapObjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects/{id}][%d] deleteLdapObjectInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteLdapObjectInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteLdapObjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
