// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// CreateChargebackReversalReader is a Reader for the CreateChargebackReversal structure.
type CreateChargebackReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateChargebackReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateChargebackReversalCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateChargebackReversalBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateChargebackReversalNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateChargebackReversalCreated creates a CreateChargebackReversalCreated with default headers values
func NewCreateChargebackReversalCreated() *CreateChargebackReversalCreated {
	return &CreateChargebackReversalCreated{}
}

/*
CreateChargebackReversalCreated describes a response with status code 201, with default header values.

Created chargeback reversal successfully
*/
type CreateChargebackReversalCreated struct {
	Payload *kbmodel.InvoicePayment
}

// IsSuccess returns true when this create chargeback reversal created response has a 2xx status code
func (o *CreateChargebackReversalCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create chargeback reversal created response has a 3xx status code
func (o *CreateChargebackReversalCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chargeback reversal created response has a 4xx status code
func (o *CreateChargebackReversalCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create chargeback reversal created response has a 5xx status code
func (o *CreateChargebackReversalCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create chargeback reversal created response a status code equal to that given
func (o *CreateChargebackReversalCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create chargeback reversal created response
func (o *CreateChargebackReversalCreated) Code() int {
	return 201
}

func (o *CreateChargebackReversalCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalCreated  %+v", 201, o.Payload)
}

func (o *CreateChargebackReversalCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalCreated  %+v", 201, o.Payload)
}

func (o *CreateChargebackReversalCreated) GetPayload() *kbmodel.InvoicePayment {
	return o.Payload
}

func (o *CreateChargebackReversalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoicePayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChargebackReversalBadRequest creates a CreateChargebackReversalBadRequest with default headers values
func NewCreateChargebackReversalBadRequest() *CreateChargebackReversalBadRequest {
	return &CreateChargebackReversalBadRequest{}
}

/*
CreateChargebackReversalBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type CreateChargebackReversalBadRequest struct {
}

// IsSuccess returns true when this create chargeback reversal bad request response has a 2xx status code
func (o *CreateChargebackReversalBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create chargeback reversal bad request response has a 3xx status code
func (o *CreateChargebackReversalBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chargeback reversal bad request response has a 4xx status code
func (o *CreateChargebackReversalBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create chargeback reversal bad request response has a 5xx status code
func (o *CreateChargebackReversalBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create chargeback reversal bad request response a status code equal to that given
func (o *CreateChargebackReversalBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create chargeback reversal bad request response
func (o *CreateChargebackReversalBadRequest) Code() int {
	return 400
}

func (o *CreateChargebackReversalBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalBadRequest ", 400)
}

func (o *CreateChargebackReversalBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalBadRequest ", 400)
}

func (o *CreateChargebackReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateChargebackReversalNotFound creates a CreateChargebackReversalNotFound with default headers values
func NewCreateChargebackReversalNotFound() *CreateChargebackReversalNotFound {
	return &CreateChargebackReversalNotFound{}
}

/*
CreateChargebackReversalNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type CreateChargebackReversalNotFound struct {
}

// IsSuccess returns true when this create chargeback reversal not found response has a 2xx status code
func (o *CreateChargebackReversalNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create chargeback reversal not found response has a 3xx status code
func (o *CreateChargebackReversalNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chargeback reversal not found response has a 4xx status code
func (o *CreateChargebackReversalNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create chargeback reversal not found response has a 5xx status code
func (o *CreateChargebackReversalNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create chargeback reversal not found response a status code equal to that given
func (o *CreateChargebackReversalNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create chargeback reversal not found response
func (o *CreateChargebackReversalNotFound) Code() int {
	return 404
}

func (o *CreateChargebackReversalNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalNotFound ", 404)
}

func (o *CreateChargebackReversalNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalNotFound ", 404)
}

func (o *CreateChargebackReversalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
