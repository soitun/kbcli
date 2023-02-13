// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// NewGetPaymentParams creates a new GetPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPaymentParams() *GetPaymentParams {
	return &GetPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentParamsWithTimeout creates a new GetPaymentParams object
// with the ability to set a timeout on a request.
func NewGetPaymentParamsWithTimeout(timeout time.Duration) *GetPaymentParams {
	return &GetPaymentParams{
		timeout: timeout,
	}
}

// NewGetPaymentParamsWithContext creates a new GetPaymentParams object
// with the ability to set a context for a request.
func NewGetPaymentParamsWithContext(ctx context.Context) *GetPaymentParams {
	return &GetPaymentParams{
		Context: ctx,
	}
}

// NewGetPaymentParamsWithHTTPClient creates a new GetPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPaymentParamsWithHTTPClient(client *http.Client) *GetPaymentParams {
	return &GetPaymentParams{
		HTTPClient: client,
	}
}

/*
GetPaymentParams contains all the parameters to send to the API endpoint

	for the get payment operation.

	Typically these are written to a http.Request.
*/
type GetPaymentParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	// WithAttempts.
	WithAttempts *bool

	// WithPluginInfo.
	WithPluginInfo *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentParams) WithDefaults() *GetPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		withAttemptsDefault = bool(false)

		withPluginInfoDefault = bool(false)
	)

	val := GetPaymentParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get payment params
func (o *GetPaymentParams) WithTimeout(timeout time.Duration) *GetPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment params
func (o *GetPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment params
func (o *GetPaymentParams) WithContext(ctx context.Context) *GetPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment params
func (o *GetPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment params
func (o *GetPaymentParams) WithHTTPClient(client *http.Client) *GetPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment params
func (o *GetPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get payment params
func (o *GetPaymentParams) WithAudit(audit *string) *GetPaymentParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment params
func (o *GetPaymentParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPaymentID adds the paymentID to the get payment params
func (o *GetPaymentParams) WithPaymentID(paymentID strfmt.UUID) *GetPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the get payment params
func (o *GetPaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the get payment params
func (o *GetPaymentParams) WithPluginProperty(pluginProperty []string) *GetPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payment params
func (o *GetPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithAttempts adds the withAttempts to the get payment params
func (o *GetPaymentParams) WithWithAttempts(withAttempts *bool) *GetPaymentParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get payment params
func (o *GetPaymentParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get payment params
func (o *GetPaymentParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payment params
func (o *GetPaymentParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if o.WithAttempts != nil {

		// query param withAttempts
		var qrWithAttempts bool

		if o.WithAttempts != nil {
			qrWithAttempts = *o.WithAttempts
		}
		qWithAttempts := swag.FormatBool(qrWithAttempts)
		if qWithAttempts != "" {

			if err := r.SetQueryParam("withAttempts", qWithAttempts); err != nil {
				return err
			}
		}
	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool

		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {

			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetPayment binds the parameter pluginProperty
func (o *GetPaymentParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
