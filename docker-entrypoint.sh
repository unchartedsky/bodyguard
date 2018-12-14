#!/bin/sh
set -e

if [ "$1" != "./bodyguard" ]; then
  exec "$@"
  exit $?
fi

exec "$@" --logtostderr=true --stderrthreshold=${LOGGING_LEVEL}
