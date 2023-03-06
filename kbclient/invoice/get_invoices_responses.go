// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoicesReader is a Reader for the GetInvoices structure.
type GetInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicesOK()
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

// NewGetInvoicesOK creates a GetInvoicesOK with default headers values
func NewGetInvoicesOK() *GetInvoicesOK {
	return &GetInvoicesOK{}
}

/*
GetInvoicesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoicesOK struct {
	Payload      []*kbmodel.Invoice
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoices o k response
func (o *GetInvoicesOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoices o k response has a 2xx status code
func (o *GetInvoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoices o k response has a 3xx status code
func (o *GetInvoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoices o k response has a 4xx status code
func (o *GetInvoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoices o k response has a 5xx status code
func (o *GetInvoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoices o k response a status code equal to that given
func (o *GetInvoicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/pagination][%d] getInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/pagination][%d] getInvoicesOK  %+v", 200, o.Payload)
}

func (o *GetInvoicesOK) GetPayload() []*kbmodel.Invoice {
	return o.Payload
}

func (o *GetInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
