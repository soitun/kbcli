// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewUndoChangeSubscriptionPlanParams creates a new UndoChangeSubscriptionPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUndoChangeSubscriptionPlanParams() *UndoChangeSubscriptionPlanParams {
	return &UndoChangeSubscriptionPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUndoChangeSubscriptionPlanParamsWithTimeout creates a new UndoChangeSubscriptionPlanParams object
// with the ability to set a timeout on a request.
func NewUndoChangeSubscriptionPlanParamsWithTimeout(timeout time.Duration) *UndoChangeSubscriptionPlanParams {
	return &UndoChangeSubscriptionPlanParams{
		timeout: timeout,
	}
}

// NewUndoChangeSubscriptionPlanParamsWithContext creates a new UndoChangeSubscriptionPlanParams object
// with the ability to set a context for a request.
func NewUndoChangeSubscriptionPlanParamsWithContext(ctx context.Context) *UndoChangeSubscriptionPlanParams {
	return &UndoChangeSubscriptionPlanParams{
		Context: ctx,
	}
}

// NewUndoChangeSubscriptionPlanParamsWithHTTPClient creates a new UndoChangeSubscriptionPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewUndoChangeSubscriptionPlanParamsWithHTTPClient(client *http.Client) *UndoChangeSubscriptionPlanParams {
	return &UndoChangeSubscriptionPlanParams{
		HTTPClient: client,
	}
}

/*
UndoChangeSubscriptionPlanParams contains all the parameters to send to the API endpoint

	for the undo change subscription plan operation.

	Typically these are written to a http.Request.
*/
type UndoChangeSubscriptionPlanParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// PluginProperty.
	PluginProperty []string

	// SubscriptionID.
	//
	// Format: uuid
	SubscriptionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the undo change subscription plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UndoChangeSubscriptionPlanParams) WithDefaults() *UndoChangeSubscriptionPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the undo change subscription plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UndoChangeSubscriptionPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithTimeout(timeout time.Duration) *UndoChangeSubscriptionPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithContext(ctx context.Context) *UndoChangeSubscriptionPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithHTTPClient(client *http.Client) *UndoChangeSubscriptionPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithXKillbillComment(xKillbillComment *string) *UndoChangeSubscriptionPlanParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UndoChangeSubscriptionPlanParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithXKillbillReason(xKillbillReason *string) *UndoChangeSubscriptionPlanParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPluginProperty adds the pluginProperty to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithPluginProperty(pluginProperty []string) *UndoChangeSubscriptionPlanParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithSubscriptionID adds the subscriptionID to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) WithSubscriptionID(subscriptionID strfmt.UUID) *UndoChangeSubscriptionPlanParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the undo change subscription plan params
func (o *UndoChangeSubscriptionPlanParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *UndoChangeSubscriptionPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamUndoChangeSubscriptionPlan binds the parameter pluginProperty
func (o *UndoChangeSubscriptionPlanParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
