// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package v2

import "github.com/openconfig/gnmi/proto/gnmi_ext"

type ExtensionID = gnmi_ext.ExtensionID

const (
	TransactionInfoExtensionID     ExtensionID = 110
	TransactionStrategyExtensionID ExtensionID = 111
	TargetVersionOverridesID       ExtensionID = 112
)
