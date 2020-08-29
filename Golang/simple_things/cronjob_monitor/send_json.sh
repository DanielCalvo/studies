#!/usr/bin/env bash
#This script would:
#Send cronjob info
#Wait 10 seconds for the job to finish (even if it fails!)

JOB_NAME="danitest"
START_TIME=`date +"%Y-%m-%dT%H:%M:%S%:z"`

sleep 1
EXIT_CODE=$?
END_TIME=`date +"%Y-%m-%dT%H:%M:%S%:z"`

JSON_STRING=$( jq -n \
                  --arg jb "$JOB_NAME" \
                  --arg st "$START_TIME" \
                  --arg ec "$EXIT_CODE" \
                  --arg et "$END_TIME" \
                  '{job_name: $jb, start_time: $st, exit_code: $ec, end_time: $et}' )

echo $JSON_STRING

curl localhost:9999 || true #fails
echo "Script finishes succesfully despite a curl command that fails!"

JSON_STRING=$(jq -n --arg st 2020-04-23T12:03:27+00:00 --arg es 0 --arg et 2020-04-23T12:03:58+00:00 --arg na e51d6910-4b6d-7708-0e27-c7aacf74197e --arg nj banana '{start_time: $st, exit_status: $es, end_time: $et, nomad_alloc_id: $na, nomad_job_name: $nj}')