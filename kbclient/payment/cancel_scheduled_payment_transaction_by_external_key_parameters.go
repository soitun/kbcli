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
)

// NewCancelScheduledPaymentTransactionByExternalKeyParams creates a new CancelScheduledPaymentTransactionByExternalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelScheduledPaymentTransactionByExternalKeyParams() *CancelScheduledPaymentTransactionByExternalKeyParams {
	return &CancelScheduledPaymentTransactionByExternalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelScheduledPaymentTransactionByExternalKeyParamsWithTimeout creates a new CancelScheduledPaymentTransactionByExternalKeyParams object
// with the ability to set a timeout on a request.
func NewCancelScheduledPaymentTransactionByExternalKeyParamsWithTimeout(timeout time.Duration) *CancelScheduledPaymentTransactionByExternalKeyParams {
	return &CancelScheduledPaymentTransactionByExternalKeyParams{
		timeout: timeout,
	}
}

// NewCancelScheduledPaymentTransactionByExternalKeyParamsWithContext creates a new CancelScheduledPaymentTransactionByExternalKeyParams object
// with the ability to set a context for a request.
func NewCancelScheduledPaymentTransactionByExternalKeyParamsWithContext(ctx context.Context) *CancelScheduledPaymentTransactionByExternalKeyParams {
	return &CancelScheduledPaymentTransactionByExternalKeyParams{
		Context: ctx,
	}
}

// NewCancelScheduledPaymentTransactionByExternalKeyParamsWithHTTPClient creates a new CancelScheduledPaymentTransactionByExternalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelScheduledPaymentTransactionByExternalKeyParamsWithHTTPClient(client *http.Client) *CancelScheduledPaymentTransactionByExternalKeyParams {
	return &CancelScheduledPaymentTransactionByExternalKeyParams{
		HTTPClient: client,
	}
}

/*
CancelScheduledPaymentTransactionByExternalKeyParams contains all the parameters to send to the API endpoint

	for the cancel scheduled payment transaction by external key operation.

	Typically these are written to a http.Request.
*/
type CancelScheduledPaymentTransactionByExternalKeyParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// TransactionExternalKey.
	TransactionExternalKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cancel scheduled payment transaction by external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithDefaults() *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel scheduled payment transaction by external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithTimeout(timeout time.Duration) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithContext(ctx context.Context) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithHTTPClient(client *http.Client) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithXKillbillComment(xKillbillComment *string) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithXKillbillReason(xKillbillReason *string) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithTransactionExternalKey adds the transactionExternalKey to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WithTransactionExternalKey(transactionExternalKey string) *CancelScheduledPaymentTransactionByExternalKeyParams {
	o.SetTransactionExternalKey(transactionExternalKey)
	return o
}

// SetTransactionExternalKey adds the transactionExternalKey to the cancel scheduled payment transaction by external key params
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) SetTransactionExternalKey(transactionExternalKey string) {
	o.TransactionExternalKey = transactionExternalKey
}

// WriteToRequest writes these params to a swagger request
func (o *CancelScheduledPaymentTransactionByExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param transactionExternalKey
	qrTransactionExternalKey := o.TransactionExternalKey
	qTransactionExternalKey := qrTransactionExternalKey
	if qTransactionExternalKey != "" {

		if err := r.SetQueryParam("transactionExternalKey", qTransactionExternalKey); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
