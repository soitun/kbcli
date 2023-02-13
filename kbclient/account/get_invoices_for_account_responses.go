// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetInvoicesForAccountReader is a Reader for the GetInvoicesForAccount structure.
type GetInvoicesForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicesForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoicesForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInvoicesForAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInvoicesForAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoicesForAccountOK creates a GetInvoicesForAccountOK with default headers values
func NewGetInvoicesForAccountOK() *GetInvoicesForAccountOK {
	return &GetInvoicesForAccountOK{}
}

/*
GetInvoicesForAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoicesForAccountOK struct {
	Payload []*kbmodel.Invoice
}

// IsSuccess returns true when this get invoices for account o k response has a 2xx status code
func (o *GetInvoicesForAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoices for account o k response has a 3xx status code
func (o *GetInvoicesForAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices for account o k response has a 4xx status code
func (o *GetInvoicesForAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices for account o k response has a 5xx status code
func (o *GetInvoicesForAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices for account o k response a status code equal to that given
func (o *GetInvoicesForAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get invoices for account o k response
func (o *GetInvoicesForAccountOK) Code() int {
	return 200
}

func (o *GetInvoicesForAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesForAccountOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesForAccountOK) GetPayload() []*kbmodel.Invoice {
	return o.Payload
}

func (o *GetInvoicesForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicesForAccountBadRequest creates a GetInvoicesForAccountBadRequest with default headers values
func NewGetInvoicesForAccountBadRequest() *GetInvoicesForAccountBadRequest {
	return &GetInvoicesForAccountBadRequest{}
}

/*
GetInvoicesForAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetInvoicesForAccountBadRequest struct {
}

// IsSuccess returns true when this get invoices for account bad request response has a 2xx status code
func (o *GetInvoicesForAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices for account bad request response has a 3xx status code
func (o *GetInvoicesForAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices for account bad request response has a 4xx status code
func (o *GetInvoicesForAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices for account bad request response has a 5xx status code
func (o *GetInvoicesForAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices for account bad request response a status code equal to that given
func (o *GetInvoicesForAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get invoices for account bad request response
func (o *GetInvoicesForAccountBadRequest) Code() int {
	return 400
}

func (o *GetInvoicesForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountBadRequest ", 400)
}

func (o *GetInvoicesForAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountBadRequest ", 400)
}

func (o *GetInvoicesForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInvoicesForAccountNotFound creates a GetInvoicesForAccountNotFound with default headers values
func NewGetInvoicesForAccountNotFound() *GetInvoicesForAccountNotFound {
	return &GetInvoicesForAccountNotFound{}
}

/*
GetInvoicesForAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetInvoicesForAccountNotFound struct {
}

// IsSuccess returns true when this get invoices for account not found response has a 2xx status code
func (o *GetInvoicesForAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoices for account not found response has a 3xx status code
func (o *GetInvoicesForAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices for account not found response has a 4xx status code
func (o *GetInvoicesForAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoices for account not found response has a 5xx status code
func (o *GetInvoicesForAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices for account not found response a status code equal to that given
func (o *GetInvoicesForAccountNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get invoices for account not found response
func (o *GetInvoicesForAccountNotFound) Code() int {
	return 404
}

func (o *GetInvoicesForAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountNotFound ", 404)
}

func (o *GetInvoicesForAccountNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountNotFound ", 404)
}

func (o *GetInvoicesForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
