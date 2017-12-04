#!/usr/bin/env bash

set -exu -o pipefail

export BOSH_LOG_LEVEL=debug
export BOSH_LOG_PATH="$PWD/bosh.log"
version=$(cat kubo-version/version)

cp -r git-kubo-release/. git-kubo-release-output

pushd git-kubo-release-output

cat <<EOF > "config/private.yml"
blobstore:
  options:
    access_key_id: ${ACCESS_KEY_ID}
    secret_access_key: ${SECRET_ACCESS_KEY}
EOF

bosh-cli finalize-release ../gcs-kubo-release-tarball/kubo-release-*.tgz --version=${version}

git add .
git config --global user.name "Kubo CI"
git config --global user.email "cf-london-eng@pivotal.io"
git commit -m "Final release for ${version}"
git tag -a ${version} -m "Tagging for version ${version}"