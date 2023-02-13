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

// NewCreateAccountTagsParams creates a new CreateAccountTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAccountTagsParams() *CreateAccountTagsParams {
	return &CreateAccountTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountTagsParamsWithTimeout creates a new CreateAccountTagsParams object
// with the ability to set a timeout on a request.
func NewCreateAccountTagsParamsWithTimeout(timeout time.Duration) *CreateAccountTagsParams {
	return &CreateAccountTagsParams{
		timeout: timeout,
	}
}

// NewCreateAccountTagsParamsWithContext creates a new CreateAccountTagsParams object
// with the ability to set a context for a request.
func NewCreateAccountTagsParamsWithContext(ctx context.Context) *CreateAccountTagsParams {
	return &CreateAccountTagsParams{
		Context: ctx,
	}
}

// NewCreateAccountTagsParamsWithHTTPClient creates a new CreateAccountTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAccountTagsParamsWithHTTPClient(client *http.Client) *CreateAccountTagsParams {
	return &CreateAccountTagsParams{
		HTTPClient: client,
	}
}

/*
CreateAccountTagsParams contains all the parameters to send to the API endpoint

	for the create account tags operation.

	Typically these are written to a http.Request.
*/
type CreateAccountTagsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	// Body.
	Body []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create account tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccountTagsParams) WithDefaults() *CreateAccountTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create account tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccountTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create account tags params
func (o *CreateAccountTagsParams) WithTimeout(timeout time.Duration) *CreateAccountTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account tags params
func (o *CreateAccountTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account tags params
func (o *CreateAccountTagsParams) WithContext(ctx context.Context) *CreateAccountTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account tags params
func (o *CreateAccountTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create account tags params
func (o *CreateAccountTagsParams) WithHTTPClient(client *http.Client) *CreateAccountTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create account tags params
func (o *CreateAccountTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillComment(xKillbillComment *string) *CreateAccountTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateAccountTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillReason(xKillbillReason *string) *CreateAccountTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create account tags params
func (o *CreateAccountTagsParams) WithAccountID(accountID strfmt.UUID) *CreateAccountTagsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create account tags params
func (o *CreateAccountTagsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the create account tags params
func (o *CreateAccountTagsParams) WithBody(body []strfmt.UUID) *CreateAccountTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create account tags params
func (o *CreateAccountTagsParams) SetBody(body []strfmt.UUID) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
