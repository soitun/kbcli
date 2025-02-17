// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// GetPaymentTagsReader is a Reader for the GetPaymentTags structure.
type GetPaymentTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentTagsOK()
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

// NewGetPaymentTagsOK creates a GetPaymentTagsOK with default headers values
func NewGetPaymentTagsOK() *GetPaymentTagsOK {
	return &GetPaymentTagsOK{}
}

/*GetPaymentTagsOK handles this case with default header values.

successful operation
*/
type GetPaymentTagsOK struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/tags][%d] getPaymentTagsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentTagsOK) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *GetPaymentTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentTagsBadRequest creates a GetPaymentTagsBadRequest with default headers values
func NewGetPaymentTagsBadRequest() *GetPaymentTagsBadRequest {
	return &GetPaymentTagsBadRequest{}
}

/*GetPaymentTagsBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type GetPaymentTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/tags][%d] getPaymentTagsBadRequest ", 400)
}

func (o *GetPaymentTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentTagsNotFound creates a GetPaymentTagsNotFound with default headers values
func NewGetPaymentTagsNotFound() *GetPaymentTagsNotFound {
	return &GetPaymentTagsNotFound{}
}

/*GetPaymentTagsNotFound handles this case with default header values.

Invoice not found
*/
type GetPaymentTagsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}/tags][%d] getPaymentTagsNotFound ", 404)
}

func (o *GetPaymentTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
