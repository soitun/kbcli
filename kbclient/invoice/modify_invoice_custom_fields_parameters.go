// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewModifyInvoiceCustomFieldsParams creates a new ModifyInvoiceCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyInvoiceCustomFieldsParams() *ModifyInvoiceCustomFieldsParams {
	return &ModifyInvoiceCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyInvoiceCustomFieldsParamsWithTimeout creates a new ModifyInvoiceCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewModifyInvoiceCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyInvoiceCustomFieldsParams {
	return &ModifyInvoiceCustomFieldsParams{
		timeout: timeout,
	}
}

// NewModifyInvoiceCustomFieldsParamsWithContext creates a new ModifyInvoiceCustomFieldsParams object
// with the ability to set a context for a request.
func NewModifyInvoiceCustomFieldsParamsWithContext(ctx context.Context) *ModifyInvoiceCustomFieldsParams {
	return &ModifyInvoiceCustomFieldsParams{
		Context: ctx,
	}
}

// NewModifyInvoiceCustomFieldsParamsWithHTTPClient creates a new ModifyInvoiceCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyInvoiceCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyInvoiceCustomFieldsParams {
	return &ModifyInvoiceCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
ModifyInvoiceCustomFieldsParams contains all the parameters to send to the API endpoint

	for the modify invoice custom fields operation.

	Typically these are written to a http.Request.
*/
type ModifyInvoiceCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []*kbmodel.CustomField

	// InvoiceID.
	//
	// Format: uuid
	InvoiceID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the modify invoice custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyInvoiceCustomFieldsParams) WithDefaults() *ModifyInvoiceCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify invoice custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyInvoiceCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyInvoiceCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithContext(ctx context.Context) *ModifyInvoiceCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyInvoiceCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyInvoiceCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyInvoiceCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyInvoiceCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyInvoiceCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithInvoiceID adds the invoiceID to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) WithInvoiceID(invoiceID strfmt.UUID) *ModifyInvoiceCustomFieldsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the modify invoice custom fields params
func (o *ModifyInvoiceCustomFieldsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyInvoiceCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
		return err
	}

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
