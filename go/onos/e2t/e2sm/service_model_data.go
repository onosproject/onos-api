// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

// ServiceModelData service model related information
type ServiceModelData struct {
	Name       ShortName
	Version    Version
	ModuleName ModuleName
	OID        OID
}
