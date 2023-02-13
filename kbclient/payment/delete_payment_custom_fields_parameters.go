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

// NewDeletePaymentCustomFieldsParams creates a new DeletePaymentCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePaymentCustomFieldsParams() *DeletePaymentCustomFieldsParams {
	return &DeletePaymentCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentCustomFieldsParamsWithTimeout creates a new DeletePaymentCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewDeletePaymentCustomFieldsParamsWithTimeout(timeout time.Duration) *DeletePaymentCustomFieldsParams {
	return &DeletePaymentCustomFieldsParams{
		timeout: timeout,
	}
}

// NewDeletePaymentCustomFieldsParamsWithContext creates a new DeletePaymentCustomFieldsParams object
// with the ability to set a context for a request.
func NewDeletePaymentCustomFieldsParamsWithContext(ctx context.Context) *DeletePaymentCustomFieldsParams {
	return &DeletePaymentCustomFieldsParams{
		Context: ctx,
	}
}

// NewDeletePaymentCustomFieldsParamsWithHTTPClient creates a new DeletePaymentCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePaymentCustomFieldsParamsWithHTTPClient(client *http.Client) *DeletePaymentCustomFieldsParams {
	return &DeletePaymentCustomFieldsParams{
		HTTPClient: client,
	}
}

/*
DeletePaymentCustomFieldsParams contains all the parameters to send to the API endpoint

	for the delete payment custom fields operation.

	Typically these are written to a http.Request.
*/
type DeletePaymentCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// CustomField.
	CustomField []strfmt.UUID

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete payment custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePaymentCustomFieldsParams) WithDefaults() *DeletePaymentCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete payment custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePaymentCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithTimeout(timeout time.Duration) *DeletePaymentCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithContext(ctx context.Context) *DeletePaymentCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithHTTPClient(client *http.Client) *DeletePaymentCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeletePaymentCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePaymentCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeletePaymentCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithCustomField adds the customField to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeletePaymentCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WithPaymentID adds the paymentID to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) WithPaymentID(paymentID strfmt.UUID) *DeletePaymentCustomFieldsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the delete payment custom fields params
func (o *DeletePaymentCustomFieldsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CustomField != nil {

		// binding items for customField
		joinedCustomField := o.bindParamCustomField(reg)

		// query array param customField
		if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
			return err
		}
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeletePaymentCustomFields binds the parameter customField
func (o *DeletePaymentCustomFieldsParams) bindParamCustomField(formats strfmt.Registry) []string {
	customFieldIR := o.CustomField

	var customFieldIC []string
	for _, customFieldIIR := range customFieldIR { // explode []strfmt.UUID

		customFieldIIV := customFieldIIR.String() // strfmt.UUID as string
		customFieldIC = append(customFieldIC, customFieldIIV)
	}

	// items.CollectionFormat: "multi"
	customFieldIS := swag.JoinByFormat(customFieldIC, "multi")

	return customFieldIS
}
