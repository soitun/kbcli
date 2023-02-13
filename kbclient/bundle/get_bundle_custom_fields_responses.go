// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetBundleCustomFieldsReader is a Reader for the GetBundleCustomFields structure.
type GetBundleCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundleCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBundleCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBundleCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBundleCustomFieldsOK creates a GetBundleCustomFieldsOK with default headers values
func NewGetBundleCustomFieldsOK() *GetBundleCustomFieldsOK {
	return &GetBundleCustomFieldsOK{}
}

/*
GetBundleCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBundleCustomFieldsOK struct {
	Payload []*kbmodel.CustomField
}

// IsSuccess returns true when this get bundle custom fields o k response has a 2xx status code
func (o *GetBundleCustomFieldsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get bundle custom fields o k response has a 3xx status code
func (o *GetBundleCustomFieldsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundle custom fields o k response has a 4xx status code
func (o *GetBundleCustomFieldsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get bundle custom fields o k response has a 5xx status code
func (o *GetBundleCustomFieldsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get bundle custom fields o k response a status code equal to that given
func (o *GetBundleCustomFieldsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get bundle custom fields o k response
func (o *GetBundleCustomFieldsOK) Code() int {
	return 200
}

func (o *GetBundleCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/customFields][%d] getBundleCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetBundleCustomFieldsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/customFields][%d] getBundleCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetBundleCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetBundleCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleCustomFieldsBadRequest creates a GetBundleCustomFieldsBadRequest with default headers values
func NewGetBundleCustomFieldsBadRequest() *GetBundleCustomFieldsBadRequest {
	return &GetBundleCustomFieldsBadRequest{}
}

/*
GetBundleCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid bundle id supplied
*/
type GetBundleCustomFieldsBadRequest struct {
}

// IsSuccess returns true when this get bundle custom fields bad request response has a 2xx status code
func (o *GetBundleCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get bundle custom fields bad request response has a 3xx status code
func (o *GetBundleCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get bundle custom fields bad request response has a 4xx status code
func (o *GetBundleCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get bundle custom fields bad request response has a 5xx status code
func (o *GetBundleCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get bundle custom fields bad request response a status code equal to that given
func (o *GetBundleCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get bundle custom fields bad request response
func (o *GetBundleCustomFieldsBadRequest) Code() int {
	return 400
}

func (o *GetBundleCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/customFields][%d] getBundleCustomFieldsBadRequest ", 400)
}

func (o *GetBundleCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/customFields][%d] getBundleCustomFieldsBadRequest ", 400)
}

func (o *GetBundleCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
