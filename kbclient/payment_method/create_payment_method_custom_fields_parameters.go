// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewCreatePaymentMethodCustomFieldsParams creates a new CreatePaymentMethodCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePaymentMethodCustomFieldsParams() *CreatePaymentMethodCustomFieldsParams {
	return &CreatePaymentMethodCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePaymentMethodCustomFieldsParamsWithTimeout creates a new CreatePaymentMethodCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewCreatePaymentMethodCustomFieldsParamsWithTimeout(timeout time.Duration) *CreatePaymentMethodCustomFieldsParams {
	return &CreatePaymentMethodCustomFieldsParams{
		timeout: timeout,
	}
}

// NewCreatePaymentMethodCustomFieldsParamsWithContext creates a new CreatePaymentMethodCustomFieldsParams object
// with the ability to set a context for a request.
func NewCreatePaymentMethodCustomFieldsParamsWithContext(ctx context.Context) *CreatePaymentMethodCustomFieldsParams {
	return &CreatePaymentMethodCustomFieldsParams{
		Context: ctx,
	}
}

// NewCreatePaymentMethodCustomFieldsParamsWithHTTPClient creates a new CreatePaymentMethodCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePaymentMethodCustomFieldsParamsWithHTTPClient(client *http.Client) *CreatePaymentMethodCustomFieldsParams {
	return &CreatePaymentMethodCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
CreatePaymentMethodCustomFieldsParams contains all the parameters to send to the API endpoint

	for the create payment method custom fields operation.

	Typically these are written to a http.Request.
*/
type CreatePaymentMethodCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []*kbmodel.CustomField

	// PaymentMethodID.
	//
	// Format: uuid
	PaymentMethodID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create payment method custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePaymentMethodCustomFieldsParams) WithDefaults() *CreatePaymentMethodCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create payment method custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePaymentMethodCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithTimeout(timeout time.Duration) *CreatePaymentMethodCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithContext(ctx context.Context) *CreatePaymentMethodCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithHTTPClient(client *http.Client) *CreatePaymentMethodCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreatePaymentMethodCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreatePaymentMethodCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreatePaymentMethodCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreatePaymentMethodCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithPaymentMethodID adds the paymentMethodID to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *CreatePaymentMethodCustomFieldsParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the create payment method custom fields params
func (o *CreatePaymentMethodCustomFieldsParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePaymentMethodCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
