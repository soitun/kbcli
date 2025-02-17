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

// GetAllPluginConfigurationReader is a Reader for the GetAllPluginConfiguration structure.
type GetAllPluginConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllPluginConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllPluginConfigurationOK()
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

// NewGetAllPluginConfigurationOK creates a GetAllPluginConfigurationOK with default headers values
func NewGetAllPluginConfigurationOK() *GetAllPluginConfigurationOK {
	return &GetAllPluginConfigurationOK{}
}

/*GetAllPluginConfigurationOK handles this case with default header values.

successful operation
*/
type GetAllPluginConfigurationOK struct {
	Payload []*kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
}

func (o *GetAllPluginConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPerTenantConfig/{keyPrefix}/search][%d] getAllPluginConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetAllPluginConfigurationOK) GetPayload() []*kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *GetAllPluginConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllPluginConfigurationBadRequest creates a GetAllPluginConfigurationBadRequest with default headers values
func NewGetAllPluginConfigurationBadRequest() *GetAllPluginConfigurationBadRequest {
	return &GetAllPluginConfigurationBadRequest{}
}

/*GetAllPluginConfigurationBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type GetAllPluginConfigurationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAllPluginConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPerTenantConfig/{keyPrefix}/search][%d] getAllPluginConfigurationBadRequest ", 400)
}

func (o *GetAllPluginConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
