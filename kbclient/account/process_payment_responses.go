// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// ProcessPaymentReader is a Reader for the ProcessPayment structure.
type ProcessPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProcessPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewProcessPaymentCreated()
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

// NewProcessPaymentCreated creates a ProcessPaymentCreated with default headers values
func NewProcessPaymentCreated() *ProcessPaymentCreated {
	return &ProcessPaymentCreated{}
}

/*ProcessPaymentCreated handles this case with default header values.

Payment transaction created successfully
*/
type ProcessPaymentCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentCreated  %+v", 201, o.Payload)
}

func (o *ProcessPaymentCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ProcessPaymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProcessPaymentBadRequest creates a ProcessPaymentBadRequest with default headers values
func NewProcessPaymentBadRequest() *ProcessPaymentBadRequest {
	return &ProcessPaymentBadRequest{}
}

/*ProcessPaymentBadRequest handles this case with default header values.

Invalid account id supplied
*/
type ProcessPaymentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentBadRequest ", 400)
}

func (o *ProcessPaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentPaymentRequired creates a ProcessPaymentPaymentRequired with default headers values
func NewProcessPaymentPaymentRequired() *ProcessPaymentPaymentRequired {
	return &ProcessPaymentPaymentRequired{}
}

/*ProcessPaymentPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type ProcessPaymentPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentPaymentRequired ", 402)
}

func (o *ProcessPaymentPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentNotFound creates a ProcessPaymentNotFound with default headers values
func NewProcessPaymentNotFound() *ProcessPaymentNotFound {
	return &ProcessPaymentNotFound{}
}

/*ProcessPaymentNotFound handles this case with default header values.

Account not found
*/
type ProcessPaymentNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentNotFound ", 404)
}

func (o *ProcessPaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentUnprocessableEntity creates a ProcessPaymentUnprocessableEntity with default headers values
func NewProcessPaymentUnprocessableEntity() *ProcessPaymentUnprocessableEntity {
	return &ProcessPaymentUnprocessableEntity{}
}

/*ProcessPaymentUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type ProcessPaymentUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentUnprocessableEntity ", 422)
}

func (o *ProcessPaymentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentBadGateway creates a ProcessPaymentBadGateway with default headers values
func NewProcessPaymentBadGateway() *ProcessPaymentBadGateway {
	return &ProcessPaymentBadGateway{}
}

/*ProcessPaymentBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type ProcessPaymentBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentBadGateway ", 502)
}

func (o *ProcessPaymentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentServiceUnavailable creates a ProcessPaymentServiceUnavailable with default headers values
func NewProcessPaymentServiceUnavailable() *ProcessPaymentServiceUnavailable {
	return &ProcessPaymentServiceUnavailable{}
}

/*ProcessPaymentServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ProcessPaymentServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentServiceUnavailable ", 503)
}

func (o *ProcessPaymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewProcessPaymentGatewayTimeout creates a ProcessPaymentGatewayTimeout with default headers values
func NewProcessPaymentGatewayTimeout() *ProcessPaymentGatewayTimeout {
	return &ProcessPaymentGatewayTimeout{}
}

/*ProcessPaymentGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type ProcessPaymentGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *ProcessPaymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/payments][%d] processPaymentGatewayTimeout ", 504)
}

func (o *ProcessPaymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
