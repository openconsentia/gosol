#!/bin/bash

COMMAND="$1"
SOLC_VER=${2:-0.4.26}

export ABIGEN_BUILDER_IMAGE=oc/abigenbuilder:current
export ABIGEN_TOOL_IMAGE=oc/abigentool

function usage(){
    echo "Usage: $COMMAND command <solc version>"
    echo " command:"
    echo "   image   create docker image"
    echo "   clean   remove images locally"
    echo " <solc version> default 0.4.24"
}


function image(){
    local ver="$1"
    docker build -f ./tools/abigenbuilder.dockerfile -t $ABIGEN_BUILDER_IMAGE .
    docker build --build-arg SOLC_VER=$ver --build-arg BUILDER_IMAGE=$ABIGEN_BUILDER_IMAGE -f ./tools/abigentool.dockerfile -t $ABIGEN_TOOL_IMAGE:$ver .
}

function clean(){
    local ver="$1"
    docker rmi -f $ABIGEN_BUILDER_IMAGE
    docker rmi -f $ABIGEN_TOOL_IMAGE:$ver
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

if [ "$#" -ne 2 ]; then
    usage
    exit 1
fi

case $COMMAND in
    "image")
        image $SOLC_VER
        ;;
    "clean")
        clean $SOLC_VER
        ;;
    *)
        usage
        ;;
esac



