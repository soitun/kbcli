// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveEmailReader is a Reader for the RemoveEmail structure.
type RemoveEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveEmailNoContent creates a RemoveEmailNoContent with default headers values
func NewRemoveEmailNoContent() *RemoveEmailNoContent {
	return &RemoveEmailNoContent{}
}

/*
RemoveEmailNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type RemoveEmailNoContent struct {
}

// IsSuccess returns true when this remove email no content response has a 2xx status code
func (o *RemoveEmailNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove email no content response has a 3xx status code
func (o *RemoveEmailNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove email no content response has a 4xx status code
func (o *RemoveEmailNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove email no content response has a 5xx status code
func (o *RemoveEmailNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove email no content response a status code equal to that given
func (o *RemoveEmailNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove email no content response
func (o *RemoveEmailNoContent) Code() int {
	return 204
}

func (o *RemoveEmailNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/emails/{email}][%d] removeEmailNoContent ", 204)
}

func (o *RemoveEmailNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/emails/{email}][%d] removeEmailNoContent ", 204)
}

func (o *RemoveEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveEmailBadRequest creates a RemoveEmailBadRequest with default headers values
func NewRemoveEmailBadRequest() *RemoveEmailBadRequest {
	return &RemoveEmailBadRequest{}
}

/*
RemoveEmailBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type RemoveEmailBadRequest struct {
}

// IsSuccess returns true when this remove email bad request response has a 2xx status code
func (o *RemoveEmailBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove email bad request response has a 3xx status code
func (o *RemoveEmailBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove email bad request response has a 4xx status code
func (o *RemoveEmailBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove email bad request response has a 5xx status code
func (o *RemoveEmailBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove email bad request response a status code equal to that given
func (o *RemoveEmailBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove email bad request response
func (o *RemoveEmailBadRequest) Code() int {
	return 400
}

func (o *RemoveEmailBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/emails/{email}][%d] removeEmailBadRequest ", 400)
}

func (o *RemoveEmailBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/emails/{email}][%d] removeEmailBadRequest ", 400)
}

func (o *RemoveEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
