// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SimplePlan simple plan
// swagger:model SimplePlan
type SimplePlan struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// available base products
	AvailableBaseProducts []string `json:"availableBaseProducts"`

	// billing period
	// Enum: [DAILY WEEKLY BIWEEKLY THIRTY_DAYS THIRTY_ONE_DAYS SIXTY_DAYS NINETY_DAYS MONTHLY BIMESTRIAL QUARTERLY TRIANNUAL BIANNUAL ANNUAL SESQUIENNIAL BIENNIAL TRIENNIAL NO_BILLING_PERIOD]
	BillingPeriod SimplePlanBillingPeriodEnum `json:"billingPeriod,omitempty"`

	// currency
	// Enum: [AED AFN ALL AMD ANG AOA ARS AUD AWG AZN BAM BBD BDT BGN BHD BIF BMD BND BOB BRL BSD BTN BWP BYR BZD CAD CDF CHF CLP CNY COP CRC CUC CUP CVE CZK DJF DKK DOP DZD EGP ERN ETB EUR FJD FKP GBP GEL GGP GHS GIP GMD GNF GTQ GYD HKD HNL HRK HTG HUF IDR ILS IMP INR IQD IRR ISK JEP JMD JOD JPY KES KGS KHR KMF KPW KRW KWD KYD KZT LAK LBP LKR LRD LSL LTL LVL LYD MAD MDL MGA MKD MMK MNT MOP MRO MUR MVR MWK MXN MYR MZN NAD NGN NIO NOK NPR NZD OMR PAB PEN PGK PHP PKR PLN PYG QAR RON RSD RUB RWF SAR SBD SCR SDG SEK SGD SHP SLL SOS SPL SRD STD SVC SYP SZL THB TJS TMT TND TOP TRY TTD TVD TWD TZS UAH UGX USD UYU UZS VEF VND VUV WST XAF XCD XDR XOF XPF YER ZAR ZMW ZWD BTC]
	Currency SimplePlanCurrencyEnum `json:"currency,omitempty"`

	// plan Id
	PlanID string `json:"planId,omitempty"`

	// product category
	// Enum: [BASE ADD_ON STANDALONE]
	ProductCategory SimplePlanProductCategoryEnum `json:"productCategory,omitempty"`

	// product name
	ProductName string `json:"productName,omitempty"`

	// trial length
	TrialLength int32 `json:"trialLength,omitempty"`

	// trial time unit
	// Enum: [DAYS WEEKS MONTHS YEARS UNLIMITED]
	TrialTimeUnit SimplePlanTrialTimeUnitEnum `json:"trialTimeUnit,omitempty"`
}

