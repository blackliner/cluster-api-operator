#!/bin/bash

# Copyright 2023 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

VERSION=${1}
OUTPUT_PATH=${2}

# Ensure the output folder exists
mkdir -p "${OUTPUT_PATH}"

# Install cargo
curl https://sh.rustup.rs -sSf | sh -s -- -y
. "$HOME/.cargo/env"

# Install mdbook and dependencies
cargo install mdbook --version "$VERSION" --root "$OUTPUT_PATH"
cargo install mdbook-fs-summary --root "$OUTPUT_PATH"
cargo install mdbook-toc --root "$OUTPUT_PATH"
