// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// CreateTaxItemsReader is a Reader for the CreateTaxItems structure.
type CreateTaxItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTaxItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTaxItemsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTaxItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateTaxItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTaxItemsCreated creates a CreateTaxItemsCreated with default headers values
func NewCreateTaxItemsCreated() *CreateTaxItemsCreated {
	return &CreateTaxItemsCreated{}
}

/*
CreateTaxItemsCreated describes a response with status code 201, with default header values.

Create tax items successfully
*/
type CreateTaxItemsCreated struct {
	Payload []*kbmodel.InvoiceItem
}

// IsSuccess returns true when this create tax items created response has a 2xx status code
func (o *CreateTaxItemsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tax items created response has a 3xx status code
func (o *CreateTaxItemsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tax items created response has a 4xx status code
func (o *CreateTaxItemsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tax items created response has a 5xx status code
func (o *CreateTaxItemsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create tax items created response a status code equal to that given
func (o *CreateTaxItemsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create tax items created response
func (o *CreateTaxItemsCreated) Code() int {
	return 201
}

func (o *CreateTaxItemsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsCreated  %+v", 201, o.Payload)
}

func (o *CreateTaxItemsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsCreated  %+v", 201, o.Payload)
}

func (o *CreateTaxItemsCreated) GetPayload() []*kbmodel.InvoiceItem {
	return o.Payload
}

func (o *CreateTaxItemsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTaxItemsBadRequest creates a CreateTaxItemsBadRequest with default headers values
func NewCreateTaxItemsBadRequest() *CreateTaxItemsBadRequest {
	return &CreateTaxItemsBadRequest{}
}

/*
CreateTaxItemsBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type CreateTaxItemsBadRequest struct {
}

// IsSuccess returns true when this create tax items bad request response has a 2xx status code
func (o *CreateTaxItemsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tax items bad request response has a 3xx status code
func (o *CreateTaxItemsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tax items bad request response has a 4xx status code
func (o *CreateTaxItemsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tax items bad request response has a 5xx status code
func (o *CreateTaxItemsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create tax items bad request response a status code equal to that given
func (o *CreateTaxItemsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create tax items bad request response
func (o *CreateTaxItemsBadRequest) Code() int {
	return 400
}

func (o *CreateTaxItemsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsBadRequest ", 400)
}

func (o *CreateTaxItemsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsBadRequest ", 400)
}

func (o *CreateTaxItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTaxItemsNotFound creates a CreateTaxItemsNotFound with default headers values
func NewCreateTaxItemsNotFound() *CreateTaxItemsNotFound {
	return &CreateTaxItemsNotFound{}
}

/*
CreateTaxItemsNotFound describes a response with status code 404, with default header values.

Account not found
*/
type CreateTaxItemsNotFound struct {
}

// IsSuccess returns true when this create tax items not found response has a 2xx status code
func (o *CreateTaxItemsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tax items not found response has a 3xx status code
func (o *CreateTaxItemsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tax items not found response has a 4xx status code
func (o *CreateTaxItemsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tax items not found response has a 5xx status code
func (o *CreateTaxItemsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create tax items not found response a status code equal to that given
func (o *CreateTaxItemsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create tax items not found response
func (o *CreateTaxItemsNotFound) Code() int {
	return 404
}

func (o *CreateTaxItemsNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsNotFound ", 404)
}

func (o *CreateTaxItemsNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/taxes/{accountId}][%d] createTaxItemsNotFound ", 404)
}

func (o *CreateTaxItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
