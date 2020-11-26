#!/bin/sh

proto_imports=".:${GOPATH}/src/github.com/gogo/protobuf/protobuf:${GOPATH}/src/github.com/gogo/protobuf:${GOPATH}/src":"${GOPATH}/src/github.com/onosproject/onos-api/api"

protoc -I=$proto_imports --doc_out=docs/api  --doc_opt=markdown,topo.md --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,import_path=topo,plugins=grpc:. api/topo/*.proto

protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,headers.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=ricapi/e2,plugins=grpc:. api/ricapi/e2/headers/v1beta1/headers.proto
protoc -I=$proto_imports  --doc_out=docs/api  --doc_opt=markdown,ricapi.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=ricapi/e2,plugins=grpc:. api/ricapi/e2/v1beta1/ricapi_e2.proto

cp -r github.com/onosproject/onos-api/* .
rm -rf github.com

### onos-e2t related APIs ###
protoc -I=$proto_imports  --doc_out=docs/api/e2t  --doc_opt=markdown,admin.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-api/api,plugins=grpc:. api/e2t/admin.proto
protoc -I=$proto_imports  --doc_out=docs/api/e2t  --doc_opt=markdown,e2.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-api/api,plugins=grpc:. api/e2t/e2.proto

### onos-e2sub related APIs ###
protoc -I=$proto_imports  --doc_out=docs/api/e2sub  --doc_opt=markdown,subscription.md --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-api/api,plugins=grpc:. api/e2sub/subscription.proto
protoc -I=$proto_imports  --doc_out=docs/api/e2sub  --doc_opt=markdown,endpoint.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-api/api,plugins=grpc:. api/e2sub/endpoint.proto
protoc -I=$proto_imports  --doc_out=docs/api/e2sub  --doc_opt=markdown,task.md  --gogo_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,import_path=github.com/onosproject/onos-api/api,plugins=grpc:. api/e2sub/task.proto




### onos-topo related APIs ###
protoc -I=$proto_imports --doc_out=docs/api/topo  --doc_opt=markdown,topo.md --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,import_path=topo,plugins=grpc:. api/topo/*.proto

