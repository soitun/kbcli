// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// GetUserKeyValueReader is a Reader for the GetUserKeyValue structure.
type GetUserKeyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserKeyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserKeyValueOK()
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

// NewGetUserKeyValueOK creates a GetUserKeyValueOK with default headers values
func NewGetUserKeyValueOK() *GetUserKeyValueOK {
	return &GetUserKeyValueOK{}
}

/*GetUserKeyValueOK handles this case with default header values.

successful operation
*/
type GetUserKeyValueOK struct {
	Payload *kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
}

func (o *GetUserKeyValueOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/userKeyValue/{keyName}][%d] getUserKeyValueOK  %+v", 200, o.Payload)
}

func (o *GetUserKeyValueOK) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *GetUserKeyValueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserKeyValueBadRequest creates a GetUserKeyValueBadRequest with default headers values
func NewGetUserKeyValueBadRequest() *GetUserKeyValueBadRequest {
	return &GetUserKeyValueBadRequest{}
}

/*GetUserKeyValueBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type GetUserKeyValueBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetUserKeyValueBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/userKeyValue/{keyName}][%d] getUserKeyValueBadRequest ", 400)
}

func (o *GetUserKeyValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
