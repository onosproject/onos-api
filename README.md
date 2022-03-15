<!--
SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
SPDX-License-Identifier: Apache-2.0
-->

# onos-api
gRPC API definitions for the µONOS platform

## Overview
This repository houses not only the `.proto` files, but also a number of automatically generated artifacts, e.g. `.pb.go`, `.py` files and similarly for other languages that may be added in the future.

The source tree structure is paritioned into `proto`, which contains the canonical protobuf definitions and a top-level directory for each of the supported languages. The structure within each of the language-specific directories reflects the idoms and conventions appropriate to each language.

## Proto
The top-level package for the protobuf definitions is `onos` and the next level subpackage is the name of the particular platform subsystem, such as `config`, `topo`, etc. This directory tree should contain exclusively the `.proto` files and not be tainted by any other artifacts, especially any language-specific ones.

The proto files are compiled and processed via `build/bin/compile-protos.sh` script, which is invoked by the `Makefile`. All protofiles here should follow the established guidelines and must pass the protobuf lint checker enforcing these conventions.

## Golang
The `go` source tree holds the automatically-generated `.pb.go` artifacts and also any manually authored `.go` source files written in support of the Golang bindings. To minimize the churn and to exercise tighter control over versioning, the generated files are also versioned and maintained in the SCM repo.

The root package of the module is `github.com/onosproject/onos-api/go/onos`, with subpackages being named after each of the platform subsystems, mirroring the structure of the `proto` packages. Golang projects that wish to import µONOS API packages should include the following in the requirements section of their `go.mod` file:

```go
require (
  ...
	github.com/onosproject/onos-api/go v0.6.1
  ...
)

```

Additionally, Go bindings are generated with mocks for testing with [gomock]. Mocks of Protobuf interfaces can be constructed via the same package as the interfaces they mock:

```go
import topoapi "github.com/onosproject/onos-api/go/onos/topo"
...
mockClient := topoapi.NewMockTopoClient(ctrl)
```

## Python

The `python` source tree holds gRPC bindings are generated for [Python]. Pyhton bindings are generated with the [python-betterproto] protoc plugin and support Python3's `asyncio` framework.

[gomock]: https://github.com/golang/mock
[Go]: https://golang.org/
[Protobuf]: https://developers.google.com/protocol-buffers
[Python]: https://www.python.org
[python-betterproto]: https://github.com/danielgtaylor/python-betterproto
