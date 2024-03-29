#!/bin/bash

function_name="wtf2"

zip -r main-go.zip main.go

version_id=$(yc serverless function version create \
    --function-name $function_name \
    --runtime golang113 \
    --source-path main-go.zip \
    --entrypoint main.Handler \
    --format json | jq -r '.id')

yc serverless function version set-tag \
    --id $version_id \
    --tag public
