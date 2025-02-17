// Code generated by go-swagger; DO NOT EDIT.

package usage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// RecordUsageReader is a Reader for the RecordUsage structure.
type RecordUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RecordUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRecordUsageOK()
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

// NewRecordUsageOK creates a RecordUsageOK with default headers values
func NewRecordUsageOK() *RecordUsageOK {
	return &RecordUsageOK{}
}

/*RecordUsageOK handles this case with default header values.

Successfully recorded usage data change
*/
type RecordUsageOK struct {
	HttpResponse runtime.ClientResponse
}

func (o *RecordUsageOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/usages][%d] recordUsageOK ", 200)
}

func (o *RecordUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRecordUsageBadRequest creates a RecordUsageBadRequest with default headers values
func NewRecordUsageBadRequest() *RecordUsageBadRequest {
	return &RecordUsageBadRequest{}
}

/*RecordUsageBadRequest handles this case with default header values.

Invalid subscription (e.g. inactive)
*/
type RecordUsageBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *RecordUsageBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/usages][%d] recordUsageBadRequest ", 400)
}

func (o *RecordUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
