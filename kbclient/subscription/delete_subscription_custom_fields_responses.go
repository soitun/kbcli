// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// DeleteSubscriptionCustomFieldsReader is a Reader for the DeleteSubscriptionCustomFields structure.
type DeleteSubscriptionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubscriptionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSubscriptionCustomFieldsNoContent()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewDeleteSubscriptionCustomFieldsNoContent creates a DeleteSubscriptionCustomFieldsNoContent with default headers values
func NewDeleteSubscriptionCustomFieldsNoContent() *DeleteSubscriptionCustomFieldsNoContent {
	return &DeleteSubscriptionCustomFieldsNoContent{}
}

/*
DeleteSubscriptionCustomFieldsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteSubscriptionCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete subscription custom fields no content response
func (o *DeleteSubscriptionCustomFieldsNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete subscription custom fields no content response has a 2xx status code
func (o *DeleteSubscriptionCustomFieldsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete subscription custom fields no content response has a 3xx status code
func (o *DeleteSubscriptionCustomFieldsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subscription custom fields no content response has a 4xx status code
func (o *DeleteSubscriptionCustomFieldsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete subscription custom fields no content response has a 5xx status code
func (o *DeleteSubscriptionCustomFieldsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subscription custom fields no content response a status code equal to that given
func (o *DeleteSubscriptionCustomFieldsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteSubscriptionCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] deleteSubscriptionCustomFieldsNoContent ", 204)
}

func (o *DeleteSubscriptionCustomFieldsNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] deleteSubscriptionCustomFieldsNoContent ", 204)
}

func (o *DeleteSubscriptionCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubscriptionCustomFieldsBadRequest creates a DeleteSubscriptionCustomFieldsBadRequest with default headers values
func NewDeleteSubscriptionCustomFieldsBadRequest() *DeleteSubscriptionCustomFieldsBadRequest {
	return &DeleteSubscriptionCustomFieldsBadRequest{}
}

/*
DeleteSubscriptionCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid subscription id supplied
*/
type DeleteSubscriptionCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete subscription custom fields bad request response
func (o *DeleteSubscriptionCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete subscription custom fields bad request response has a 2xx status code
func (o *DeleteSubscriptionCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete subscription custom fields bad request response has a 3xx status code
func (o *DeleteSubscriptionCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete subscription custom fields bad request response has a 4xx status code
func (o *DeleteSubscriptionCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete subscription custom fields bad request response has a 5xx status code
func (o *DeleteSubscriptionCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete subscription custom fields bad request response a status code equal to that given
func (o *DeleteSubscriptionCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteSubscriptionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] deleteSubscriptionCustomFieldsBadRequest ", 400)
}

func (o *DeleteSubscriptionCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] deleteSubscriptionCustomFieldsBadRequest ", 400)
}

func (o *DeleteSubscriptionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
