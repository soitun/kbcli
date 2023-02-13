// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetTransactionTagsParams creates a new GetTransactionTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTransactionTagsParams() *GetTransactionTagsParams {
	return &GetTransactionTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionTagsParamsWithTimeout creates a new GetTransactionTagsParams object
// with the ability to set a timeout on a request.
func NewGetTransactionTagsParamsWithTimeout(timeout time.Duration) *GetTransactionTagsParams {
	return &GetTransactionTagsParams{
		timeout: timeout,
	}
}

// NewGetTransactionTagsParamsWithContext creates a new GetTransactionTagsParams object
// with the ability to set a context for a request.
func NewGetTransactionTagsParamsWithContext(ctx context.Context) *GetTransactionTagsParams {
	return &GetTransactionTagsParams{
		Context: ctx,
	}
}

// NewGetTransactionTagsParamsWithHTTPClient creates a new GetTransactionTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTransactionTagsParamsWithHTTPClient(client *http.Client) *GetTransactionTagsParams {
	return &GetTransactionTagsParams{
		HTTPClient: client,
	}
}

/*
GetTransactionTagsParams contains all the parameters to send to the API endpoint

	for the get transaction tags operation.

	Typically these are written to a http.Request.
*/
type GetTransactionTagsParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// IncludedDeleted.
	IncludedDeleted *bool

	// TransactionID.
	//
	// Format: uuid
	TransactionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get transaction tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionTagsParams) WithDefaults() *GetTransactionTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get transaction tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTransactionTagsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		includedDeletedDefault = bool(false)
	)

	val := GetTransactionTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get transaction tags params
func (o *GetTransactionTagsParams) WithTimeout(timeout time.Duration) *GetTransactionTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transaction tags params
func (o *GetTransactionTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transaction tags params
func (o *GetTransactionTagsParams) WithContext(ctx context.Context) *GetTransactionTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transaction tags params
func (o *GetTransactionTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transaction tags params
func (o *GetTransactionTagsParams) WithHTTPClient(client *http.Client) *GetTransactionTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transaction tags params
func (o *GetTransactionTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get transaction tags params
func (o *GetTransactionTagsParams) WithAudit(audit *string) *GetTransactionTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get transaction tags params
func (o *GetTransactionTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get transaction tags params
func (o *GetTransactionTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetTransactionTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get transaction tags params
func (o *GetTransactionTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithTransactionID adds the transactionID to the get transaction tags params
func (o *GetTransactionTagsParams) WithTransactionID(transactionID strfmt.UUID) *GetTransactionTagsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get transaction tags params
func (o *GetTransactionTagsParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit string

		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {

			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}
	}

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool

		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {

			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
		}
	}

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
