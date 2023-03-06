// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetAccountsReader is a Reader for the GetAccounts structure.
type GetAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountsOK()
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

// NewGetAccountsOK creates a GetAccountsOK with default headers values
func NewGetAccountsOK() *GetAccountsOK {
	return &GetAccountsOK{}
}

/*
GetAccountsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccountsOK struct {
	Payload      []*kbmodel.Account
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get accounts o k response
func (o *GetAccountsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get accounts o k response has a 2xx status code
func (o *GetAccountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get accounts o k response has a 3xx status code
func (o *GetAccountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get accounts o k response has a 4xx status code
func (o *GetAccountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get accounts o k response has a 5xx status code
func (o *GetAccountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get accounts o k response a status code equal to that given
func (o *GetAccountsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAccountsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/pagination][%d] getAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/pagination][%d] getAccountsOK  %+v", 200, o.Payload)
}

func (o *GetAccountsOK) GetPayload() []*kbmodel.Account {
	return o.Payload
}

func (o *GetAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
