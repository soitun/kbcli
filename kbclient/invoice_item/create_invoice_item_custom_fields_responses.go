// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// CreateInvoiceItemCustomFieldsReader is a Reader for the CreateInvoiceItemCustomFields structure.
type CreateInvoiceItemCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoiceItemCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInvoiceItemCustomFieldsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInvoiceItemCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInvoiceItemCustomFieldsCreated creates a CreateInvoiceItemCustomFieldsCreated with default headers values
func NewCreateInvoiceItemCustomFieldsCreated() *CreateInvoiceItemCustomFieldsCreated {
	return &CreateInvoiceItemCustomFieldsCreated{}
}

/*
CreateInvoiceItemCustomFieldsCreated describes a response with status code 201, with default header values.

Custom field created successfully
*/
type CreateInvoiceItemCustomFieldsCreated struct {
	Payload []*kbmodel.CustomField
}

// IsSuccess returns true when this create invoice item custom fields created response has a 2xx status code
func (o *CreateInvoiceItemCustomFieldsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create invoice item custom fields created response has a 3xx status code
func (o *CreateInvoiceItemCustomFieldsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create invoice item custom fields created response has a 4xx status code
func (o *CreateInvoiceItemCustomFieldsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create invoice item custom fields created response has a 5xx status code
func (o *CreateInvoiceItemCustomFieldsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create invoice item custom fields created response a status code equal to that given
func (o *CreateInvoiceItemCustomFieldsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create invoice item custom fields created response
func (o *CreateInvoiceItemCustomFieldsCreated) Code() int {
	return 201
}

func (o *CreateInvoiceItemCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] createInvoiceItemCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoiceItemCustomFieldsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] createInvoiceItemCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoiceItemCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreateInvoiceItemCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInvoiceItemCustomFieldsBadRequest creates a CreateInvoiceItemCustomFieldsBadRequest with default headers values
func NewCreateInvoiceItemCustomFieldsBadRequest() *CreateInvoiceItemCustomFieldsBadRequest {
	return &CreateInvoiceItemCustomFieldsBadRequest{}
}

/*
CreateInvoiceItemCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid invoice item id supplied
*/
type CreateInvoiceItemCustomFieldsBadRequest struct {
}

// IsSuccess returns true when this create invoice item custom fields bad request response has a 2xx status code
func (o *CreateInvoiceItemCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create invoice item custom fields bad request response has a 3xx status code
func (o *CreateInvoiceItemCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create invoice item custom fields bad request response has a 4xx status code
func (o *CreateInvoiceItemCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create invoice item custom fields bad request response has a 5xx status code
func (o *CreateInvoiceItemCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create invoice item custom fields bad request response a status code equal to that given
func (o *CreateInvoiceItemCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create invoice item custom fields bad request response
func (o *CreateInvoiceItemCustomFieldsBadRequest) Code() int {
	return 400
}

func (o *CreateInvoiceItemCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] createInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoiceItemCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] createInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *CreateInvoiceItemCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
