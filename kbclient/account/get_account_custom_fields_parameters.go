// Code generated by go-swagger; DO NOT EDIT.

package account

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
)

// NewGetAccountCustomFieldsParams creates a new GetAccountCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountCustomFieldsParams() *GetAccountCustomFieldsParams {
	return &GetAccountCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountCustomFieldsParamsWithTimeout creates a new GetAccountCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewGetAccountCustomFieldsParamsWithTimeout(timeout time.Duration) *GetAccountCustomFieldsParams {
	return &GetAccountCustomFieldsParams{
		timeout: timeout,
	}
}

// NewGetAccountCustomFieldsParamsWithContext creates a new GetAccountCustomFieldsParams object
// with the ability to set a context for a request.
func NewGetAccountCustomFieldsParamsWithContext(ctx context.Context) *GetAccountCustomFieldsParams {
	return &GetAccountCustomFieldsParams{
		Context: ctx,
	}
}

// NewGetAccountCustomFieldsParamsWithHTTPClient creates a new GetAccountCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountCustomFieldsParamsWithHTTPClient(client *http.Client) *GetAccountCustomFieldsParams {
	return &GetAccountCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
GetAccountCustomFieldsParams contains all the parameters to send to the API endpoint

	for the get account custom fields operation.

	Typically these are written to a http.Request.
*/
type GetAccountCustomFieldsParams struct {

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountCustomFieldsParams) WithDefaults() *GetAccountCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountCustomFieldsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")
	)

	val := GetAccountCustomFieldsParams{
		Audit: &auditDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithTimeout(timeout time.Duration) *GetAccountCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithContext(ctx context.Context) *GetAccountCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithHTTPClient(client *http.Client) *GetAccountCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithAccountID(accountID strfmt.UUID) *GetAccountCustomFieldsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get account custom fields params
func (o *GetAccountCustomFieldsParams) WithAudit(audit *string) *GetAccountCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get account custom fields params
func (o *GetAccountCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
