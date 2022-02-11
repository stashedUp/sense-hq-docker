#!/bin/sh

dirname=$1
if [[ $dirname == "" ]];then
    echo "Dirname cannot be blank"
fi

docker run -v ${dirname}:/dist/tmp sense-docker