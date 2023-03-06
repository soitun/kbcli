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
	"github.com/go-openapi/swag"

	"github.com/killbill/kbcli/v3/kbmodel"
)

// NewAddAccountBlockingStateParams creates a new AddAccountBlockingStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAccountBlockingStateParams() *AddAccountBlockingStateParams {
	return &AddAccountBlockingStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAccountBlockingStateParamsWithTimeout creates a new AddAccountBlockingStateParams object
// with the ability to set a timeout on a request.
func NewAddAccountBlockingStateParamsWithTimeout(timeout time.Duration) *AddAccountBlockingStateParams {
	return &AddAccountBlockingStateParams{
		timeout: timeout,
	}
}

// NewAddAccountBlockingStateParamsWithContext creates a new AddAccountBlockingStateParams object
// with the ability to set a context for a request.
func NewAddAccountBlockingStateParamsWithContext(ctx context.Context) *AddAccountBlockingStateParams {
	return &AddAccountBlockingStateParams{
		Context: ctx,
	}
}

// NewAddAccountBlockingStateParamsWithHTTPClient creates a new AddAccountBlockingStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAccountBlockingStateParamsWithHTTPClient(client *http.Client) *AddAccountBlockingStateParams {
	return &AddAccountBlockingStateParams{
		HTTPClient: client,
	}
}

/*
AddAccountBlockingStateParams contains all the parameters to send to the API endpoint

	for the add account blocking state operation.

	Typically these are written to a http.Request.
*/
type AddAccountBlockingStateParams struct {

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
	Body *kbmodel.BlockingState

	// PluginProperty.
	PluginProperty []string

	// RequestedDate.
	//
	// Format: date
	RequestedDate *strfmt.Date

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the add account blocking state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAccountBlockingStateParams) WithDefaults() *AddAccountBlockingStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add account blocking state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAccountBlockingStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithTimeout(timeout time.Duration) *AddAccountBlockingStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithContext(ctx context.Context) *AddAccountBlockingStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithHTTPClient(client *http.Client) *AddAccountBlockingStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithXKillbillComment(xKillbillComment *string) *AddAccountBlockingStateParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddAccountBlockingStateParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithXKillbillReason(xKillbillReason *string) *AddAccountBlockingStateParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithAccountID(accountID strfmt.UUID) *AddAccountBlockingStateParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithBody(body *kbmodel.BlockingState) *AddAccountBlockingStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetBody(body *kbmodel.BlockingState) {
	o.Body = body
}

// WithPluginProperty adds the pluginProperty to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithPluginProperty(pluginProperty []string) *AddAccountBlockingStateParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the add account blocking state params
func (o *AddAccountBlockingStateParams) WithRequestedDate(requestedDate *strfmt.Date) *AddAccountBlockingStateParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the add account blocking state params
func (o *AddAccountBlockingStateParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *AddAccountBlockingStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date

		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {

			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
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

// bindParamAddAccountBlockingState binds the parameter pluginProperty
func (o *AddAccountBlockingStateParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
