// Code generated by go-swagger; DO NOT EDIT.

package payment_gateway

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

// BuildFormDescriptorReader is a Reader for the BuildFormDescriptor structure.
type BuildFormDescriptorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BuildFormDescriptorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBuildFormDescriptorOK()
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

// NewBuildFormDescriptorOK creates a BuildFormDescriptorOK with default headers values
func NewBuildFormDescriptorOK() *BuildFormDescriptorOK {
	return &BuildFormDescriptorOK{}
}

/*BuildFormDescriptorOK handles this case with default header values.

successful operation
*/
type BuildFormDescriptorOK struct {
	Payload *kbmodel.HostedPaymentPageFormDescriptor

	HttpResponse runtime.ClientResponse
}

func (o *BuildFormDescriptorOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentGateways/hosted/form/{accountId}][%d] buildFormDescriptorOK  %+v", 200, o.Payload)
}

func (o *BuildFormDescriptorOK) GetPayload() *kbmodel.HostedPaymentPageFormDescriptor {
	return o.Payload
}

func (o *BuildFormDescriptorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.HostedPaymentPageFormDescriptor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBuildFormDescriptorNotFound creates a BuildFormDescriptorNotFound with default headers values
func NewBuildFormDescriptorNotFound() *BuildFormDescriptorNotFound {
	return &BuildFormDescriptorNotFound{}
}

/*BuildFormDescriptorNotFound handles this case with default header values.

Account not found
*/
type BuildFormDescriptorNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *BuildFormDescriptorNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentGateways/hosted/form/{accountId}][%d] buildFormDescriptorNotFound ", 404)
}

func (o *BuildFormDescriptorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
