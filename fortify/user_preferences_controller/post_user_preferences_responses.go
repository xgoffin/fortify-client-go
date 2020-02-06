// Code generated by go-swagger; DO NOT EDIT.

package user_preferences_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// PostUserPreferencesReader is a Reader for the PostUserPreferences structure.
type PostUserPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserPreferencesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserPreferencesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserPreferencesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserPreferencesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostUserPreferencesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserPreferencesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUserPreferencesOK creates a PostUserPreferencesOK with default headers values
func NewPostUserPreferencesOK() *PostUserPreferencesOK {
	return &PostUserPreferencesOK{}
}

/*PostUserPreferencesOK handles this case with default header values.

OK
*/
type PostUserPreferencesOK struct {
	Payload *models.APIResultUserPreferences
}

func (o *PostUserPreferencesOK) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesOK  %+v", 200, o.Payload)
}

func (o *PostUserPreferencesOK) GetPayload() *models.APIResultUserPreferences {
	return o.Payload
}

func (o *PostUserPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultUserPreferences)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPreferencesBadRequest creates a PostUserPreferencesBadRequest with default headers values
func NewPostUserPreferencesBadRequest() *PostUserPreferencesBadRequest {
	return &PostUserPreferencesBadRequest{}
}

/*PostUserPreferencesBadRequest handles this case with default header values.

Bad Request
*/
type PostUserPreferencesBadRequest struct {
	Payload *models.APIResult
}

func (o *PostUserPreferencesBadRequest) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserPreferencesBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserPreferencesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPreferencesUnauthorized creates a PostUserPreferencesUnauthorized with default headers values
func NewPostUserPreferencesUnauthorized() *PostUserPreferencesUnauthorized {
	return &PostUserPreferencesUnauthorized{}
}

/*PostUserPreferencesUnauthorized handles this case with default header values.

Unauthorized
*/
type PostUserPreferencesUnauthorized struct {
	Payload *models.APIResult
}

func (o *PostUserPreferencesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserPreferencesUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserPreferencesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPreferencesForbidden creates a PostUserPreferencesForbidden with default headers values
func NewPostUserPreferencesForbidden() *PostUserPreferencesForbidden {
	return &PostUserPreferencesForbidden{}
}

/*PostUserPreferencesForbidden handles this case with default header values.

Forbidden
*/
type PostUserPreferencesForbidden struct {
	Payload *models.APIResult
}

func (o *PostUserPreferencesForbidden) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesForbidden  %+v", 403, o.Payload)
}

func (o *PostUserPreferencesForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserPreferencesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPreferencesNotFound creates a PostUserPreferencesNotFound with default headers values
func NewPostUserPreferencesNotFound() *PostUserPreferencesNotFound {
	return &PostUserPreferencesNotFound{}
}

/*PostUserPreferencesNotFound handles this case with default header values.

Not Found
*/
type PostUserPreferencesNotFound struct {
	Payload *models.APIResult
}

func (o *PostUserPreferencesNotFound) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesNotFound  %+v", 404, o.Payload)
}

func (o *PostUserPreferencesNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserPreferencesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPreferencesConflict creates a PostUserPreferencesConflict with default headers values
func NewPostUserPreferencesConflict() *PostUserPreferencesConflict {
	return &PostUserPreferencesConflict{}
}

/*PostUserPreferencesConflict handles this case with default header values.

Conflict
*/
type PostUserPreferencesConflict struct {
	Payload *models.APIResult
}

func (o *PostUserPreferencesConflict) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesConflict  %+v", 409, o.Payload)
}

func (o *PostUserPreferencesConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserPreferencesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserPreferencesInternalServerError creates a PostUserPreferencesInternalServerError with default headers values
func NewPostUserPreferencesInternalServerError() *PostUserPreferencesInternalServerError {
	return &PostUserPreferencesInternalServerError{}
}

/*PostUserPreferencesInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostUserPreferencesInternalServerError struct {
	Payload *models.APIResult
}

func (o *PostUserPreferencesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /userSession/prefs][%d] postUserPreferencesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserPreferencesInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserPreferencesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
