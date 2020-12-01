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
    proto/onos/config/types/change/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,device_change.md \
    proto/onos/config/types/change/device/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,network_change.md \
    proto/onos/config/types/change/network/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,snapshot_types.md \
    proto/onos/config/types/snapshot/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,device_snapshot.md \
    proto/onos/config/types/snapshot/device/types.proto
protoc --proto_path=$proto_path \
    --doc_out=docs/onos/config \
    --doc_opt=markdown,network_snapshot.md \
    proto/onos/config/types/snapshot/network/types.proto



### Go Protobuf code generation
go_import_paths="Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types"

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
    --gogofaster_out=$go_import_paths,import_path=onos/config/types/change,plugins=grpc:./go \
    proto/onos/config/types/change/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/types/change/device,plugins=grpc:./go \
    proto/onos/config/types/change/device/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/types/change/network,plugins=grpc:./go \
    proto/onos/config/types/change/network/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/types/snapshot,plugins=grpc:./go \
    proto/onos/config/types/snapshot/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/types/snapshot/device,plugins=grpc:./go \
    proto/onos/config/types/snapshot/device/*.proto
protoc --proto_path=$proto_path \
    --gogofaster_out=$go_import_paths,import_path=onos/config/types/snapshot/network,plugins=grpc:./go \
    proto/onos/config/types/snapshot/network/*.proto
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
    proto/onos/config/types/change/types.proto \
    proto/onos/config/types/change/device/types.proto \
    proto/onos/config/types/change/network/types.proto \
    proto/onos/config/types/snapshot/types.proto \
    proto/onos/config/types/snapshot/device/types.proto \
    proto/onos/config/types/snapshot/network/types.proto

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
