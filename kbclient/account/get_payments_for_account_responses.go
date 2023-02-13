// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetPaymentsForAccountReader is a Reader for the GetPaymentsForAccount structure.
type GetPaymentsForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentsForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentsForAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPaymentsForAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPaymentsForAccountOK creates a GetPaymentsForAccountOK with default headers values
func NewGetPaymentsForAccountOK() *GetPaymentsForAccountOK {
	return &GetPaymentsForAccountOK{}
}

/*
GetPaymentsForAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPaymentsForAccountOK struct {
	Payload []*kbmodel.Payment
}

// IsSuccess returns true when this get payments for account o k response has a 2xx status code
func (o *GetPaymentsForAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payments for account o k response has a 3xx status code
func (o *GetPaymentsForAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payments for account o k response has a 4xx status code
func (o *GetPaymentsForAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payments for account o k response has a 5xx status code
func (o *GetPaymentsForAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payments for account o k response a status code equal to that given
func (o *GetPaymentsForAccountOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payments for account o k response
func (o *GetPaymentsForAccountOK) Code() int {
	return 200
}

func (o *GetPaymentsForAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/payments][%d] getPaymentsForAccountOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsForAccountOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/payments][%d] getPaymentsForAccountOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsForAccountOK) GetPayload() []*kbmodel.Payment {
	return o.Payload
}

func (o *GetPaymentsForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsForAccountBadRequest creates a GetPaymentsForAccountBadRequest with default headers values
func NewGetPaymentsForAccountBadRequest() *GetPaymentsForAccountBadRequest {
	return &GetPaymentsForAccountBadRequest{}
}

/*
GetPaymentsForAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetPaymentsForAccountBadRequest struct {
}

// IsSuccess returns true when this get payments for account bad request response has a 2xx status code
func (o *GetPaymentsForAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payments for account bad request response has a 3xx status code
func (o *GetPaymentsForAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payments for account bad request response has a 4xx status code
func (o *GetPaymentsForAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payments for account bad request response has a 5xx status code
func (o *GetPaymentsForAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get payments for account bad request response a status code equal to that given
func (o *GetPaymentsForAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get payments for account bad request response
func (o *GetPaymentsForAccountBadRequest) Code() int {
	return 400
}

func (o *GetPaymentsForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/payments][%d] getPaymentsForAccountBadRequest ", 400)
}

func (o *GetPaymentsForAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/payments][%d] getPaymentsForAccountBadRequest ", 400)
}

func (o *GetPaymentsForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
