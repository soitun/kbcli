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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// NewCreateTaxItemsParams creates a new CreateTaxItemsParams object
// with the default values initialized.
func NewCreateTaxItemsParams() *CreateTaxItemsParams {
	var (
		autoCommitDefault = bool(false)
	)
	return &CreateTaxItemsParams{
		AutoCommit: &autoCommitDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTaxItemsParamsWithTimeout creates a new CreateTaxItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTaxItemsParamsWithTimeout(timeout time.Duration) *CreateTaxItemsParams {
	var (
		autoCommitDefault = bool(false)
	)
	return &CreateTaxItemsParams{
		AutoCommit: &autoCommitDefault,

		timeout: timeout,
	}
}

// NewCreateTaxItemsParamsWithContext creates a new CreateTaxItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTaxItemsParamsWithContext(ctx context.Context) *CreateTaxItemsParams {
	var (
		autoCommitDefault = bool(false)
	)
	return &CreateTaxItemsParams{
		AutoCommit: &autoCommitDefault,

		Context: ctx,
	}
}

// NewCreateTaxItemsParamsWithHTTPClient creates a new CreateTaxItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTaxItemsParamsWithHTTPClient(client *http.Client) *CreateTaxItemsParams {
	var (
		autoCommitDefault = bool(false)
	)
	return &CreateTaxItemsParams{
		AutoCommit: &autoCommitDefault,
		HTTPClient: client,
	}
}

/*CreateTaxItemsParams contains all the parameters to send to the API endpoint
for the create tax items operation typically these are written to a http.Request
*/
type CreateTaxItemsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*AutoCommit*/
	AutoCommit *bool
	/*Body*/
	Body []*kbmodel.InvoiceItem
	/*PluginProperty*/
	PluginProperty []string
	/*RequestedDate*/
	RequestedDate *strfmt.Date

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create tax items params
func (o *CreateTaxItemsParams) WithTimeout(timeout time.Duration) *CreateTaxItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create tax items params
func (o *CreateTaxItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create tax items params
func (o *CreateTaxItemsParams) WithContext(ctx context.Context) *CreateTaxItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create tax items params
func (o *CreateTaxItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create tax items params
func (o *CreateTaxItemsParams) WithHTTPClient(client *http.Client) *CreateTaxItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create tax items params
func (o *CreateTaxItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create tax items params
func (o *CreateTaxItemsParams) WithXKillbillComment(xKillbillComment *string) *CreateTaxItemsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create tax items params
func (o *CreateTaxItemsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create tax items params
func (o *CreateTaxItemsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateTaxItemsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create tax items params
func (o *CreateTaxItemsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create tax items params
func (o *CreateTaxItemsParams) WithXKillbillReason(xKillbillReason *string) *CreateTaxItemsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create tax items params
func (o *CreateTaxItemsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create tax items params
func (o *CreateTaxItemsParams) WithAccountID(accountID strfmt.UUID) *CreateTaxItemsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create tax items params
func (o *CreateTaxItemsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAutoCommit adds the autoCommit to the create tax items params
func (o *CreateTaxItemsParams) WithAutoCommit(autoCommit *bool) *CreateTaxItemsParams {
	o.SetAutoCommit(autoCommit)
	return o
}

// SetAutoCommit adds the autoCommit to the create tax items params
func (o *CreateTaxItemsParams) SetAutoCommit(autoCommit *bool) {
	o.AutoCommit = autoCommit
}

// WithBody adds the body to the create tax items params
func (o *CreateTaxItemsParams) WithBody(body []*kbmodel.InvoiceItem) *CreateTaxItemsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create tax items params
func (o *CreateTaxItemsParams) SetBody(body []*kbmodel.InvoiceItem) {
	o.Body = body
}

// WithPluginProperty adds the pluginProperty to the create tax items params
func (o *CreateTaxItemsParams) WithPluginProperty(pluginProperty []string) *CreateTaxItemsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create tax items params
func (o *CreateTaxItemsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the create tax items params
func (o *CreateTaxItemsParams) WithRequestedDate(requestedDate *strfmt.Date) *CreateTaxItemsParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the create tax items params
func (o *CreateTaxItemsParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTaxItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AutoCommit != nil {

		// query param autoCommit
		var qrAutoCommit bool
		if o.AutoCommit != nil {
			qrAutoCommit = *o.AutoCommit
		}
		qAutoCommit := swag.FormatBool(qrAutoCommit)
		if qAutoCommit != "" {
			if err := r.SetQueryParam("autoCommit", qAutoCommit); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
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
