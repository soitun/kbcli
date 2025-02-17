// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

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

// NewCompleteInvoicePaymentTransactionParams creates a new CompleteInvoicePaymentTransactionParams object
// with the default values initialized.
func NewCompleteInvoicePaymentTransactionParams() *CompleteInvoicePaymentTransactionParams {
	var ()
	return &CompleteInvoicePaymentTransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteInvoicePaymentTransactionParamsWithTimeout creates a new CompleteInvoicePaymentTransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCompleteInvoicePaymentTransactionParamsWithTimeout(timeout time.Duration) *CompleteInvoicePaymentTransactionParams {
	var ()
	return &CompleteInvoicePaymentTransactionParams{

		timeout: timeout,
	}
}

// NewCompleteInvoicePaymentTransactionParamsWithContext creates a new CompleteInvoicePaymentTransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCompleteInvoicePaymentTransactionParamsWithContext(ctx context.Context) *CompleteInvoicePaymentTransactionParams {
	var ()
	return &CompleteInvoicePaymentTransactionParams{

		Context: ctx,
	}
}

// NewCompleteInvoicePaymentTransactionParamsWithHTTPClient creates a new CompleteInvoicePaymentTransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCompleteInvoicePaymentTransactionParamsWithHTTPClient(client *http.Client) *CompleteInvoicePaymentTransactionParams {
	var ()
	return &CompleteInvoicePaymentTransactionParams{
		HTTPClient: client,
	}
}

/*CompleteInvoicePaymentTransactionParams contains all the parameters to send to the API endpoint
for the complete invoice payment transaction operation typically these are written to a http.Request
*/
type CompleteInvoicePaymentTransactionParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.PaymentTransaction
	/*ControlPluginName*/
	ControlPluginName []string
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithTimeout(timeout time.Duration) *CompleteInvoicePaymentTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithContext(ctx context.Context) *CompleteInvoicePaymentTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithHTTPClient(client *http.Client) *CompleteInvoicePaymentTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithXKillbillComment(xKillbillComment *string) *CompleteInvoicePaymentTransactionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CompleteInvoicePaymentTransactionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithXKillbillReason(xKillbillReason *string) *CompleteInvoicePaymentTransactionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithBody(body *kbmodel.PaymentTransaction) *CompleteInvoicePaymentTransactionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithControlPluginName(controlPluginName []string) *CompleteInvoicePaymentTransactionParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithPaymentID(paymentID strfmt.UUID) *CompleteInvoicePaymentTransactionParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) WithPluginProperty(pluginProperty []string) *CompleteInvoicePaymentTransactionParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the complete invoice payment transaction params
func (o *CompleteInvoicePaymentTransactionParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteInvoicePaymentTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesControlPluginName := o.ControlPluginName

	joinedControlPluginName := swag.JoinByFormat(valuesControlPluginName, "multi")
	// query array param controlPluginName
	if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
		return err
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
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
