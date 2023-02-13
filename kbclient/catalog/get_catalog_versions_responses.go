// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetCatalogVersionsReader is a Reader for the GetCatalogVersions structure.
type GetCatalogVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCatalogVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCatalogVersionsOK creates a GetCatalogVersionsOK with default headers values
func NewGetCatalogVersionsOK() *GetCatalogVersionsOK {
	return &GetCatalogVersionsOK{}
}

/*
GetCatalogVersionsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCatalogVersionsOK struct {
	Payload []strfmt.DateTime
}

// IsSuccess returns true when this get catalog versions o k response has a 2xx status code
func (o *GetCatalogVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get catalog versions o k response has a 3xx status code
func (o *GetCatalogVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get catalog versions o k response has a 4xx status code
func (o *GetCatalogVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get catalog versions o k response has a 5xx status code
func (o *GetCatalogVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get catalog versions o k response a status code equal to that given
func (o *GetCatalogVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get catalog versions o k response
func (o *GetCatalogVersionsOK) Code() int {
	return 200
}

func (o *GetCatalogVersionsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/versions][%d] getCatalogVersionsOK  %+v", 200, o.Payload)
}

func (o *GetCatalogVersionsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/versions][%d] getCatalogVersionsOK  %+v", 200, o.Payload)
}

func (o *GetCatalogVersionsOK) GetPayload() []strfmt.DateTime {
	return o.Payload
}

func (o *GetCatalogVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
