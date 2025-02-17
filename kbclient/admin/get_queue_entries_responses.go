// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// GetQueueEntriesReader is a Reader for the GetQueueEntries structure.
type GetQueueEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQueueEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetQueueEntriesOK()
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

// NewGetQueueEntriesOK creates a GetQueueEntriesOK with default headers values
func NewGetQueueEntriesOK() *GetQueueEntriesOK {
	return &GetQueueEntriesOK{}
}

/*GetQueueEntriesOK handles this case with default header values.

Success
*/
type GetQueueEntriesOK struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetQueueEntriesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesOK ", 200)
}

func (o *GetQueueEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQueueEntriesBadRequest creates a GetQueueEntriesBadRequest with default headers values
func NewGetQueueEntriesBadRequest() *GetQueueEntriesBadRequest {
	return &GetQueueEntriesBadRequest{}
}

/*GetQueueEntriesBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetQueueEntriesBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetQueueEntriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesBadRequest ", 400)
}

func (o *GetQueueEntriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQueueEntriesNotFound creates a GetQueueEntriesNotFound with default headers values
func NewGetQueueEntriesNotFound() *GetQueueEntriesNotFound {
	return &GetQueueEntriesNotFound{}
}

/*GetQueueEntriesNotFound handles this case with default header values.

Account not found
*/
type GetQueueEntriesNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetQueueEntriesNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/admin/queues][%d] getQueueEntriesNotFound ", 404)
}

func (o *GetQueueEntriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
