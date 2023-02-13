// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ModifySubscriptionCustomFieldsReader is a Reader for the ModifySubscriptionCustomFields structure.
type ModifySubscriptionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifySubscriptionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewModifySubscriptionCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewModifySubscriptionCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifySubscriptionCustomFieldsNoContent creates a ModifySubscriptionCustomFieldsNoContent with default headers values
func NewModifySubscriptionCustomFieldsNoContent() *ModifySubscriptionCustomFieldsNoContent {
	return &ModifySubscriptionCustomFieldsNoContent{}
}

/*
ModifySubscriptionCustomFieldsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type ModifySubscriptionCustomFieldsNoContent struct {
}

// IsSuccess returns true when this modify subscription custom fields no content response has a 2xx status code
func (o *ModifySubscriptionCustomFieldsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify subscription custom fields no content response has a 3xx status code
func (o *ModifySubscriptionCustomFieldsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription custom fields no content response has a 4xx status code
func (o *ModifySubscriptionCustomFieldsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify subscription custom fields no content response has a 5xx status code
func (o *ModifySubscriptionCustomFieldsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this modify subscription custom fields no content response a status code equal to that given
func (o *ModifySubscriptionCustomFieldsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the modify subscription custom fields no content response
func (o *ModifySubscriptionCustomFieldsNoContent) Code() int {
	return 204
}

func (o *ModifySubscriptionCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] modifySubscriptionCustomFieldsNoContent ", 204)
}

func (o *ModifySubscriptionCustomFieldsNoContent) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] modifySubscriptionCustomFieldsNoContent ", 204)
}

func (o *ModifySubscriptionCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifySubscriptionCustomFieldsBadRequest creates a ModifySubscriptionCustomFieldsBadRequest with default headers values
func NewModifySubscriptionCustomFieldsBadRequest() *ModifySubscriptionCustomFieldsBadRequest {
	return &ModifySubscriptionCustomFieldsBadRequest{}
}

/*
ModifySubscriptionCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid subscription id supplied
*/
type ModifySubscriptionCustomFieldsBadRequest struct {
}

// IsSuccess returns true when this modify subscription custom fields bad request response has a 2xx status code
func (o *ModifySubscriptionCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this modify subscription custom fields bad request response has a 3xx status code
func (o *ModifySubscriptionCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify subscription custom fields bad request response has a 4xx status code
func (o *ModifySubscriptionCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this modify subscription custom fields bad request response has a 5xx status code
func (o *ModifySubscriptionCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this modify subscription custom fields bad request response a status code equal to that given
func (o *ModifySubscriptionCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the modify subscription custom fields bad request response
func (o *ModifySubscriptionCustomFieldsBadRequest) Code() int {
	return 400
}

func (o *ModifySubscriptionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] modifySubscriptionCustomFieldsBadRequest ", 400)
}

func (o *ModifySubscriptionCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[PUT /1.0/kb/subscriptions/{subscriptionId}/customFields][%d] modifySubscriptionCustomFieldsBadRequest ", 400)
}

func (o *ModifySubscriptionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
