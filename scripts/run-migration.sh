#!/bin/bash
set -e

if [ -f .env.local ]; then
  export $(cat .env.local | grep -v '#' | awk '/=/ {print $1}')
fi

./bin/migrate -command=$1 -version=$2
