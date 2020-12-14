#!/bin/sh

files=$(find go -name "*.pb.go")

for f in $files; do
    dir=$(dirname ${f})
    pkg=$(basename ${dir})
    root=$(echo $dir | cut -d '/' -f-2)
    path=$(echo $dir | cut -d '/' -f3-)
    if [[ $root != "go/test" ]]; then
        echo "Generating go/test/${path}/${pkg}_mock.go"
        mockgen -source ${f} -destination go/test/${path}/${pkg}_mock.go -package ${pkg}
    fi
done
