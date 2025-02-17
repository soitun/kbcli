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

// AddAccountBlockingStateReader is a Reader for the AddAccountBlockingState structure.
type AddAccountBlockingStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAccountBlockingStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewAddAccountBlockingStateCreated()
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

// NewAddAccountBlockingStateCreated creates a AddAccountBlockingStateCreated with default headers values
func NewAddAccountBlockingStateCreated() *AddAccountBlockingStateCreated {
	return &AddAccountBlockingStateCreated{}
}

/*AddAccountBlockingStateCreated handles this case with default header values.

Blocking state created successfully
*/
type AddAccountBlockingStateCreated struct {
	Payload []*kbmodel.BlockingState

	HttpResponse runtime.ClientResponse
}

func (o *AddAccountBlockingStateCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateCreated  %+v", 201, o.Payload)
}

func (o *AddAccountBlockingStateCreated) GetPayload() []*kbmodel.BlockingState {
	return o.Payload
}

func (o *AddAccountBlockingStateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAccountBlockingStateBadRequest creates a AddAccountBlockingStateBadRequest with default headers values
func NewAddAccountBlockingStateBadRequest() *AddAccountBlockingStateBadRequest {
	return &AddAccountBlockingStateBadRequest{}
}

/*AddAccountBlockingStateBadRequest handles this case with default header values.

Invalid account id supplied
*/
type AddAccountBlockingStateBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *AddAccountBlockingStateBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateBadRequest ", 400)
}

func (o *AddAccountBlockingStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAccountBlockingStateNotFound creates a AddAccountBlockingStateNotFound with default headers values
func NewAddAccountBlockingStateNotFound() *AddAccountBlockingStateNotFound {
	return &AddAccountBlockingStateNotFound{}
}

/*AddAccountBlockingStateNotFound handles this case with default header values.

Account not found
*/
type AddAccountBlockingStateNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *AddAccountBlockingStateNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateNotFound ", 404)
}

func (o *AddAccountBlockingStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
