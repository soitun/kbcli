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

// NewDeletePaymentTagsParams creates a new DeletePaymentTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePaymentTagsParams() *DeletePaymentTagsParams {
	return &DeletePaymentTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentTagsParamsWithTimeout creates a new DeletePaymentTagsParams object
// with the ability to set a timeout on a request.
func NewDeletePaymentTagsParamsWithTimeout(timeout time.Duration) *DeletePaymentTagsParams {
	return &DeletePaymentTagsParams{
		timeout: timeout,
	}
}

// NewDeletePaymentTagsParamsWithContext creates a new DeletePaymentTagsParams object
// with the ability to set a context for a request.
func NewDeletePaymentTagsParamsWithContext(ctx context.Context) *DeletePaymentTagsParams {
	return &DeletePaymentTagsParams{
		Context: ctx,
	}
}

// NewDeletePaymentTagsParamsWithHTTPClient creates a new DeletePaymentTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePaymentTagsParamsWithHTTPClient(client *http.Client) *DeletePaymentTagsParams {
	return &DeletePaymentTagsParams{
		HTTPClient: client,
	}
}

/*
DeletePaymentTagsParams contains all the parameters to send to the API endpoint

	for the delete payment tags operation.

	Typically these are written to a http.Request.
*/
type DeletePaymentTagsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	// TagDef.
	TagDef []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete payment tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePaymentTagsParams) WithDefaults() *DeletePaymentTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete payment tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePaymentTagsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete payment tags params
func (o *DeletePaymentTagsParams) WithTimeout(timeout time.Duration) *DeletePaymentTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete payment tags params
func (o *DeletePaymentTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete payment tags params
func (o *DeletePaymentTagsParams) WithContext(ctx context.Context) *DeletePaymentTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete payment tags params
func (o *DeletePaymentTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete payment tags params
func (o *DeletePaymentTagsParams) WithHTTPClient(client *http.Client) *DeletePaymentTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete payment tags params
func (o *DeletePaymentTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete payment tags params
func (o *DeletePaymentTagsParams) WithXKillbillComment(xKillbillComment *string) *DeletePaymentTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete payment tags params
func (o *DeletePaymentTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment tags params
func (o *DeletePaymentTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePaymentTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment tags params
func (o *DeletePaymentTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete payment tags params
func (o *DeletePaymentTagsParams) WithXKillbillReason(xKillbillReason *string) *DeletePaymentTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete payment tags params
func (o *DeletePaymentTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPaymentID adds the paymentID to the delete payment tags params
func (o *DeletePaymentTagsParams) WithPaymentID(paymentID strfmt.UUID) *DeletePaymentTagsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the delete payment tags params
func (o *DeletePaymentTagsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithTagDef adds the tagDef to the delete payment tags params
func (o *DeletePaymentTagsParams) WithTagDef(tagDef []strfmt.UUID) *DeletePaymentTagsParams {
	o.SetTagDef(tagDef)
	return o
}

// SetTagDef adds the tagDef to the delete payment tags params
func (o *DeletePaymentTagsParams) SetTagDef(tagDef []strfmt.UUID) {
	o.TagDef = tagDef
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if o.TagDef != nil {

		// binding items for tagDef
		joinedTagDef := o.bindParamTagDef(reg)

		// query array param tagDef
		if err := r.SetQueryParam("tagDef", joinedTagDef...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeletePaymentTags binds the parameter tagDef
func (o *DeletePaymentTagsParams) bindParamTagDef(formats strfmt.Registry) []string {
	tagDefIR := o.TagDef

	var tagDefIC []string
	for _, tagDefIIR := range tagDefIR { // explode []strfmt.UUID

		tagDefIIV := tagDefIIR.String() // strfmt.UUID as string
		tagDefIC = append(tagDefIC, tagDefIIV)
	}

	// items.CollectionFormat: "multi"
	tagDefIS := swag.JoinByFormat(tagDefIC, "multi")

	return tagDefIS
}
