// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// SetClockReader is a Reader for the SetClock structure.
type SetClockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetClockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetClockOK()
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

// NewSetClockOK creates a SetClockOK with default headers values
func NewSetClockOK() *SetClockOK {
	return &SetClockOK{}
}

/*SetClockOK handles this case with default header values.

successful operation
*/
type SetClockOK struct {
	Payload *kbmodel.Clock

	HttpResponse runtime.ClientResponse
}

func (o *SetClockOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/test/clock][%d] setClockOK  %+v", 200, o.Payload)
}

func (o *SetClockOK) GetPayload() *kbmodel.Clock {
	return o.Payload
}

func (o *SetClockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Clock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
