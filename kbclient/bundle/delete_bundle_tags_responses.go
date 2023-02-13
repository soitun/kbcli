// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBundleTagsReader is a Reader for the DeleteBundleTags structure.
type DeleteBundleTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBundleTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteBundleTagsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteBundleTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteBundleTagsNoContent creates a DeleteBundleTagsNoContent with default headers values
func NewDeleteBundleTagsNoContent() *DeleteBundleTagsNoContent {
	return &DeleteBundleTagsNoContent{}
}

/*
DeleteBundleTagsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteBundleTagsNoContent struct {
}

// IsSuccess returns true when this delete bundle tags no content response has a 2xx status code
func (o *DeleteBundleTagsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete bundle tags no content response has a 3xx status code
func (o *DeleteBundleTagsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete bundle tags no content response has a 4xx status code
func (o *DeleteBundleTagsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete bundle tags no content response has a 5xx status code
func (o *DeleteBundleTagsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete bundle tags no content response a status code equal to that given
func (o *DeleteBundleTagsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete bundle tags no content response
func (o *DeleteBundleTagsNoContent) Code() int {
	return 204
}

func (o *DeleteBundleTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/bundles/{bundleId}/tags][%d] deleteBundleTagsNoContent ", 204)
}

func (o *DeleteBundleTagsNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/bundles/{bundleId}/tags][%d] deleteBundleTagsNoContent ", 204)
}

func (o *DeleteBundleTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBundleTagsBadRequest creates a DeleteBundleTagsBadRequest with default headers values
func NewDeleteBundleTagsBadRequest() *DeleteBundleTagsBadRequest {
	return &DeleteBundleTagsBadRequest{}
}

/*
DeleteBundleTagsBadRequest describes a response with status code 400, with default header values.

Invalid bundle id supplied
*/
type DeleteBundleTagsBadRequest struct {
}

// IsSuccess returns true when this delete bundle tags bad request response has a 2xx status code
func (o *DeleteBundleTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete bundle tags bad request response has a 3xx status code
func (o *DeleteBundleTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete bundle tags bad request response has a 4xx status code
func (o *DeleteBundleTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete bundle tags bad request response has a 5xx status code
func (o *DeleteBundleTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete bundle tags bad request response a status code equal to that given
func (o *DeleteBundleTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete bundle tags bad request response
func (o *DeleteBundleTagsBadRequest) Code() int {
	return 400
}

func (o *DeleteBundleTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/bundles/{bundleId}/tags][%d] deleteBundleTagsBadRequest ", 400)
}

func (o *DeleteBundleTagsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/bundles/{bundleId}/tags][%d] deleteBundleTagsBadRequest ", 400)
}

func (o *DeleteBundleTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
