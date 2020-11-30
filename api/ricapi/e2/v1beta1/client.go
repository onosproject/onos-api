// Copyright 2020-present Open Networking Foundation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import "google.golang.org/grpc"

// E2TClientFactory : Default E2TClient creation.
var E2TClientFactory = func(cc *grpc.ClientConn) E2TServiceClient {
	return NewE2TServiceClient(cc)
}

// CreateE2ServiceClient creates and returns a new config admin client
func CreateE2TServiceClient(cc *grpc.ClientConn) E2TServiceClient {
	return E2TClientFactory(cc)
}
