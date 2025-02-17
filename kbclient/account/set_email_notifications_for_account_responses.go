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
)

// SetEmailNotificationsForAccountReader is a Reader for the SetEmailNotificationsForAccount structure.
type SetEmailNotificationsForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetEmailNotificationsForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSetEmailNotificationsForAccountNoContent()
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

// NewSetEmailNotificationsForAccountNoContent creates a SetEmailNotificationsForAccountNoContent with default headers values
func NewSetEmailNotificationsForAccountNoContent() *SetEmailNotificationsForAccountNoContent {
	return &SetEmailNotificationsForAccountNoContent{}
}

/*SetEmailNotificationsForAccountNoContent handles this case with default header values.

Successful operation
*/
type SetEmailNotificationsForAccountNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *SetEmailNotificationsForAccountNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/emailNotifications][%d] setEmailNotificationsForAccountNoContent ", 204)
}

func (o *SetEmailNotificationsForAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetEmailNotificationsForAccountBadRequest creates a SetEmailNotificationsForAccountBadRequest with default headers values
func NewSetEmailNotificationsForAccountBadRequest() *SetEmailNotificationsForAccountBadRequest {
	return &SetEmailNotificationsForAccountBadRequest{}
}

/*SetEmailNotificationsForAccountBadRequest handles this case with default header values.

Invalid account id supplied
*/
type SetEmailNotificationsForAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *SetEmailNotificationsForAccountBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/emailNotifications][%d] setEmailNotificationsForAccountBadRequest ", 400)
}

func (o *SetEmailNotificationsForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSetEmailNotificationsForAccountNotFound creates a SetEmailNotificationsForAccountNotFound with default headers values
func NewSetEmailNotificationsForAccountNotFound() *SetEmailNotificationsForAccountNotFound {
	return &SetEmailNotificationsForAccountNotFound{}
}

/*SetEmailNotificationsForAccountNotFound handles this case with default header values.

Account not found
*/
type SetEmailNotificationsForAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *SetEmailNotificationsForAccountNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/emailNotifications][%d] setEmailNotificationsForAccountNotFound ", 404)
}

func (o *SetEmailNotificationsForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
