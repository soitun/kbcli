// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ModifyPaymentCustomFieldsReader is a Reader for the ModifyPaymentCustomFields structure.
type ModifyPaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyPaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewModifyPaymentCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewModifyPaymentCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyPaymentCustomFieldsNoContent creates a ModifyPaymentCustomFieldsNoContent with default headers values
func NewModifyPaymentCustomFieldsNoContent() *ModifyPaymentCustomFieldsNoContent {
	return &ModifyPaymentCustomFieldsNoContent{}
}

/*
ModifyPaymentCustomFieldsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type ModifyPaymentCustomFieldsNoContent struct {
}

// IsSuccess returns true when this modify payment custom fields no content response has a 2xx status code
func (o *ModifyPaymentCustomFieldsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify payment custom fields no content response has a 3xx status code
func (o *ModifyPaymentCustomFieldsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify payment custom fields no content response has a 4xx status code
func (o *ModifyPaymentCustomFieldsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify payment custom fields no content response has a 5xx status code
func (o *ModifyPaymentCustomFieldsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this modify payment custom fields no content response a status code equal to that given
func (o *ModifyPaymentCustomFieldsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the modify payment custom fields no content response
func (o *ModifyPaymentCustomFieldsNoContent) Code() int {
	return 204
}

func (o *ModifyPaymentCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}/customFields][%d] modifyPaymentCustomFieldsNoContent ", 204)
}

func (o *ModifyPaymentCustomFieldsNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}/customFields][%d] modifyPaymentCustomFieldsNoContent ", 204)
}

func (o *ModifyPaymentCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPaymentCustomFieldsBadRequest creates a ModifyPaymentCustomFieldsBadRequest with default headers values
func NewModifyPaymentCustomFieldsBadRequest() *ModifyPaymentCustomFieldsBadRequest {
	return &ModifyPaymentCustomFieldsBadRequest{}
}

/*
ModifyPaymentCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type ModifyPaymentCustomFieldsBadRequest struct {
}

// IsSuccess returns true when this modify payment custom fields bad request response has a 2xx status code
func (o *ModifyPaymentCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify payment custom fields bad request response has a 3xx status code
func (o *ModifyPaymentCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify payment custom fields bad request response has a 4xx status code
func (o *ModifyPaymentCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify payment custom fields bad request response has a 5xx status code
func (o *ModifyPaymentCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this modify payment custom fields bad request response a status code equal to that given
func (o *ModifyPaymentCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the modify payment custom fields bad request response
func (o *ModifyPaymentCustomFieldsBadRequest) Code() int {
	return 400
}

func (o *ModifyPaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}/customFields][%d] modifyPaymentCustomFieldsBadRequest ", 400)
}

func (o *ModifyPaymentCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}/customFields][%d] modifyPaymentCustomFieldsBadRequest ", 400)
}

func (o *ModifyPaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
