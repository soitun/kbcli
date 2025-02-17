// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// GetPushNotificationCallbacksReader is a Reader for the GetPushNotificationCallbacks structure.
type GetPushNotificationCallbacksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPushNotificationCallbacksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPushNotificationCallbacksOK()
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

// NewGetPushNotificationCallbacksOK creates a GetPushNotificationCallbacksOK with default headers values
func NewGetPushNotificationCallbacksOK() *GetPushNotificationCallbacksOK {
	return &GetPushNotificationCallbacksOK{}
}

/*GetPushNotificationCallbacksOK handles this case with default header values.

successful operation
*/
type GetPushNotificationCallbacksOK struct {
	Payload *kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
}

func (o *GetPushNotificationCallbacksOK) Error() string {
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

/*GetPushNotificationCallbacksBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type GetPushNotificationCallbacksBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPushNotificationCallbacksBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/registerNotificationCallback][%d] getPushNotificationCallbacksBadRequest ", 400)
}

func (o *GetPushNotificationCallbacksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
