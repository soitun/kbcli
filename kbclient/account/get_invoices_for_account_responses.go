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

// GetInvoicesForAccountReader is a Reader for the GetInvoicesForAccount structure.
type GetInvoicesForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicesForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicesForAccountOK()
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

// NewGetInvoicesForAccountOK creates a GetInvoicesForAccountOK with default headers values
func NewGetInvoicesForAccountOK() *GetInvoicesForAccountOK {
	return &GetInvoicesForAccountOK{}
}

/*GetInvoicesForAccountOK handles this case with default header values.

successful operation
*/
type GetInvoicesForAccountOK struct {
	Payload []*kbmodel.Invoice

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoicesForAccountOK) Error() string {
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

/*GetInvoicesForAccountBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetInvoicesForAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoicesForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountBadRequest ", 400)
}

func (o *GetInvoicesForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInvoicesForAccountNotFound creates a GetInvoicesForAccountNotFound with default headers values
func NewGetInvoicesForAccountNotFound() *GetInvoicesForAccountNotFound {
	return &GetInvoicesForAccountNotFound{}
}

/*GetInvoicesForAccountNotFound handles this case with default header values.

Account not found
*/
type GetInvoicesForAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoicesForAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/invoices][%d] getInvoicesForAccountNotFound ", 404)
}

func (o *GetInvoicesForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
