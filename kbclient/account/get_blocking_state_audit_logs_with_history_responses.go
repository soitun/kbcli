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

// GetBlockingStateAuditLogsWithHistoryReader is a Reader for the GetBlockingStateAuditLogsWithHistory structure.
type GetBlockingStateAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlockingStateAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlockingStateAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBlockingStateAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBlockingStateAuditLogsWithHistoryOK creates a GetBlockingStateAuditLogsWithHistoryOK with default headers values
func NewGetBlockingStateAuditLogsWithHistoryOK() *GetBlockingStateAuditLogsWithHistoryOK {
	return &GetBlockingStateAuditLogsWithHistoryOK{}
}

/*
GetBlockingStateAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBlockingStateAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

// IsSuccess returns true when this get blocking state audit logs with history o k response has a 2xx status code
func (o *GetBlockingStateAuditLogsWithHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get blocking state audit logs with history o k response has a 3xx status code
func (o *GetBlockingStateAuditLogsWithHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blocking state audit logs with history o k response has a 4xx status code
func (o *GetBlockingStateAuditLogsWithHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get blocking state audit logs with history o k response has a 5xx status code
func (o *GetBlockingStateAuditLogsWithHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get blocking state audit logs with history o k response a status code equal to that given
func (o *GetBlockingStateAuditLogsWithHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get blocking state audit logs with history o k response
func (o *GetBlockingStateAuditLogsWithHistoryOK) Code() int {
	return 200
}

func (o *GetBlockingStateAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/block/{blockingId}/auditLogsWithHistory][%d] getBlockingStateAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetBlockingStateAuditLogsWithHistoryOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/block/{blockingId}/auditLogsWithHistory][%d] getBlockingStateAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetBlockingStateAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetBlockingStateAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlockingStateAuditLogsWithHistoryNotFound creates a GetBlockingStateAuditLogsWithHistoryNotFound with default headers values
func NewGetBlockingStateAuditLogsWithHistoryNotFound() *GetBlockingStateAuditLogsWithHistoryNotFound {
	return &GetBlockingStateAuditLogsWithHistoryNotFound{}
}

/*
GetBlockingStateAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Blocking state  not found
*/
type GetBlockingStateAuditLogsWithHistoryNotFound struct {
}

// IsSuccess returns true when this get blocking state audit logs with history not found response has a 2xx status code
func (o *GetBlockingStateAuditLogsWithHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get blocking state audit logs with history not found response has a 3xx status code
func (o *GetBlockingStateAuditLogsWithHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get blocking state audit logs with history not found response has a 4xx status code
func (o *GetBlockingStateAuditLogsWithHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get blocking state audit logs with history not found response has a 5xx status code
func (o *GetBlockingStateAuditLogsWithHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get blocking state audit logs with history not found response a status code equal to that given
func (o *GetBlockingStateAuditLogsWithHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get blocking state audit logs with history not found response
func (o *GetBlockingStateAuditLogsWithHistoryNotFound) Code() int {
	return 404
}

func (o *GetBlockingStateAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/block/{blockingId}/auditLogsWithHistory][%d] getBlockingStateAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetBlockingStateAuditLogsWithHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/block/{blockingId}/auditLogsWithHistory][%d] getBlockingStateAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetBlockingStateAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
