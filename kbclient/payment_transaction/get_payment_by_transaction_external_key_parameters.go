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

// NewGetPaymentByTransactionExternalKeyParams creates a new GetPaymentByTransactionExternalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPaymentByTransactionExternalKeyParams() *GetPaymentByTransactionExternalKeyParams {
	return &GetPaymentByTransactionExternalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentByTransactionExternalKeyParamsWithTimeout creates a new GetPaymentByTransactionExternalKeyParams object
// with the ability to set a timeout on a request.
func NewGetPaymentByTransactionExternalKeyParamsWithTimeout(timeout time.Duration) *GetPaymentByTransactionExternalKeyParams {
	return &GetPaymentByTransactionExternalKeyParams{
		timeout: timeout,
	}
}

// NewGetPaymentByTransactionExternalKeyParamsWithContext creates a new GetPaymentByTransactionExternalKeyParams object
// with the ability to set a context for a request.
func NewGetPaymentByTransactionExternalKeyParamsWithContext(ctx context.Context) *GetPaymentByTransactionExternalKeyParams {
	return &GetPaymentByTransactionExternalKeyParams{
		Context: ctx,
	}
}

// NewGetPaymentByTransactionExternalKeyParamsWithHTTPClient creates a new GetPaymentByTransactionExternalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPaymentByTransactionExternalKeyParamsWithHTTPClient(client *http.Client) *GetPaymentByTransactionExternalKeyParams {
	return &GetPaymentByTransactionExternalKeyParams{
		HTTPClient: client,
	}
}

/*
GetPaymentByTransactionExternalKeyParams contains all the parameters to send to the API endpoint

	for the get payment by transaction external key operation.

	Typically these are written to a http.Request.
*/
type GetPaymentByTransactionExternalKeyParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// PluginProperty.
	PluginProperty []string

	// TransactionExternalKey.
	TransactionExternalKey string

	// WithAttempts.
	WithAttempts *bool

	// WithPluginInfo.
	WithPluginInfo *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get payment by transaction external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentByTransactionExternalKeyParams) WithDefaults() *GetPaymentByTransactionExternalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get payment by transaction external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentByTransactionExternalKeyParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		withAttemptsDefault = bool(false)

		withPluginInfoDefault = bool(false)
	)

	val := GetPaymentByTransactionExternalKeyParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithTimeout(timeout time.Duration) *GetPaymentByTransactionExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithContext(ctx context.Context) *GetPaymentByTransactionExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithHTTPClient(client *http.Client) *GetPaymentByTransactionExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithAudit(audit *string) *GetPaymentByTransactionExternalKeyParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPluginProperty adds the pluginProperty to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithPluginProperty(pluginProperty []string) *GetPaymentByTransactionExternalKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithTransactionExternalKey adds the transactionExternalKey to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithTransactionExternalKey(transactionExternalKey string) *GetPaymentByTransactionExternalKeyParams {
	o.SetTransactionExternalKey(transactionExternalKey)
	return o
}

// SetTransactionExternalKey adds the transactionExternalKey to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetTransactionExternalKey(transactionExternalKey string) {
	o.TransactionExternalKey = transactionExternalKey
}

// WithWithAttempts adds the withAttempts to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithWithAttempts(withAttempts *bool) *GetPaymentByTransactionExternalKeyParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) WithWithPluginInfo(withPluginInfo *bool) *GetPaymentByTransactionExternalKeyParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get payment by transaction external key params
func (o *GetPaymentByTransactionExternalKeyParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentByTransactionExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	// query param transactionExternalKey
	qrTransactionExternalKey := o.TransactionExternalKey
	qTransactionExternalKey := qrTransactionExternalKey
	if qTransactionExternalKey != "" {

		if err := r.SetQueryParam("transactionExternalKey", qTransactionExternalKey); err != nil {
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

// bindParamGetPaymentByTransactionExternalKey binds the parameter pluginProperty
func (o *GetPaymentByTransactionExternalKeyParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
