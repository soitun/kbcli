// Code generated by go-swagger; DO NOT EDIT.

package overdue

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

// NewUploadOverdueConfigJSONParams creates a new UploadOverdueConfigJSONParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadOverdueConfigJSONParams() *UploadOverdueConfigJSONParams {
	return &UploadOverdueConfigJSONParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadOverdueConfigJSONParamsWithTimeout creates a new UploadOverdueConfigJSONParams object
// with the ability to set a timeout on a request.
func NewUploadOverdueConfigJSONParamsWithTimeout(timeout time.Duration) *UploadOverdueConfigJSONParams {
	return &UploadOverdueConfigJSONParams{
		timeout: timeout,
	}
}

// NewUploadOverdueConfigJSONParamsWithContext creates a new UploadOverdueConfigJSONParams object
// with the ability to set a context for a request.
func NewUploadOverdueConfigJSONParamsWithContext(ctx context.Context) *UploadOverdueConfigJSONParams {
	return &UploadOverdueConfigJSONParams{
		Context: ctx,
	}
}

// NewUploadOverdueConfigJSONParamsWithHTTPClient creates a new UploadOverdueConfigJSONParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadOverdueConfigJSONParamsWithHTTPClient(client *http.Client) *UploadOverdueConfigJSONParams {
	return &UploadOverdueConfigJSONParams{
		HTTPClient: client,
	}
}

/*
UploadOverdueConfigJSONParams contains all the parameters to send to the API endpoint

	for the upload overdue config Json operation.

	Typically these are written to a http.Request.
*/
type UploadOverdueConfigJSONParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.Overdue

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the upload overdue config Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadOverdueConfigJSONParams) WithDefaults() *UploadOverdueConfigJSONParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload overdue config Json params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadOverdueConfigJSONParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithTimeout(timeout time.Duration) *UploadOverdueConfigJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithContext(ctx context.Context) *UploadOverdueConfigJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithHTTPClient(client *http.Client) *UploadOverdueConfigJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithXKillbillComment(xKillbillComment *string) *UploadOverdueConfigJSONParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadOverdueConfigJSONParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithXKillbillReason(xKillbillReason *string) *UploadOverdueConfigJSONParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) WithBody(body *kbmodel.Overdue) *UploadOverdueConfigJSONParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload overdue config Json params
func (o *UploadOverdueConfigJSONParams) SetBody(body *kbmodel.Overdue) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UploadOverdueConfigJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
