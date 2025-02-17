// Code generated by go-swagger; DO NOT EDIT.

package export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ExportDataForAccountReader is a Reader for the ExportDataForAccount structure.
type ExportDataForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExportDataForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExportDataForAccountOK()
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

// NewExportDataForAccountOK creates a ExportDataForAccountOK with default headers values
func NewExportDataForAccountOK() *ExportDataForAccountOK {
	return &ExportDataForAccountOK{}
}

/*ExportDataForAccountOK handles this case with default header values.

Success
*/
type ExportDataForAccountOK struct {
	HttpResponse runtime.ClientResponse
}

func (o *ExportDataForAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountOK ", 200)
}

func (o *ExportDataForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportDataForAccountBadRequest creates a ExportDataForAccountBadRequest with default headers values
func NewExportDataForAccountBadRequest() *ExportDataForAccountBadRequest {
	return &ExportDataForAccountBadRequest{}
}

/*ExportDataForAccountBadRequest handles this case with default header values.

Invalid account id supplied
*/
type ExportDataForAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ExportDataForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountBadRequest ", 400)
}

func (o *ExportDataForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExportDataForAccountNotFound creates a ExportDataForAccountNotFound with default headers values
func NewExportDataForAccountNotFound() *ExportDataForAccountNotFound {
	return &ExportDataForAccountNotFound{}
}

/*ExportDataForAccountNotFound handles this case with default header values.

Account not found
*/
type ExportDataForAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ExportDataForAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/export/{accountId}][%d] exportDataForAccountNotFound ", 404)
}

func (o *ExportDataForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
