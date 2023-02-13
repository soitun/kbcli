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

// CreateRefundWithAdjustmentsReader is a Reader for the CreateRefundWithAdjustments structure.
type CreateRefundWithAdjustmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRefundWithAdjustmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRefundWithAdjustmentsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRefundWithAdjustmentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRefundWithAdjustmentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRefundWithAdjustmentsCreated creates a CreateRefundWithAdjustmentsCreated with default headers values
func NewCreateRefundWithAdjustmentsCreated() *CreateRefundWithAdjustmentsCreated {
	return &CreateRefundWithAdjustmentsCreated{}
}

/*
CreateRefundWithAdjustmentsCreated describes a response with status code 201, with default header values.

Created refund successfully
*/
type CreateRefundWithAdjustmentsCreated struct {
	Payload *kbmodel.InvoicePayment
}

// IsSuccess returns true when this create refund with adjustments created response has a 2xx status code
func (o *CreateRefundWithAdjustmentsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create refund with adjustments created response has a 3xx status code
func (o *CreateRefundWithAdjustmentsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create refund with adjustments created response has a 4xx status code
func (o *CreateRefundWithAdjustmentsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create refund with adjustments created response has a 5xx status code
func (o *CreateRefundWithAdjustmentsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create refund with adjustments created response a status code equal to that given
func (o *CreateRefundWithAdjustmentsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create refund with adjustments created response
func (o *CreateRefundWithAdjustmentsCreated) Code() int {
	return 201
}

func (o *CreateRefundWithAdjustmentsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsCreated  %+v", 201, o.Payload)
}

func (o *CreateRefundWithAdjustmentsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsCreated  %+v", 201, o.Payload)
}

func (o *CreateRefundWithAdjustmentsCreated) GetPayload() *kbmodel.InvoicePayment {
	return o.Payload
}

func (o *CreateRefundWithAdjustmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoicePayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRefundWithAdjustmentsBadRequest creates a CreateRefundWithAdjustmentsBadRequest with default headers values
func NewCreateRefundWithAdjustmentsBadRequest() *CreateRefundWithAdjustmentsBadRequest {
	return &CreateRefundWithAdjustmentsBadRequest{}
}

/*
CreateRefundWithAdjustmentsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type CreateRefundWithAdjustmentsBadRequest struct {
}

// IsSuccess returns true when this create refund with adjustments bad request response has a 2xx status code
func (o *CreateRefundWithAdjustmentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create refund with adjustments bad request response has a 3xx status code
func (o *CreateRefundWithAdjustmentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create refund with adjustments bad request response has a 4xx status code
func (o *CreateRefundWithAdjustmentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create refund with adjustments bad request response has a 5xx status code
func (o *CreateRefundWithAdjustmentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create refund with adjustments bad request response a status code equal to that given
func (o *CreateRefundWithAdjustmentsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create refund with adjustments bad request response
func (o *CreateRefundWithAdjustmentsBadRequest) Code() int {
	return 400
}

func (o *CreateRefundWithAdjustmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsBadRequest ", 400)
}

func (o *CreateRefundWithAdjustmentsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsBadRequest ", 400)
}

func (o *CreateRefundWithAdjustmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRefundWithAdjustmentsNotFound creates a CreateRefundWithAdjustmentsNotFound with default headers values
func NewCreateRefundWithAdjustmentsNotFound() *CreateRefundWithAdjustmentsNotFound {
	return &CreateRefundWithAdjustmentsNotFound{}
}

/*
CreateRefundWithAdjustmentsNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type CreateRefundWithAdjustmentsNotFound struct {
}

// IsSuccess returns true when this create refund with adjustments not found response has a 2xx status code
func (o *CreateRefundWithAdjustmentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create refund with adjustments not found response has a 3xx status code
func (o *CreateRefundWithAdjustmentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create refund with adjustments not found response has a 4xx status code
func (o *CreateRefundWithAdjustmentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create refund with adjustments not found response has a 5xx status code
func (o *CreateRefundWithAdjustmentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create refund with adjustments not found response a status code equal to that given
func (o *CreateRefundWithAdjustmentsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create refund with adjustments not found response
func (o *CreateRefundWithAdjustmentsNotFound) Code() int {
	return 404
}

func (o *CreateRefundWithAdjustmentsNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsNotFound ", 404)
}

func (o *CreateRefundWithAdjustmentsNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsNotFound ", 404)
}

func (o *CreateRefundWithAdjustmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
