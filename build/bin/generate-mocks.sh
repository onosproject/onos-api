#!/bin/sh

files=$(find go -name "*.pb.go")

for f in $files; do
    dir=$(dirname ${f})
    pkg=$(basename ${dir})
    path=$(echo $f | cut -d '/' -f3-)
    echo "Generating go/test/${path}/${pkg}_mock.go"
    mockgen -source ${f} -destination go/test/${path}/${pkg}_mock.go -package ${pkg}
done
