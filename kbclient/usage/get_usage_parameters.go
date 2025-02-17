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

// NewGetUsageParams creates a new GetUsageParams object
// with the default values initialized.
func NewGetUsageParams() *GetUsageParams {
	var ()
	return &GetUsageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsageParamsWithTimeout creates a new GetUsageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsageParamsWithTimeout(timeout time.Duration) *GetUsageParams {
	var ()
	return &GetUsageParams{

		timeout: timeout,
	}
}

// NewGetUsageParamsWithContext creates a new GetUsageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsageParamsWithContext(ctx context.Context) *GetUsageParams {
	var ()
	return &GetUsageParams{

		Context: ctx,
	}
}

// NewGetUsageParamsWithHTTPClient creates a new GetUsageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsageParamsWithHTTPClient(client *http.Client) *GetUsageParams {
	var ()
	return &GetUsageParams{
		HTTPClient: client,
	}
}

/*GetUsageParams contains all the parameters to send to the API endpoint
for the get usage operation typically these are written to a http.Request
*/
type GetUsageParams struct {

	/*EndDate*/
	EndDate *strfmt.Date
	/*PluginProperty*/
	PluginProperty []string
	/*StartDate*/
	StartDate *strfmt.Date
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID
	/*UnitType*/
	UnitType string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get usage params
func (o *GetUsageParams) WithTimeout(timeout time.Duration) *GetUsageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get usage params
func (o *GetUsageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get usage params
func (o *GetUsageParams) WithContext(ctx context.Context) *GetUsageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get usage params
func (o *GetUsageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get usage params
func (o *GetUsageParams) WithHTTPClient(client *http.Client) *GetUsageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get usage params
func (o *GetUsageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndDate adds the endDate to the get usage params
func (o *GetUsageParams) WithEndDate(endDate *strfmt.Date) *GetUsageParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get usage params
func (o *GetUsageParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithPluginProperty adds the pluginProperty to the get usage params
func (o *GetUsageParams) WithPluginProperty(pluginProperty []string) *GetUsageParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get usage params
func (o *GetUsageParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithStartDate adds the startDate to the get usage params
func (o *GetUsageParams) WithStartDate(startDate *strfmt.Date) *GetUsageParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get usage params
func (o *GetUsageParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithSubscriptionID adds the subscriptionID to the get usage params
func (o *GetUsageParams) WithSubscriptionID(subscriptionID strfmt.UUID) *GetUsageParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the get usage params
func (o *GetUsageParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WithUnitType adds the unitType to the get usage params
func (o *GetUsageParams) WithUnitType(unitType string) *GetUsageParams {
	o.SetUnitType(unitType)
	return o
}

// SetUnitType adds the unitType to the get usage params
func (o *GetUsageParams) SetUnitType(unitType string) {
	o.UnitType = unitType
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param unitType
	if err := r.SetPathParam("unitType", o.UnitType); err != nil {
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
