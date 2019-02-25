#!/bin/sh

# github-newbranch-publish.bash <RELEASE_ID> - release from the security repository

set -euo pipefail

BRANCH="${CIRCLE_TAG:-v0.0.0}-security-`date -u +%Y%m%d%H%M%S`"
# Check if the patch release exist already as a branch
if [ -n "`git branch | grep '${BRANCH}'`" ]; then
  echo "WARNING: Branch ${BRANCH} already exists."
else
  echo "Creating branch ${BRANCH}."
  git branch "${BRANCH}"
fi

git checkout "${BRANCH}"

git remote add tendermint-origin git@github.com:tendermint/tendermint.git

git push tendermint-origin
git push tendermint-origin --tags

#Now the tag exists on the public repository
python -u scripts/release_management/github-publish.py --id "${1}"

python -u scripts/release_management/github-openpr.py --head "${BRANCH}" --base "${BRANCH:%.*}"
