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

// UploadPerTenantConfigurationReader is a Reader for the UploadPerTenantConfiguration structure.
type UploadPerTenantConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadPerTenantConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUploadPerTenantConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadPerTenantConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadPerTenantConfigurationCreated creates a UploadPerTenantConfigurationCreated with default headers values
func NewUploadPerTenantConfigurationCreated() *UploadPerTenantConfigurationCreated {
	return &UploadPerTenantConfigurationCreated{}
}

/*
UploadPerTenantConfigurationCreated describes a response with status code 201, with default header values.

Per tenant configuration uploaded successfully
*/
type UploadPerTenantConfigurationCreated struct {
	Payload *kbmodel.TenantKeyValue
}

// IsSuccess returns true when this upload per tenant configuration created response has a 2xx status code
func (o *UploadPerTenantConfigurationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload per tenant configuration created response has a 3xx status code
func (o *UploadPerTenantConfigurationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload per tenant configuration created response has a 4xx status code
func (o *UploadPerTenantConfigurationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload per tenant configuration created response has a 5xx status code
func (o *UploadPerTenantConfigurationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this upload per tenant configuration created response a status code equal to that given
func (o *UploadPerTenantConfigurationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the upload per tenant configuration created response
func (o *UploadPerTenantConfigurationCreated) Code() int {
	return 201
}

func (o *UploadPerTenantConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPerTenantConfig][%d] uploadPerTenantConfigurationCreated  %+v", 201, o.Payload)
}

func (o *UploadPerTenantConfigurationCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPerTenantConfig][%d] uploadPerTenantConfigurationCreated  %+v", 201, o.Payload)
}

func (o *UploadPerTenantConfigurationCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *UploadPerTenantConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPerTenantConfigurationBadRequest creates a UploadPerTenantConfigurationBadRequest with default headers values
func NewUploadPerTenantConfigurationBadRequest() *UploadPerTenantConfigurationBadRequest {
	return &UploadPerTenantConfigurationBadRequest{}
}

/*
UploadPerTenantConfigurationBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type UploadPerTenantConfigurationBadRequest struct {
}

// IsSuccess returns true when this upload per tenant configuration bad request response has a 2xx status code
func (o *UploadPerTenantConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload per tenant configuration bad request response has a 3xx status code
func (o *UploadPerTenantConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload per tenant configuration bad request response has a 4xx status code
func (o *UploadPerTenantConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload per tenant configuration bad request response has a 5xx status code
func (o *UploadPerTenantConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload per tenant configuration bad request response a status code equal to that given
func (o *UploadPerTenantConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upload per tenant configuration bad request response
func (o *UploadPerTenantConfigurationBadRequest) Code() int {
	return 400
}

func (o *UploadPerTenantConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPerTenantConfig][%d] uploadPerTenantConfigurationBadRequest ", 400)
}

func (o *UploadPerTenantConfigurationBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPerTenantConfig][%d] uploadPerTenantConfigurationBadRequest ", 400)
}

func (o *UploadPerTenantConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
