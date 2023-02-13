// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetPushNotificationCallbacksReader is a Reader for the GetPushNotificationCallbacks structure.
type GetPushNotificationCallbacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPushNotificationCallbacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPushNotificationCallbacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPushNotificationCallbacksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPushNotificationCallbacksOK creates a GetPushNotificationCallbacksOK with default headers values
func NewGetPushNotificationCallbacksOK() *GetPushNotificationCallbacksOK {
	return &GetPushNotificationCallbacksOK{}
}

/*
GetPushNotificationCallbacksOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPushNotificationCallbacksOK struct {
	Payload *kbmodel.TenantKeyValue
}

// IsSuccess returns true when this get push notification callbacks o k response has a 2xx status code
func (o *GetPushNotificationCallbacksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get push notification callbacks o k response has a 3xx status code
func (o *GetPushNotificationCallbacksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get push notification callbacks o k response has a 4xx status code
func (o *GetPushNotificationCallbacksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get push notification callbacks o k response has a 5xx status code
func (o *GetPushNotificationCallbacksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get push notification callbacks o k response a status code equal to that given
func (o *GetPushNotificationCallbacksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get push notification callbacks o k response
func (o *GetPushNotificationCallbacksOK) Code() int {
	return 200
}

func (o *GetPushNotificationCallbacksOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/registerNotificationCallback][%d] getPushNotificationCallbacksOK  %+v", 200, o.Payload)
}

func (o *GetPushNotificationCallbacksOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/registerNotificationCallback][%d] getPushNotificationCallbacksOK  %+v", 200, o.Payload)
}

func (o *GetPushNotificationCallbacksOK) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *GetPushNotificationCallbacksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPushNotificationCallbacksBadRequest creates a GetPushNotificationCallbacksBadRequest with default headers values
func NewGetPushNotificationCallbacksBadRequest() *GetPushNotificationCallbacksBadRequest {
	return &GetPushNotificationCallbacksBadRequest{}
}

/*
GetPushNotificationCallbacksBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type GetPushNotificationCallbacksBadRequest struct {
}

// IsSuccess returns true when this get push notification callbacks bad request response has a 2xx status code
func (o *GetPushNotificationCallbacksBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get push notification callbacks bad request response has a 3xx status code
func (o *GetPushNotificationCallbacksBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get push notification callbacks bad request response has a 4xx status code
func (o *GetPushNotificationCallbacksBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get push notification callbacks bad request response has a 5xx status code
func (o *GetPushNotificationCallbacksBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get push notification callbacks bad request response a status code equal to that given
func (o *GetPushNotificationCallbacksBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get push notification callbacks bad request response
func (o *GetPushNotificationCallbacksBadRequest) Code() int {
	return 400
}

func (o *GetPushNotificationCallbacksBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/registerNotificationCallback][%d] getPushNotificationCallbacksBadRequest ", 400)
}

func (o *GetPushNotificationCallbacksBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/registerNotificationCallback][%d] getPushNotificationCallbacksBadRequest ", 400)
}

func (o *GetPushNotificationCallbacksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
