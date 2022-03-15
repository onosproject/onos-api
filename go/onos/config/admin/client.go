// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package admin

import (
	"google.golang.org/grpc"
)

// ConfigAdminClientFactory : Default ConfigAdminClient creation.
var ConfigAdminClientFactory = func(cc *grpc.ClientConn) ConfigAdminServiceClient {
	return NewConfigAdminServiceClient(cc)
}

// CreateConfigAdminServiceClient creates and returns a new config admin client
func CreateConfigAdminServiceClient(cc *grpc.ClientConn) ConfigAdminServiceClient {
	return ConfigAdminClientFactory(cc)
}
