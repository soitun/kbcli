// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new invoice payment API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoice payment API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CompleteInvoicePaymentTransaction(params *CompleteInvoicePaymentTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CompleteInvoicePaymentTransactionNoContent, error)

	CreateChargeback(params *CreateChargebackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChargebackCreated, error)

	CreateChargebackReversal(params *CreateChargebackReversalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChargebackReversalCreated, error)

	CreateInvoicePaymentCustomFields(params *CreateInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoicePaymentCustomFieldsCreated, error)

	CreateInvoicePaymentTags(params *CreateInvoicePaymentTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoicePaymentTagsCreated, error)

	CreateRefundWithAdjustments(params *CreateRefundWithAdjustmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRefundWithAdjustmentsCreated, error)

	DeleteInvoicePaymentCustomFields(params *DeleteInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoicePaymentCustomFieldsNoContent, error)

	DeleteInvoicePaymentTags(params *DeleteInvoicePaymentTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoicePaymentTagsNoContent, error)

	GetInvoicePayment(params *GetInvoicePaymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentOK, error)

	GetInvoicePaymentAuditLogsWithHistory(params *GetInvoicePaymentAuditLogsWithHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentAuditLogsWithHistoryOK, error)

	GetInvoicePaymentCustomFields(params *GetInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentCustomFieldsOK, error)

	GetInvoicePaymentTags(params *GetInvoicePaymentTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentTagsOK, error)

	ModifyInvoicePaymentCustomFields(params *ModifyInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyInvoicePaymentCustomFieldsNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CompleteInvoicePaymentTransaction completes an existing transaction
*/
func (a *Client) CompleteInvoicePaymentTransaction(params *CompleteInvoicePaymentTransactionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CompleteInvoicePaymentTransactionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCompleteInvoicePaymentTransactionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "completeInvoicePaymentTransaction",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CompleteInvoicePaymentTransactionReader{formats: a.formats},
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
	success, ok := result.(*CompleteInvoicePaymentTransactionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for completeInvoicePaymentTransaction: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateChargeback records a chargeback
*/
func (a *Client) CreateChargeback(params *CreateChargebackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChargebackCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateChargebackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createChargeback",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/chargebacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChargebackReader{formats: a.formats},
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
	success, ok := result.(*CreateChargebackCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createChargeback: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateChargebackReversal records a chargeback reversal
*/
func (a *Client) CreateChargebackReversal(params *CreateChargebackReversalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateChargebackReversalCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateChargebackReversalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createChargebackReversal",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/chargebackReversals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateChargebackReversalReader{formats: a.formats},
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
	success, ok := result.(*CreateChargebackReversalCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createChargebackReversal: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateInvoicePaymentCustomFields adds custom fields to payment
*/
func (a *Client) CreateInvoicePaymentCustomFields(params *CreateInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoicePaymentCustomFieldsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoicePaymentCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInvoicePaymentCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoicePaymentCustomFieldsReader{formats: a.formats},
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
	success, ok := result.(*CreateInvoicePaymentCustomFieldsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateInvoicePaymentTags adds tags to payment
*/
func (a *Client) CreateInvoicePaymentTags(params *CreateInvoicePaymentTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateInvoicePaymentTagsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateInvoicePaymentTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createInvoicePaymentTags",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateInvoicePaymentTagsReader{formats: a.formats},
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
	success, ok := result.(*CreateInvoicePaymentTagsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createInvoicePaymentTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateRefundWithAdjustments refunds a payment and adjust the invoice if needed
*/
func (a *Client) CreateRefundWithAdjustments(params *CreateRefundWithAdjustmentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateRefundWithAdjustmentsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRefundWithAdjustmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createRefundWithAdjustments",
		Method:             "POST",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/refunds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRefundWithAdjustmentsReader{formats: a.formats},
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
	success, ok := result.(*CreateRefundWithAdjustmentsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createRefundWithAdjustments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteInvoicePaymentCustomFields removes custom fields from payment
*/
func (a *Client) DeleteInvoicePaymentCustomFields(params *DeleteInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoicePaymentCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvoicePaymentCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInvoicePaymentCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoicePaymentCustomFieldsReader{formats: a.formats},
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
	success, ok := result.(*DeleteInvoicePaymentCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteInvoicePaymentTags removes tags from payment
*/
func (a *Client) DeleteInvoicePaymentTags(params *DeleteInvoicePaymentTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteInvoicePaymentTagsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteInvoicePaymentTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteInvoicePaymentTags",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteInvoicePaymentTagsReader{formats: a.formats},
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
	success, ok := result.(*DeleteInvoicePaymentTagsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteInvoicePaymentTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoicePayment retrieves a payment by id
*/
func (a *Client) GetInvoicePayment(params *GetInvoicePaymentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoicePaymentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoicePayment",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentReader{formats: a.formats},
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
	success, ok := result.(*GetInvoicePaymentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePayment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoicePaymentAuditLogsWithHistory retrieves invoice payment audit logs with history by id
*/
func (a *Client) GetInvoicePaymentAuditLogsWithHistory(params *GetInvoicePaymentAuditLogsWithHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentAuditLogsWithHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoicePaymentAuditLogsWithHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoicePaymentAuditLogsWithHistory",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{invoicePaymentId}/auditLogsWithHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentAuditLogsWithHistoryReader{formats: a.formats},
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
	success, ok := result.(*GetInvoicePaymentAuditLogsWithHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePaymentAuditLogsWithHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoicePaymentCustomFields retrieves payment custom fields
*/
func (a *Client) GetInvoicePaymentCustomFields(params *GetInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoicePaymentCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoicePaymentCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentCustomFieldsReader{formats: a.formats},
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
	success, ok := result.(*GetInvoicePaymentCustomFieldsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetInvoicePaymentTags retrieves payment tags
*/
func (a *Client) GetInvoicePaymentTags(params *GetInvoicePaymentTagsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetInvoicePaymentTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetInvoicePaymentTagsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getInvoicePaymentTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/tags",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetInvoicePaymentTagsReader{formats: a.formats},
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
	success, ok := result.(*GetInvoicePaymentTagsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getInvoicePaymentTags: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ModifyInvoicePaymentCustomFields modifies custom fields to payment
*/
func (a *Client) ModifyInvoicePaymentCustomFields(params *ModifyInvoicePaymentCustomFieldsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ModifyInvoicePaymentCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyInvoicePaymentCustomFieldsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyInvoicePaymentCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/invoicePayments/{paymentId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyInvoicePaymentCustomFieldsReader{formats: a.formats},
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
	success, ok := result.(*ModifyInvoicePaymentCustomFieldsNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyInvoicePaymentCustomFields: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
