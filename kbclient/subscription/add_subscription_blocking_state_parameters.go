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

// NewAddSubscriptionBlockingStateParams creates a new AddSubscriptionBlockingStateParams object
// with the default values initialized.
func NewAddSubscriptionBlockingStateParams() *AddSubscriptionBlockingStateParams {
	var ()
	return &AddSubscriptionBlockingStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddSubscriptionBlockingStateParamsWithTimeout creates a new AddSubscriptionBlockingStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddSubscriptionBlockingStateParamsWithTimeout(timeout time.Duration) *AddSubscriptionBlockingStateParams {
	var ()
	return &AddSubscriptionBlockingStateParams{

		timeout: timeout,
	}
}

// NewAddSubscriptionBlockingStateParamsWithContext creates a new AddSubscriptionBlockingStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddSubscriptionBlockingStateParamsWithContext(ctx context.Context) *AddSubscriptionBlockingStateParams {
	var ()
	return &AddSubscriptionBlockingStateParams{

		Context: ctx,
	}
}

// NewAddSubscriptionBlockingStateParamsWithHTTPClient creates a new AddSubscriptionBlockingStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddSubscriptionBlockingStateParamsWithHTTPClient(client *http.Client) *AddSubscriptionBlockingStateParams {
	var ()
	return &AddSubscriptionBlockingStateParams{
		HTTPClient: client,
	}
}

/*AddSubscriptionBlockingStateParams contains all the parameters to send to the API endpoint
for the add subscription blocking state operation typically these are written to a http.Request
*/
type AddSubscriptionBlockingStateParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.BlockingState
	/*PluginProperty*/
	PluginProperty []string
	/*RequestedDate*/
	RequestedDate *strfmt.Date
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithTimeout(timeout time.Duration) *AddSubscriptionBlockingStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithContext(ctx context.Context) *AddSubscriptionBlockingStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithHTTPClient(client *http.Client) *AddSubscriptionBlockingStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithXKillbillComment(xKillbillComment *string) *AddSubscriptionBlockingStateParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddSubscriptionBlockingStateParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithXKillbillReason(xKillbillReason *string) *AddSubscriptionBlockingStateParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithBody(body *kbmodel.BlockingState) *AddSubscriptionBlockingStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetBody(body *kbmodel.BlockingState) {
	o.Body = body
}

// WithPluginProperty adds the pluginProperty to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithPluginProperty(pluginProperty []string) *AddSubscriptionBlockingStateParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithRequestedDate(requestedDate *strfmt.Date) *AddSubscriptionBlockingStateParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WithSubscriptionID adds the subscriptionID to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) WithSubscriptionID(subscriptionID strfmt.UUID) *AddSubscriptionBlockingStateParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the add subscription blocking state params
func (o *AddSubscriptionBlockingStateParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddSubscriptionBlockingStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
