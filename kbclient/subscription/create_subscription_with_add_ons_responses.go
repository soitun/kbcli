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

// CreateSubscriptionWithAddOnsReader is a Reader for the CreateSubscriptionWithAddOns structure.
type CreateSubscriptionWithAddOnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionWithAddOnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateSubscriptionWithAddOnsCreated()
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

// NewCreateSubscriptionWithAddOnsCreated creates a CreateSubscriptionWithAddOnsCreated with default headers values
func NewCreateSubscriptionWithAddOnsCreated() *CreateSubscriptionWithAddOnsCreated {
	return &CreateSubscriptionWithAddOnsCreated{}
}

/*
CreateSubscriptionWithAddOnsCreated describes a response with status code 201, with default header values.

Subscriptions created successfully
*/
type CreateSubscriptionWithAddOnsCreated struct {
	Payload      *kbmodel.Bundle
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create subscription with add ons created response
func (o *CreateSubscriptionWithAddOnsCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create subscription with add ons created response has a 2xx status code
func (o *CreateSubscriptionWithAddOnsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create subscription with add ons created response has a 3xx status code
func (o *CreateSubscriptionWithAddOnsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create subscription with add ons created response has a 4xx status code
func (o *CreateSubscriptionWithAddOnsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create subscription with add ons created response has a 5xx status code
func (o *CreateSubscriptionWithAddOnsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create subscription with add ons created response a status code equal to that given
func (o *CreateSubscriptionWithAddOnsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateSubscriptionWithAddOnsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/createSubscriptionWithAddOns][%d] createSubscriptionWithAddOnsCreated  %+v", 201, o.Payload)
}

func (o *CreateSubscriptionWithAddOnsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/createSubscriptionWithAddOns][%d] createSubscriptionWithAddOnsCreated  %+v", 201, o.Payload)
}

func (o *CreateSubscriptionWithAddOnsCreated) GetPayload() *kbmodel.Bundle {
	return o.Payload
}

func (o *CreateSubscriptionWithAddOnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Bundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
