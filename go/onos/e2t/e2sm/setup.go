// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package types

import (
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

// OnSetupRequest on E2 setup request
type OnSetupRequest struct {
	ServiceModels          map[string]*topoapi.ServiceModelInfo
	E2Cells                *[]*topoapi.E2Cell
	RANFunctionDescription []byte
}
