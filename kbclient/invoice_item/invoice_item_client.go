// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new invoice item API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoice item API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateInvoiceItemCustomFields(params *CreateInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoiceItemCustomFieldsCreated, error)

	CreateInvoiceItemTags(params *CreateInvoiceItemTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoiceItemTagsCreated, error)

	DeleteInvoiceItemCustomFields(params *DeleteInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoiceItemCustomFieldsNoContent, error)

	DeleteInvoiceItemTags(params *DeleteInvoiceItemTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoiceItemTagsNoContent, error)

	GetInvoiceItemAuditLogsWithHistory(params *GetInvoiceItemAuditLogsWithHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoiceItemAuditLogsWithHistoryOK, error)

	GetInvoiceItemCustomFields(params *GetInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoiceItemCustomFieldsOK, error)

	GetInvoiceItemTags(params *GetInvoiceItemTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoiceItemTagsOK, error)

	ModifyInvoiceItemCustomFields(params *ModifyInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyInvoiceItemCustomFieldsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateInvoiceItemCustomFields adds custom fields to invoice item
*/
func (a *Client) CreateInvoiceItemCustomFields(params *CreateInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoiceItemCustomFieldsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoiceItemCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInvoiceItemCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateInvoiceItemCustomFieldsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateInvoiceItemTags adds tags to invoice item
*/
func (a *Client) CreateInvoiceItemTags(params *CreateInvoiceItemTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoiceItemTagsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoiceItemTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInvoiceItemTags",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateInvoiceItemTagsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInvoiceItemTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteInvoiceItemCustomFields removes custom fields from invoice item
*/
func (a *Client) DeleteInvoiceItemCustomFields(params *DeleteInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoiceItemCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvoiceItemCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInvoiceItemCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvoiceItemCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteInvoiceItemTags removes tags from invoice item
*/
func (a *Client) DeleteInvoiceItemTags(params *DeleteInvoiceItemTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoiceItemTagsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvoiceItemTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInvoiceItemTags",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteInvoiceItemTagsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoiceItemTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoiceItemAuditLogsWithHistory retrieves invoice item audit logs with history by id
*/
func (a *Client) GetInvoiceItemAuditLogsWithHistory(params *GetInvoiceItemAuditLogsWithHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoiceItemAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceItemAuditLogsWithHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoiceItemAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoiceItemAuditLogsWithHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceItemAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoiceItemAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoiceItemCustomFields retrieves invoice item custom fields
*/
func (a *Client) GetInvoiceItemCustomFields(params *GetInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoiceItemCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceItemCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoiceItemCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceItemCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoiceItemTags retrieves invoice item tags
*/
func (a *Client) GetInvoiceItemTags(params *GetInvoiceItemTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoiceItemTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoiceItemTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoiceItemTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoiceItemTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetInvoiceItemTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoiceItemTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ModifyInvoiceItemCustomFields modifies custom fields to invoice item
*/
func (a *Client) ModifyInvoiceItemCustomFields(params *ModifyInvoiceItemCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyInvoiceItemCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyInvoiceItemCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyInvoiceItemCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/invoiceItems/{invoiceItemId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyInvoiceItemCustomFieldsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyInvoiceItemCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyInvoiceItemCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
