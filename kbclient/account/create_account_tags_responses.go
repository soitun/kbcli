// Code generated by go-swagger; DO NOT EDIT.

package account

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

// CreateAccountTagsReader is a Reader for the CreateAccountTags structure.
type CreateAccountTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateAccountTagsCreated()
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

// NewCreateAccountTagsCreated creates a CreateAccountTagsCreated with default headers values
func NewCreateAccountTagsCreated() *CreateAccountTagsCreated {
	return &CreateAccountTagsCreated{}
}

/*CreateAccountTagsCreated handles this case with default header values.

Tag created successfully
*/
type CreateAccountTagsCreated struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *CreateAccountTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/tags][%d] createAccountTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateAccountTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreateAccountTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountTagsBadRequest creates a CreateAccountTagsBadRequest with default headers values
func NewCreateAccountTagsBadRequest() *CreateAccountTagsBadRequest {
	return &CreateAccountTagsBadRequest{}
}

/*CreateAccountTagsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type CreateAccountTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateAccountTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/tags][%d] createAccountTagsBadRequest ", 400)
}

func (o *CreateAccountTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
