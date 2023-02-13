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

// NewGetEmailsParams creates a new GetEmailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEmailsParams() *GetEmailsParams {
	return &GetEmailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailsParamsWithTimeout creates a new GetEmailsParams object
// with the ability to set a timeout on a request.
func NewGetEmailsParamsWithTimeout(timeout time.Duration) *GetEmailsParams {
	return &GetEmailsParams{
		timeout: timeout,
	}
}

// NewGetEmailsParamsWithContext creates a new GetEmailsParams object
// with the ability to set a context for a request.
func NewGetEmailsParamsWithContext(ctx context.Context) *GetEmailsParams {
	return &GetEmailsParams{
		Context: ctx,
	}
}

// NewGetEmailsParamsWithHTTPClient creates a new GetEmailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEmailsParamsWithHTTPClient(client *http.Client) *GetEmailsParams {
	return &GetEmailsParams{
		HTTPClient: client,
	}
}

/*
GetEmailsParams contains all the parameters to send to the API endpoint

	for the get emails operation.

	Typically these are written to a http.Request.
*/
type GetEmailsParams struct {

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmailsParams) WithDefaults() *GetEmailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get emails params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get emails params
func (o *GetEmailsParams) WithTimeout(timeout time.Duration) *GetEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get emails params
func (o *GetEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get emails params
func (o *GetEmailsParams) WithContext(ctx context.Context) *GetEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get emails params
func (o *GetEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get emails params
func (o *GetEmailsParams) WithHTTPClient(client *http.Client) *GetEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get emails params
func (o *GetEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get emails params
func (o *GetEmailsParams) WithAccountID(accountID strfmt.UUID) *GetEmailsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get emails params
func (o *GetEmailsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
