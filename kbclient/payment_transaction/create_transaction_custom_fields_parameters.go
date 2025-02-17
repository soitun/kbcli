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

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// NewCreateTransactionCustomFieldsParams creates a new CreateTransactionCustomFieldsParams object
// with the default values initialized.
func NewCreateTransactionCustomFieldsParams() *CreateTransactionCustomFieldsParams {
	var ()
	return &CreateTransactionCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateTransactionCustomFieldsParamsWithTimeout creates a new CreateTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateTransactionCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateTransactionCustomFieldsParams {
	var ()
	return &CreateTransactionCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreateTransactionCustomFieldsParamsWithContext creates a new CreateTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateTransactionCustomFieldsParamsWithContext(ctx context.Context) *CreateTransactionCustomFieldsParams {
	var ()
	return &CreateTransactionCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreateTransactionCustomFieldsParamsWithHTTPClient creates a new CreateTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateTransactionCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateTransactionCustomFieldsParams {
	var ()
	return &CreateTransactionCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreateTransactionCustomFieldsParams contains all the parameters to send to the API endpoint
for the create transaction custom fields operation typically these are written to a http.Request
*/
type CreateTransactionCustomFieldsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []*kbmodel.CustomField
	/*TransactionID*/
	TransactionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateTransactionCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithContext(ctx context.Context) *CreateTransactionCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateTransactionCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateTransactionCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateTransactionCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateTransactionCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateTransactionCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithTransactionID adds the transactionID to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) WithTransactionID(transactionID strfmt.UUID) *CreateTransactionCustomFieldsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the create transaction custom fields params
func (o *CreateTransactionCustomFieldsParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateTransactionCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID.String()); err != nil {
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