// Validate validates this simple plan
func (m *SimplePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrialTimeUnit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var simplePlanTypeBillingPeriodPropEnum []interface{}

func init() {
	var res []SimplePlanBillingPeriodEnum
	if err := json.Unmarshal([]byte(`["DAILY","WEEKLY","BIWEEKLY","THIRTY_DAYS","THIRTY_ONE_DAYS","SIXTY_DAYS","NINETY_DAYS","MONTHLY","BIMESTRIAL","QUARTERLY","TRIANNUAL","BIANNUAL","ANNUAL","SESQUIENNIAL","BIENNIAL","TRIENNIAL","NO_BILLING_PERIOD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeBillingPeriodPropEnum = append(simplePlanTypeBillingPeriodPropEnum, v)
	}
}

type SimplePlanBillingPeriodEnum string

const (

	// SimplePlanBillingPeriodDAILY captures enum value "DAILY"
	SimplePlanBillingPeriodDAILY SimplePlanBillingPeriodEnum = "DAILY"

	// SimplePlanBillingPeriodWEEKLY captures enum value "WEEKLY"
	SimplePlanBillingPeriodWEEKLY SimplePlanBillingPeriodEnum = "WEEKLY"

	// SimplePlanBillingPeriodBIWEEKLY captures enum value "BIWEEKLY"
	SimplePlanBillingPeriodBIWEEKLY SimplePlanBillingPeriodEnum = "BIWEEKLY"

	// SimplePlanBillingPeriodTHIRTYDAYS captures enum value "THIRTY_DAYS"
	SimplePlanBillingPeriodTHIRTYDAYS SimplePlanBillingPeriodEnum = "THIRTY_DAYS"

	// SimplePlanBillingPeriodTHIRTYONEDAYS captures enum value "THIRTY_ONE_DAYS"
	SimplePlanBillingPeriodTHIRTYONEDAYS SimplePlanBillingPeriodEnum = "THIRTY_ONE_DAYS"

	// SimplePlanBillingPeriodSIXTYDAYS captures enum value "SIXTY_DAYS"
	SimplePlanBillingPeriodSIXTYDAYS SimplePlanBillingPeriodEnum = "SIXTY_DAYS"

	// SimplePlanBillingPeriodNINETYDAYS captures enum value "NINETY_DAYS"
	SimplePlanBillingPeriodNINETYDAYS SimplePlanBillingPeriodEnum = "NINETY_DAYS"

	// SimplePlanBillingPeriodMONTHLY captures enum value "MONTHLY"
	SimplePlanBillingPeriodMONTHLY SimplePlanBillingPeriodEnum = "MONTHLY"

	// SimplePlanBillingPeriodBIMESTRIAL captures enum value "BIMESTRIAL"
	SimplePlanBillingPeriodBIMESTRIAL SimplePlanBillingPeriodEnum = "BIMESTRIAL"

	// SimplePlanBillingPeriodQUARTERLY captures enum value "QUARTERLY"
	SimplePlanBillingPeriodQUARTERLY SimplePlanBillingPeriodEnum = "QUARTERLY"

	// SimplePlanBillingPeriodTRIANNUAL captures enum value "TRIANNUAL"
	SimplePlanBillingPeriodTRIANNUAL SimplePlanBillingPeriodEnum = "TRIANNUAL"

	// SimplePlanBillingPeriodBIANNUAL captures enum value "BIANNUAL"
	SimplePlanBillingPeriodBIANNUAL SimplePlanBillingPeriodEnum = "BIANNUAL"

	// SimplePlanBillingPeriodANNUAL captures enum value "ANNUAL"
	SimplePlanBillingPeriodANNUAL SimplePlanBillingPeriodEnum = "ANNUAL"

	// SimplePlanBillingPeriodSESQUIENNIAL captures enum value "SESQUIENNIAL"
	SimplePlanBillingPeriodSESQUIENNIAL SimplePlanBillingPeriodEnum = "SESQUIENNIAL"

	// SimplePlanBillingPeriodBIENNIAL captures enum value "BIENNIAL"
	SimplePlanBillingPeriodBIENNIAL SimplePlanBillingPeriodEnum = "BIENNIAL"

	// SimplePlanBillingPeriodTRIENNIAL captures enum value "TRIENNIAL"
	SimplePlanBillingPeriodTRIENNIAL SimplePlanBillingPeriodEnum = "TRIENNIAL"

	// SimplePlanBillingPeriodNOBILLINGPERIOD captures enum value "NO_BILLING_PERIOD"
	SimplePlanBillingPeriodNOBILLINGPERIOD SimplePlanBillingPeriodEnum = "NO_BILLING_PERIOD"
)

var SimplePlanBillingPeriodEnumValues = []string{
	"DAILY",
	"WEEKLY",
	"BIWEEKLY",
	"THIRTY_DAYS",
	"THIRTY_ONE_DAYS",
	"SIXTY_DAYS",
	"NINETY_DAYS",
	"MONTHLY",
	"BIMESTRIAL",
	"QUARTERLY",
	"TRIANNUAL",
	"BIANNUAL",
	"ANNUAL",
	"SESQUIENNIAL",
	"BIENNIAL",
	"TRIENNIAL",
	"NO_BILLING_PERIOD",
}

func (e SimplePlanBillingPeriodEnum) IsValid() bool {
	for _, v := range SimplePlanBillingPeriodEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *SimplePlan) validateBillingPeriodEnum(path, location string, value SimplePlanBillingPeriodEnum) error {
	if err := validate.Enum(path, location, value, simplePlanTypeBillingPeriodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateBillingPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingPeriod) { // not required
		return nil
	}

	// value enum
	if err := m.validateBillingPeriodEnum("billingPeriod", "body", m.BillingPeriod); err != nil {
		return err
	}

	return nil
}

var simplePlanTypeCurrencyPropEnum []interface{}

func init() {
	var res []SimplePlanCurrencyEnum
	if err := json.Unmarshal([]byte(`["AED","AFN","ALL","AMD","ANG","AOA","ARS","AUD","AWG","AZN","BAM","BBD","BDT","BGN","BHD","BIF","BMD","BND","BOB","BRL","BSD","BTN","BWP","BYR","BZD","CAD","CDF","CHF","CLP","CNY","COP","CRC","CUC","CUP","CVE","CZK","DJF","DKK","DOP","DZD","EGP","ERN","ETB","EUR","FJD","FKP","GBP","GEL","GGP","GHS","GIP","GMD","GNF","GTQ","GYD","HKD","HNL","HRK","HTG","HUF","IDR","ILS","IMP","INR","IQD","IRR","ISK","JEP","JMD","JOD","JPY","KES","KGS","KHR","KMF","KPW","KRW","KWD","KYD","KZT","LAK","LBP","LKR","LRD","LSL","LTL","LVL","LYD","MAD","MDL","MGA","MKD","MMK","MNT","MOP","MRO","MUR","MVR","MWK","MXN","MYR","MZN","NAD","NGN","NIO","NOK","NPR","NZD","OMR","PAB","PEN","PGK","PHP","PKR","PLN","PYG","QAR","RON","RSD","RUB","RWF","SAR","SBD","SCR","SDG","SEK","SGD","SHP","SLL","SOS","SPL","SRD","STD","SVC","SYP","SZL","THB","TJS","TMT","TND","TOP","TRY","TTD","TVD","TWD","TZS","UAH","UGX","USD","UYU","UZS","VEF","VND","VUV","WST","XAF","XCD","XDR","XOF","XPF","YER","ZAR","ZMW","ZWD","BTC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeCurrencyPropEnum = append(simplePlanTypeCurrencyPropEnum, v)
	}
}

type SimplePlanCurrencyEnum string

const (

	// SimplePlanCurrencyAED captures enum value "AED"
	SimplePlanCurrencyAED SimplePlanCurrencyEnum = "AED"

	// SimplePlanCurrencyAFN captures enum value "AFN"
	SimplePlanCurrencyAFN SimplePlanCurrencyEnum = "AFN"

	// SimplePlanCurrencyALL captures enum value "ALL"
	SimplePlanCurrencyALL SimplePlanCurrencyEnum = "ALL"

	// SimplePlanCurrencyAMD captures enum value "AMD"
	SimplePlanCurrencyAMD SimplePlanCurrencyEnum = "AMD"

	// SimplePlanCurrencyANG captures enum value "ANG"
	SimplePlanCurrencyANG SimplePlanCurrencyEnum = "ANG"

	// SimplePlanCurrencyAOA captures enum value "AOA"
	SimplePlanCurrencyAOA SimplePlanCurrencyEnum = "AOA"

	// SimplePlanCurrencyARS captures enum value "ARS"
	SimplePlanCurrencyARS SimplePlanCurrencyEnum = "ARS"

	// SimplePlanCurrencyAUD captures enum value "AUD"
	SimplePlanCurrencyAUD SimplePlanCurrencyEnum = "AUD"

	// SimplePlanCurrencyAWG captures enum value "AWG"
	SimplePlanCurrencyAWG SimplePlanCurrencyEnum = "AWG"

	// SimplePlanCurrencyAZN captures enum value "AZN"
	SimplePlanCurrencyAZN SimplePlanCurrencyEnum = "AZN"

	// SimplePlanCurrencyBAM captures enum value "BAM"
	SimplePlanCurrencyBAM SimplePlanCurrencyEnum = "BAM"

	// SimplePlanCurrencyBBD captures enum value "BBD"
	SimplePlanCurrencyBBD SimplePlanCurrencyEnum = "BBD"

	// SimplePlanCurrencyBDT captures enum value "BDT"
	SimplePlanCurrencyBDT SimplePlanCurrencyEnum = "BDT"

	// SimplePlanCurrencyBGN captures enum value "BGN"
	SimplePlanCurrencyBGN SimplePlanCurrencyEnum = "BGN"

	// SimplePlanCurrencyBHD captures enum value "BHD"
	SimplePlanCurrencyBHD SimplePlanCurrencyEnum = "BHD"

	// SimplePlanCurrencyBIF captures enum value "BIF"
	SimplePlanCurrencyBIF SimplePlanCurrencyEnum = "BIF"

	// SimplePlanCurrencyBMD captures enum value "BMD"
	SimplePlanCurrencyBMD SimplePlanCurrencyEnum = "BMD"

	// SimplePlanCurrencyBND captures enum value "BND"
	SimplePlanCurrencyBND SimplePlanCurrencyEnum = "BND"

	// SimplePlanCurrencyBOB captures enum value "BOB"
	SimplePlanCurrencyBOB SimplePlanCurrencyEnum = "BOB"

	// SimplePlanCurrencyBRL captures enum value "BRL"
	SimplePlanCurrencyBRL SimplePlanCurrencyEnum = "BRL"

	// SimplePlanCurrencyBSD captures enum value "BSD"
	SimplePlanCurrencyBSD SimplePlanCurrencyEnum = "BSD"

	// SimplePlanCurrencyBTN captures enum value "BTN"
	SimplePlanCurrencyBTN SimplePlanCurrencyEnum = "BTN"

	// SimplePlanCurrencyBWP captures enum value "BWP"
	SimplePlanCurrencyBWP SimplePlanCurrencyEnum = "BWP"

	// SimplePlanCurrencyBYR captures enum value "BYR"
	SimplePlanCurrencyBYR SimplePlanCurrencyEnum = "BYR"

	// SimplePlanCurrencyBZD captures enum value "BZD"
	SimplePlanCurrencyBZD SimplePlanCurrencyEnum = "BZD"

	// SimplePlanCurrencyCAD captures enum value "CAD"
	SimplePlanCurrencyCAD SimplePlanCurrencyEnum = "CAD"

	// SimplePlanCurrencyCDF captures enum value "CDF"
	SimplePlanCurrencyCDF SimplePlanCurrencyEnum = "CDF"

	// SimplePlanCurrencyCHF captures enum value "CHF"
	SimplePlanCurrencyCHF SimplePlanCurrencyEnum = "CHF"

	// SimplePlanCurrencyCLP captures enum value "CLP"
	SimplePlanCurrencyCLP SimplePlanCurrencyEnum = "CLP"

	// SimplePlanCurrencyCNY captures enum value "CNY"
	SimplePlanCurrencyCNY SimplePlanCurrencyEnum = "CNY"

	// SimplePlanCurrencyCOP captures enum value "COP"
	SimplePlanCurrencyCOP SimplePlanCurrencyEnum = "COP"

	// SimplePlanCurrencyCRC captures enum value "CRC"
	SimplePlanCurrencyCRC SimplePlanCurrencyEnum = "CRC"

	// SimplePlanCurrencyCUC captures enum value "CUC"
	SimplePlanCurrencyCUC SimplePlanCurrencyEnum = "CUC"

	// SimplePlanCurrencyCUP captures enum value "CUP"
	SimplePlanCurrencyCUP SimplePlanCurrencyEnum = "CUP"

	// SimplePlanCurrencyCVE captures enum value "CVE"
	SimplePlanCurrencyCVE SimplePlanCurrencyEnum = "CVE"

	// SimplePlanCurrencyCZK captures enum value "CZK"
	SimplePlanCurrencyCZK SimplePlanCurrencyEnum = "CZK"

	// SimplePlanCurrencyDJF captures enum value "DJF"
	SimplePlanCurrencyDJF SimplePlanCurrencyEnum = "DJF"

	// SimplePlanCurrencyDKK captures enum value "DKK"
	SimplePlanCurrencyDKK SimplePlanCurrencyEnum = "DKK"

	// SimplePlanCurrencyDOP captures enum value "DOP"
	SimplePlanCurrencyDOP SimplePlanCurrencyEnum = "DOP"

	// SimplePlanCurrencyDZD captures enum value "DZD"
	SimplePlanCurrencyDZD SimplePlanCurrencyEnum = "DZD"

	// SimplePlanCurrencyEGP captures enum value "EGP"
	SimplePlanCurrencyEGP SimplePlanCurrencyEnum = "EGP"

	// SimplePlanCurrencyERN captures enum value "ERN"
	SimplePlanCurrencyERN SimplePlanCurrencyEnum = "ERN"

	// SimplePlanCurrencyETB captures enum value "ETB"
	SimplePlanCurrencyETB SimplePlanCurrencyEnum = "ETB"

	// SimplePlanCurrencyEUR captures enum value "EUR"
	SimplePlanCurrencyEUR SimplePlanCurrencyEnum = "EUR"

	// SimplePlanCurrencyFJD captures enum value "FJD"
	SimplePlanCurrencyFJD SimplePlanCurrencyEnum = "FJD"

	// SimplePlanCurrencyFKP captures enum value "FKP"
	SimplePlanCurrencyFKP SimplePlanCurrencyEnum = "FKP"

	// SimplePlanCurrencyGBP captures enum value "GBP"
	SimplePlanCurrencyGBP SimplePlanCurrencyEnum = "GBP"

	// SimplePlanCurrencyGEL captures enum value "GEL"
	SimplePlanCurrencyGEL SimplePlanCurrencyEnum = "GEL"

	// SimplePlanCurrencyGGP captures enum value "GGP"
	SimplePlanCurrencyGGP SimplePlanCurrencyEnum = "GGP"

	// SimplePlanCurrencyGHS captures enum value "GHS"
	SimplePlanCurrencyGHS SimplePlanCurrencyEnum = "GHS"

	// SimplePlanCurrencyGIP captures enum value "GIP"
	SimplePlanCurrencyGIP SimplePlanCurrencyEnum = "GIP"

	// SimplePlanCurrencyGMD captures enum value "GMD"
	SimplePlanCurrencyGMD SimplePlanCurrencyEnum = "GMD"

	// SimplePlanCurrencyGNF captures enum value "GNF"
	SimplePlanCurrencyGNF SimplePlanCurrencyEnum = "GNF"

	// SimplePlanCurrencyGTQ captures enum value "GTQ"
	SimplePlanCurrencyGTQ SimplePlanCurrencyEnum = "GTQ"

	// SimplePlanCurrencyGYD captures enum value "GYD"
	SimplePlanCurrencyGYD SimplePlanCurrencyEnum = "GYD"

	// SimplePlanCurrencyHKD captures enum value "HKD"
	SimplePlanCurrencyHKD SimplePlanCurrencyEnum = "HKD"

	// SimplePlanCurrencyHNL captures enum value "HNL"
	SimplePlanCurrencyHNL SimplePlanCurrencyEnum = "HNL"

	// SimplePlanCurrencyHRK captures enum value "HRK"
	SimplePlanCurrencyHRK SimplePlanCurrencyEnum = "HRK"

	// SimplePlanCurrencyHTG captures enum value "HTG"
	SimplePlanCurrencyHTG SimplePlanCurrencyEnum = "HTG"

	// SimplePlanCurrencyHUF captures enum value "HUF"
	SimplePlanCurrencyHUF SimplePlanCurrencyEnum = "HUF"

	// SimplePlanCurrencyIDR captures enum value "IDR"
	SimplePlanCurrencyIDR SimplePlanCurrencyEnum = "IDR"

	// SimplePlanCurrencyILS captures enum value "ILS"
	SimplePlanCurrencyILS SimplePlanCurrencyEnum = "ILS"

	// SimplePlanCurrencyIMP captures enum value "IMP"
	SimplePlanCurrencyIMP SimplePlanCurrencyEnum = "IMP"

	// SimplePlanCurrencyINR captures enum value "INR"
	SimplePlanCurrencyINR SimplePlanCurrencyEnum = "INR"

	// SimplePlanCurrencyIQD captures enum value "IQD"
	SimplePlanCurrencyIQD SimplePlanCurrencyEnum = "IQD"

	// SimplePlanCurrencyIRR captures enum value "IRR"
	SimplePlanCurrencyIRR SimplePlanCurrencyEnum = "IRR"

	// SimplePlanCurrencyISK captures enum value "ISK"
	SimplePlanCurrencyISK SimplePlanCurrencyEnum = "ISK"

	// SimplePlanCurrencyJEP captures enum value "JEP"
	SimplePlanCurrencyJEP SimplePlanCurrencyEnum = "JEP"

	// SimplePlanCurrencyJMD captures enum value "JMD"
	SimplePlanCurrencyJMD SimplePlanCurrencyEnum = "JMD"

	// SimplePlanCurrencyJOD captures enum value "JOD"
	SimplePlanCurrencyJOD SimplePlanCurrencyEnum = "JOD"

	// SimplePlanCurrencyJPY captures enum value "JPY"
	SimplePlanCurrencyJPY SimplePlanCurrencyEnum = "JPY"

	// SimplePlanCurrencyKES captures enum value "KES"
	SimplePlanCurrencyKES SimplePlanCurrencyEnum = "KES"

	// SimplePlanCurrencyKGS captures enum value "KGS"
	SimplePlanCurrencyKGS SimplePlanCurrencyEnum = "KGS"

	// SimplePlanCurrencyKHR captures enum value "KHR"
	SimplePlanCurrencyKHR SimplePlanCurrencyEnum = "KHR"

	// SimplePlanCurrencyKMF captures enum value "KMF"
	SimplePlanCurrencyKMF SimplePlanCurrencyEnum = "KMF"

	// SimplePlanCurrencyKPW captures enum value "KPW"
	SimplePlanCurrencyKPW SimplePlanCurrencyEnum = "KPW"

	// SimplePlanCurrencyKRW captures enum value "KRW"
	SimplePlanCurrencyKRW SimplePlanCurrencyEnum = "KRW"

	// SimplePlanCurrencyKWD captures enum value "KWD"
	SimplePlanCurrencyKWD SimplePlanCurrencyEnum = "KWD"

	// SimplePlanCurrencyKYD captures enum value "KYD"
	SimplePlanCurrencyKYD SimplePlanCurrencyEnum = "KYD"

	// SimplePlanCurrencyKZT captures enum value "KZT"
	SimplePlanCurrencyKZT SimplePlanCurrencyEnum = "KZT"

	// SimplePlanCurrencyLAK captures enum value "LAK"
	SimplePlanCurrencyLAK SimplePlanCurrencyEnum = "LAK"

	// SimplePlanCurrencyLBP captures enum value "LBP"
	SimplePlanCurrencyLBP SimplePlanCurrencyEnum = "LBP"

	// SimplePlanCurrencyLKR captures enum value "LKR"
	SimplePlanCurrencyLKR SimplePlanCurrencyEnum = "LKR"

	// SimplePlanCurrencyLRD captures enum value "LRD"
	SimplePlanCurrencyLRD SimplePlanCurrencyEnum = "LRD"

	// SimplePlanCurrencyLSL captures enum value "LSL"
	SimplePlanCurrencyLSL SimplePlanCurrencyEnum = "LSL"

	// SimplePlanCurrencyLTL captures enum value "LTL"
	SimplePlanCurrencyLTL SimplePlanCurrencyEnum = "LTL"

	// SimplePlanCurrencyLVL captures enum value "LVL"
	SimplePlanCurrencyLVL SimplePlanCurrencyEnum = "LVL"

	// SimplePlanCurrencyLYD captures enum value "LYD"
	SimplePlanCurrencyLYD SimplePlanCurrencyEnum = "LYD"

	// SimplePlanCurrencyMAD captures enum value "MAD"
	SimplePlanCurrencyMAD SimplePlanCurrencyEnum = "MAD"

	// SimplePlanCurrencyMDL captures enum value "MDL"
	SimplePlanCurrencyMDL SimplePlanCurrencyEnum = "MDL"

	// SimplePlanCurrencyMGA captures enum value "MGA"
	SimplePlanCurrencyMGA SimplePlanCurrencyEnum = "MGA"

	// SimplePlanCurrencyMKD captures enum value "MKD"
	SimplePlanCurrencyMKD SimplePlanCurrencyEnum = "MKD"

	// SimplePlanCurrencyMMK captures enum value "MMK"
	SimplePlanCurrencyMMK SimplePlanCurrencyEnum = "MMK"

	// SimplePlanCurrencyMNT captures enum value "MNT"
	SimplePlanCurrencyMNT SimplePlanCurrencyEnum = "MNT"

	// SimplePlanCurrencyMOP captures enum value "MOP"
	SimplePlanCurrencyMOP SimplePlanCurrencyEnum = "MOP"

	// SimplePlanCurrencyMRO captures enum value "MRO"
	SimplePlanCurrencyMRO SimplePlanCurrencyEnum = "MRO"

	// SimplePlanCurrencyMUR captures enum value "MUR"
	SimplePlanCurrencyMUR SimplePlanCurrencyEnum = "MUR"

	// SimplePlanCurrencyMVR captures enum value "MVR"
	SimplePlanCurrencyMVR SimplePlanCurrencyEnum = "MVR"

	// SimplePlanCurrencyMWK captures enum value "MWK"
	SimplePlanCurrencyMWK SimplePlanCurrencyEnum = "MWK"

	// SimplePlanCurrencyMXN captures enum value "MXN"
	SimplePlanCurrencyMXN SimplePlanCurrencyEnum = "MXN"

	// SimplePlanCurrencyMYR captures enum value "MYR"
	SimplePlanCurrencyMYR SimplePlanCurrencyEnum = "MYR"

	// SimplePlanCurrencyMZN captures enum value "MZN"
	SimplePlanCurrencyMZN SimplePlanCurrencyEnum = "MZN"

	// SimplePlanCurrencyNAD captures enum value "NAD"
	SimplePlanCurrencyNAD SimplePlanCurrencyEnum = "NAD"

	// SimplePlanCurrencyNGN captures enum value "NGN"
	SimplePlanCurrencyNGN SimplePlanCurrencyEnum = "NGN"

	// SimplePlanCurrencyNIO captures enum value "NIO"
	SimplePlanCurrencyNIO SimplePlanCurrencyEnum = "NIO"

	// SimplePlanCurrencyNOK captures enum value "NOK"
	SimplePlanCurrencyNOK SimplePlanCurrencyEnum = "NOK"

	// SimplePlanCurrencyNPR captures enum value "NPR"
	SimplePlanCurrencyNPR SimplePlanCurrencyEnum = "NPR"

	// SimplePlanCurrencyNZD captures enum value "NZD"
	SimplePlanCurrencyNZD SimplePlanCurrencyEnum = "NZD"

	// SimplePlanCurrencyOMR captures enum value "OMR"
	SimplePlanCurrencyOMR SimplePlanCurrencyEnum = "OMR"

	// SimplePlanCurrencyPAB captures enum value "PAB"
	SimplePlanCurrencyPAB SimplePlanCurrencyEnum = "PAB"

	// SimplePlanCurrencyPEN captures enum value "PEN"
	SimplePlanCurrencyPEN SimplePlanCurrencyEnum = "PEN"

	// SimplePlanCurrencyPGK captures enum value "PGK"
	SimplePlanCurrencyPGK SimplePlanCurrencyEnum = "PGK"

	// SimplePlanCurrencyPHP captures enum value "PHP"
	SimplePlanCurrencyPHP SimplePlanCurrencyEnum = "PHP"

	// SimplePlanCurrencyPKR captures enum value "PKR"
	SimplePlanCurrencyPKR SimplePlanCurrencyEnum = "PKR"

	// SimplePlanCurrencyPLN captures enum value "PLN"
	SimplePlanCurrencyPLN SimplePlanCurrencyEnum = "PLN"

	// SimplePlanCurrencyPYG captures enum value "PYG"
	SimplePlanCurrencyPYG SimplePlanCurrencyEnum = "PYG"

	// SimplePlanCurrencyQAR captures enum value "QAR"
	SimplePlanCurrencyQAR SimplePlanCurrencyEnum = "QAR"

	// SimplePlanCurrencyRON captures enum value "RON"
	SimplePlanCurrencyRON SimplePlanCurrencyEnum = "RON"

	// SimplePlanCurrencyRSD captures enum value "RSD"
	SimplePlanCurrencyRSD SimplePlanCurrencyEnum = "RSD"

	// SimplePlanCurrencyRUB captures enum value "RUB"
	SimplePlanCurrencyRUB SimplePlanCurrencyEnum = "RUB"

	// SimplePlanCurrencyRWF captures enum value "RWF"
	SimplePlanCurrencyRWF SimplePlanCurrencyEnum = "RWF"

	// SimplePlanCurrencySAR captures enum value "SAR"
	SimplePlanCurrencySAR SimplePlanCurrencyEnum = "SAR"

	// SimplePlanCurrencySBD captures enum value "SBD"
	SimplePlanCurrencySBD SimplePlanCurrencyEnum = "SBD"

	// SimplePlanCurrencySCR captures enum value "SCR"
	SimplePlanCurrencySCR SimplePlanCurrencyEnum = "SCR"

	// SimplePlanCurrencySDG captures enum value "SDG"
	SimplePlanCurrencySDG SimplePlanCurrencyEnum = "SDG"

	// SimplePlanCurrencySEK captures enum value "SEK"
	SimplePlanCurrencySEK SimplePlanCurrencyEnum = "SEK"

	// SimplePlanCurrencySGD captures enum value "SGD"
	SimplePlanCurrencySGD SimplePlanCurrencyEnum = "SGD"

	// SimplePlanCurrencySHP captures enum value "SHP"
	SimplePlanCurrencySHP SimplePlanCurrencyEnum = "SHP"

	// SimplePlanCurrencySLL captures enum value "SLL"
	SimplePlanCurrencySLL SimplePlanCurrencyEnum = "SLL"

	// SimplePlanCurrencySOS captures enum value "SOS"
	SimplePlanCurrencySOS SimplePlanCurrencyEnum = "SOS"

	// SimplePlanCurrencySPL captures enum value "SPL"
	SimplePlanCurrencySPL SimplePlanCurrencyEnum = "SPL"

	// SimplePlanCurrencySRD captures enum value "SRD"
	SimplePlanCurrencySRD SimplePlanCurrencyEnum = "SRD"

	// SimplePlanCurrencySTD captures enum value "STD"
	SimplePlanCurrencySTD SimplePlanCurrencyEnum = "STD"

	// SimplePlanCurrencySVC captures enum value "SVC"
	SimplePlanCurrencySVC SimplePlanCurrencyEnum = "SVC"

	// SimplePlanCurrencySYP captures enum value "SYP"
	SimplePlanCurrencySYP SimplePlanCurrencyEnum = "SYP"

	// SimplePlanCurrencySZL captures enum value "SZL"
	SimplePlanCurrencySZL SimplePlanCurrencyEnum = "SZL"

	// SimplePlanCurrencyTHB captures enum value "THB"
	SimplePlanCurrencyTHB SimplePlanCurrencyEnum = "THB"

	// SimplePlanCurrencyTJS captures enum value "TJS"
	SimplePlanCurrencyTJS SimplePlanCurrencyEnum = "TJS"

	// SimplePlanCurrencyTMT captures enum value "TMT"
	SimplePlanCurrencyTMT SimplePlanCurrencyEnum = "TMT"

	// SimplePlanCurrencyTND captures enum value "TND"
	SimplePlanCurrencyTND SimplePlanCurrencyEnum = "TND"

	// SimplePlanCurrencyTOP captures enum value "TOP"
	SimplePlanCurrencyTOP SimplePlanCurrencyEnum = "TOP"

	// SimplePlanCurrencyTRY captures enum value "TRY"
	SimplePlanCurrencyTRY SimplePlanCurrencyEnum = "TRY"

	// SimplePlanCurrencyTTD captures enum value "TTD"
	SimplePlanCurrencyTTD SimplePlanCurrencyEnum = "TTD"

	// SimplePlanCurrencyTVD captures enum value "TVD"
	SimplePlanCurrencyTVD SimplePlanCurrencyEnum = "TVD"

	// SimplePlanCurrencyTWD captures enum value "TWD"
	SimplePlanCurrencyTWD SimplePlanCurrencyEnum = "TWD"

	// SimplePlanCurrencyTZS captures enum value "TZS"
	SimplePlanCurrencyTZS SimplePlanCurrencyEnum = "TZS"

	// SimplePlanCurrencyUAH captures enum value "UAH"
	SimplePlanCurrencyUAH SimplePlanCurrencyEnum = "UAH"

	// SimplePlanCurrencyUGX captures enum value "UGX"
	SimplePlanCurrencyUGX SimplePlanCurrencyEnum = "UGX"

	// SimplePlanCurrencyUSD captures enum value "USD"
	SimplePlanCurrencyUSD SimplePlanCurrencyEnum = "USD"

	// SimplePlanCurrencyUYU captures enum value "UYU"
	SimplePlanCurrencyUYU SimplePlanCurrencyEnum = "UYU"

	// SimplePlanCurrencyUZS captures enum value "UZS"
	SimplePlanCurrencyUZS SimplePlanCurrencyEnum = "UZS"

	// SimplePlanCurrencyVEF captures enum value "VEF"
	SimplePlanCurrencyVEF SimplePlanCurrencyEnum = "VEF"

	// SimplePlanCurrencyVND captures enum value "VND"
	SimplePlanCurrencyVND SimplePlanCurrencyEnum = "VND"

	// SimplePlanCurrencyVUV captures enum value "VUV"
	SimplePlanCurrencyVUV SimplePlanCurrencyEnum = "VUV"

	// SimplePlanCurrencyWST captures enum value "WST"
	SimplePlanCurrencyWST SimplePlanCurrencyEnum = "WST"

	// SimplePlanCurrencyXAF captures enum value "XAF"
	SimplePlanCurrencyXAF SimplePlanCurrencyEnum = "XAF"

	// SimplePlanCurrencyXCD captures enum value "XCD"
	SimplePlanCurrencyXCD SimplePlanCurrencyEnum = "XCD"

	// SimplePlanCurrencyXDR captures enum value "XDR"
	SimplePlanCurrencyXDR SimplePlanCurrencyEnum = "XDR"

	// SimplePlanCurrencyXOF captures enum value "XOF"
	SimplePlanCurrencyXOF SimplePlanCurrencyEnum = "XOF"

	// SimplePlanCurrencyXPF captures enum value "XPF"
	SimplePlanCurrencyXPF SimplePlanCurrencyEnum = "XPF"

	// SimplePlanCurrencyYER captures enum value "YER"
	SimplePlanCurrencyYER SimplePlanCurrencyEnum = "YER"

	// SimplePlanCurrencyZAR captures enum value "ZAR"
	SimplePlanCurrencyZAR SimplePlanCurrencyEnum = "ZAR"

	// SimplePlanCurrencyZMW captures enum value "ZMW"
	SimplePlanCurrencyZMW SimplePlanCurrencyEnum = "ZMW"

	// SimplePlanCurrencyZWD captures enum value "ZWD"
	SimplePlanCurrencyZWD SimplePlanCurrencyEnum = "ZWD"

	// SimplePlanCurrencyBTC captures enum value "BTC"
	SimplePlanCurrencyBTC SimplePlanCurrencyEnum = "BTC"
)

var SimplePlanCurrencyEnumValues = []string{
	"AED",
	"AFN",
	"ALL",
	"AMD",
	"ANG",
	"AOA",
	"ARS",
	"AUD",
	"AWG",
	"AZN",
	"BAM",
	"BBD",
	"BDT",
	"BGN",
	"BHD",
	"BIF",
	"BMD",
	"BND",
	"BOB",
	"BRL",
	"BSD",
	"BTN",
	"BWP",
	"BYR",
	"BZD",
	"CAD",
	"CDF",
	"CHF",
	"CLP",
	"CNY",
	"COP",
	"CRC",
	"CUC",
	"CUP",
	"CVE",
	"CZK",
	"DJF",
	"DKK",
	"DOP",
	"DZD",
	"EGP",
	"ERN",
	"ETB",
	"EUR",
	"FJD",
	"FKP",
	"GBP",
	"GEL",
	"GGP",
	"GHS",
	"GIP",
	"GMD",
	"GNF",
	"GTQ",
	"GYD",
	"HKD",
	"HNL",
	"HRK",
	"HTG",
	"HUF",
	"IDR",
	"ILS",
	"IMP",
	"INR",
	"IQD",
	"IRR",
	"ISK",
	"JEP",
	"JMD",
	"JOD",
	"JPY",
	"KES",
	"KGS",
	"KHR",
	"KMF",
	"KPW",
	"KRW",
	"KWD",
	"KYD",
	"KZT",
	"LAK",
	"LBP",
	"LKR",
	"LRD",
	"LSL",
	"LTL",
	"LVL",
	"LYD",
	"MAD",
	"MDL",
	"MGA",
	"MKD",
	"MMK",
	"MNT",
	"MOP",
	"MRO",
	"MUR",
	"MVR",
	"MWK",
	"MXN",
	"MYR",
	"MZN",
	"NAD",
	"NGN",
	"NIO",
	"NOK",
	"NPR",
	"NZD",
	"OMR",
	"PAB",
	"PEN",
	"PGK",
	"PHP",
	"PKR",
	"PLN",
	"PYG",
	"QAR",
	"RON",
	"RSD",
	"RUB",
	"RWF",
	"SAR",
	"SBD",
	"SCR",
	"SDG",
	"SEK",
	"SGD",
	"SHP",
	"SLL",
	"SOS",
	"SPL",
	"SRD",
	"STD",
	"SVC",
	"SYP",
	"SZL",
	"THB",
	"TJS",
	"TMT",
	"TND",
	"TOP",
	"TRY",
	"TTD",
	"TVD",
	"TWD",
	"TZS",
	"UAH",
	"UGX",
	"USD",
	"UYU",
	"UZS",
	"VEF",
	"VND",
	"VUV",
	"WST",
	"XAF",
	"XCD",
	"XDR",
	"XOF",
	"XPF",
	"YER",
	"ZAR",
	"ZMW",
	"ZWD",
	"BTC",
}

func (e SimplePlanCurrencyEnum) IsValid() bool {
	for _, v := range SimplePlanCurrencyEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *SimplePlan) validateCurrencyEnum(path, location string, value SimplePlanCurrencyEnum) error {
	if err := validate.Enum(path, location, value, simplePlanTypeCurrencyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("currency", "body", m.Currency); err != nil {
		return err
	}

	return nil
}

var simplePlanTypeProductCategoryPropEnum []interface{}

func init() {
	var res []SimplePlanProductCategoryEnum
	if err := json.Unmarshal([]byte(`["BASE","ADD_ON","STANDALONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeProductCategoryPropEnum = append(simplePlanTypeProductCategoryPropEnum, v)
	}
}

type SimplePlanProductCategoryEnum string

const (

	// SimplePlanProductCategoryBASE captures enum value "BASE"
	SimplePlanProductCategoryBASE SimplePlanProductCategoryEnum = "BASE"

	// SimplePlanProductCategoryADDON captures enum value "ADD_ON"
	SimplePlanProductCategoryADDON SimplePlanProductCategoryEnum = "ADD_ON"

	// SimplePlanProductCategorySTANDALONE captures enum value "STANDALONE"
	SimplePlanProductCategorySTANDALONE SimplePlanProductCategoryEnum = "STANDALONE"
)

var SimplePlanProductCategoryEnumValues = []string{
	"BASE",
	"ADD_ON",
	"STANDALONE",
}

func (e SimplePlanProductCategoryEnum) IsValid() bool {
	for _, v := range SimplePlanProductCategoryEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *SimplePlan) validateProductCategoryEnum(path, location string, value SimplePlanProductCategoryEnum) error {
	if err := validate.Enum(path, location, value, simplePlanTypeProductCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateProductCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateProductCategoryEnum("productCategory", "body", m.ProductCategory); err != nil {
		return err
	}

	return nil
}

var simplePlanTypeTrialTimeUnitPropEnum []interface{}

func init() {
	var res []SimplePlanTrialTimeUnitEnum
	if err := json.Unmarshal([]byte(`["DAYS","WEEKS","MONTHS","YEARS","UNLIMITED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		simplePlanTypeTrialTimeUnitPropEnum = append(simplePlanTypeTrialTimeUnitPropEnum, v)
	}
}

type SimplePlanTrialTimeUnitEnum string

const (

	// SimplePlanTrialTimeUnitDAYS captures enum value "DAYS"
	SimplePlanTrialTimeUnitDAYS SimplePlanTrialTimeUnitEnum = "DAYS"

	// SimplePlanTrialTimeUnitWEEKS captures enum value "WEEKS"
	SimplePlanTrialTimeUnitWEEKS SimplePlanTrialTimeUnitEnum = "WEEKS"

	// SimplePlanTrialTimeUnitMONTHS captures enum value "MONTHS"
	SimplePlanTrialTimeUnitMONTHS SimplePlanTrialTimeUnitEnum = "MONTHS"

	// SimplePlanTrialTimeUnitYEARS captures enum value "YEARS"
	SimplePlanTrialTimeUnitYEARS SimplePlanTrialTimeUnitEnum = "YEARS"

	// SimplePlanTrialTimeUnitUNLIMITED captures enum value "UNLIMITED"
	SimplePlanTrialTimeUnitUNLIMITED SimplePlanTrialTimeUnitEnum = "UNLIMITED"
)

var SimplePlanTrialTimeUnitEnumValues = []string{
	"DAYS",
	"WEEKS",
	"MONTHS",
	"YEARS",
	"UNLIMITED",
}

func (e SimplePlanTrialTimeUnitEnum) IsValid() bool {
	for _, v := range SimplePlanTrialTimeUnitEnumValues {
		if v == string(e) {
			return true
		}
	}
	return false
}

// prop value enum
func (m *SimplePlan) validateTrialTimeUnitEnum(path, location string, value SimplePlanTrialTimeUnitEnum) error {
	if err := validate.Enum(path, location, value, simplePlanTypeTrialTimeUnitPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SimplePlan) validateTrialTimeUnit(formats strfmt.Registry) error {

	if swag.IsZero(m.TrialTimeUnit) { // not required
		return nil
	}

	// value enum
	if err := m.validateTrialTimeUnitEnum("trialTimeUnit", "body", m.TrialTimeUnit); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimplePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimplePlan) UnmarshalBinary(b []byte) error {
	var res SimplePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
