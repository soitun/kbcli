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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionOK()
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

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*
GetSubscriptionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSubscriptionOK struct {
	Payload      *kbmodel.Subscription
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription o k response
func (o *GetSubscriptionOK) Code() int {
	return 200
}

// IsSuccess returns true when this get subscription o k response has a 2xx status code
func (o *GetSubscriptionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subscription o k response has a 3xx status code
func (o *GetSubscriptionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription o k response has a 4xx status code
func (o *GetSubscriptionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subscription o k response has a 5xx status code
func (o *GetSubscriptionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription o k response a status code equal to that given
func (o *GetSubscriptionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionOK) GetPayload() *kbmodel.Subscription {
	return o.Payload
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionBadRequest creates a GetSubscriptionBadRequest with default headers values
func NewGetSubscriptionBadRequest() *GetSubscriptionBadRequest {
	return &GetSubscriptionBadRequest{}
}

/*
GetSubscriptionBadRequest describes a response with status code 400, with default header values.

Invalid subscription id supplied
*/
type GetSubscriptionBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription bad request response
func (o *GetSubscriptionBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get subscription bad request response has a 2xx status code
func (o *GetSubscriptionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription bad request response has a 3xx status code
func (o *GetSubscriptionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription bad request response has a 4xx status code
func (o *GetSubscriptionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription bad request response has a 5xx status code
func (o *GetSubscriptionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription bad request response a status code equal to that given
func (o *GetSubscriptionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionBadRequest ", 400)
}

func (o *GetSubscriptionBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionBadRequest ", 400)
}

func (o *GetSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionNotFound creates a GetSubscriptionNotFound with default headers values
func NewGetSubscriptionNotFound() *GetSubscriptionNotFound {
	return &GetSubscriptionNotFound{}
}

/*
GetSubscriptionNotFound describes a response with status code 404, with default header values.

Subscription not found
*/
type GetSubscriptionNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription not found response
func (o *GetSubscriptionNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get subscription not found response has a 2xx status code
func (o *GetSubscriptionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription not found response has a 3xx status code
func (o *GetSubscriptionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription not found response has a 4xx status code
func (o *GetSubscriptionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription not found response has a 5xx status code
func (o *GetSubscriptionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription not found response a status code equal to that given
func (o *GetSubscriptionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionNotFound ", 404)
}

func (o *GetSubscriptionNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionNotFound ", 404)
}

func (o *GetSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
