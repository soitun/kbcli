// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// TransferBundleReader is a Reader for the TransferBundle structure.
type TransferBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewTransferBundleCreated()
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

// NewTransferBundleCreated creates a TransferBundleCreated with default headers values
func NewTransferBundleCreated() *TransferBundleCreated {
	return &TransferBundleCreated{}
}

/*TransferBundleCreated handles this case with default header values.

Bundle transferred successfully
*/
type TransferBundleCreated struct {
	Payload *kbmodel.Bundle

	HttpResponse runtime.ClientResponse
}

func (o *TransferBundleCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}][%d] transferBundleCreated  %+v", 201, o.Payload)
}

func (o *TransferBundleCreated) GetPayload() *kbmodel.Bundle {
	return o.Payload
}

func (o *TransferBundleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Bundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransferBundleBadRequest creates a TransferBundleBadRequest with default headers values
func NewTransferBundleBadRequest() *TransferBundleBadRequest {
	return &TransferBundleBadRequest{}
}

/*TransferBundleBadRequest handles this case with default header values.

Invalid bundle id, requested date or policy supplied
*/
type TransferBundleBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *TransferBundleBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}][%d] transferBundleBadRequest ", 400)
}

func (o *TransferBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferBundleNotFound creates a TransferBundleNotFound with default headers values
func NewTransferBundleNotFound() *TransferBundleNotFound {
	return &TransferBundleNotFound{}
}

/*TransferBundleNotFound handles this case with default header values.

Bundle not found
*/
type TransferBundleNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *TransferBundleNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}][%d] transferBundleNotFound ", 404)
}

func (o *TransferBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
