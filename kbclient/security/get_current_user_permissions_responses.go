// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetCurrentUserPermissionsReader is a Reader for the GetCurrentUserPermissions structure.
type GetCurrentUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentUserPermissionsOK creates a GetCurrentUserPermissionsOK with default headers values
func NewGetCurrentUserPermissionsOK() *GetCurrentUserPermissionsOK {
	return &GetCurrentUserPermissionsOK{}
}

/*
GetCurrentUserPermissionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCurrentUserPermissionsOK struct {
	Payload []string
}

// IsSuccess returns true when this get current user permissions o k response has a 2xx status code
func (o *GetCurrentUserPermissionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get current user permissions o k response has a 3xx status code
func (o *GetCurrentUserPermissionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get current user permissions o k response has a 4xx status code
func (o *GetCurrentUserPermissionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get current user permissions o k response has a 5xx status code
func (o *GetCurrentUserPermissionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get current user permissions o k response a status code equal to that given
func (o *GetCurrentUserPermissionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get current user permissions o k response
func (o *GetCurrentUserPermissionsOK) Code() int {
	return 200
}

func (o *GetCurrentUserPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/security/permissions][%d] getCurrentUserPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserPermissionsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/security/permissions][%d] getCurrentUserPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetCurrentUserPermissionsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetCurrentUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
