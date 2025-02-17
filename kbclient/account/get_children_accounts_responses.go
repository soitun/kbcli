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

// GetChildrenAccountsReader is a Reader for the GetChildrenAccounts structure.
type GetChildrenAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetChildrenAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetChildrenAccountsOK()
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

// NewGetChildrenAccountsOK creates a GetChildrenAccountsOK with default headers values
func NewGetChildrenAccountsOK() *GetChildrenAccountsOK {
	return &GetChildrenAccountsOK{}
}

/*GetChildrenAccountsOK handles this case with default header values.

successful operation
*/
type GetChildrenAccountsOK struct {
	Payload []*kbmodel.Account

	HttpResponse runtime.ClientResponse
}

func (o *GetChildrenAccountsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/children][%d] getChildrenAccountsOK  %+v", 200, o.Payload)
}

func (o *GetChildrenAccountsOK) GetPayload() []*kbmodel.Account {
	return o.Payload
}

func (o *GetChildrenAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetChildrenAccountsBadRequest creates a GetChildrenAccountsBadRequest with default headers values
func NewGetChildrenAccountsBadRequest() *GetChildrenAccountsBadRequest {
	return &GetChildrenAccountsBadRequest{}
}

/*GetChildrenAccountsBadRequest handles this case with default header values.

Invalid parent account id supplied
*/
type GetChildrenAccountsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetChildrenAccountsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/children][%d] getChildrenAccountsBadRequest ", 400)
}

func (o *GetChildrenAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetChildrenAccountsNotFound creates a GetChildrenAccountsNotFound with default headers values
func NewGetChildrenAccountsNotFound() *GetChildrenAccountsNotFound {
	return &GetChildrenAccountsNotFound{}
}

/*GetChildrenAccountsNotFound handles this case with default header values.

Parent Account not found
*/
type GetChildrenAccountsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetChildrenAccountsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/children][%d] getChildrenAccountsNotFound ", 404)
}

func (o *GetChildrenAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
