// Code generated by go-swagger; DO NOT EDIT.

package usage

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
)

// NewGetAllUsageParams creates a new GetAllUsageParams object
// with the default values initialized.
func NewGetAllUsageParams() *GetAllUsageParams {
	var ()
	return &GetAllUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllUsageParamsWithTimeout creates a new GetAllUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllUsageParamsWithTimeout(timeout time.Duration) *GetAllUsageParams {
	var ()
	return &GetAllUsageParams{

		timeout: timeout,
	}
}

// NewGetAllUsageParamsWithContext creates a new GetAllUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllUsageParamsWithContext(ctx context.Context) *GetAllUsageParams {
	var ()
	return &GetAllUsageParams{

		Context: ctx,
	}
}

// NewGetAllUsageParamsWithHTTPClient creates a new GetAllUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllUsageParamsWithHTTPClient(client *http.Client) *GetAllUsageParams {
	var ()
	return &GetAllUsageParams{
		HTTPClient: client,
	}
}

/*GetAllUsageParams contains all the parameters to send to the API endpoint
for the get all usage operation typically these are written to a http.Request
*/
type GetAllUsageParams struct {

	/*EndDate*/
	EndDate *strfmt.Date
	/*PluginProperty*/
	PluginProperty []string
	/*StartDate*/
	StartDate *strfmt.Date
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get all usage params
func (o *GetAllUsageParams) WithTimeout(timeout time.Duration) *GetAllUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all usage params
func (o *GetAllUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all usage params
func (o *GetAllUsageParams) WithContext(ctx context.Context) *GetAllUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all usage params
func (o *GetAllUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all usage params
func (o *GetAllUsageParams) WithHTTPClient(client *http.Client) *GetAllUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all usage params
func (o *GetAllUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get all usage params
func (o *GetAllUsageParams) WithEndDate(endDate *strfmt.Date) *GetAllUsageParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get all usage params
func (o *GetAllUsageParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithPluginProperty adds the pluginProperty to the get all usage params
func (o *GetAllUsageParams) WithPluginProperty(pluginProperty []string) *GetAllUsageParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get all usage params
func (o *GetAllUsageParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithStartDate adds the startDate to the get all usage params
func (o *GetAllUsageParams) WithStartDate(startDate *strfmt.Date) *GetAllUsageParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get all usage params
func (o *GetAllUsageParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithSubscriptionID adds the subscriptionID to the get all usage params
func (o *GetAllUsageParams) WithSubscriptionID(subscriptionID strfmt.UUID) *GetAllUsageParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the get all usage params
func (o *GetAllUsageParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}

	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
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
