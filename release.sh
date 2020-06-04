#!/usr/bin/env bash

# Copyright 2020 Ray Holder
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o nounset
set -o errexit
set -o pipefail

# just a few common OS_ARCH combinations, not everything
OS_ARCHS=(
    darwin/amd64
    linux/amd64
    linux/arm
    linux/arm64
    windows/amd64
)

function build_for() {
    local os=${1}
    local arch=${2}
    local name=${3}
    local bin_name=${name}-${os}_${arch}
    if [[ "${os}" == "windows" ]]; then
        bin_name=${bin_name}.exe
    fi
    GOOS=${os} GOARCH=${arch} make build BIN_NAME=${bin_name}
}

function build_checksums() {
    local name=${1}
    cd build/
    file ${name}*
    sha256sum ${name}* | tee sha256sums
    sha256sum -c sha256sums
}

function build_release() {
    local name=${1}

    make clean
    for os_arch in "${OS_ARCHS[@]}"
    do
        goos=${os_arch%/*}
        goarch=${os_arch#*/}
        build_for ${goos} ${goarch} ${name}
    done

    build_checksums ${name}
}

build_release "${1}"