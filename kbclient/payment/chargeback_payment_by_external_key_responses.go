// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// ChargebackPaymentByExternalKeyReader is a Reader for the ChargebackPaymentByExternalKey structure.
type ChargebackPaymentByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargebackPaymentByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewChargebackPaymentByExternalKeyCreated()
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

// NewChargebackPaymentByExternalKeyCreated creates a ChargebackPaymentByExternalKeyCreated with default headers values
func NewChargebackPaymentByExternalKeyCreated() *ChargebackPaymentByExternalKeyCreated {
	return &ChargebackPaymentByExternalKeyCreated{}
}

/*ChargebackPaymentByExternalKeyCreated handles this case with default header values.

Payment transaction created successfully
*/
type ChargebackPaymentByExternalKeyCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyCreated  %+v", 201, o.Payload)
}

func (o *ChargebackPaymentByExternalKeyCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ChargebackPaymentByExternalKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargebackPaymentByExternalKeyPaymentRequired creates a ChargebackPaymentByExternalKeyPaymentRequired with default headers values
func NewChargebackPaymentByExternalKeyPaymentRequired() *ChargebackPaymentByExternalKeyPaymentRequired {
	return &ChargebackPaymentByExternalKeyPaymentRequired{}
}

/*ChargebackPaymentByExternalKeyPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type ChargebackPaymentByExternalKeyPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyPaymentRequired ", 402)
}

func (o *ChargebackPaymentByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyNotFound creates a ChargebackPaymentByExternalKeyNotFound with default headers values
func NewChargebackPaymentByExternalKeyNotFound() *ChargebackPaymentByExternalKeyNotFound {
	return &ChargebackPaymentByExternalKeyNotFound{}
}

/*ChargebackPaymentByExternalKeyNotFound handles this case with default header values.

Account or payment not found
*/
type ChargebackPaymentByExternalKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyNotFound ", 404)
}

func (o *ChargebackPaymentByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyUnprocessableEntity creates a ChargebackPaymentByExternalKeyUnprocessableEntity with default headers values
func NewChargebackPaymentByExternalKeyUnprocessableEntity() *ChargebackPaymentByExternalKeyUnprocessableEntity {
	return &ChargebackPaymentByExternalKeyUnprocessableEntity{}
}

/*ChargebackPaymentByExternalKeyUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type ChargebackPaymentByExternalKeyUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyUnprocessableEntity ", 422)
}

func (o *ChargebackPaymentByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyBadGateway creates a ChargebackPaymentByExternalKeyBadGateway with default headers values
func NewChargebackPaymentByExternalKeyBadGateway() *ChargebackPaymentByExternalKeyBadGateway {
	return &ChargebackPaymentByExternalKeyBadGateway{}
}

/*ChargebackPaymentByExternalKeyBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type ChargebackPaymentByExternalKeyBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyBadGateway ", 502)
}

func (o *ChargebackPaymentByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyServiceUnavailable creates a ChargebackPaymentByExternalKeyServiceUnavailable with default headers values
func NewChargebackPaymentByExternalKeyServiceUnavailable() *ChargebackPaymentByExternalKeyServiceUnavailable {
	return &ChargebackPaymentByExternalKeyServiceUnavailable{}
}

/*ChargebackPaymentByExternalKeyServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ChargebackPaymentByExternalKeyServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyServiceUnavailable ", 503)
}

func (o *ChargebackPaymentByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyGatewayTimeout creates a ChargebackPaymentByExternalKeyGatewayTimeout with default headers values
func NewChargebackPaymentByExternalKeyGatewayTimeout() *ChargebackPaymentByExternalKeyGatewayTimeout {
	return &ChargebackPaymentByExternalKeyGatewayTimeout{}
}

/*ChargebackPaymentByExternalKeyGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type ChargebackPaymentByExternalKeyGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackPaymentByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyGatewayTimeout ", 504)
}

func (o *ChargebackPaymentByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
