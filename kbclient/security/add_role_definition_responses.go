// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// AddRoleDefinitionReader is a Reader for the AddRoleDefinition structure.
type AddRoleDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddRoleDefinitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddRoleDefinitionCreated creates a AddRoleDefinitionCreated with default headers values
func NewAddRoleDefinitionCreated() *AddRoleDefinitionCreated {
	return &AddRoleDefinitionCreated{}
}

/*
AddRoleDefinitionCreated describes a response with status code 201, with default header values.

Role definition created successfully
*/
type AddRoleDefinitionCreated struct {
	Payload *kbmodel.RoleDefinition
}

// IsSuccess returns true when this add role definition created response has a 2xx status code
func (o *AddRoleDefinitionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add role definition created response has a 3xx status code
func (o *AddRoleDefinitionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add role definition created response has a 4xx status code
func (o *AddRoleDefinitionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add role definition created response has a 5xx status code
func (o *AddRoleDefinitionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add role definition created response a status code equal to that given
func (o *AddRoleDefinitionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add role definition created response
func (o *AddRoleDefinitionCreated) Code() int {
	return 201
}

func (o *AddRoleDefinitionCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/security/roles][%d] addRoleDefinitionCreated  %+v", 201, o.Payload)
}

func (o *AddRoleDefinitionCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/security/roles][%d] addRoleDefinitionCreated  %+v", 201, o.Payload)
}

func (o *AddRoleDefinitionCreated) GetPayload() *kbmodel.RoleDefinition {
	return o.Payload
}

func (o *AddRoleDefinitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.RoleDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
