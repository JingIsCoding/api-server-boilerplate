#!/bin/bash
set -e

if [ -f .env.local ]; then
  export $(cat .env.local | grep -v '#' | awk '/=/ {print $1}')
fi

if [ "$ENV" = "production" ]
then
  echo Starting web..
  exec ./bin/web
else
  echo Start web server in dev mode
  make web-dev
fi
