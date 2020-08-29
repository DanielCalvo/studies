#!/bin/bash

echo "0: $0"
echo "1: $1"
echo "2: $2"
echo "3: $3"
echo "4: $4"
echo "${1#-}"

if [[ "$1" == "sh" ]]; then
    echo "This is a batch job"
    START_TIME=`date`
    "$@" #actually runs the command
    EXIT_STATUS=$?
    END_TIME=`date`
    echo "start time: $START_TIME"
    echo "exit status: $EXIT_STATUS"
    echo "end time: $END_TIME"
fi