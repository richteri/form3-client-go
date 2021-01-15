#!/bin/sh

ENDPOINT=$API_ADDR/v1/organisation/accounts

until curl --output /dev/null --silent --fail "$ENDPOINT"; do
    echo Waiting for API
    sleep 5
done

make test
