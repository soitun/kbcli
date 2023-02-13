// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// NewGetBundleCustomFieldsParams creates a new GetBundleCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBundleCustomFieldsParams() *GetBundleCustomFieldsParams {
	return &GetBundleCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBundleCustomFieldsParamsWithTimeout creates a new GetBundleCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewGetBundleCustomFieldsParamsWithTimeout(timeout time.Duration) *GetBundleCustomFieldsParams {
	return &GetBundleCustomFieldsParams{
		timeout: timeout,
	}
}

// NewGetBundleCustomFieldsParamsWithContext creates a new GetBundleCustomFieldsParams object
// with the ability to set a context for a request.
func NewGetBundleCustomFieldsParamsWithContext(ctx context.Context) *GetBundleCustomFieldsParams {
	return &GetBundleCustomFieldsParams{
		Context: ctx,
	}
}

// NewGetBundleCustomFieldsParamsWithHTTPClient creates a new GetBundleCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBundleCustomFieldsParamsWithHTTPClient(client *http.Client) *GetBundleCustomFieldsParams {
	return &GetBundleCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
GetBundleCustomFieldsParams contains all the parameters to send to the API endpoint

	for the get bundle custom fields operation.

	Typically these are written to a http.Request.
*/
type GetBundleCustomFieldsParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// BundleID.
	//
	// Format: uuid
	BundleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get bundle custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundleCustomFieldsParams) WithDefaults() *GetBundleCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get bundle custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBundleCustomFieldsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")
	)

	val := GetBundleCustomFieldsParams{
		Audit: &auditDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) WithTimeout(timeout time.Duration) *GetBundleCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) WithContext(ctx context.Context) *GetBundleCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) WithHTTPClient(client *http.Client) *GetBundleCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) WithAudit(audit *string) *GetBundleCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithBundleID adds the bundleID to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) WithBundleID(bundleID strfmt.UUID) *GetBundleCustomFieldsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the get bundle custom fields params
func (o *GetBundleCustomFieldsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetBundleCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
