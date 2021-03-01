#!/bin/sh

proto_path="./proto:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src"

### Documentation generation

# e2sub
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2sub \
    --doc_opt=markdown,endpoint.md \
    proto/onos/e2sub/endpoint/endpoint.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2sub \
    --doc_opt=markdown,subscription.md \
    proto/onos/e2sub/subscription/subscription.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2sub \
    --doc_opt=markdown,task.md \
    proto/onos/e2sub/task/task.proto

# e2t
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2t \
    --doc_opt=markdown,admin.md \
    proto/onos/e2t/admin/admin.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2t \
    --doc_opt=markdown,e2.md \
    proto/onos/e2t/e2/e2.proto

# topo
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/topo \
    --doc_opt=markdown,topo.md \
    proto/onos/topo/topo.proto

# config
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,admin.md \
    proto/onos/config/admin/admin.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,diags.md \
    proto/onos/config/diags/diags.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,change_types.md \
    proto/onos/config/change/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,device_change.md \
    proto/onos/config/change/device/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,network_change.md \
    proto/onos/config/change/network/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,snapshot_types.md \
    proto/onos/config/snapshot/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,device_snapshot.md \
    proto/onos/config/snapshot/device/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,network_snapshot.md \
    proto/onos/config/snapshot/network/types.proto

# kpimon
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/kpimon \
    --doc_opt=markdown,kpimon.md \
    proto/onos/kpimon/kpimon.proto

# ransim
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/ransim \
    --doc_opt=markdown,metrics.md \
    proto/onos/ransim/metrics/metrics.proto


protoc --proto_path=$proto_path \
    --doc_out=docs/onos/ransim \
    --doc_opt=markdown,model.md \
    proto/onos/ransim/model/model.proto

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/ransim \
    --doc_opt=markdown,trafficsim.md \
    proto/onos/ransim/trafficsim/trafficsim.proto

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/ransim \
    --doc_opt=markdown,types.md \
    proto/onos/ransim/types/types.proto


### Go Protobuf code generation
go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Monos/config/device/types.proto=github.com/onosproject/onos-api/go/onos/config/device"
go_import_paths="${go_import_paths},Monos/config/admin/admin.proto=github.com/onosproject/onos-api/go/onos/config/admin"
go_import_paths="${go_import_paths},Monos/config/change/types.proto=github.com/onosproject/onos-api/go/onos/config/change"
go_import_paths="${go_import_paths},Monos/config/change/device/types.proto=github.com/onosproject/onos-api/go/onos/config/change/device"
go_import_paths="${go_import_paths},Monos/config/change/network/types.proto=github.com/onosproject/onos-api/go/onos/config/change/network"
go_import_paths="${go_import_paths},Monos/config/snapshot/types.proto=github.com/onosproject/onos-api/go/onos/config/snapshot"
go_import_paths="${go_import_paths},Monos/config/snapshot/device/types.proto=github.com/onosproject/onos-api/go/onos/config/snapshot/device"
go_import_paths="${go_import_paths},Monos/ransim/types/types.proto=github.com/onosproject/onos-api/go/onos/ransim/types"
# e2sub
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2sub/endpoint,plugins=grpc:./go \
    proto/onos/e2sub/endpoint/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2sub/subscription,plugins=grpc:./go \
    proto/onos/e2sub/subscription/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2sub/task,plugins=grpc:./go \
    proto/onos/e2sub/task/*.proto

# e2t
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2t/admin,plugins=grpc:./go \
    proto/onos/e2t/admin/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2t/e2,plugins=grpc:./go \
    proto/onos/e2t/e2/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/topo,plugins=grpc:./go \
    proto/onos/topo/*.proto

# config
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/change,plugins=grpc:./go \
    proto/onos/config/change/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/change/device,plugins=grpc:./go \
    proto/onos/config/change/device/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/change/network,plugins=grpc:./go \
    proto/onos/config/change/network/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/snapshot,plugins=grpc:./go \
    proto/onos/config/snapshot/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/snapshot/device,plugins=grpc:./go \
    proto/onos/config/snapshot/device/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/snapshot/network,plugins=grpc:./go \
    proto/onos/config/snapshot/network/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/diags,plugins=grpc:./go \
    proto/onos/config/diags/*.proto

# admin.proto cannot be generated with fast marshaler/unmarshaler because it uses gnmi.ModelData
protoc --proto_path=$proto_path \
    --gogo_out=$go_import_paths,import_path=onos/config/admin,plugins=grpc:./go \
    proto/onos/config/admin/*.proto


# kpimon
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/kpimon,plugins=grpc:./go \
    proto/onos/kpimon/*.proto

# ransim
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/ransim/metrics,plugins=grpc:./go \
    proto/onos/ransim/metrics/*.proto

protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/ransim/model,plugins=grpc:./go \
    proto/onos/ransim/model/*.proto

protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/ransim/trafficsim,plugins=grpc:./go \
    proto/onos/ransim/trafficsim/*.proto

protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/ransim/types,plugins=grpc:./go \
    proto/onos/ransim/types/*.proto


### Python Protobuf code generation
mkdir -p ./python
protoc --proto_path=$proto_path \
    --python_betterproto_out=./python \
    $(find proto -name "*.proto")
