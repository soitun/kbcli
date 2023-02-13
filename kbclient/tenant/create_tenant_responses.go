// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// CreateTenantReader is a Reader for the CreateTenant structure.
type CreateTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTenantCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewCreateTenantConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTenantCreated creates a CreateTenantCreated with default headers values
func NewCreateTenantCreated() *CreateTenantCreated {
	return &CreateTenantCreated{}
}

/*
CreateTenantCreated describes a response with status code 201, with default header values.

Tenant created successfully
*/
type CreateTenantCreated struct {
	Payload *kbmodel.Tenant
}

// IsSuccess returns true when this create tenant created response has a 2xx status code
func (o *CreateTenantCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create tenant created response has a 3xx status code
func (o *CreateTenantCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tenant created response has a 4xx status code
func (o *CreateTenantCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create tenant created response has a 5xx status code
func (o *CreateTenantCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create tenant created response a status code equal to that given
func (o *CreateTenantCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create tenant created response
func (o *CreateTenantCreated) Code() int {
	return 201
}

func (o *CreateTenantCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants][%d] createTenantCreated  %+v", 201, o.Payload)
}

func (o *CreateTenantCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants][%d] createTenantCreated  %+v", 201, o.Payload)
}

func (o *CreateTenantCreated) GetPayload() *kbmodel.Tenant {
	return o.Payload
}

func (o *CreateTenantCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTenantConflict creates a CreateTenantConflict with default headers values
func NewCreateTenantConflict() *CreateTenantConflict {
	return &CreateTenantConflict{}
}

/*
CreateTenantConflict describes a response with status code 409, with default header values.

Tenant already exists
*/
type CreateTenantConflict struct {
}

// IsSuccess returns true when this create tenant conflict response has a 2xx status code
func (o *CreateTenantConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create tenant conflict response has a 3xx status code
func (o *CreateTenantConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create tenant conflict response has a 4xx status code
func (o *CreateTenantConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create tenant conflict response has a 5xx status code
func (o *CreateTenantConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create tenant conflict response a status code equal to that given
func (o *CreateTenantConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create tenant conflict response
func (o *CreateTenantConflict) Code() int {
	return 409
}

func (o *CreateTenantConflict) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants][%d] createTenantConflict ", 409)
}

func (o *CreateTenantConflict) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants][%d] createTenantConflict ", 409)
}

func (o *CreateTenantConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
