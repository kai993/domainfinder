#!/bin/bash

function go_build() {
  output=$1
  go build -o $output
}

echo "build domainfinder..." 
go_build domainfinder

cd ./synonyms
echo "build synonyms..."
go_build ../lib/synonyms

cd ../available
echo "build available..."
go_build ../lib/available

cd ../sprinkle
echo "build sprinkle..."
go_build ../lib/sprinkle

cd ../coolify
echo "build coolify..."
go_build ../lib/coolify

cd ../domainify
echo "build domainify..."
go_build ../lib/domainify

