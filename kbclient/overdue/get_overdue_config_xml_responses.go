// Code generated by go-swagger; DO NOT EDIT.

package overdue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v3/kbcommon"
)

// GetOverdueConfigXMLReader is a Reader for the GetOverdueConfigXML structure.
type GetOverdueConfigXMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOverdueConfigXMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOverdueConfigXMLOK()
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

// NewGetOverdueConfigXMLOK creates a GetOverdueConfigXMLOK with default headers values
func NewGetOverdueConfigXMLOK() *GetOverdueConfigXMLOK {
	return &GetOverdueConfigXMLOK{}
}

/*
GetOverdueConfigXMLOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOverdueConfigXMLOK struct {
	Payload      string
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get overdue config Xml o k response
func (o *GetOverdueConfigXMLOK) Code() int {
	return 200
}

// IsSuccess returns true when this get overdue config Xml o k response has a 2xx status code
func (o *GetOverdueConfigXMLOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get overdue config Xml o k response has a 3xx status code
func (o *GetOverdueConfigXMLOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get overdue config Xml o k response has a 4xx status code
func (o *GetOverdueConfigXMLOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get overdue config Xml o k response has a 5xx status code
func (o *GetOverdueConfigXMLOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get overdue config Xml o k response a status code equal to that given
func (o *GetOverdueConfigXMLOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOverdueConfigXMLOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/overdue/xml][%d] getOverdueConfigXmlOK  %+v", 200, o.Payload)
}

func (o *GetOverdueConfigXMLOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/overdue/xml][%d] getOverdueConfigXmlOK  %+v", 200, o.Payload)
}

func (o *GetOverdueConfigXMLOK) GetPayload() string {
	return o.Payload
}

func (o *GetOverdueConfigXMLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
