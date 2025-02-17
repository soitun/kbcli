// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewGetInvoicesForAccountParams creates a new GetInvoicesForAccountParams object
// with the default values initialized.
func NewGetInvoicesForAccountParams() *GetInvoicesForAccountParams {
	var (
		auditDefault                    = string("NONE")
		includeInvoiceComponentsDefault = bool(false)
		includeVoidedInvoicesDefault    = bool(false)
		unpaidInvoicesOnlyDefault       = bool(false)
		withMigrationInvoicesDefault    = bool(false)
	)
	return &GetInvoicesForAccountParams{
		Audit:                    &auditDefault,
		IncludeInvoiceComponents: &includeInvoiceComponentsDefault,
		IncludeVoidedInvoices:    &includeVoidedInvoicesDefault,
		UnpaidInvoicesOnly:       &unpaidInvoicesOnlyDefault,
		WithMigrationInvoices:    &withMigrationInvoicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoicesForAccountParamsWithTimeout creates a new GetInvoicesForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoicesForAccountParamsWithTimeout(timeout time.Duration) *GetInvoicesForAccountParams {
	var (
		auditDefault                    = string("NONE")
		includeInvoiceComponentsDefault = bool(false)
		includeVoidedInvoicesDefault    = bool(false)
		unpaidInvoicesOnlyDefault       = bool(false)
		withMigrationInvoicesDefault    = bool(false)
	)
	return &GetInvoicesForAccountParams{
		Audit:                    &auditDefault,
		IncludeInvoiceComponents: &includeInvoiceComponentsDefault,
		IncludeVoidedInvoices:    &includeVoidedInvoicesDefault,
		UnpaidInvoicesOnly:       &unpaidInvoicesOnlyDefault,
		WithMigrationInvoices:    &withMigrationInvoicesDefault,

		timeout: timeout,
	}
}

// NewGetInvoicesForAccountParamsWithContext creates a new GetInvoicesForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoicesForAccountParamsWithContext(ctx context.Context) *GetInvoicesForAccountParams {
	var (
		auditDefault                    = string("NONE")
		includeInvoiceComponentsDefault = bool(false)
		includeVoidedInvoicesDefault    = bool(false)
		unpaidInvoicesOnlyDefault       = bool(false)
		withMigrationInvoicesDefault    = bool(false)
	)
	return &GetInvoicesForAccountParams{
		Audit:                    &auditDefault,
		IncludeInvoiceComponents: &includeInvoiceComponentsDefault,
		IncludeVoidedInvoices:    &includeVoidedInvoicesDefault,
		UnpaidInvoicesOnly:       &unpaidInvoicesOnlyDefault,
		WithMigrationInvoices:    &withMigrationInvoicesDefault,

		Context: ctx,
	}
}

// NewGetInvoicesForAccountParamsWithHTTPClient creates a new GetInvoicesForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoicesForAccountParamsWithHTTPClient(client *http.Client) *GetInvoicesForAccountParams {
	var (
		auditDefault                    = string("NONE")
		includeInvoiceComponentsDefault = bool(false)
		includeVoidedInvoicesDefault    = bool(false)
		unpaidInvoicesOnlyDefault       = bool(false)
		withMigrationInvoicesDefault    = bool(false)
	)
	return &GetInvoicesForAccountParams{
		Audit:                    &auditDefault,
		IncludeInvoiceComponents: &includeInvoiceComponentsDefault,
		IncludeVoidedInvoices:    &includeVoidedInvoicesDefault,
		UnpaidInvoicesOnly:       &unpaidInvoicesOnlyDefault,
		WithMigrationInvoices:    &withMigrationInvoicesDefault,
		HTTPClient:               client,
	}
}

/*GetInvoicesForAccountParams contains all the parameters to send to the API endpoint
for the get invoices for account operation typically these are written to a http.Request
*/
type GetInvoicesForAccountParams struct {

	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string
	/*EndDate*/
	EndDate *strfmt.Date
	/*IncludeInvoiceComponents*/
	IncludeInvoiceComponents *bool
	/*IncludeVoidedInvoices*/
	IncludeVoidedInvoices *bool
	/*InvoicesFilter*/
	InvoicesFilter *string
	/*StartDate*/
	StartDate *strfmt.Date
	/*UnpaidInvoicesOnly*/
	UnpaidInvoicesOnly *bool
	/*WithMigrationInvoices*/
	WithMigrationInvoices *bool

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithTimeout(timeout time.Duration) *GetInvoicesForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithContext(ctx context.Context) *GetInvoicesForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithHTTPClient(client *http.Client) *GetInvoicesForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithAccountID(accountID strfmt.UUID) *GetInvoicesForAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithAudit(audit *string) *GetInvoicesForAccountParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithEndDate adds the endDate to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithEndDate(endDate *strfmt.Date) *GetInvoicesForAccountParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithIncludeInvoiceComponents adds the includeInvoiceComponents to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithIncludeInvoiceComponents(includeInvoiceComponents *bool) *GetInvoicesForAccountParams {
	o.SetIncludeInvoiceComponents(includeInvoiceComponents)
	return o
}

// SetIncludeInvoiceComponents adds the includeInvoiceComponents to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetIncludeInvoiceComponents(includeInvoiceComponents *bool) {
	o.IncludeInvoiceComponents = includeInvoiceComponents
}

