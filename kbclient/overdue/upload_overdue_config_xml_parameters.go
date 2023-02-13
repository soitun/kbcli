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
)

// NewUploadOverdueConfigXMLParams creates a new UploadOverdueConfigXMLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadOverdueConfigXMLParams() *UploadOverdueConfigXMLParams {
	return &UploadOverdueConfigXMLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadOverdueConfigXMLParamsWithTimeout creates a new UploadOverdueConfigXMLParams object
// with the ability to set a timeout on a request.
func NewUploadOverdueConfigXMLParamsWithTimeout(timeout time.Duration) *UploadOverdueConfigXMLParams {
	return &UploadOverdueConfigXMLParams{
		timeout: timeout,
	}
}

// NewUploadOverdueConfigXMLParamsWithContext creates a new UploadOverdueConfigXMLParams object
// with the ability to set a context for a request.
func NewUploadOverdueConfigXMLParamsWithContext(ctx context.Context) *UploadOverdueConfigXMLParams {
	return &UploadOverdueConfigXMLParams{
		Context: ctx,
	}
}

// NewUploadOverdueConfigXMLParamsWithHTTPClient creates a new UploadOverdueConfigXMLParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadOverdueConfigXMLParamsWithHTTPClient(client *http.Client) *UploadOverdueConfigXMLParams {
	return &UploadOverdueConfigXMLParams{
		HTTPClient: client,
	}
}

/*
UploadOverdueConfigXMLParams contains all the parameters to send to the API endpoint

	for the upload overdue config Xml operation.

	Typically these are written to a http.Request.
*/
type UploadOverdueConfigXMLParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upload overdue config Xml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadOverdueConfigXMLParams) WithDefaults() *UploadOverdueConfigXMLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload overdue config Xml params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadOverdueConfigXMLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithTimeout(timeout time.Duration) *UploadOverdueConfigXMLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithContext(ctx context.Context) *UploadOverdueConfigXMLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithHTTPClient(client *http.Client) *UploadOverdueConfigXMLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillComment(xKillbillComment *string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillReason(xKillbillReason *string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithBody(body string) *UploadOverdueConfigXMLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UploadOverdueConfigXMLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
