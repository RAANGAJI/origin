#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

# Generates types_swagger_doc_generated file for the given group version.
# $1: Name of the group version
# $2: Path to the directory where types.go for that group version exists. This
# is the directory where the file will be generated.
kube::swagger::gen_types_swagger_doc() {
  local group_version=$1
  local gv_dir=$2
  local TMPFILE="${TMPDIR:-/tmp}/types_swagger_doc_generated.$(date +%s).go"

  echo "Generating swagger type docs for ${group_version} at ${gv_dir}"

  sed 's/YEAR/2017/' hack/empty.txt > "$TMPFILE"
  echo "package ${group_version##*/}" >> "$TMPFILE"
  cat >> "$TMPFILE" <<EOF

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
EOF

  go run tools/genswaggertypedocs/swagger_type_docs.go -s \
    "${gv_dir}/types.go" \
    -f - \
    >>  "$TMPFILE"

  echo "// AUTO-GENERATED FUNCTIONS END HERE" >> "$TMPFILE"

  gofmt -w -s "$TMPFILE"
  mv "$TMPFILE" ""${gv_dir}"/types_swagger_doc_generated.go"
}

rm -f "${SCRIPT_ROOT}/apps/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "apps/v1" "apps/v1"
rm -f "${SCRIPT_ROOT}/authorization/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "authorization/v1" "authorization/v1"
rm -f "${SCRIPT_ROOT}/build/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "build/v1" "build/v1"
rm -f "${SCRIPT_ROOT}/image/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "image/v1" "image/v1"
rm -f "${SCRIPT_ROOT}/network/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "network/v1" "network/v1"
rm -f "${SCRIPT_ROOT}/oauth/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "oauth/v1" "oauth/v1"
rm -f "${SCRIPT_ROOT}/project/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "project/v1" "project/v1"
rm -f "${SCRIPT_ROOT}/quota/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "quota/v1" "quota/v1"
rm -f "${SCRIPT_ROOT}/route/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "route/v1" "route/v1"
rm -f "${SCRIPT_ROOT}/security/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "security/v1" "security/v1"
rm -f "${SCRIPT_ROOT}/template/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "template/v1" "template/v1"
rm -f "${SCRIPT_ROOT}/user/v1/types_swagger_doc_generated.go"
kube::swagger::gen_types_swagger_doc "user/v1" "user/v1"