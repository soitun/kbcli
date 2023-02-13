// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvoiceItem invoice item
//
// swagger:model InvoiceItem
type InvoiceItem struct {

	// account Id
	// Required: true
	// Format: uuid
	AccountID *strfmt.UUID `json:"accountId"`

	// amount
	Amount float64 `json:"amount,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// bundle Id
	// Format: uuid
	BundleID strfmt.UUID `json:"bundleId,omitempty"`

	// catalog effective date
	// Format: date-time
	CatalogEffectiveDate strfmt.DateTime `json:"catalogEffectiveDate,omitempty"`

	// child account Id
	// Format: uuid
	ChildAccountID strfmt.UUID `json:"childAccountId,omitempty"`

	// child items
	ChildItems []*InvoiceItem `json:"childItems"`

	// currency
	// Enum: [AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYR BZD CAD CDF CHF CLP CNY COP CRC CUC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GGP GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS IMP INR IQD IRR ISK JEP JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SPL SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TVD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XDR XOF XPF YER ZAR ZMW ZWD BTC]
	Currency string `json:"currency,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// end date
	// Format: date
	EndDate strfmt.Date `json:"endDate,omitempty"`

	// invoice Id
	// Format: uuid
	InvoiceID strfmt.UUID `json:"invoiceId,omitempty"`

	// invoice item Id
	// Required: true
	// Format: uuid
	InvoiceItemID *strfmt.UUID `json:"invoiceItemId"`

	// item details
	ItemDetails string `json:"itemDetails,omitempty"`

	// item type
	// Enum: [EXTERNAL_CHARGE FIXED RECURRING REPAIR_ADJ CBA_ADJ CREDIT_ADJ ITEM_ADJ USAGE TAX PARENT_SUMMARY]
	ItemType string `json:"itemType,omitempty"`

	// linked invoice item Id
	// Format: uuid
	LinkedInvoiceItemID strfmt.UUID `json:"linkedInvoiceItemId,omitempty"`

	// phase name
	PhaseName string `json:"phaseName,omitempty"`

	// plan name
	PlanName string `json:"planName,omitempty"`

	// pretty phase name
	PrettyPhaseName string `json:"prettyPhaseName,omitempty"`

	// pretty plan name
	PrettyPlanName string `json:"prettyPlanName,omitempty"`

	// pretty product name
	PrettyProductName string `json:"prettyProductName,omitempty"`

	// pretty usage name
	PrettyUsageName string `json:"prettyUsageName,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// quantity
	Quantity float64 `json:"quantity,omitempty"`

	// rate
	Rate float64 `json:"rate,omitempty"`

	// start date
	// Format: date
	StartDate strfmt.Date `json:"startDate,omitempty"`

	// subscription Id
	// Format: uuid
	SubscriptionID strfmt.UUID `json:"subscriptionId,omitempty"`

	// usage name
	UsageName string `json:"usageName,omitempty"`
}

// Validate validates this invoice item
func (m *InvoiceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCatalogEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChildItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItemType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinkedInvoiceItemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceItem) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceItem) validateBundleID(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleID) { // not required
		return nil
	}

	if err := validate.FormatOf("bundleId", "body", "uuid", m.BundleID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateCatalogEffectiveDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogEffectiveDate) { // not required
		return nil
	}

	if err := validate.FormatOf("catalogEffectiveDate", "body", "date-time", m.CatalogEffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateChildAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.ChildAccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("childAccountId", "body", "uuid", m.ChildAccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateChildItems(formats strfmt.Registry) error {
	if swag.IsZero(m.ChildItems) { // not required
		return nil
	}

	for i := 0; i < len(m.ChildItems); i++ {
		if swag.IsZero(m.ChildItems[i]) { // not required
			continue
		}

		if m.ChildItems[i] != nil {
			if err := m.ChildItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var invoiceItemTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceItemTypeCurrencyPropEnum = append(invoiceItemTypeCurrencyPropEnum, v)
	}
}

const (

	// InvoiceItemCurrencyAED captures enum value "AED"
	InvoiceItemCurrencyAED string = "AED"

	// InvoiceItemCurrencyAFN captures enum value "AFN"
	InvoiceItemCurrencyAFN string = "AFN"

	// InvoiceItemCurrencyALL captures enum value "ALL"
	InvoiceItemCurrencyALL string = "ALL"

	// InvoiceItemCurrencyAMD captures enum value "AMD"
	InvoiceItemCurrencyAMD string = "AMD"

	// InvoiceItemCurrencyANG captures enum value "ANG"
	InvoiceItemCurrencyANG string = "ANG"

	// InvoiceItemCurrencyAOA captures enum value "AOA"
	InvoiceItemCurrencyAOA string = "AOA"

	// InvoiceItemCurrencyARS captures enum value "ARS"
	InvoiceItemCurrencyARS string = "ARS"

	// InvoiceItemCurrencyAUD captures enum value "AUD"
	InvoiceItemCurrencyAUD string = "AUD"

	// InvoiceItemCurrencyAWG captures enum value "AWG"
	InvoiceItemCurrencyAWG string = "AWG"

	// InvoiceItemCurrencyAZN captures enum value "AZN"
	InvoiceItemCurrencyAZN string = "AZN"

	// InvoiceItemCurrencyBAM captures enum value "BAM"
	InvoiceItemCurrencyBAM string = "BAM"

	// InvoiceItemCurrencyBBD captures enum value "BBD"
	InvoiceItemCurrencyBBD string = "BBD"

	// InvoiceItemCurrencyBDT captures enum value "BDT"
	InvoiceItemCurrencyBDT string = "BDT"

	// InvoiceItemCurrencyBGN captures enum value "BGN"
	InvoiceItemCurrencyBGN string = "BGN"

	// InvoiceItemCurrencyBHD captures enum value "BHD"
	InvoiceItemCurrencyBHD string = "BHD"

	// InvoiceItemCurrencyBIF captures enum value "BIF"
	InvoiceItemCurrencyBIF string = "BIF"

	// InvoiceItemCurrencyBMD captures enum value "BMD"
	InvoiceItemCurrencyBMD string = "BMD"

	// InvoiceItemCurrencyBND captures enum value "BND"
	InvoiceItemCurrencyBND string = "BND"

	// InvoiceItemCurrencyBOB captures enum value "BOB"
	InvoiceItemCurrencyBOB string = "BOB"

	// InvoiceItemCurrencyBRL captures enum value "BRL"
	InvoiceItemCurrencyBRL string = "BRL"

	// InvoiceItemCurrencyBSD captures enum value "BSD"
	InvoiceItemCurrencyBSD string = "BSD"

	// InvoiceItemCurrencyBTN captures enum value "BTN"
	InvoiceItemCurrencyBTN string = "BTN"

	// InvoiceItemCurrencyBWP captures enum value "BWP"
	InvoiceItemCurrencyBWP string = "BWP"

	// InvoiceItemCurrencyBYR captures enum value "BYR"
	InvoiceItemCurrencyBYR string = "BYR"

	// InvoiceItemCurrencyBZD captures enum value "BZD"
	InvoiceItemCurrencyBZD string = "BZD"

	// InvoiceItemCurrencyCAD captures enum value "CAD"
	InvoiceItemCurrencyCAD string = "CAD"

	// InvoiceItemCurrencyCDF captures enum value "CDF"
	InvoiceItemCurrencyCDF string = "CDF"

	// InvoiceItemCurrencyCHF captures enum value "CHF"
	InvoiceItemCurrencyCHF string = "CHF"

	// InvoiceItemCurrencyCLP captures enum value "CLP"
	InvoiceItemCurrencyCLP string = "CLP"

	// InvoiceItemCurrencyCNY captures enum value "CNY"
	InvoiceItemCurrencyCNY string = "CNY"

	// InvoiceItemCurrencyCOP captures enum value "COP"
	InvoiceItemCurrencyCOP string = "COP"

	// InvoiceItemCurrencyCRC captures enum value "CRC"
	InvoiceItemCurrencyCRC string = "CRC"

	// InvoiceItemCurrencyCUC captures enum value "CUC"
	InvoiceItemCurrencyCUC string = "CUC"

	// InvoiceItemCurrencyCUP captures enum value "CUP"
	InvoiceItemCurrencyCUP string = "CUP"

	// InvoiceItemCurrencyCVE captures enum value "CVE"
	InvoiceItemCurrencyCVE string = "CVE"

	// InvoiceItemCurrencyCZK captures enum value "CZK"
	InvoiceItemCurrencyCZK string = "CZK"

	// InvoiceItemCurrencyDJF captures enum value "DJF"
	InvoiceItemCurrencyDJF string = "DJF"

	// InvoiceItemCurrencyDKK captures enum value "DKK"
	InvoiceItemCurrencyDKK string = "DKK"

	// InvoiceItemCurrencyDOP captures enum value "DOP"
	InvoiceItemCurrencyDOP string = "DOP"

	// InvoiceItemCurrencyDZD captures enum value "DZD"
	InvoiceItemCurrencyDZD string = "DZD"

	// InvoiceItemCurrencyEGP captures enum value "EGP"
	InvoiceItemCurrencyEGP string = "EGP"

	// InvoiceItemCurrencyERN captures enum value "ERN"
	InvoiceItemCurrencyERN string = "ERN"

	// InvoiceItemCurrencyETB captures enum value "ETB"
	InvoiceItemCurrencyETB string = "ETB"

	// InvoiceItemCurrencyEUR captures enum value "EUR"
	InvoiceItemCurrencyEUR string = "EUR"

	// InvoiceItemCurrencyFJD captures enum value "FJD"
	InvoiceItemCurrencyFJD string = "FJD"

	// InvoiceItemCurrencyFKP captures enum value "FKP"
	InvoiceItemCurrencyFKP string = "FKP"

	// InvoiceItemCurrencyGBP captures enum value "GBP"
	InvoiceItemCurrencyGBP string = "GBP"

	// InvoiceItemCurrencyGEL captures enum value "GEL"
	InvoiceItemCurrencyGEL string = "GEL"

	// InvoiceItemCurrencyGGP captures enum value "GGP"
	InvoiceItemCurrencyGGP string = "GGP"

	// InvoiceItemCurrencyGHS captures enum value "GHS"
	InvoiceItemCurrencyGHS string = "GHS"

	// InvoiceItemCurrencyGIP captures enum value "GIP"
	InvoiceItemCurrencyGIP string = "GIP"

	// InvoiceItemCurrencyGMD captures enum value "GMD"
	InvoiceItemCurrencyGMD string = "GMD"

	// InvoiceItemCurrencyGNF captures enum value "GNF"
	InvoiceItemCurrencyGNF string = "GNF"

	// InvoiceItemCurrencyGTQ captures enum value "GTQ"
	InvoiceItemCurrencyGTQ string = "GTQ"

	// InvoiceItemCurrencyGYD captures enum value "GYD"
	InvoiceItemCurrencyGYD string = "GYD"

	// InvoiceItemCurrencyHKD captures enum value "HKD"
	InvoiceItemCurrencyHKD string = "HKD"

	// InvoiceItemCurrencyHNL captures enum value "HNL"
	InvoiceItemCurrencyHNL string = "HNL"

	// InvoiceItemCurrencyHRK captures enum value "HRK"
	InvoiceItemCurrencyHRK string = "HRK"

	// InvoiceItemCurrencyHTG captures enum value "HTG"
	InvoiceItemCurrencyHTG string = "HTG"

	// InvoiceItemCurrencyHUF captures enum value "HUF"
	InvoiceItemCurrencyHUF string = "HUF"

	// InvoiceItemCurrencyIDR captures enum value "IDR"
	InvoiceItemCurrencyIDR string = "IDR"

	// InvoiceItemCurrencyILS captures enum value "ILS"
	InvoiceItemCurrencyILS string = "ILS"

	// InvoiceItemCurrencyIMP captures enum value "IMP"
	InvoiceItemCurrencyIMP string = "IMP"

	// InvoiceItemCurrencyINR captures enum value "INR"
	InvoiceItemCurrencyINR string = "INR"

	// InvoiceItemCurrencyIQD captures enum value "IQD"
	InvoiceItemCurrencyIQD string = "IQD"

	// InvoiceItemCurrencyIRR captures enum value "IRR"
	InvoiceItemCurrencyIRR string = "IRR"

	// InvoiceItemCurrencyISK captures enum value "ISK"
	InvoiceItemCurrencyISK string = "ISK"

	// InvoiceItemCurrencyJEP captures enum value "JEP"
	InvoiceItemCurrencyJEP string = "JEP"

	// InvoiceItemCurrencyJMD captures enum value "JMD"
	InvoiceItemCurrencyJMD string = "JMD"

	// InvoiceItemCurrencyJOD captures enum value "JOD"
	InvoiceItemCurrencyJOD string = "JOD"

	// InvoiceItemCurrencyJPY captures enum value "JPY"
	InvoiceItemCurrencyJPY string = "JPY"

	// InvoiceItemCurrencyKES captures enum value "KES"
	InvoiceItemCurrencyKES string = "KES"

	// InvoiceItemCurrencyKGS captures enum value "KGS"
	InvoiceItemCurrencyKGS string = "KGS"

	// InvoiceItemCurrencyKHR captures enum value "KHR"
	InvoiceItemCurrencyKHR string = "KHR"

	// InvoiceItemCurrencyKMF captures enum value "KMF"
	InvoiceItemCurrencyKMF string = "KMF"

	// InvoiceItemCurrencyKPW captures enum value "KPW"
	InvoiceItemCurrencyKPW string = "KPW"

	// InvoiceItemCurrencyKRW captures enum value "KRW"
	InvoiceItemCurrencyKRW string = "KRW"

	// InvoiceItemCurrencyKWD captures enum value "KWD"
	InvoiceItemCurrencyKWD string = "KWD"

	// InvoiceItemCurrencyKYD captures enum value "KYD"
	InvoiceItemCurrencyKYD string = "KYD"

	// InvoiceItemCurrencyKZT captures enum value "KZT"
	InvoiceItemCurrencyKZT string = "KZT"

	// InvoiceItemCurrencyLAK captures enum value "LAK"
	InvoiceItemCurrencyLAK string = "LAK"

	// InvoiceItemCurrencyLBP captures enum value "LBP"
	InvoiceItemCurrencyLBP string = "LBP"

	// InvoiceItemCurrencyLKR captures enum value "LKR"
	InvoiceItemCurrencyLKR string = "LKR"

	// InvoiceItemCurrencyLRD captures enum value "LRD"
	InvoiceItemCurrencyLRD string = "LRD"

	// InvoiceItemCurrencyLSL captures enum value "LSL"
	InvoiceItemCurrencyLSL string = "LSL"

	// InvoiceItemCurrencyLTL captures enum value "LTL"
	InvoiceItemCurrencyLTL string = "LTL"

	// InvoiceItemCurrencyLVL captures enum value "LVL"
	InvoiceItemCurrencyLVL string = "LVL"

	// InvoiceItemCurrencyLYD captures enum value "LYD"
	InvoiceItemCurrencyLYD string = "LYD"

	// InvoiceItemCurrencyMAD captures enum value "MAD"
	InvoiceItemCurrencyMAD string = "MAD"

	// InvoiceItemCurrencyMDL captures enum value "MDL"
	InvoiceItemCurrencyMDL string = "MDL"

	// InvoiceItemCurrencyMGA captures enum value "MGA"
	InvoiceItemCurrencyMGA string = "MGA"

	// InvoiceItemCurrencyMKD captures enum value "MKD"
	InvoiceItemCurrencyMKD string = "MKD"

	// InvoiceItemCurrencyMMK captures enum value "MMK"
	InvoiceItemCurrencyMMK string = "MMK"

	// InvoiceItemCurrencyMNT captures enum value "MNT"
	InvoiceItemCurrencyMNT string = "MNT"

	// InvoiceItemCurrencyMOP captures enum value "MOP"
	InvoiceItemCurrencyMOP string = "MOP"

	// InvoiceItemCurrencyMRO captures enum value "MRO"
	InvoiceItemCurrencyMRO string = "MRO"

	// InvoiceItemCurrencyMUR captures enum value "MUR"
	InvoiceItemCurrencyMUR string = "MUR"

	// InvoiceItemCurrencyMVR captures enum value "MVR"
	InvoiceItemCurrencyMVR string = "MVR"

	// InvoiceItemCurrencyMWK captures enum value "MWK"
	InvoiceItemCurrencyMWK string = "MWK"

	// InvoiceItemCurrencyMXN captures enum value "MXN"
	InvoiceItemCurrencyMXN string = "MXN"

	// InvoiceItemCurrencyMYR captures enum value "MYR"
	InvoiceItemCurrencyMYR string = "MYR"

	// InvoiceItemCurrencyMZN captures enum value "MZN"
	InvoiceItemCurrencyMZN string = "MZN"

	// InvoiceItemCurrencyNAD captures enum value "NAD"
	InvoiceItemCurrencyNAD string = "NAD"

	// InvoiceItemCurrencyNGN captures enum value "NGN"
	InvoiceItemCurrencyNGN string = "NGN"

	// InvoiceItemCurrencyNIO captures enum value "NIO"
	InvoiceItemCurrencyNIO string = "NIO"

	// InvoiceItemCurrencyNOK captures enum value "NOK"
	InvoiceItemCurrencyNOK string = "NOK"

	// InvoiceItemCurrencyNPR captures enum value "NPR"
	InvoiceItemCurrencyNPR string = "NPR"

	// InvoiceItemCurrencyNZD captures enum value "NZD"
	InvoiceItemCurrencyNZD string = "NZD"

	// InvoiceItemCurrencyOMR captures enum value "OMR"
	InvoiceItemCurrencyOMR string = "OMR"

	// InvoiceItemCurrencyPAB captures enum value "PAB"
	InvoiceItemCurrencyPAB string = "PAB"

	// InvoiceItemCurrencyPEN captures enum value "PEN"
	InvoiceItemCurrencyPEN string = "PEN"

	// InvoiceItemCurrencyPGK captures enum value "PGK"
	InvoiceItemCurrencyPGK string = "PGK"

	// InvoiceItemCurrencyPHP captures enum value "PHP"
	InvoiceItemCurrencyPHP string = "PHP"

	// InvoiceItemCurrencyPKR captures enum value "PKR"
	InvoiceItemCurrencyPKR string = "PKR"

	// InvoiceItemCurrencyPLN captures enum value "PLN"
	InvoiceItemCurrencyPLN string = "PLN"

	// InvoiceItemCurrencyPYG captures enum value "PYG"
	InvoiceItemCurrencyPYG string = "PYG"

	// InvoiceItemCurrencyQAR captures enum value "QAR"
	InvoiceItemCurrencyQAR string = "QAR"

	// InvoiceItemCurrencyRON captures enum value "RON"
	InvoiceItemCurrencyRON string = "RON"

	// InvoiceItemCurrencyRSD captures enum value "RSD"
	InvoiceItemCurrencyRSD string = "RSD"

	// InvoiceItemCurrencyRUB captures enum value "RUB"
	InvoiceItemCurrencyRUB string = "RUB"

	// InvoiceItemCurrencyRWF captures enum value "RWF"
	InvoiceItemCurrencyRWF string = "RWF"

	// InvoiceItemCurrencySAR captures enum value "SAR"
	InvoiceItemCurrencySAR string = "SAR"

	// InvoiceItemCurrencySBD captures enum value "SBD"
	InvoiceItemCurrencySBD string = "SBD"

	// InvoiceItemCurrencySCR captures enum value "SCR"
	InvoiceItemCurrencySCR string = "SCR"

	// InvoiceItemCurrencySDG captures enum value "SDG"
	InvoiceItemCurrencySDG string = "SDG"

	// InvoiceItemCurrencySEK captures enum value "SEK"
	InvoiceItemCurrencySEK string = "SEK"

	// InvoiceItemCurrencySGD captures enum value "SGD"
	InvoiceItemCurrencySGD string = "SGD"

	// InvoiceItemCurrencySHP captures enum value "SHP"
	InvoiceItemCurrencySHP string = "SHP"

	// InvoiceItemCurrencySLL captures enum value "SLL"
	InvoiceItemCurrencySLL string = "SLL"

	// InvoiceItemCurrencySOS captures enum value "SOS"
	InvoiceItemCurrencySOS string = "SOS"

	// InvoiceItemCurrencySPL captures enum value "SPL"
	InvoiceItemCurrencySPL string = "SPL"

	// InvoiceItemCurrencySRD captures enum value "SRD"
	InvoiceItemCurrencySRD string = "SRD"

	// InvoiceItemCurrencySTD captures enum value "STD"
	InvoiceItemCurrencySTD string = "STD"

	// InvoiceItemCurrencySVC captures enum value "SVC"
	InvoiceItemCurrencySVC string = "SVC"

	// InvoiceItemCurrencySYP captures enum value "SYP"
	InvoiceItemCurrencySYP string = "SYP"

	// InvoiceItemCurrencySZL captures enum value "SZL"
	InvoiceItemCurrencySZL string = "SZL"

	// InvoiceItemCurrencyTHB captures enum value "THB"
	InvoiceItemCurrencyTHB string = "THB"

	// InvoiceItemCurrencyTJS captures enum value "TJS"
	InvoiceItemCurrencyTJS string = "TJS"

	// InvoiceItemCurrencyTMT captures enum value "TMT"
	InvoiceItemCurrencyTMT string = "TMT"

	// InvoiceItemCurrencyTND captures enum value "TND"
	InvoiceItemCurrencyTND string = "TND"

	// InvoiceItemCurrencyTOP captures enum value "TOP"
	InvoiceItemCurrencyTOP string = "TOP"

	// InvoiceItemCurrencyTRY captures enum value "TRY"
	InvoiceItemCurrencyTRY string = "TRY"

	// InvoiceItemCurrencyTTD captures enum value "TTD"
	InvoiceItemCurrencyTTD string = "TTD"

	// InvoiceItemCurrencyTVD captures enum value "TVD"
	InvoiceItemCurrencyTVD string = "TVD"

	// InvoiceItemCurrencyTWD captures enum value "TWD"
	InvoiceItemCurrencyTWD string = "TWD"

	// InvoiceItemCurrencyTZS captures enum value "TZS"
	InvoiceItemCurrencyTZS string = "TZS"

	// InvoiceItemCurrencyUAH captures enum value "UAH"
	InvoiceItemCurrencyUAH string = "UAH"

	// InvoiceItemCurrencyUGX captures enum value "UGX"
	InvoiceItemCurrencyUGX string = "UGX"

	// InvoiceItemCurrencyUSD captures enum value "USD"
	InvoiceItemCurrencyUSD string = "USD"

	// InvoiceItemCurrencyUYU captures enum value "UYU"
	InvoiceItemCurrencyUYU string = "UYU"

	// InvoiceItemCurrencyUZS captures enum value "UZS"
	InvoiceItemCurrencyUZS string = "UZS"

	// InvoiceItemCurrencyVEF captures enum value "VEF"
	InvoiceItemCurrencyVEF string = "VEF"

	// InvoiceItemCurrencyVND captures enum value "VND"
	InvoiceItemCurrencyVND string = "VND"

	// InvoiceItemCurrencyVUV captures enum value "VUV"
	InvoiceItemCurrencyVUV string = "VUV"

	// InvoiceItemCurrencyWST captures enum value "WST"
	InvoiceItemCurrencyWST string = "WST"

	// InvoiceItemCurrencyXAF captures enum value "XAF"
	InvoiceItemCurrencyXAF string = "XAF"

	// InvoiceItemCurrencyXCD captures enum value "XCD"
	InvoiceItemCurrencyXCD string = "XCD"

	// InvoiceItemCurrencyXDR captures enum value "XDR"
	InvoiceItemCurrencyXDR string = "XDR"

	// InvoiceItemCurrencyXOF captures enum value "XOF"
	InvoiceItemCurrencyXOF string = "XOF"

	// InvoiceItemCurrencyXPF captures enum value "XPF"
	InvoiceItemCurrencyXPF string = "XPF"

	// InvoiceItemCurrencyYER captures enum value "YER"
	InvoiceItemCurrencyYER string = "YER"

	// InvoiceItemCurrencyZAR captures enum value "ZAR"
	InvoiceItemCurrencyZAR string = "ZAR"

	// InvoiceItemCurrencyZMW captures enum value "ZMW"
	InvoiceItemCurrencyZMW string = "ZMW"

	// InvoiceItemCurrencyZWD captures enum value "ZWD"
	InvoiceItemCurrencyZWD string = "ZWD"

	// InvoiceItemCurrencyBTC captures enum value "BTC"
	InvoiceItemCurrencyBTC string = "BTC"
)

// prop value enum
func (m *InvoiceItem) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceItemTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceItem) validateCurrency(formats strfmt.Registry) error {
	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateInvoiceID(formats strfmt.Registry) error {
	if swag.IsZero(m.InvoiceID) { // not required
		return nil
	}

	if err := validate.FormatOf("invoiceId", "body", "uuid", m.InvoiceID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateInvoiceItemID(formats strfmt.Registry) error {

	if err := validate.Required("invoiceItemId", "body", m.InvoiceItemID); err != nil {
		return err
	}

	if err := validate.FormatOf("invoiceItemId", "body", "uuid", m.InvoiceItemID.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceItemTypeItemTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXTERNAL_CHARGE","FIXED","RECURRING","REPAIR_ADJ","CBA_ADJ","CREDIT_ADJ","ITEM_ADJ","USAGE","TAX","PARENT_SUMMARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceItemTypeItemTypePropEnum = append(invoiceItemTypeItemTypePropEnum, v)
	}
}

const (

	// InvoiceItemItemTypeEXTERNALCHARGE captures enum value "EXTERNAL_CHARGE"
	InvoiceItemItemTypeEXTERNALCHARGE string = "EXTERNAL_CHARGE"

	// InvoiceItemItemTypeFIXED captures enum value "FIXED"
	InvoiceItemItemTypeFIXED string = "FIXED"

	// InvoiceItemItemTypeRECURRING captures enum value "RECURRING"
	InvoiceItemItemTypeRECURRING string = "RECURRING"

	// InvoiceItemItemTypeREPAIRADJ captures enum value "REPAIR_ADJ"
	InvoiceItemItemTypeREPAIRADJ string = "REPAIR_ADJ"

	// InvoiceItemItemTypeCBAADJ captures enum value "CBA_ADJ"
	InvoiceItemItemTypeCBAADJ string = "CBA_ADJ"

	// InvoiceItemItemTypeCREDITADJ captures enum value "CREDIT_ADJ"
	InvoiceItemItemTypeCREDITADJ string = "CREDIT_ADJ"

	// InvoiceItemItemTypeITEMADJ captures enum value "ITEM_ADJ"
	InvoiceItemItemTypeITEMADJ string = "ITEM_ADJ"

	// InvoiceItemItemTypeUSAGE captures enum value "USAGE"
	InvoiceItemItemTypeUSAGE string = "USAGE"

	// InvoiceItemItemTypeTAX captures enum value "TAX"
	InvoiceItemItemTypeTAX string = "TAX"

	// InvoiceItemItemTypePARENTSUMMARY captures enum value "PARENT_SUMMARY"
	InvoiceItemItemTypePARENTSUMMARY string = "PARENT_SUMMARY"
)

// prop value enum
func (m *InvoiceItem) validateItemTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceItemTypeItemTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceItem) validateItemType(formats strfmt.Registry) error {
	if swag.IsZero(m.ItemType) { // not required
		return nil
	}

	// value enum
	if err := m.validateItemTypeEnum("itemType", "body", m.ItemType); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateLinkedInvoiceItemID(formats strfmt.Registry) error {
	if swag.IsZero(m.LinkedInvoiceItemID) { // not required
		return nil
	}

	if err := validate.FormatOf("linkedInvoiceItemId", "body", "uuid", m.LinkedInvoiceItemID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceItem) validateSubscriptionID(formats strfmt.Registry) error {
	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.FormatOf("subscriptionId", "body", "uuid", m.SubscriptionID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this invoice item based on the context it is used
func (m *InvoiceItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChildItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvoiceItem) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InvoiceItem) contextValidateChildItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ChildItems); i++ {

		if m.ChildItems[i] != nil {
			if err := m.ChildItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("childItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("childItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvoiceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceItem) UnmarshalBinary(b []byte) error {
	var res InvoiceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
