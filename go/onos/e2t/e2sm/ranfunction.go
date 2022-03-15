// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

type ShortName string

type OID string

type Description string

type Instance int32

type StyleType int32

type StyleName string

type FormatType int32

type Version string

type ModuleName string

type RanfunctionNameDef struct {
	RanFunctionShortName   ShortName
	RanFunctionE2SmOid     OID
	RanFunctionDescription Description
	RanFunctionInstance    Instance
}

type RicReportStyleDef struct {
	RicReportStyleType             StyleType
	RicReportStyleName             StyleName
	RicIndicationHeaderFormatType  FormatType
	RicIndicationMessageFormatType FormatType
}

type RicEventTriggerDef struct {
	RicEventStyleType  StyleType
	RicEventStyleName  StyleName
	RicEventFormatType FormatType
}

type RicReportList map[StyleType]RicReportStyleDef
type RicEventTriggerList map[StyleType]RicEventTriggerDef
