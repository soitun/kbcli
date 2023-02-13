// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// SearchInvoicesReader is a Reader for the SearchInvoices structure.
type SearchInvoicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchInvoicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchInvoicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchInvoicesOK creates a SearchInvoicesOK with default headers values
func NewSearchInvoicesOK() *SearchInvoicesOK {
	return &SearchInvoicesOK{}
}

/*
SearchInvoicesOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchInvoicesOK struct {
	Payload []*kbmodel.Invoice
}

// IsSuccess returns true when this search invoices o k response has a 2xx status code
func (o *SearchInvoicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search invoices o k response has a 3xx status code
func (o *SearchInvoicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search invoices o k response has a 4xx status code
func (o *SearchInvoicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search invoices o k response has a 5xx status code
func (o *SearchInvoicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search invoices o k response a status code equal to that given
func (o *SearchInvoicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search invoices o k response
func (o *SearchInvoicesOK) Code() int {
	return 200
}

func (o *SearchInvoicesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/search/{searchKey}][%d] searchInvoicesOK  %+v", 200, o.Payload)
}

func (o *SearchInvoicesOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/search/{searchKey}][%d] searchInvoicesOK  %+v", 200, o.Payload)
}

func (o *SearchInvoicesOK) GetPayload() []*kbmodel.Invoice {
	return o.Payload
}

func (o *SearchInvoicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
