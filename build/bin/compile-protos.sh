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
    proto/onos/e2t/e2/v1beta1/*.proto

# a1t
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/a1t \
    --doc_opt=markdown,a1.md \
    proto/onos/a1t/a1/*.proto

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

## onos-config v2 API

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config/v2 \
    --doc_opt=markdown,value.md \
    proto/onos/config/v2/value.proto

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config/v2 \
    --doc_opt=markdown,transaction.md \
    proto/onos/config/v2/transaction.proto

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config/v2 \
    --doc_opt=markdown,proposal.md \
    proto/onos/config/v2/proposal.proto

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config/v2 \
    --doc_opt=markdown,configuration.md \
    proto/onos/config/v2/configuration.proto

protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config/v2 \
    --doc_opt=markdown,extensions.md \
    proto/onos/config/v2/extensions.proto


#configmodel
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/configmodel \
    --doc_opt=markdown,registry.md \
    proto/onos/configmodel/registry.proto

# kpimon
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/kpimon \
    --doc_opt=markdown,kpimon.md \
    proto/onos/kpimon/kpimon.proto

# pci
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/pci \
    --doc_opt=markdown,pci.md \
    proto/onos/pci/pci.proto

# mlb
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/mlb \
    --doc_opt=markdown,mlb.md \
    proto/onos/mlb/mlb.proto

# rsm
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/rsm \
    --doc_opt=markdown,rsm.md \
    proto/onos/rsm/rsm.proto

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
go_import_paths="${go_import_paths},Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types"
go_import_paths="${go_import_paths},Monos/config/device/types.proto=github.com/onosproject/onos-api/go/onos/config/device"
go_import_paths="${go_import_paths},Monos/config/admin/admin.proto=github.com/onosproject/onos-api/go/onos/config/admin"
go_import_paths="${go_import_paths},Monos/config/change/types.proto=github.com/onosproject/onos-api/go/onos/config/change"
go_import_paths="${go_import_paths},Monos/config/change/device/types.proto=github.com/onosproject/onos-api/go/onos/config/change/device"
go_import_paths="${go_import_paths},Monos/config/change/network/types.proto=github.com/onosproject/onos-api/go/onos/config/change/network"
go_import_paths="${go_import_paths},Monos/config/snapshot/types.proto=github.com/onosproject/onos-api/go/onos/config/snapshot"
go_import_paths="${go_import_paths},Monos/config/snapshot/device/types.proto=github.com/onosproject/onos-api/go/onos/config/snapshot/device"
go_import_paths="${go_import_paths},Monos/ransim/types/types.proto=github.com/onosproject/onos-api/go/onos/ransim/types"
go_import_paths="${go_import_paths},Monos/config/v2/object.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/failure.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/value.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/change.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/transaction.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/proposal.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/configuration.proto=github.com/onosproject/onos-api/go/onos/config/v2"
go_import_paths="${go_import_paths},Monos/config/v2/extensions.proto=github.com/onosproject/onos-api/go/onos/config/v2"


# topo and UE-NIB
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/topo,plugins=grpc:./go \
    proto/onos/topo/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/uenib,plugins=grpc:./go \
    proto/onos/uenib/*.proto

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
    --gogofaster_out=$go_import_paths,import_path=onos/e2t/e2/v1beta1,plugins=grpc:./go \
    proto/onos/e2t/e2/v1beta1/*.proto

# a1t
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/a1t/a1,plugins=grpc:./go \
    proto/onos/a1t/a1/*.proto

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

# onos-config v2 API

protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/v2,plugins=grpc:./go \
    proto/onos/config/v2/*.proto

#configmodel
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/configmodel,plugins=grpc:./go \
    proto/onos/configmodel/*.proto

# admin.proto cannot be generated with fast marshaler/unmarshaler because it uses gnmi.ModelData
protoc --proto_path=$proto_path \
    --gogo_out=$go_import_paths,import_path=onos/config/admin,plugins=grpc:./go \
    proto/onos/config/admin/*.proto


# kpimon
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/kpimon,plugins=grpc:./go \
    proto/onos/kpimon/*.proto

# pci
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/pci,plugins=grpc:./go \
    proto/onos/pci/*.proto

# mlb
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/mlb,plugins=grpc:./go \
    proto/onos/mlb/*.proto

# rsm
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/rsm,plugins=grpc:./go \
    proto/onos/rsm/*.proto

# mho
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/mho,plugins=grpc:./go \
    proto/onos/mho/*.proto

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

# perf
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/topo,plugins=grpc:./go \
    proto/onos/perf/perf.proto

### Python Protobuf code generation
mkdir -p ./python
protoc --proto_path=$proto_path \
    --python_betterproto_out=./python \
    $(find proto -name "*.proto" | sort)

# FIXME: come up with a better way to patch python files; this is too brittle
# git apply ./build/bin/patches/*.patch
