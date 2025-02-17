// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

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

// CreateTagDefinitionReader is a Reader for the CreateTagDefinition structure.
type CreateTagDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTagDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateTagDefinitionCreated()
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

// NewCreateTagDefinitionCreated creates a CreateTagDefinitionCreated with default headers values
func NewCreateTagDefinitionCreated() *CreateTagDefinitionCreated {
	return &CreateTagDefinitionCreated{}
}

/*CreateTagDefinitionCreated handles this case with default header values.

Tag definition created successfully
*/
type CreateTagDefinitionCreated struct {
	Payload *kbmodel.TagDefinition

	HttpResponse runtime.ClientResponse
}

func (o *CreateTagDefinitionCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tagDefinitions][%d] createTagDefinitionCreated  %+v", 201, o.Payload)
}

func (o *CreateTagDefinitionCreated) GetPayload() *kbmodel.TagDefinition {
	return o.Payload
}

func (o *CreateTagDefinitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TagDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTagDefinitionBadRequest creates a CreateTagDefinitionBadRequest with default headers values
func NewCreateTagDefinitionBadRequest() *CreateTagDefinitionBadRequest {
	return &CreateTagDefinitionBadRequest{}
}

/*CreateTagDefinitionBadRequest handles this case with default header values.

Invalid name or description supplied
*/
type CreateTagDefinitionBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateTagDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tagDefinitions][%d] createTagDefinitionBadRequest ", 400)
}

func (o *CreateTagDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
