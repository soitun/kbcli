// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// GetAvailableAddonsReader is a Reader for the GetAvailableAddons structure.
type GetAvailableAddonsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAvailableAddonsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAvailableAddonsOK()
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

// NewGetAvailableAddonsOK creates a GetAvailableAddonsOK with default headers values
func NewGetAvailableAddonsOK() *GetAvailableAddonsOK {
	return &GetAvailableAddonsOK{}
}

/*GetAvailableAddonsOK handles this case with default header values.

successful operation
*/
type GetAvailableAddonsOK struct {
	Payload []*kbmodel.PlanDetail

	HttpResponse runtime.ClientResponse
}

func (o *GetAvailableAddonsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/availableAddons][%d] getAvailableAddonsOK  %+v", 200, o.Payload)
}

func (o *GetAvailableAddonsOK) GetPayload() []*kbmodel.PlanDetail {
	return o.Payload
}

func (o *GetAvailableAddonsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
