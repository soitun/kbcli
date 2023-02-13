// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CreateSubscriptionTagsReader is a Reader for the CreateSubscriptionTags structure.
type CreateSubscriptionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSubscriptionTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateSubscriptionTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSubscriptionTagsCreated creates a CreateSubscriptionTagsCreated with default headers values
func NewCreateSubscriptionTagsCreated() *CreateSubscriptionTagsCreated {
	return &CreateSubscriptionTagsCreated{}
}

/*
CreateSubscriptionTagsCreated describes a response with status code 201, with default header values.

Tag created successfully
*/
type CreateSubscriptionTagsCreated struct {
}

// IsSuccess returns true when this create subscription tags created response has a 2xx status code
func (o *CreateSubscriptionTagsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create subscription tags created response has a 3xx status code
func (o *CreateSubscriptionTagsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create subscription tags created response has a 4xx status code
func (o *CreateSubscriptionTagsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create subscription tags created response has a 5xx status code
func (o *CreateSubscriptionTagsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create subscription tags created response a status code equal to that given
func (o *CreateSubscriptionTagsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create subscription tags created response
func (o *CreateSubscriptionTagsCreated) Code() int {
	return 201
}

func (o *CreateSubscriptionTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/tags][%d] createSubscriptionTagsCreated ", 201)
}

func (o *CreateSubscriptionTagsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/tags][%d] createSubscriptionTagsCreated ", 201)
}

func (o *CreateSubscriptionTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateSubscriptionTagsBadRequest creates a CreateSubscriptionTagsBadRequest with default headers values
func NewCreateSubscriptionTagsBadRequest() *CreateSubscriptionTagsBadRequest {
	return &CreateSubscriptionTagsBadRequest{}
}

/*
CreateSubscriptionTagsBadRequest describes a response with status code 400, with default header values.

Invalid subscription id supplied
*/
type CreateSubscriptionTagsBadRequest struct {
}

// IsSuccess returns true when this create subscription tags bad request response has a 2xx status code
func (o *CreateSubscriptionTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create subscription tags bad request response has a 3xx status code
func (o *CreateSubscriptionTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create subscription tags bad request response has a 4xx status code
func (o *CreateSubscriptionTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create subscription tags bad request response has a 5xx status code
func (o *CreateSubscriptionTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create subscription tags bad request response a status code equal to that given
func (o *CreateSubscriptionTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create subscription tags bad request response
func (o *CreateSubscriptionTagsBadRequest) Code() int {
	return 400
}

func (o *CreateSubscriptionTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/tags][%d] createSubscriptionTagsBadRequest ", 400)
}

func (o *CreateSubscriptionTagsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/tags][%d] createSubscriptionTagsBadRequest ", 400)
}

func (o *CreateSubscriptionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
