// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetAccountTimelineReader is a Reader for the GetAccountTimeline structure.
type GetAccountTimelineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountTimelineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountTimelineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAccountTimelineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAccountTimelineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountTimelineOK creates a GetAccountTimelineOK with default headers values
func NewGetAccountTimelineOK() *GetAccountTimelineOK {
	return &GetAccountTimelineOK{}
}

/*
GetAccountTimelineOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccountTimelineOK struct {
	Payload *kbmodel.AccountTimeline
}

// IsSuccess returns true when this get account timeline o k response has a 2xx status code
func (o *GetAccountTimelineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get account timeline o k response has a 3xx status code
func (o *GetAccountTimelineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account timeline o k response has a 4xx status code
func (o *GetAccountTimelineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get account timeline o k response has a 5xx status code
func (o *GetAccountTimelineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get account timeline o k response a status code equal to that given
func (o *GetAccountTimelineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get account timeline o k response
func (o *GetAccountTimelineOK) Code() int {
	return 200
}

func (o *GetAccountTimelineOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/timeline][%d] getAccountTimelineOK  %+v", 200, o.Payload)
}

func (o *GetAccountTimelineOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/timeline][%d] getAccountTimelineOK  %+v", 200, o.Payload)
}

func (o *GetAccountTimelineOK) GetPayload() *kbmodel.AccountTimeline {
	return o.Payload
}

func (o *GetAccountTimelineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.AccountTimeline)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountTimelineBadRequest creates a GetAccountTimelineBadRequest with default headers values
func NewGetAccountTimelineBadRequest() *GetAccountTimelineBadRequest {
	return &GetAccountTimelineBadRequest{}
}

/*
GetAccountTimelineBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetAccountTimelineBadRequest struct {
}

// IsSuccess returns true when this get account timeline bad request response has a 2xx status code
func (o *GetAccountTimelineBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account timeline bad request response has a 3xx status code
func (o *GetAccountTimelineBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account timeline bad request response has a 4xx status code
func (o *GetAccountTimelineBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get account timeline bad request response has a 5xx status code
func (o *GetAccountTimelineBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get account timeline bad request response a status code equal to that given
func (o *GetAccountTimelineBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get account timeline bad request response
func (o *GetAccountTimelineBadRequest) Code() int {
	return 400
}

func (o *GetAccountTimelineBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/timeline][%d] getAccountTimelineBadRequest ", 400)
}

func (o *GetAccountTimelineBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/timeline][%d] getAccountTimelineBadRequest ", 400)
}

func (o *GetAccountTimelineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountTimelineNotFound creates a GetAccountTimelineNotFound with default headers values
func NewGetAccountTimelineNotFound() *GetAccountTimelineNotFound {
	return &GetAccountTimelineNotFound{}
}

/*
GetAccountTimelineNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetAccountTimelineNotFound struct {
}

// IsSuccess returns true when this get account timeline not found response has a 2xx status code
func (o *GetAccountTimelineNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account timeline not found response has a 3xx status code
func (o *GetAccountTimelineNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account timeline not found response has a 4xx status code
func (o *GetAccountTimelineNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get account timeline not found response has a 5xx status code
func (o *GetAccountTimelineNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get account timeline not found response a status code equal to that given
func (o *GetAccountTimelineNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get account timeline not found response
func (o *GetAccountTimelineNotFound) Code() int {
	return 404
}

func (o *GetAccountTimelineNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/timeline][%d] getAccountTimelineNotFound ", 404)
}

func (o *GetAccountTimelineNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/timeline][%d] getAccountTimelineNotFound ", 404)
}

func (o *GetAccountTimelineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
