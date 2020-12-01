#!/bin/sh

proto_path="./proto:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src"

### Documentation generation
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
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2t \
    --doc_opt=markdown,admin.md \
    proto/onos/e2t/admin/admin.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/e2t \
    --doc_opt=markdown,e2.md \
    proto/onos/e2t/e2/e2.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/topo \
    --doc_opt=markdown,topo.md \
    proto/onos/topo/topo.proto

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

protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2sub/endpoint,plugins=grpc:./go \
    proto/onos/e2sub/endpoint/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2sub/subscription,plugins=grpc:./go \
    proto/onos/e2sub/subscription/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2sub/task,plugins=grpc:./go \
    proto/onos/e2sub/task/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2t/admin,plugins=grpc:./go \
    proto/onos/e2t/admin/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/e2t/e2,plugins=grpc:./go \
    proto/onos/e2t/e2/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/topo,plugins=grpc:./go \
    proto/onos/topo/*.proto

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
    --gogofaster_out=$go_import_paths,import_path=onos/config/admin,plugins=grpc:./go \
    proto/onos/config/admin/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/diags,plugins=grpc:./go \
    proto/onos/config/diags/*.proto


### Python Protobuf code generation
protoc --proto_path=$proto_path \
    --python_betterproto_out=./python \
    proto/onos/e2sub/endpoint/endpoint.proto \
    proto/onos/e2sub/subscription/subscription.proto \
    proto/onos/e2sub/task/task.proto \
    proto/onos/e2t/admin/admin.proto \
    proto/onos/e2t/e2/e2.proto \
    proto/onos/topo/topo.proto \
    \
    proto/onos/config/admin/admin.proto \
    proto/onos/config/diags/diags.proto \
    proto/onos/config/change/types.proto \
    proto/onos/config/change/device/types.proto \
    proto/onos/config/change/network/types.proto \
    proto/onos/config/snapshot/types.proto \
    proto/onos/config/snapshot/device/types.proto \
    proto/onos/config/snapshot/network/types.proto

mv ./python/onos/e2sub/endpoint/__init__.py ./python/onos/e2sub/endpoint.py && rm -r ./python/onos/e2sub/endpoint
mv ./python/onos/e2sub/subscription/__init__.py ./python/onos/e2sub/subscription.py && rm -r ./python/onos/e2sub/subscription
mv ./python/onos/e2sub/task/__init__.py ./python/onos/e2sub/task.py && rm -r ./python/onos/e2sub/task
mv ./python/onos/e2t/admin/__init__.py ./python/onos/e2t/admin.py && rm -r ./python/onos/e2t/admin
mv ./python/onos/e2t/e2/__init__.py ./python/onos/e2t/e2.py && rm -r ./python/onos/e2t/e2
mv ./python/onos/topo/__init__.py ./python/onos/topo/topo.py && touch ./python/onos/topo/__init__.py

mv ./python/onos/config/admin/__init__.py ./python/onos/config/admin/admin.py && touch ./python/onos/config/admin/__init__.py
mv ./python/onos/config/diags/__init__.py ./python/onos/config/diags/diags.py && touch ./python/onos/config/diags/__init__.py
mv ./python/onos/config/change/__init__.py ./python/onos/config/change.py && touch ./python/onos/config/change/__init__.py
mv ./python/onos/config/change/device/__init__.py ./python/onos/config/device_change.py && touch ./python/onos/config/change/device/__init__.py
mv ./python/onos/config/change/network/__init__.py ./python/onos/config/network_change.py && touch ./python/onos/config/change/network/__init__.py
mv ./python/onos/config/snapshot/__init__.py ./python/onos/config/snapshot.py && touch ./python/onos/config/snapshot/__init__.py
mv ./python/onos/config/snapshot/device/__init__.py ./python/onos/config/device_snapshot.py && touch ./python/onos/config/snapshot/device/__init__.py
mv ./python/onos/config/snapshot/network/__init__.py ./python/onos/config/network_snapshottopo.py && touch ./python/onos/config/snapshot/network/__init__.py
