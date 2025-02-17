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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// NewUpdateSubscriptionQuantityParams creates a new UpdateSubscriptionQuantityParams object
// with the default values initialized.
func NewUpdateSubscriptionQuantityParams() *UpdateSubscriptionQuantityParams {
	var (
		forceNewQuantityWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionQuantityParams{
		ForceNewQuantityWithPastEffectiveDate: &forceNewQuantityWithPastEffectiveDateDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSubscriptionQuantityParamsWithTimeout creates a new UpdateSubscriptionQuantityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSubscriptionQuantityParamsWithTimeout(timeout time.Duration) *UpdateSubscriptionQuantityParams {
	var (
		forceNewQuantityWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionQuantityParams{
		ForceNewQuantityWithPastEffectiveDate: &forceNewQuantityWithPastEffectiveDateDefault,

		timeout: timeout,
	}
}

// NewUpdateSubscriptionQuantityParamsWithContext creates a new UpdateSubscriptionQuantityParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSubscriptionQuantityParamsWithContext(ctx context.Context) *UpdateSubscriptionQuantityParams {
	var (
		forceNewQuantityWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionQuantityParams{
		ForceNewQuantityWithPastEffectiveDate: &forceNewQuantityWithPastEffectiveDateDefault,

		Context: ctx,
	}
}

// NewUpdateSubscriptionQuantityParamsWithHTTPClient creates a new UpdateSubscriptionQuantityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSubscriptionQuantityParamsWithHTTPClient(client *http.Client) *UpdateSubscriptionQuantityParams {
	var (
		forceNewQuantityWithPastEffectiveDateDefault = bool(false)
	)
	return &UpdateSubscriptionQuantityParams{
		ForceNewQuantityWithPastEffectiveDate: &forceNewQuantityWithPastEffectiveDateDefault,
		HTTPClient:                            client,
	}
}

/*UpdateSubscriptionQuantityParams contains all the parameters to send to the API endpoint
for the update subscription quantity operation typically these are written to a http.Request
*/
type UpdateSubscriptionQuantityParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.Subscription
	/*EffectiveFromDate*/
	EffectiveFromDate *strfmt.Date
	/*ForceNewQuantityWithPastEffectiveDate*/
	ForceNewQuantityWithPastEffectiveDate *bool
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithTimeout(timeout time.Duration) *UpdateSubscriptionQuantityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithContext(ctx context.Context) *UpdateSubscriptionQuantityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithHTTPClient(client *http.Client) *UpdateSubscriptionQuantityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithXKillbillComment(xKillbillComment *string) *UpdateSubscriptionQuantityParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UpdateSubscriptionQuantityParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithXKillbillReason(xKillbillReason *string) *UpdateSubscriptionQuantityParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithBody(body *kbmodel.Subscription) *UpdateSubscriptionQuantityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetBody(body *kbmodel.Subscription) {
	o.Body = body
}

// WithEffectiveFromDate adds the effectiveFromDate to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithEffectiveFromDate(effectiveFromDate *strfmt.Date) *UpdateSubscriptionQuantityParams {
	o.SetEffectiveFromDate(effectiveFromDate)
	return o
}

// SetEffectiveFromDate adds the effectiveFromDate to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetEffectiveFromDate(effectiveFromDate *strfmt.Date) {
	o.EffectiveFromDate = effectiveFromDate
}

// WithForceNewQuantityWithPastEffectiveDate adds the forceNewQuantityWithPastEffectiveDate to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithForceNewQuantityWithPastEffectiveDate(forceNewQuantityWithPastEffectiveDate *bool) *UpdateSubscriptionQuantityParams {
	o.SetForceNewQuantityWithPastEffectiveDate(forceNewQuantityWithPastEffectiveDate)
	return o
}

// SetForceNewQuantityWithPastEffectiveDate adds the forceNewQuantityWithPastEffectiveDate to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetForceNewQuantityWithPastEffectiveDate(forceNewQuantityWithPastEffectiveDate *bool) {
	o.ForceNewQuantityWithPastEffectiveDate = forceNewQuantityWithPastEffectiveDate
}

// WithSubscriptionID adds the subscriptionID to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) WithSubscriptionID(subscriptionID strfmt.UUID) *UpdateSubscriptionQuantityParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the update subscription quantity params
func (o *UpdateSubscriptionQuantityParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSubscriptionQuantityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EffectiveFromDate != nil {

		// query param effectiveFromDate
		var qrEffectiveFromDate strfmt.Date
		if o.EffectiveFromDate != nil {
			qrEffectiveFromDate = *o.EffectiveFromDate
		}
		qEffectiveFromDate := qrEffectiveFromDate.String()
		if qEffectiveFromDate != "" {
			if err := r.SetQueryParam("effectiveFromDate", qEffectiveFromDate); err != nil {
				return err
			}
		}

	}

	if o.ForceNewQuantityWithPastEffectiveDate != nil {

		// query param forceNewQuantityWithPastEffectiveDate
		var qrForceNewQuantityWithPastEffectiveDate bool
		if o.ForceNewQuantityWithPastEffectiveDate != nil {
			qrForceNewQuantityWithPastEffectiveDate = *o.ForceNewQuantityWithPastEffectiveDate
		}
		qForceNewQuantityWithPastEffectiveDate := swag.FormatBool(qrForceNewQuantityWithPastEffectiveDate)
		if qForceNewQuantityWithPastEffectiveDate != "" {
			if err := r.SetQueryParam("forceNewQuantityWithPastEffectiveDate", qForceNewQuantityWithPastEffectiveDate); err != nil {
				return err
			}
		}

	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
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
