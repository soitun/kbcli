// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RenameExternalKeyReader is a Reader for the RenameExternalKey structure.
type RenameExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenameExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRenameExternalKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRenameExternalKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRenameExternalKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRenameExternalKeyNoContent creates a RenameExternalKeyNoContent with default headers values
func NewRenameExternalKeyNoContent() *RenameExternalKeyNoContent {
	return &RenameExternalKeyNoContent{}
}

/*
RenameExternalKeyNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type RenameExternalKeyNoContent struct {
}

// IsSuccess returns true when this rename external key no content response has a 2xx status code
func (o *RenameExternalKeyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rename external key no content response has a 3xx status code
func (o *RenameExternalKeyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename external key no content response has a 4xx status code
func (o *RenameExternalKeyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this rename external key no content response has a 5xx status code
func (o *RenameExternalKeyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this rename external key no content response a status code equal to that given
func (o *RenameExternalKeyNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the rename external key no content response
func (o *RenameExternalKeyNoContent) Code() int {
	return 204
}

func (o *RenameExternalKeyNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyNoContent ", 204)
}

func (o *RenameExternalKeyNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyNoContent ", 204)
}

func (o *RenameExternalKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameExternalKeyBadRequest creates a RenameExternalKeyBadRequest with default headers values
func NewRenameExternalKeyBadRequest() *RenameExternalKeyBadRequest {
	return &RenameExternalKeyBadRequest{}
}

/*
RenameExternalKeyBadRequest describes a response with status code 400, with default header values.

Invalid argumnent supplied
*/
type RenameExternalKeyBadRequest struct {
}

// IsSuccess returns true when this rename external key bad request response has a 2xx status code
func (o *RenameExternalKeyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename external key bad request response has a 3xx status code
func (o *RenameExternalKeyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename external key bad request response has a 4xx status code
func (o *RenameExternalKeyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename external key bad request response has a 5xx status code
func (o *RenameExternalKeyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this rename external key bad request response a status code equal to that given
func (o *RenameExternalKeyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the rename external key bad request response
func (o *RenameExternalKeyBadRequest) Code() int {
	return 400
}

func (o *RenameExternalKeyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyBadRequest ", 400)
}

func (o *RenameExternalKeyBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyBadRequest ", 400)
}

func (o *RenameExternalKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameExternalKeyNotFound creates a RenameExternalKeyNotFound with default headers values
func NewRenameExternalKeyNotFound() *RenameExternalKeyNotFound {
	return &RenameExternalKeyNotFound{}
}

/*
RenameExternalKeyNotFound describes a response with status code 404, with default header values.

Bundle not found
*/
type RenameExternalKeyNotFound struct {
}

// IsSuccess returns true when this rename external key not found response has a 2xx status code
func (o *RenameExternalKeyNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this rename external key not found response has a 3xx status code
func (o *RenameExternalKeyNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rename external key not found response has a 4xx status code
func (o *RenameExternalKeyNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this rename external key not found response has a 5xx status code
func (o *RenameExternalKeyNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this rename external key not found response a status code equal to that given
func (o *RenameExternalKeyNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the rename external key not found response
func (o *RenameExternalKeyNotFound) Code() int {
	return 404
}

func (o *RenameExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyNotFound ", 404)
}

func (o *RenameExternalKeyNotFound) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyNotFound ", 404)
}

func (o *RenameExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
