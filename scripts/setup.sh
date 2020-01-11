#!/bin/bash

if [ -z ${CI_PIPELINE_ID+x} ]; then 
    echo "Not a gitlab CI run"
else
    export CC=/usr/bin/clang
    export GOPATH="$(pwd)/go"
    mkdir -p "$(pwd)/go"
    go env -w GOPATH="$(pwd)/go"
    mkdir -p $GOPATH/src/github.com/a-castellano
    ln -s ${CI_PROJECT_DIR} $GOPATH/src/github.com/a-castellano/metal-archives-wrapper
    cd $GOPATH/src/github.com/a-castellano/metal-archives-wrapper
fi
