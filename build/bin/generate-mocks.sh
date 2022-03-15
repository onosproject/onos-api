#!/bin/sh
# SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

files=$(find go -name "*.pb.go")

for f in $files; do
    dir=$(dirname ${f})
    pkg=$(basename ${dir})
    file="${dir}/${pkg}_mock.go"
    echo "Generating ${file}"
    mockgen -source ${f} -destination ${file} -package ${pkg}
    goimports -w ${file}
done
