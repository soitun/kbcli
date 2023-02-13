// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePerTenantConfigurationReader is a Reader for the DeletePerTenantConfiguration structure.
type DeletePerTenantConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePerTenantConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePerTenantConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePerTenantConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePerTenantConfigurationNoContent creates a DeletePerTenantConfigurationNoContent with default headers values
func NewDeletePerTenantConfigurationNoContent() *DeletePerTenantConfigurationNoContent {
	return &DeletePerTenantConfigurationNoContent{}
}

/*
DeletePerTenantConfigurationNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeletePerTenantConfigurationNoContent struct {
}

// IsSuccess returns true when this delete per tenant configuration no content response has a 2xx status code
func (o *DeletePerTenantConfigurationNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete per tenant configuration no content response has a 3xx status code
func (o *DeletePerTenantConfigurationNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete per tenant configuration no content response has a 4xx status code
func (o *DeletePerTenantConfigurationNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete per tenant configuration no content response has a 5xx status code
func (o *DeletePerTenantConfigurationNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete per tenant configuration no content response a status code equal to that given
func (o *DeletePerTenantConfigurationNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete per tenant configuration no content response
func (o *DeletePerTenantConfigurationNoContent) Code() int {
	return 204
}

func (o *DeletePerTenantConfigurationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPerTenantConfig][%d] deletePerTenantConfigurationNoContent ", 204)
}

func (o *DeletePerTenantConfigurationNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPerTenantConfig][%d] deletePerTenantConfigurationNoContent ", 204)
}

func (o *DeletePerTenantConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePerTenantConfigurationBadRequest creates a DeletePerTenantConfigurationBadRequest with default headers values
func NewDeletePerTenantConfigurationBadRequest() *DeletePerTenantConfigurationBadRequest {
	return &DeletePerTenantConfigurationBadRequest{}
}

/*
DeletePerTenantConfigurationBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type DeletePerTenantConfigurationBadRequest struct {
}

// IsSuccess returns true when this delete per tenant configuration bad request response has a 2xx status code
func (o *DeletePerTenantConfigurationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete per tenant configuration bad request response has a 3xx status code
func (o *DeletePerTenantConfigurationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete per tenant configuration bad request response has a 4xx status code
func (o *DeletePerTenantConfigurationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete per tenant configuration bad request response has a 5xx status code
func (o *DeletePerTenantConfigurationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete per tenant configuration bad request response a status code equal to that given
func (o *DeletePerTenantConfigurationBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete per tenant configuration bad request response
func (o *DeletePerTenantConfigurationBadRequest) Code() int {
	return 400
}

func (o *DeletePerTenantConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPerTenantConfig][%d] deletePerTenantConfigurationBadRequest ", 400)
}

func (o *DeletePerTenantConfigurationBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPerTenantConfig][%d] deletePerTenantConfigurationBadRequest ", 400)
}

func (o *DeletePerTenantConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
