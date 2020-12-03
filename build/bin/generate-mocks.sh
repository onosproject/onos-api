#!/bin/sh

files=$(find go -name "*.pb.go")

for f in $files; do
    dir=$(dirname $f)
    pkg=$(basename $dir)
    mockgen -source $f -destination $dir/mock/topo_mock.go -package $pkg
done