// WithIncludeVoidedInvoices adds the includeVoidedInvoices to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithIncludeVoidedInvoices(includeVoidedInvoices *bool) *GetInvoicesForAccountParams {
	o.SetIncludeVoidedInvoices(includeVoidedInvoices)
	return o
}

// SetIncludeVoidedInvoices adds the includeVoidedInvoices to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetIncludeVoidedInvoices(includeVoidedInvoices *bool) {
	o.IncludeVoidedInvoices = includeVoidedInvoices
}

// WithInvoicesFilter adds the invoicesFilter to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithInvoicesFilter(invoicesFilter *string) *GetInvoicesForAccountParams {
	o.SetInvoicesFilter(invoicesFilter)
	return o
}

// SetInvoicesFilter adds the invoicesFilter to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetInvoicesFilter(invoicesFilter *string) {
	o.InvoicesFilter = invoicesFilter
}

// WithStartDate adds the startDate to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithStartDate(startDate *strfmt.Date) *GetInvoicesForAccountParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithUnpaidInvoicesOnly adds the unpaidInvoicesOnly to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithUnpaidInvoicesOnly(unpaidInvoicesOnly *bool) *GetInvoicesForAccountParams {
	o.SetUnpaidInvoicesOnly(unpaidInvoicesOnly)
	return o
}

// SetUnpaidInvoicesOnly adds the unpaidInvoicesOnly to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetUnpaidInvoicesOnly(unpaidInvoicesOnly *bool) {
	o.UnpaidInvoicesOnly = unpaidInvoicesOnly
}

// WithWithMigrationInvoices adds the withMigrationInvoices to the get invoices for account params
func (o *GetInvoicesForAccountParams) WithWithMigrationInvoices(withMigrationInvoices *bool) *GetInvoicesForAccountParams {
	o.SetWithMigrationInvoices(withMigrationInvoices)
	return o
}

// SetWithMigrationInvoices adds the withMigrationInvoices to the get invoices for account params
func (o *GetInvoicesForAccountParams) SetWithMigrationInvoices(withMigrationInvoices *bool) {
	o.WithMigrationInvoices = withMigrationInvoices
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoicesForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

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

	if o.IncludeInvoiceComponents != nil {

		// query param includeInvoiceComponents
		var qrIncludeInvoiceComponents bool
		if o.IncludeInvoiceComponents != nil {
			qrIncludeInvoiceComponents = *o.IncludeInvoiceComponents
		}
		qIncludeInvoiceComponents := swag.FormatBool(qrIncludeInvoiceComponents)
		if qIncludeInvoiceComponents != "" {
			if err := r.SetQueryParam("includeInvoiceComponents", qIncludeInvoiceComponents); err != nil {
				return err
			}
		}

	}

	if o.IncludeVoidedInvoices != nil {

		// query param includeVoidedInvoices
		var qrIncludeVoidedInvoices bool
		if o.IncludeVoidedInvoices != nil {
			qrIncludeVoidedInvoices = *o.IncludeVoidedInvoices
		}
		qIncludeVoidedInvoices := swag.FormatBool(qrIncludeVoidedInvoices)
		if qIncludeVoidedInvoices != "" {
			if err := r.SetQueryParam("includeVoidedInvoices", qIncludeVoidedInvoices); err != nil {
				return err
			}
		}

	}

	if o.InvoicesFilter != nil {

		// query param invoicesFilter
		var qrInvoicesFilter string
		if o.InvoicesFilter != nil {
			qrInvoicesFilter = *o.InvoicesFilter
		}
		qInvoicesFilter := qrInvoicesFilter
		if qInvoicesFilter != "" {
			if err := r.SetQueryParam("invoicesFilter", qInvoicesFilter); err != nil {
				return err
			}
		}

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

	if o.UnpaidInvoicesOnly != nil {

		// query param unpaidInvoicesOnly
		var qrUnpaidInvoicesOnly bool
		if o.UnpaidInvoicesOnly != nil {
			qrUnpaidInvoicesOnly = *o.UnpaidInvoicesOnly
		}
		qUnpaidInvoicesOnly := swag.FormatBool(qrUnpaidInvoicesOnly)
		if qUnpaidInvoicesOnly != "" {
			if err := r.SetQueryParam("unpaidInvoicesOnly", qUnpaidInvoicesOnly); err != nil {
				return err
			}
		}

	}

	if o.WithMigrationInvoices != nil {

		// query param withMigrationInvoices
		var qrWithMigrationInvoices bool
		if o.WithMigrationInvoices != nil {
			qrWithMigrationInvoices = *o.WithMigrationInvoices
		}
		qWithMigrationInvoices := swag.FormatBool(qrWithMigrationInvoices)
		if qWithMigrationInvoices != "" {
			if err := r.SetQueryParam("withMigrationInvoices", qWithMigrationInvoices); err != nil {
				return err
			}
		}

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
