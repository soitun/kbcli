// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// GetPaymentCustomFieldsReader is a Reader for the GetPaymentCustomFields structure.
type GetPaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentCustomFieldsOK()
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

// NewGetPaymentCustomFieldsOK creates a GetPaymentCustomFieldsOK with default headers values
func NewGetPaymentCustomFieldsOK() *GetPaymentCustomFieldsOK {
	return &GetPaymentCustomFieldsOK{}
}

/*
GetPaymentCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPaymentCustomFieldsOK struct {
	Payload      []*kbmodel.CustomField
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment custom fields o k response
func (o *GetPaymentCustomFieldsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get payment custom fields o k response has a 2xx status code
func (o *GetPaymentCustomFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment custom fields o k response has a 3xx status code
func (o *GetPaymentCustomFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment custom fields o k response has a 4xx status code
func (o *GetPaymentCustomFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment custom fields o k response has a 5xx status code
func (o *GetPaymentCustomFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment custom fields o k response a status code equal to that given
func (o *GetPaymentCustomFieldsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPaymentCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/customFields][%d] getPaymentCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentCustomFieldsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/customFields][%d] getPaymentCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetPaymentCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentCustomFieldsBadRequest creates a GetPaymentCustomFieldsBadRequest with default headers values
func NewGetPaymentCustomFieldsBadRequest() *GetPaymentCustomFieldsBadRequest {
	return &GetPaymentCustomFieldsBadRequest{}
}

/*
GetPaymentCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type GetPaymentCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment custom fields bad request response
func (o *GetPaymentCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get payment custom fields bad request response has a 2xx status code
func (o *GetPaymentCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payment custom fields bad request response has a 3xx status code
func (o *GetPaymentCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment custom fields bad request response has a 4xx status code
func (o *GetPaymentCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payment custom fields bad request response has a 5xx status code
func (o *GetPaymentCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment custom fields bad request response a status code equal to that given
func (o *GetPaymentCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/customFields][%d] getPaymentCustomFieldsBadRequest ", 400)
}

func (o *GetPaymentCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/customFields][%d] getPaymentCustomFieldsBadRequest ", 400)
}

func (o *GetPaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
