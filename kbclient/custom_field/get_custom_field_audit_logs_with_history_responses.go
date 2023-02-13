// Code generated by go-swagger; DO NOT EDIT.

package custom_field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetCustomFieldAuditLogsWithHistoryReader is a Reader for the GetCustomFieldAuditLogsWithHistory structure.
type GetCustomFieldAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomFieldAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomFieldAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetCustomFieldAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomFieldAuditLogsWithHistoryOK creates a GetCustomFieldAuditLogsWithHistoryOK with default headers values
func NewGetCustomFieldAuditLogsWithHistoryOK() *GetCustomFieldAuditLogsWithHistoryOK {
	return &GetCustomFieldAuditLogsWithHistoryOK{}
}

/*
GetCustomFieldAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCustomFieldAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

// IsSuccess returns true when this get custom field audit logs with history o k response has a 2xx status code
func (o *GetCustomFieldAuditLogsWithHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get custom field audit logs with history o k response has a 3xx status code
func (o *GetCustomFieldAuditLogsWithHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom field audit logs with history o k response has a 4xx status code
func (o *GetCustomFieldAuditLogsWithHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get custom field audit logs with history o k response has a 5xx status code
func (o *GetCustomFieldAuditLogsWithHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom field audit logs with history o k response a status code equal to that given
func (o *GetCustomFieldAuditLogsWithHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get custom field audit logs with history o k response
func (o *GetCustomFieldAuditLogsWithHistoryOK) Code() int {
	return 200
}

func (o *GetCustomFieldAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/{customFieldId}/auditLogsWithHistory][%d] getCustomFieldAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetCustomFieldAuditLogsWithHistoryOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/{customFieldId}/auditLogsWithHistory][%d] getCustomFieldAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetCustomFieldAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetCustomFieldAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCustomFieldAuditLogsWithHistoryNotFound creates a GetCustomFieldAuditLogsWithHistoryNotFound with default headers values
func NewGetCustomFieldAuditLogsWithHistoryNotFound() *GetCustomFieldAuditLogsWithHistoryNotFound {
	return &GetCustomFieldAuditLogsWithHistoryNotFound{}
}

/*
GetCustomFieldAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetCustomFieldAuditLogsWithHistoryNotFound struct {
}

// IsSuccess returns true when this get custom field audit logs with history not found response has a 2xx status code
func (o *GetCustomFieldAuditLogsWithHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get custom field audit logs with history not found response has a 3xx status code
func (o *GetCustomFieldAuditLogsWithHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get custom field audit logs with history not found response has a 4xx status code
func (o *GetCustomFieldAuditLogsWithHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get custom field audit logs with history not found response has a 5xx status code
func (o *GetCustomFieldAuditLogsWithHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get custom field audit logs with history not found response a status code equal to that given
func (o *GetCustomFieldAuditLogsWithHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get custom field audit logs with history not found response
func (o *GetCustomFieldAuditLogsWithHistoryNotFound) Code() int {
	return 404
}

func (o *GetCustomFieldAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/{customFieldId}/auditLogsWithHistory][%d] getCustomFieldAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetCustomFieldAuditLogsWithHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/{customFieldId}/auditLogsWithHistory][%d] getCustomFieldAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetCustomFieldAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
