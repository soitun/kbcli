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

// SearchBundlesReader is a Reader for the SearchBundles structure.
type SearchBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchBundlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchBundlesOK creates a SearchBundlesOK with default headers values
func NewSearchBundlesOK() *SearchBundlesOK {
	return &SearchBundlesOK{}
}

/*
SearchBundlesOK describes a response with status code 200, with default header values.

successful operation
*/
type SearchBundlesOK struct {
	Payload []*kbmodel.Bundle
}

// IsSuccess returns true when this search bundles o k response has a 2xx status code
func (o *SearchBundlesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search bundles o k response has a 3xx status code
func (o *SearchBundlesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search bundles o k response has a 4xx status code
func (o *SearchBundlesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search bundles o k response has a 5xx status code
func (o *SearchBundlesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search bundles o k response a status code equal to that given
func (o *SearchBundlesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the search bundles o k response
func (o *SearchBundlesOK) Code() int {
	return 200
}

func (o *SearchBundlesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/search/{searchKey}][%d] searchBundlesOK  %+v", 200, o.Payload)
}

func (o *SearchBundlesOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/search/{searchKey}][%d] searchBundlesOK  %+v", 200, o.Payload)
}

func (o *SearchBundlesOK) GetPayload() []*kbmodel.Bundle {
	return o.Payload
}

func (o *SearchBundlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
