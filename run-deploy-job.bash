#!/bin/bash
# Script to run deploy_job against current binaries.
# usage: run-deploy-job.bash {deploy,upgrade} SERIES BASE_ENVIRONMENT TIMEOUT
set -eu
: ${SCRIPTS=$(readlink -f $(dirname $0))}
export SCRIPTS
export USER=jenkins
export JUJU_REPOSITORY=$HOME/repository
export JUJU_HOME=$HOME/cloud-city
export ENV=$3
source $JUJU_HOME/juju-qa.jujuci
set -x
if [ "$1" = "upgrade" ]; then
  extra_args="--upgrade"
elif [ "$1" = "deploy" ]; then
  extra_args=""
else
  echo "Unknown action $1"
  exit 1
fi
series=$2
timeout=$4
shift 4
$SCRIPTS/jujuci.py -v setup-workspace --clean-env $JOB_NAME $WORKSPACE
JUJU_BIN=$(dirname $($SCRIPTS/jujuci.py get-juju-bin))
$SCRIPTS/jujuci.py get-build-vars --summary --env $ENV $REVISION_BUILD
VERSION=$($SCRIPTS/jujuci.py get-build-vars --version $REVISION_BUILD)
if [[ $VERSION =~ ^1\.2[1-2].*$ ]]; then
    echo "Setting the defaul juju to 1.20.11."
    export PATH="$HOME/old-juju/1.20.11/usr/lib/juju-1.20.11/bin:$PATH"
fi

timeout -s INT $timeout $SCRIPTS/deploy_job.py --new-juju-bin $JUJU_BIN\
  --series $series $ENV $WORKSPACE/artifacts $JOB_NAME $extra_args "$@"
