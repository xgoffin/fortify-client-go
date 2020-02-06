// Code generated by go-swagger; DO NOT EDIT.

package parser_plugin_image_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// GetParserPluginImageReader is a Reader for the GetParserPluginImage structure.
type GetParserPluginImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParserPluginImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParserPluginImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetParserPluginImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetParserPluginImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetParserPluginImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetParserPluginImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetParserPluginImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetParserPluginImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParserPluginImageOK creates a GetParserPluginImageOK with default headers values
func NewGetParserPluginImageOK() *GetParserPluginImageOK {
	return &GetParserPluginImageOK{}
}

/*GetParserPluginImageOK handles this case with default header values.

OK
*/
type GetParserPluginImageOK struct {
	Payload strfmt.Base64
}

func (o *GetParserPluginImageOK) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageOK  %+v", 200, o.Payload)
}

func (o *GetParserPluginImageOK) GetPayload() strfmt.Base64 {
	return o.Payload
}

func (o *GetParserPluginImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParserPluginImageBadRequest creates a GetParserPluginImageBadRequest with default headers values
func NewGetParserPluginImageBadRequest() *GetParserPluginImageBadRequest {
	return &GetParserPluginImageBadRequest{}
}

/*GetParserPluginImageBadRequest handles this case with default header values.

Bad Request
*/
type GetParserPluginImageBadRequest struct {
	Payload *models.APIResult
}

func (o *GetParserPluginImageBadRequest) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageBadRequest  %+v", 400, o.Payload)
}

func (o *GetParserPluginImageBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetParserPluginImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParserPluginImageUnauthorized creates a GetParserPluginImageUnauthorized with default headers values
func NewGetParserPluginImageUnauthorized() *GetParserPluginImageUnauthorized {
	return &GetParserPluginImageUnauthorized{}
}

/*GetParserPluginImageUnauthorized handles this case with default header values.

Unauthorized
*/
type GetParserPluginImageUnauthorized struct {
	Payload *models.APIResult
}

func (o *GetParserPluginImageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetParserPluginImageUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetParserPluginImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParserPluginImageForbidden creates a GetParserPluginImageForbidden with default headers values
func NewGetParserPluginImageForbidden() *GetParserPluginImageForbidden {
	return &GetParserPluginImageForbidden{}
}

/*GetParserPluginImageForbidden handles this case with default header values.

Forbidden
*/
type GetParserPluginImageForbidden struct {
	Payload *models.APIResult
}

func (o *GetParserPluginImageForbidden) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageForbidden  %+v", 403, o.Payload)
}

func (o *GetParserPluginImageForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetParserPluginImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParserPluginImageNotFound creates a GetParserPluginImageNotFound with default headers values
func NewGetParserPluginImageNotFound() *GetParserPluginImageNotFound {
	return &GetParserPluginImageNotFound{}
}

/*GetParserPluginImageNotFound handles this case with default header values.

Not Found
*/
type GetParserPluginImageNotFound struct {
	Payload *models.APIResult
}

func (o *GetParserPluginImageNotFound) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageNotFound  %+v", 404, o.Payload)
}

func (o *GetParserPluginImageNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetParserPluginImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParserPluginImageConflict creates a GetParserPluginImageConflict with default headers values
func NewGetParserPluginImageConflict() *GetParserPluginImageConflict {
	return &GetParserPluginImageConflict{}
}

/*GetParserPluginImageConflict handles this case with default header values.

Conflict
*/
type GetParserPluginImageConflict struct {
	Payload *models.APIResult
}

func (o *GetParserPluginImageConflict) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageConflict  %+v", 409, o.Payload)
}

func (o *GetParserPluginImageConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetParserPluginImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParserPluginImageInternalServerError creates a GetParserPluginImageInternalServerError with default headers values
func NewGetParserPluginImageInternalServerError() *GetParserPluginImageInternalServerError {
	return &GetParserPluginImageInternalServerError{}
}

/*GetParserPluginImageInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetParserPluginImageInternalServerError struct {
	Payload *models.APIResult
}

func (o *GetParserPluginImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pluginimage/parser][%d] getParserPluginImageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetParserPluginImageInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetParserPluginImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
