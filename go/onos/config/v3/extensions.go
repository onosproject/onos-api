// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v3

import "github.com/openconfig/gnmi/proto/gnmi_ext"

type ExtensionID = gnmi_ext.ExtensionID

const (
	TransactionIDExtensionID       ExtensionID = 120
	TransactionStrategyExtensionID ExtensionID = 121
	TargetExtensionID              ExtensionID = 122
)
