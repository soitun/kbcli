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

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// NewCreateInvoicePaymentCustomFieldsParams creates a new CreateInvoicePaymentCustomFieldsParams object
// with the default values initialized.
func NewCreateInvoicePaymentCustomFieldsParams() *CreateInvoicePaymentCustomFieldsParams {
	var ()
	return &CreateInvoicePaymentCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoicePaymentCustomFieldsParamsWithTimeout creates a new CreateInvoicePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInvoicePaymentCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateInvoicePaymentCustomFieldsParams {
	var ()
	return &CreateInvoicePaymentCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreateInvoicePaymentCustomFieldsParamsWithContext creates a new CreateInvoicePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInvoicePaymentCustomFieldsParamsWithContext(ctx context.Context) *CreateInvoicePaymentCustomFieldsParams {
	var ()
	return &CreateInvoicePaymentCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreateInvoicePaymentCustomFieldsParamsWithHTTPClient creates a new CreateInvoicePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInvoicePaymentCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateInvoicePaymentCustomFieldsParams {
	var ()
	return &CreateInvoicePaymentCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreateInvoicePaymentCustomFieldsParams contains all the parameters to send to the API endpoint
for the create invoice payment custom fields operation typically these are written to a http.Request
*/
type CreateInvoicePaymentCustomFieldsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []*kbmodel.CustomField
	/*PaymentID*/
	PaymentID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateInvoicePaymentCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithContext(ctx context.Context) *CreateInvoicePaymentCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateInvoicePaymentCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateInvoicePaymentCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInvoicePaymentCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateInvoicePaymentCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateInvoicePaymentCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) WithPaymentID(paymentID strfmt.UUID) *CreateInvoicePaymentCustomFieldsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create invoice payment custom fields params
func (o *CreateInvoicePaymentCustomFieldsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoicePaymentCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
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
